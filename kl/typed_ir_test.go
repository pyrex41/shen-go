package kl

import (
	"math"
	"strings"
	"testing"
)

func readTypedForm(t *testing.T, src string) Obj {
	t.Helper()
	o, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		t.Fatalf("read %q: %v", src, err)
	}
	return o
}

func TestTypeHintsPreserveAdvisoryAnnotations(t *testing.T) {
	fn := mustBytecodeFunc(CompileFunc("annotated", []Obj{MakeSymbol("x")}, readTypedForm(t, "(type x number)"))).fn
	if len(fn.TypeHints) != 1 {
		t.Fatalf("top-level hints = %#v, want one", fn.TypeHints)
	}
	h := fn.TypeHints[0]
	if h.Name != "number" || !h.Kinds.Contains(KindNumber) || h.Source != TypeHintSourceAnnotation {
		t.Fatalf("top-level hint = %#v", h)
	}

	outer := mustBytecodeFunc(CompileFunc("lambda-holder", nil, readTypedForm(t, "(lambda y (type y string))"))).fn
	var inner *BytecodeFunc
	for _, c := range outer.Consts {
		if c != nil && *c == scmHeadBytecodeFunc {
			inner = mustBytecodeFunc(c).fn
			break
		}
	}
	if inner == nil || len(inner.TypeHints) != 1 {
		t.Fatalf("lambda hints not preserved: %#v", outer.Consts)
	}
	if h := inner.TypeHints[0]; h.Name != "string" || !h.Kinds.Contains(KindString) || h.Source != TypeHintSourceAnnotation {
		t.Fatalf("lambda hint = %#v", inner.TypeHints[0])
	}
}

func TestUnknownTypeHintRemainsNonEnforcing(t *testing.T) {
	fn := mustBytecodeFunc(CompileFunc("unknown-annotation", nil, readTypedForm(t, "(type 7 imaginary-type)"))).fn
	if len(fn.TypeHints) != 1 || !fn.TypeHints[0].Kinds.Contains(KindUnknown) {
		t.Fatalf("unknown annotation metadata = %#v", fn.TypeHints)
	}
	var ctl ControlFlow
	if got := Call(&ctl, makeBytecodeObj(fn, nil)); IsError(got) || mustNumber(got) != 7 {
		t.Fatalf("annotation changed runtime result: %s", ObjString(got))
	}
}

func TestNumericIntrinsicOpcodeSelection(t *testing.T) {
	params := []Obj{MakeSymbol("x"), MakeSymbol("y")}
	cases := []struct {
		name string
		op   uint8
		form string
	}{
		{"add", OP_ADD, "(+ x y)"},
		{"sub", OP_SUB, "(- x y)"},
		{"mul", OP_MUL, "(* x y)"},
		{"lt", OP_LT, "(< x y)"},
		{"le", OP_LE, "(<= x y)"},
		{"gt", OP_GT, "(> x y)"},
		{"ge", OP_GE, "(>= x y)"},
		{"eq", OP_EQ, "(= x y)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bf := mustBytecodeFunc(CompileFunc(tc.name, params, readTypedForm(t, tc.form))).fn
			found := false
			for _, ins := range bf.Code {
				if ins.Op == tc.op {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("compiled %s without intrinsic opcode %d: %#v", tc.form, tc.op, bf.Code)
			}
		})
	}
}

func TestCompilerLiteralFixnumFolding(t *testing.T) {
	cases := []struct {
		name string
		form string
		want int
	}{
		{"add", "(+ 40 2)", 42},
		{"sub", "(- 40 2)", 38},
		{"mul", "(* 7 6)", 42},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bf := mustBytecodeFunc(CompileFunc(tc.name, nil, readTypedForm(t, tc.form))).fn
			for _, ins := range bf.Code {
				if ins.Op == OP_ADD || ins.Op == OP_SUB || ins.Op == OP_MUL {
					t.Fatalf("literal expression retained arithmetic opcode: %#v", bf.Code)
				}
			}
			if len(bf.Consts) != 1 || mustNumber(bf.Consts[0]) != float64(tc.want) {
				t.Fatalf("folded constants = %#v, want [%d]", bf.Consts, tc.want)
			}
		})
	}
}

func TestCompilerFoldingRetainsDynamicAndFloatArithmetic(t *testing.T) {
	for _, tc := range []struct {
		name string
		form string
		op   uint8
	}{
		{"float-add", "(+ 1.5 2.5)", OP_ADD},
		{"local-sub", "(- x 2)", OP_SUB},
		{"local-mul", "(* x 2)", OP_MUL},
	} {
		t.Run(tc.name, func(t *testing.T) {
			params := []Obj(nil)
			if tc.name[:5] == "local" {
				params = []Obj{MakeSymbol("x")}
			}
			bf := mustBytecodeFunc(CompileFunc(tc.name, params, readTypedForm(t, tc.form))).fn
			found := false
			for _, ins := range bf.Code {
				if ins.Op == tc.op {
					found = true
				}
			}
			if !found {
				t.Fatalf("dynamic expression lost opcode %d: %#v", tc.op, bf.Code)
			}
		})
	}
}

func TestNumericFastPathBoundariesAndFloatFallback(t *testing.T) {
	cases := []struct {
		name       string
		got        Obj
		want       float64
		wantFixnum bool
	}{
		{"fixnum", numAdd(MakeInteger(40), MakeInteger(2)), 42, true},
		{"upper-bound", numAdd(MakeInteger(fixnumMax-1), MakeInteger(1)), float64(fixnumMax), false},
		{"lower-bound", numSub(MakeInteger(fixnumMin+1), MakeInteger(1)), float64(fixnumMin), true},
		{"float-fallback", numAdd(MakeNumber(1.25), MakeInteger(2)), 3.25, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := mustNumber(tc.got); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if isFixnum(tc.got) != tc.wantFixnum {
				t.Fatalf("fixnum classification = %v, want %v", isFixnum(tc.got), tc.wantFixnum)
			}
		})
	}
}

func TestNumericOpcodeTypeErrorRemainsShenError(t *testing.T) {
	var ctl ControlFlow
	if got := evalString(&ctl, "(defun typed-add (x y) (+ x y))"); IsError(got) {
		t.Fatalf("defining typed-add: %s", ObjString(got))
	}
	got := Eval(&ctl, readTypedForm(t, "(typed-add true 1)"))
	if !IsError(got) {
		t.Fatalf("typed numeric mismatch returned %s, want Shen error", ObjString(got))
	}
}

func TestCompiledNumericRoundingAndOverflow(t *testing.T) {
	var ctl ControlFlow
	if IsError(evalString(&ctl, "(defun chain (x) (+ (+ x 1) 1))")) {
		t.Fatal("defining chain")
	}
	got := Eval(&ctl, readTypedForm(t, "(chain 9007199254740992)"))
	if IsError(got) || mustNumber(got) != 9007199254740992 {
		t.Fatalf("2^53 rounding chain = %s", ObjString(got))
	}
	if IsError(evalString(&ctl, "(defun overflow (x) (+ x 1))")) {
		t.Fatal("defining overflow")
	}
	got = Eval(&ctl, readTypedForm(t, "(overflow 9223372036854775807)"))
	if IsError(got) || mustNumber(got) != float64(9223372036854775808.0) {
		t.Fatalf("int64 overflow = %s", ObjString(got))
	}
}

func TestCompiledNonFiniteNumbersAndNaNIdentity(t *testing.T) {
	if equal(MakeNumber(math.NaN()), MakeNumber(math.NaN())) != False {
		t.Fatal("distinct NaNs should not compare equal")
	}
	nan := MakeNumber(math.NaN())
	if equal(nan, nan) != True {
		t.Fatal("same-object NaN equality must remain identity-true")
	}
	addFn := CompileFunc("nonfinite-add", []Obj{MakeSymbol("x"), MakeSymbol("y")}, readTypedForm(t, "(+ x y)"))
	var ctl ControlFlow
	if !math.IsNaN(mustNumber(Call(&ctl, addFn, nan, MakeInteger(0)))) {
		t.Fatal("NaN arithmetic lost NaN")
	}
	if !math.IsInf(mustNumber(Call(&ctl, addFn, MakeNumber(math.Inf(1)), MakeInteger(1))), 1) {
		t.Fatal("+Inf arithmetic lost infinity")
	}
	eqFn := CompileFunc("nan-eq", []Obj{MakeSymbol("x")}, readTypedForm(t, "(= x x)"))
	if Call(&ctl, eqFn, nan) != True {
		t.Fatal("compiled same-object NaN equality must remain identity-true")
	}
}

func TestCompiledClosurePartialAndOverApplication(t *testing.T) {
	var ctl ControlFlow
	if IsError(evalString(&ctl, "(defun make-adder (x) (lambda y (+ x y)))")) {
		t.Fatal("defining make-adder")
	}
	got := Eval(&ctl, readTypedForm(t, "((make-adder 3) 4)"))
	if IsError(got) || mustNumber(got) != 7 {
		t.Fatalf("closure result = %s", ObjString(got))
	}
	if IsError(evalString(&ctl, "(defun add2 (x y) (+ x y))")) {
		t.Fatal("defining add2")
	}
	got = Eval(&ctl, readTypedForm(t, "((add2 1) 2)"))
	if IsError(got) || mustNumber(got) != 3 {
		t.Fatalf("partial application = %s", ObjString(got))
	}
	got = Eval(&ctl, readTypedForm(t, "(add2 1 2 3)"))
	if !IsError(got) {
		t.Fatalf("over-application should be catchable error, got %s", ObjString(got))
	}
}
