package kl

import (
	"strings"
	"testing"
)

func TestGuardedPrimitiveOpcodeSelection(t *testing.T) {
	for _, tc := range []struct {
		name  string
		form  string
		arity int32
	}{
		{"divide", "(/ x y)", 2},
		{"concat", "(cn x y)", 2},
		{"tail-string", "(tlstr x)", 1},
		{"head", "(hd x)", 1},
		{"cons", "(cons x y)", 2},
		{"number-predicate", "(number? x)", 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r, err := NewSexpReader(strings.NewReader(tc.form), true).Read()
			if err != nil {
				t.Fatal(err)
			}
			bf := mustBytecodeFunc(CompileFunc(tc.name, []Obj{MakeSymbol("x"), MakeSymbol("y")}, r)).fn
			found := false
			for _, ins := range bf.Code {
				if (ins.Op == OP_GUARDED_PRIM && ins.A == tc.arity) || (ins.Op == OP_GP1_LOCAL && tc.arity == 1) || (tc.name == "divide" && ins.Op == OP_DIV) {
					found = true
					if ins.B < 0 || int(ins.B) >= len(bf.Consts) || !IsSymbol(bf.Consts[ins.B]) {
						t.Fatalf("guard missing symbol const: %#v", ins)
					}
				}
			}
			if !found {
				t.Fatalf("no guarded opcode in %#v", bf.Code)
			}
		})
	}
}

func TestGuardedPrimitiveSemantics(t *testing.T) {
	ctl := ControlFlow{}
	for _, tc := range []struct{ form, want string }{
		{"(cn \"a\" \"b\")", `"ab"`},
		{"(tlstr \"λx\")", `"x"`},
		{"(hd (cons 7 8))", "7"},
		{"(tl (cons 7 8))", "8"},
		{"(number? 7)", "true"},
	} {
		got := Eval(&ctl, readFormForVMTest(t, tc.form))
		if ObjString(got) != tc.want {
			t.Errorf("%s = %s, want %s", tc.form, ObjString(got), tc.want)
		}
	}
}

func TestGuardedPrimitiveTypeMissUsesCatchableFallback(t *testing.T) {
	ctl := ControlFlow{}
	for _, tc := range []struct{ form, want string }{
		{`(trap-error (cn 1 "a") (lambda E 42))`, "42"},
		{`(trap-error (hd 1) (lambda E 43))`, "43"},
		{`(trap-error (/ 1 0) (lambda E 44))`, "44"},
	} {
		got := Eval(&ctl, readFormForVMTest(t, tc.form))
		if ObjString(got) != tc.want {
			t.Errorf("%s = %s, want %s", tc.form, ObjString(got), tc.want)
		}
	}
}

func readFormForVMTest(t *testing.T, src string) Obj {
	t.Helper()
	o, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	return o
}

// TestFusedInstructionsKeepTheirGuards covers OP_GP1_LOCAL and OP_EQ_JF:
// the same answers as the unfused pair on the typed path, the same catchable
// errors on a type miss, the dynamic binding after a primitive is redefined,
// and no fusion across a jump that lands between the two parts.
func TestFusedInstructionsKeepTheirGuards(t *testing.T) {
	var ctl ControlFlow
	defs := []string{
		`(defun fz-hd (X) (hd X))`,
		`(defun fz-tl (X) (tl X))`,
		`(defun fz-consp (X) (if (cons? X) yes no))`,
		`(defun fz-str (X) (string? X))`,
		`(defun fz-eq (X Y) (if (= X Y) same different))`,
		`(defun fz-or (X Y) (if (or (= X 0) (= Y 0)) zero nonzero))`,
		`(defun fz-and (X Y) (if (and (= X 1) (= Y 1)) ones other))`,
		`(defun fz-cond (X) (cond ((= X 1) one) ((= X 2) two) (true many)))`,
	}
	for _, d := range defs {
		if r := evalString(&ctl, d); IsError(r) {
			t.Fatal(ObjString(r))
		}
	}
	fused := map[uint8]bool{}
	for _, name := range []string{"fz-hd", "fz-consp", "fz-eq", "fz-cond"} {
		for _, ins := range mustBytecodeFunc(PrimFunc(MakeSymbol(name))).fn.Code {
			fused[ins.Op] = true
		}
	}
	if !fused[OP_GP1_LOCAL] || !fused[OP_EQ_JF] {
		t.Fatalf("expected both fused instructions to be emitted; saw ops %v", fused)
	}
	for _, c := range [][2]string{
		{`(fz-hd (cons 1 2))`, "1"},
		{`(fz-tl (cons 1 2))`, "2"},
		{`(fz-consp (cons 1 2))`, "yes"},
		{`(fz-consp ())`, "no"},
		{`(fz-consp 3)`, "no"},
		{`(fz-consp 2.5)`, "no"},
		{`(fz-str "s")`, "true"},
		{`(fz-eq 1 1)`, "same"},
		{`(fz-eq 1 1.0)`, "same"},
		{`(fz-eq (cons 1 ()) (cons 1 ()))`, "same"},
		{`(fz-eq 1 2)`, "different"},
		{`(fz-or 0 5)`, "zero"},
		{`(fz-or 5 0)`, "zero"},
		{`(fz-or 5 5)`, "nonzero"},
		{`(fz-and 1 1)`, "ones"},
		{`(fz-and 1 2)`, "other"},
		{`(fz-and 2 1)`, "other"},
		{`(fz-cond 1)`, "one"},
		{`(fz-cond 2)`, "two"},
		{`(fz-cond 3)`, "many"},
	} {
		if got := ObjString(evalString(&ctl, c[0])); got != c[1] {
			t.Errorf("%s = %s, want %s", c[0], got, c[1])
		}
	}
	for _, bad := range []string{`(fz-hd 5)`, `(fz-tl "s")`, `(fz-hd 2.5)`} {
		if r := evalString(&ctl, `(trap-error `+bad+` (lambda E caught))`); ObjString(r) != "caught" {
			t.Errorf("%s under trap-error = %s, want caught", bad, ObjString(r))
		}
	}
	// Redefine the primitives after compilation: the fused sites must take
	// the new bindings.
	hd, eq := MakeSymbol("hd"), MakeSymbol("=")
	oldHd, oldEq := PrimFunc(hd), PrimFunc(eq)
	defer BindSymbolFunc(hd, oldHd)
	defer BindSymbolFunc(eq, oldEq)
	BindSymbolFunc(hd, MakeNative(func(e *ControlFlow) { e.Return(MakeSymbol("my-hd")) }, 1))
	if got := ObjString(evalString(&ctl, `(fz-hd (cons 1 2))`)); got != "my-hd" {
		t.Errorf("after redefining hd, (fz-hd (cons 1 2)) = %s, want my-hd", got)
	}
	BindSymbolFunc(eq, MakeNative(func(e *ControlFlow) { e.Return(True) }, 2))
	if got := ObjString(evalString(&ctl, `(fz-eq 1 2)`)); got != "same" {
		t.Errorf("after redefining =, (fz-eq 1 2) = %s, want same", got)
	}
	BindSymbolFunc(eq, MakeNative(func(e *ControlFlow) { e.Return(MakeSymbol("maybe")) }, 2))
	if r := evalString(&ctl, `(trap-error (fz-eq 1 2) (lambda E caught))`); ObjString(r) != "caught" {
		t.Errorf("a non-boolean = under if gave %s, want the catchable if error", ObjString(r))
	}
}
