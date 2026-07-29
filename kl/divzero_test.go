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

// The guard has to hold for every shape of zero divisor, not just the literal
// integer 0 that issue #10/#16 quoted. A divisor that is only zero after
// evaluation (or that is a negative/float zero) reaches PrimNumberDivide the
// same way, and each one used to produce a different bogus saturated value
// (+maxint, -maxint, or a NaN that printed as maxint) rather than an error.
func TestDivideByZeroDivisorShapes(t *testing.T) {
	var ctx ControlFlow

	for _, expr := range []string{
		`(/ 1 0)`,       // literal integer zero
		`(/ -1 0)`,      // negative numerator -> used to saturate to -maxint
		`(/ 0 0)`,       // 0/0 -> used to be NaN
		`(/ 1.5 0)`,     // float numerator, integer zero divisor
		`(/ 1 0.0)`,     // float zero divisor
		`(/ 1 -0.0)`,    // negative float zero
		`(/ 1 (- 5 5))`, // divisor only zero after evaluation
	} {
		src := `(trap-error ` + expr + ` (lambda E (error-to-string E)))`
		if got := ObjString(evalString(&ctx, src)); got != `"division by zero"` {
			t.Errorf("%s: got %s, want %q", expr, got, "division by zero")
		}
	}
}

// Issue #10's fix only guarded the tree-walking evaluator's primitive. A defun
// is compiled to bytecode and runs on the VM, which dispatches arithmetic
// through its own fast paths -- so the VM path needs its own proof that the
// divisor guard is not bypassed, both when the zero is a constant folded into
// the function body and when it arrives as an argument at runtime.
func TestDivideByZeroOnBytecodeVM(t *testing.T) {
	var ctx ControlFlow

	if res := evalString(&ctx, `(defun dz-const () (/ 1 0))`); IsError(res) {
		t.Fatalf("defun dz-const: %s", ObjString(res))
	}
	if res := evalString(&ctx, `(defun dz-arg (X Y) (/ X Y))`); IsError(res) {
		t.Fatalf("defun dz-arg: %s", ObjString(res))
	}

	for _, tc := range []struct{ name, call string }{
		{"constant zero divisor in compiled body", `(dz-const)`},
		{"integer zero divisor passed as argument", `(dz-arg 7 0)`},
		{"float zero divisor passed as argument", `(dz-arg 7.5 0.0)`},
		{"divisor computed to zero at runtime", `(dz-arg 7 (- 3 3))`},
	} {
		src := `(trap-error ` + tc.call + ` (lambda E (error-to-string E)))`
		if got := ObjString(evalString(&ctx, src)); got != `"division by zero"` {
			t.Errorf("%s: %s got %s, want %q", tc.name, tc.call, got, "division by zero")
		}
	}

	// The compiled function is still usable for non-zero divisors afterwards.
	if got := ObjString(evalString(&ctx, `(dz-arg 10 4)`)); got != "2.5" {
		t.Errorf("(dz-arg 10 4): got %q, want 2.5", got)
	}
}

// An untrapped divide by zero must surface as an error object, not as a number.
// The original bug was precisely that it came back as a perfectly ordinary
// (but wrong) number, so nothing downstream could tell the difference.
func TestDivideByZeroIsNotANumber(t *testing.T) {
	var ctx ControlFlow

	res := evalString(&ctx, `(/ 1 0)`)
	if !IsError(res) {
		t.Fatalf("(/ 1 0) returned a non-error %s; want an error object", ObjString(res))
	}
	if got := ObjString(res); got == "9223372036854775807" || got == "-9223372036854775808" {
		t.Fatalf("(/ 1 0) regressed to the saturated maxint value %s", got)
	}
}
