package codegen

import (
	"bytes"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

func readCodegenForm(t *testing.T, src string) kl.Obj {
	t.Helper()
	o, err := kl.NewSexpReader(strings.NewReader(src), false).Read()
	if err != nil {
		t.Fatalf("read %q: %v", src, err)
	}
	return o
}

func TestGenerateExprAcceptsTypedWrapper(t *testing.T) {
	for _, src := range []string{"($type number ($const 1.5))", "($type ($const 1.5) number)"} {
		var got bytes.Buffer
		if err := New().generateExpr(&got, readCodegenForm(t, src)); err != nil {
			t.Fatal(err)
		}
		if got.String() != "MakeNumber(1.5)" {
			t.Fatalf("%s generated %q", src, got.String())
		}
	}
}

func TestTypedWrapperStillEnablesGuardedScalarLowering(t *testing.T) {
	var got bytes.Buffer
	form := readCodegenForm(t, "($type number (call ($global +) ($const 1) ($const 2)))")
	if err := New().generateExpr(&got, form); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got.String(), "TypedMaterializeNumber((__typedN0 + __typedN1))") {
		t.Fatalf("typed wrapper blocked scalar lowering:\n%s", got.String())
	}
}

func TestScalarPrimitiveLoweringHasDynamicFallback(t *testing.T) {
	var got bytes.Buffer
	form := readCodegenForm(t, "(call ($global +) ($const 1) ($const 2))")
	if err := New().generateExpr(&got, form); err != nil {
		t.Fatal(err)
	}
	s := got.String()
	for _, want := range []string{
		"TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7)",
		"TypedMaterializeNumber((__typedN0 + __typedN1))",
		"Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)",
	} {
		if !strings.Contains(s, want) {
			t.Errorf("generated code missing %q:\n%s", want, s)
		}
	}
}

func TestScalarLoweringStringsAndBooleans(t *testing.T) {
	cases := []struct{ src, want string }{
		{"(call ($global /) (call ($global /) ($const 8) ($const 2)) ($const 2))", "TypedDivideValue(TypedDivideValue(__typedN0, __typedN1), __typedN2)"},
		{"(call ($global cn) ($const \"a\") ($const \"🙂\"))", "TypedMaterializeString((__typedS0 + __typedS1))"},
		{"(call ($global tlstr) ($const \"🙂x\"))", "TypedMaterializeString(TypedStringTailValue(__typedS0))"},
		{"(call ($global not) (call ($global <) ($const 1) ($const 2)))", "TypedMaterializeBoolean((!(__typedN0 < __typedN1)))"},
	}
	for _, tc := range cases {
		var got bytes.Buffer
		if err := New().generateExpr(&got, readCodegenForm(t, tc.src)); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(got.String(), tc.want) {
			t.Errorf("generated code missing %q:\n%s", tc.want, got.String())
		}
	}
}

func TestBlockScalarFusionAcrossTemps(t *testing.T) {
	form := readCodegenForm(t, `(block
	(<= x (call ($global +) ($const 1) ($const 2)))
	(<= y (call ($global *) x ($const 4)))
	(return y))`)
	var got bytes.Buffer
	if err := New().generateExpr(&got, form); err != nil {
		t.Fatal(err)
	}
	s := got.String()
	if !strings.Contains(s, "TypedMaterializeNumber(((__typedN0 + __typedN1) * __typedN2))") {
		t.Fatalf("flattened scalar chain was not fused:\n%s", s)
	}
	if strings.Contains(s, "x :=") || strings.Contains(s, "y :=") {
		t.Fatalf("single-use scalar temporaries leaked into generated code:\n%s", s)
	}
}
