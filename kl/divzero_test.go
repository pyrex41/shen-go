// Issue #10: integer (and float) division by zero must raise a catchable
// kernel error, not surface as a bogus maxint (9223372036854775807).
//
// Go's float64 `x / 0` yields +Inf/-Inf/NaN; the kernel number->string path
// then truncated that to int, printing 9223372036854775807. That value was
// neither a real number nor catchable by (trap-error ...). PrimNumberDivide
// now panics with a standard MakeError, matching shen-cl's divide-by-zero.
package kl

import "testing"

func TestDivideByZeroIsCatchable(t *testing.T) {
	var ctx ControlFlow

	// (trap-error (/ 1 0) ...) must catch the error and never see maxint.
	got := ObjString(evalString(&ctx,
		`(trap-error (/ 1 0) (lambda E (error-to-string E)))`))
	if got != `"division by zero"` {
		t.Fatalf("trap-error (/ 1 0): got %s, want %q", got, "division by zero")
	}

	// Float zero divisor is guarded the same way.
	gotF := ObjString(evalString(&ctx,
		`(trap-error (/ 1.0 0.0) (lambda E (error-to-string E)))`))
	if gotF != `"division by zero"` {
		t.Fatalf("trap-error (/ 1.0 0.0): got %s, want %q", gotF, "division by zero")
	}

	// Evaluator stays usable after the trapped error.
	if cont := ObjString(evalString(&ctx, "(+ 40 2)")); cont != "42" {
		t.Fatalf("post-error eval: got %q, want 42", cont)
	}

	// Ordinary division still works (and prints in shortest round-trip form).
	if ok := ObjString(evalString(&ctx, "(/ 10 4)")); ok != "2.5" {
		t.Fatalf("(/ 10 4): got %q, want 2.5", ok)
	}
}
