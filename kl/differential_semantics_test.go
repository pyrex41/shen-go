package kl

// These tests compare the tree evaluator with the bytecode VM.  The two
// implementations are intentionally exercised through the same Shen source:
// the VM side wraps the source in a freshly compiled zero-argument function.
// Keeping this oracle in tests makes optimisations (typed slots, constant
// folding, and future typed IR) prove that they preserve observable Shen
// behaviour.

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
)

func evalDifferentialSource(t *testing.T, body string) (direct, compiled Obj) {
	t.Helper()
	read := func(src string) Obj {
		r := NewSexpReader(strings.NewReader(src), true)
		sexp, err := r.Read()
		if err != nil {
			t.Fatalf("read %q: %v", src, err)
		}
		var ctl ControlFlow
		return Eval(&ctl, sexp)
	}
	direct = read(body)
	compiled = read(fmt.Sprintf("(do (defun __differential_probe () %s) (__differential_probe))", body))
	return direct, compiled
}

func evalCompiledTypedMode(t *testing.T, body, mode string) Obj {
	t.Helper()
	old, had := os.LookupEnv("SHEN_GO_TYPED_IR")
	if err := os.Setenv("SHEN_GO_TYPED_IR", mode); err != nil {
		t.Fatalf("set SHEN_GO_TYPED_IR=%s: %v", mode, err)
	}
	ResetTypedIRModeForTest()
	defer func() {
		if had {
			_ = os.Setenv("SHEN_GO_TYPED_IR", old)
		} else {
			_ = os.Unsetenv("SHEN_GO_TYPED_IR")
		}
		ResetTypedIRModeForTest()
	}()
	_, compiled := evalDifferentialSource(t, body)
	return compiled
}

func TestDifferentialShenSemantics(t *testing.T) {
	cases := []struct {
		name string
		body string
	}{
		{"annotation-preserves-value", `(type (+ 20 22) number)`},
		{"annotation-mismatch-is-catchable", `(trap-error (type (+ true 1) number) (lambda E (error-to-string E)))`},
		{"rounding-past-2^53", `(+ (+ 9007199254740992 1) 1)`},
		{"integer-range-overflow", `(+ 9223372036854775807 1)`},
		{"mixed-fixnum-float-equality", `(= 1 1.0)`},
		{"divide-by-zero", `(trap-error (/ 1 0) (lambda E (error-to-string E)))`},
		{"infinity", `(* 1e308 1e308)`},
		{"boolean-branch", `(if (and true (not false)) 11 22)`},
		{"non-boolean-condition", `(trap-error (if 1 2 3) (lambda E (error-to-string E)))`},
		{"unicode-concatenation", `(cn "λ" "雪")`},
		{"unicode-string-predicate", `(string? "λ雪")`},
		{"symbol-identity", `(= foo foo)`},
		{"pair-and-nil", `(do (set P (cons 7 (cons 8 ()))) (= (hd (value P)) 7))`},
		{"vector-read-write", `(do (set V (absvector 1)) (address-> (value V) 0 7) (= (<-address (value V) 0) 7))`},
		{"closure-upvalue", `(do (defun make (X) (lambda Y (+ X Y))) ((make 3) 4))`},
		{"partial-application", `(do (defun add (X Y) (+ X Y)) ((add 3) 4))`},
		{"over-application-error", `(do (defun add (X Y) (+ X Y)) (trap-error (add 1 2 3) (lambda E (error-to-string E))))`},
		{"tail-recursion", `(do (defun down (N) (if (= N 0) true (down (- N 1)))) (down 3000))`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			direct, compiled := evalDifferentialSource(t, tc.body)
			if ObjString(direct) != ObjString(compiled) {
				t.Fatalf("tree=%q (%T), VM=%q (%T)", ObjString(direct), direct, ObjString(compiled), compiled)
			}
			on := evalCompiledTypedMode(t, tc.body, "on")
			off := evalCompiledTypedMode(t, tc.body, "off")
			if ObjString(on) != ObjString(off) {
				t.Fatalf("typed-ir mode diverged: on=%q, off=%q", ObjString(on), ObjString(off))
			}
		})
	}
}

// Primitive rebinding is a particularly important differential case: an
// intrinsic fast path must not silently bypass a dynamically replaced global.
func TestDifferentialPrimitiveRebinding(t *testing.T) {
	plus := mustSymbol(MakeSymbol("+"))
	old := plus.function
	defer func() { plus.function = old }()
	body := `(do (defun + (X Y) 99) (+ 1 2))`
	direct, compiled := evalDifferentialSource(t, body)
	if ObjString(direct) != ObjString(compiled) {
		t.Fatalf("rebinding diverged: tree=%q, VM=%q", ObjString(direct), ObjString(compiled))
	}
}

func TestDifferentialPrimitiveRebindingAfterCompile(t *testing.T) {
	plus := mustSymbol(MakeSymbol("+"))
	old := plus.function
	defer func() { plus.function = old }()
	fn := CompileFunc("rebind-after-compile", nil, readTypedForm(t, `(+ 1 2)`))
	var ctl ControlFlow
	if got := Eval(&ctl, readTypedForm(t, `(defun + (X Y) 99)`)); IsError(got) {
		t.Fatalf("rebind +: %s", ObjString(got))
	}
	got := Call(&ctl, fn)
	if ObjString(got) != "99" {
		t.Fatalf("compiled function bypassed rebind: got %q", ObjString(got))
	}
}

func TestTypedRuntimeNonFiniteAndNaNIdentity(t *testing.T) {
	// These helpers are the scalar boundary used by AOT typed regions. Keep
	// IEEE-754 payloads and Shen's identity-sensitive NaN equality intact.
	nan := TypedMaterializeNumber(math.NaN())
	if !math.IsNaN(GetNumber(TypedDivide(math.NaN(), 1))) {
		t.Fatal("typed divide lost NaN")
	}
	if equal(nan, nan) != True {
		t.Fatal("same-object NaN equality must remain true")
	}
	if equal(TypedMaterializeNumber(math.NaN()), TypedMaterializeNumber(math.NaN())) != False {
		t.Fatal("distinct NaNs must remain unequal")
	}
	if !math.IsInf(GetNumber(TypedMaterializeNumber(math.Inf(1))), 1) {
		t.Fatal("typed materialization lost +Inf")
	}
}
