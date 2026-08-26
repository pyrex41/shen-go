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
				if (ins.Op == OP_GUARDED_PRIM && ins.A == tc.arity) || (tc.name == "divide" && ins.Op == OP_DIV) {
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
