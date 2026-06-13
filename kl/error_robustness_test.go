// Adversarial robustness tests for the kl evaluator.
//
// These freeze the contract that every documented kernel error path:
//
//  1. raises a Shen-catchable error (so (trap-error ...) handles it),
//  2. surfaces a stable, informative message (so callers can match on it),
//  3. leaves the evaluator state clean enough that the next eval succeeds.
//
// Both evaluation paths must satisfy the contract — top-level forms run
// through the tree-walker (evalExp), defun bodies run through the bytecode
// VM (vmExec). The two paths emit different error messages for the same
// trigger (e.g. "can't apply non function" vs "function not bound"), so the
// table carries a `want` per path.
//
// Background: the canonical Shen kernel suite (134/134 in
// ./certification) exercises only well-formed code, so it never catches
// regressions on adversarial inputs. The pre-existing trap-error tests in
// eval_test.go put `(undefined-fn ...)` only in *dead* and/or short-circuited
// branches — they prove the branch isn't taken, not that the error would be
// catchable if it were. A previous bug let `evalFunction` panic with a bare
// Go string, which Try's recovery couldn't classify; the result cascaded
// into IsError(nil) and SIGSEGV'd the process. These tests make that class
// of regression visible.
package kl

import (
	"fmt"
	"strings"
	"testing"
)

// TestErrorIsCatchable verifies that user-reachable error paths are
// catchable via (trap-error ...) on both the tree-walked and bytecode-compiled
// evaluation paths, and that the evaluator continues to work afterwards.
func TestErrorIsCatchable(t *testing.T) {
	type want struct {
		// wantTree is the (error-to-string E) result when the trigger is
		// evaluated at the top level (tree-walker).
		wantTree string
		// wantVM is the same when the trigger is wrapped in a defun and
		// invoked, forcing the bytecode VM path.
		// Empty string means "skip the VM variant" — used for cases where
		// the VM path hits a sibling uncatchable-panic site that has not
		// yet been fixed (tracked separately; this test must stay green).
		wantVM string
	}
	cases := []struct {
		name    string
		trigger string
		want
	}{
		{
			name:    "apply unbound symbol",
			trigger: "(overflow->str)",
			want: want{
				wantTree: `"can't apply non function: overflow->str"`,
				wantVM:   `"function overflow->str not bound"`,
			},
		},
		{
			name:    "apply non-function literal",
			trigger: "(42 1)",
			want: want{
				wantTree: `"can't apply non function: 42"`,
				// VM path lowers `(42 1)` through apply() and hits
				// eval.go:242, which now raises a catchable MakeError with
				// the same message as the tree-walker rather than a bare
				// Go-string panic (the old "trap-error result is not Obj"
				// hole). Both paths agree.
				wantVM: `"can't apply non function: 42"`,
			},
		},
		{
			name:    "value of unbound variable",
			trigger: "(value never-bound-xyz)",
			want: want{
				wantTree: `"variable never-bound-xyz not bound"`,
				wantVM:   `"variable never-bound-xyz not bound"`,
			},
		},
		{
			name:    "if requires boolean",
			trigger: "(if 42 1 2)",
			want: want{
				wantTree: `"if requires a boolean"`,
				wantVM:   `"if requires a boolean"`,
			},
		},
		{
			name:    "simple-error explicit",
			trigger: `(simple-error "oops")`,
			want: want{
				wantTree: `"oops"`,
				wantVM:   `"oops"`,
			},
		},
	}

	for _, c := range cases {
		c := c
		// Each subtest gets a fresh ControlFlow so a regression in one
		// case can't poison the rest. A clean trap-error must also leave
		// no residue on the stack — the assertion at the end of the
		// subtest catches that.
		t.Run("tree/"+c.name, func(t *testing.T) {
			var ctx ControlFlow
			expr := fmt.Sprintf(`(trap-error %s (lambda E (error-to-string E)))`, c.trigger)
			got := ObjString(evalString(&ctx, expr))
			if got != c.wantTree {
				t.Fatalf("tree-walker: got %q, want %q", got, c.wantTree)
			}
			// State is clean enough for the next eval to succeed.
			if cont := ObjString(evalString(&ctx, "(+ 40 2)")); cont != "42" {
				t.Fatalf("post-error eval: got %q, want 42", cont)
			}
			if ctx.pos != 0 {
				t.Fatalf("residual sp after caught error: %d", ctx.pos)
			}
		})

		if c.wantVM == "" {
			continue
		}
		t.Run("vm/"+c.name, func(t *testing.T) {
			var ctx ControlFlow
			evalString(&ctx, fmt.Sprintf(`(defun bad-fn () %s)`, c.trigger))
			got := ObjString(evalString(&ctx,
				`(trap-error (bad-fn) (lambda E (error-to-string E)))`))
			if got != c.wantVM {
				t.Fatalf("bytecode VM: got %q, want %q", got, c.wantVM)
			}
			if cont := ObjString(evalString(&ctx, "(+ 40 2)")); cont != "42" {
				t.Fatalf("post-error eval: got %q, want 42", cont)
			}
			if ctx.pos != 0 {
				t.Fatalf("residual sp after caught error: %d", ctx.pos)
			}
		})
	}
}

// TestTryRecoversNonObjPanic pins the contract added to Try in eval.go:
// when a callee panics with a payload that isn't a Shen error Obj
// (e.g. a bare Go string from a still-rough panic site), Try must produce
// a catchable error rather than the zero-value tryResult{data: nil}.
//
// Why this matters: a nil tryResult.data flows into Catch -> IsError(nil),
// and IsError used to deref its arg unconditionally — that was the SIGSEGV
// that killed the process. With both fixes in place (the bare panic site
// now uses MakeError, and Try's recover wraps stray panics in MakeError),
// either layer alone is enough to keep the process alive, and both
// together is the belt-and-braces.
func TestTryRecoversNonObjPanic(t *testing.T) {
	var ctl ControlFlow
	bomb := MakeNative(func(*ControlFlow) {
		panic("synthetic bare-string panic")
	}, 0)

	res := Try(&ctl, bomb)
	if res.data == nil {
		t.Fatal("Try.data is nil — recovery should always publish a real Obj")
	}
	if !IsError(res.data) {
		t.Fatalf("Try.data is %q, expected an error Obj", ObjString(res.data))
	}
	// Catch must accept the recovered error and call the handler.
	handler := evalString(&ctl, `(lambda E "caught")`)
	if got := ObjString(res.Catch(handler)); got != `"caught"` {
		t.Fatalf("Catch: got %q, want %q", got, `"caught"`)
	}

	// The recovered message should carry the original payload somewhere
	// in its body. We don't pin the exact format (it goes through
	// fmt.Sprintf "%v") — just verify the payload survives, so future
	// readers can recognize what panicked.
	errMsg := ObjString(res.data)
	if !strings.Contains(errMsg, "synthetic bare-string panic") {
		t.Fatalf("recovered error %q does not mention the original panic payload", errMsg)
	}
}

// TestEvalSurvivesAdversarialSequence drives the evaluator through several
// errors in a row on the same ControlFlow and asserts that state stays
// clean. The original bug only surfaced when a malformed form fired *and*
// the cascade reached the next eval; locking this end-to-end keeps that
// regression visible even if individual layers regress separately.
func TestEvalSurvivesAdversarialSequence(t *testing.T) {
	var ctx ControlFlow
	seq := []struct {
		expr string
		want string
	}{
		{`(trap-error (overflow->str) (lambda E (error-to-string E)))`,
			`"can't apply non function: overflow->str"`},
		{`(trap-error (value not-bound-1) (lambda E (error-to-string E)))`,
			`"variable not-bound-1 not bound"`},
		{`(trap-error (if 42 1 2) (lambda E (error-to-string E)))`,
			`"if requires a boolean"`},
		{`(trap-error (simple-error "boom") (lambda E (error-to-string E)))`,
			`"boom"`},
		{`(+ 40 2)`, `42`},
	}
	for i, step := range seq {
		got := ObjString(evalString(&ctx, step.expr))
		if got != step.want {
			t.Fatalf("step %d %q: got %q, want %q", i, step.expr, got, step.want)
		}
		if ctx.pos != 0 {
			t.Fatalf("step %d: residual sp %d", i, ctx.pos)
		}
	}
}
