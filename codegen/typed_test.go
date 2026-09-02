package codegen

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// nestedCnChainSrc returns a right-nested `cn` message chain with n string
// leaves, e.g. for n=3: (call ($global cn) ($const "a0") (call ($global cn)
// ($const "a1") ($const "a2"))).
func nestedCnChainSrc(n int) string {
	if n < 2 {
		panic("nestedCnChainSrc: need at least 2 leaves")
	}
	var build func(i int) string
	build = func(i int) string {
		if i == n-1 {
			return fmt.Sprintf("($const %q)", fmt.Sprintf("a%d", i))
		}
		return fmt.Sprintf("(call ($global cn) ($const %q) %s)", fmt.Sprintf("a%d", i), build(i+1))
	}
	return build(0)
}

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

// TestNestedCnChainTenPlusArgsTypedTempsMatch is a regression test for a
// declared-vs-referenced __typedS<N> temporary mismatch in nested `cn`
// (string-concat) chains with 10+ arguments.
//
// Mechanism: scalarExpr fuses a nested cn chain into one native Go
// expression using placeholders __typedV<N> (N = leaf index), then merges a
// nested subexpression's *local* placeholders into the outer expression's
// combined index space by shifting each local index up by `off`
// (codegen.go's scalarExpr, the "off := len(leaves)" block). The prior code
// did this shift with one strings.ReplaceAll call per index. That is unsafe:
// the search key for a low index is a literal-string *prefix* of any
// placeholder whose index starts with the same digit(s) -- "__typedV1" is a
// prefix of "__typedV10" and "__typedV11". Once a chain accumulates 10+
// leaves, shifting index 1 also mangles the already-shifted "__typedV10"
// token embedded in the same expression, corrupting it into e.g.
// "__typedV20". The corrupted index then survives the later V-to-kind-letter
// rename (primitiveCallOptimize, same bare-ReplaceAll pattern) verbatim,
// producing generated Go that declares __typedS10/__typedS11 but references
// __typedS20/__typedS30 -- an "undefined: __typedS20" compile error.
//
// Before the fix, this test failed with:
//
//	generated source references undeclared temporary __typedS20:
//	...__typedS9 + (__typedS20 + __typedS30))...
//
// (declared set only went up to __typedS11; observed while diagnosing this
// bug on the pre-fix tree at commit 4b93aa0).
//
// The fix (rewriteTypedVPlaceholders in codegen.go) replaces both
// sequential-ReplaceAll call sites with a single regexp pass keyed on a full
// digit run, so indices can never bleed into one another.
func TestNestedCnChainTenPlusArgsTypedTempsMatch(t *testing.T) {
	const nArgs = 12 // 11 concatenations; below 10 leaves this bug does not reproduce.
	form := readCodegenForm(t, nestedCnChainSrc(nArgs))
	var got bytes.Buffer
	if err := New().generateExpr(&got, form); err != nil {
		t.Fatal(err)
	}
	s := got.String()

	declRe := regexp.MustCompile(`(__typedS\d+), __typedOK\d+ :=`)
	declared := map[string]bool{}
	for _, m := range declRe.FindAllStringSubmatch(s, -1) {
		declared[m[1]] = true
	}
	refRe := regexp.MustCompile(`__typedS\d+`)
	for _, ref := range refRe.FindAllString(s, -1) {
		if !declared[ref] {
			t.Fatalf("generated source references undeclared temporary %s:\n%s", ref, s)
		}
	}
}

// TestNestedCnChainTenPlusArgsCompiles is the end-to-end counterpart of
// TestNestedCnChainTenPlusArgsTypedTempsMatch: it writes a full generated
// module (via HandleBodyObj, the same path cmd/kl and yggdrasil-build use)
// to a temporary package below this one (so its go.mod supplies the kl
// import) and asks the Go toolchain to actually compile it. This is the
// most direct proof that the undefined-temporary bug is fixed: before the
// fix, this failed with "undefined: __typedS20" (and __typedS30) from the
// Go compiler itself.
func TestNestedCnChainTenPlusArgsCompiles(t *testing.T) {
	const nArgs = 14 // comfortably into the double-digit collision zone.
	bc := readCodegenForm(t, "(return "+nestedCnChainSrc(nArgs)+")")

	cg := New()
	var src bytes.Buffer
	if err := cg.HandleBodyObj(bc, "Generated", &src); err != nil {
		t.Fatal(err)
	}
	cg.HandleSymbol(&src)

	dir, err := os.MkdirTemp(".", "_typedcnchain-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	if err := os.WriteFile(filepath.Join(dir, "generated.go"), src.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated Go failed to compile: %v\n%s\nsource:\n%s", err, out, src.String())
	}
}
