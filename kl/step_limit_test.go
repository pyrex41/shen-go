// Step-limit tests for the evaluator.
//
// Background: a KLambda program may legally never terminate. The fuzzer
// (FuzzReaderEval) readily synthesises one — `(do (defun f (X) (f X)) (f 0))`
// tail-recurses forever through the trampoline without growing the Go stack,
// so no panic ever fires and the worker just spins. That is the halting
// problem, not a memory-safety bug, but to a fuzzer it looks like a stuck
// worker that deadlines the run and hides genuine crashers.
//
// ControlFlow.stepLimit gives callers (the fuzz harness) an opt-in budget:
// when > 0, evaluation aborts with a panic(MakeError(...)) after that many
// steps. The default zero is unlimited, so production is unaffected.
//
// The budget is a deliberately *non-resumable* circuit breaker: the step
// counter is shared across the whole Eval and is never reset, so once it
// trips it stays tripped. Resetting it on a caught error would re-open the
// hang — `(do (defun l (X) (trap-error (l X) (lambda E (l X)))) (l 0))` would
// reset the budget on every catch and run forever again. The point is a hard
// ceiling on total work, not a catchable-and-resumable error. At the top
// level the breaker reaches Eval's recover, which (as for any escaping error)
// returns a real, ObjString-able error Obj — exactly what the fuzzer needs to
// keep moving.
package kl

import (
	"strings"
	"testing"
	"time"
)

// evalStringWithLimit evaluates exp on a fresh ControlFlow carrying the given
// step budget.
func evalStringWithLimit(exp string, limit int64) Obj {
	r := NewSexpReader(strings.NewReader(exp), true)
	sexp, err := r.Read()
	if err != nil {
		panic(err)
	}
	ctx := ControlFlow{stepLimit: limit}
	return Eval(&ctx, sexp)
}

// runWithin runs fn, reporting whether it finished inside d. It exists so a
// regression that re-introduces an unbounded loop fails as a timeout instead
// of hanging the whole test binary.
func runWithin(d time.Duration, fn func()) bool {
	done := make(chan struct{})
	go func() {
		defer func() { recover(); close(done) }()
		fn()
	}()
	select {
	case <-done:
		return true
	case <-time.After(d):
		return false
	}
}

// TestStepLimitHaltsNonTerminating proves the budget actually stops the two
// shapes of legal non-termination the evaluator supports, and that Eval
// returns a usable error Obj rather than hanging — the exact contract the
// fuzz harness depends on.
func TestStepLimitHaltsNonTerminating(t *testing.T) {
	cases := []struct {
		name  string
		defs  []string
		final string
	}{
		// Self-tail-call: compiles to OP_SELF_TAIL_CALL and loops on pc=0
		// inside vmExec's for{} — never returns to the trampoline. Guards
		// the vm.go tick().
		{"self-tail-call", []string{"(defun f (X) (f X))"}, "(f 0)"},
		// Mutual recursion bounces f->g->f through OP_TAIL_CALL and the
		// trampoline. Guards the eval.go trampoline tick().
		{"mutual-tail-call",
			[]string{"(defun f (X) (g X))", "(defun g (X) (f X))"}, "(f 0)"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			var got Obj
			finished := runWithin(5*time.Second, func() {
				ctx := ControlFlow{stepLimit: 1_000_000}
				for _, d := range c.defs {
					evalString(&ctx, d)
				}
				got = evalString(&ctx, c.final)
			})
			if !finished {
				t.Fatalf("step limit did not halt %q within 5s", c.final)
			}
			// The fuzzer's contract: a real, renderable Obj (not nil), and
			// it must be an error (the program never produced a value).
			if got == nil {
				t.Fatalf("Eval returned nil for %q", c.final)
			}
			if !IsError(got) {
				t.Fatalf("expected an error Obj, got %q", ObjString(got))
			}
			_ = ObjString(got) // must not panic
		})
	}
}

// TestStepLimitMessageViaTry pins the breaker's actual payload. At the top
// level Eval's recover replaces escaping errors with a stack trace, so to see
// the raw error we go through Try, whose recover preserves a panicked error
// Obj verbatim — the same path the compiled/certification runtime uses.
func TestStepLimitMessageViaTry(t *testing.T) {
	ctx := ControlFlow{stepLimit: 1_000_000}
	// Define a 0-arity spinner and grab its compiled function object so Try
	// can invoke it directly (Try calls its argument with no args).
	evalString(&ctx, "(defun spin () (spin))")
	sym := mustSymbol(MakeSymbol("spin"))
	if sym.function == nil {
		t.Fatal("spin did not bind a function")
	}

	var res tryResult
	finished := runWithin(5*time.Second, func() {
		res = Try(&ctx, sym.function)
	})
	if !finished {
		t.Fatal("Try did not halt the spinner under a step limit")
	}
	if !IsError(res.data) {
		t.Fatalf("Try.data is %q, expected an error Obj", ObjString(res.data))
	}
	if msg := ObjString(res.data); !strings.Contains(msg, "step limit exceeded") {
		t.Fatalf("recovered error %q, want it to mention the step limit", msg)
	}
}

// TestZeroStepLimitIsUnlimited guards the production contract: the default
// budget of 0 must not interfere with an ordinary terminating program.
func TestZeroStepLimitIsUnlimited(t *testing.T) {
	if got := ObjString(evalStringWithLimit("(+ 40 2)", 0)); got != "42" {
		t.Fatalf("zero step limit changed a normal eval: got %q, want 42", got)
	}
}

// TestStepLimitAllowsHonestWork checks the budget is a ceiling on runaway
// loops, not a tax on legitimate computation: a real (terminating) recursive
// function well under the limit must still return its true result.
func TestStepLimitAllowsHonestWork(t *testing.T) {
	prog := `(do (defun countdown (N) (if (= N 0) done (countdown (- N 1))))
	            (countdown 1000))`
	if got := ObjString(evalStringWithLimit(prog, 1_000_000)); got != "done" {
		t.Fatalf("budgeted but terminating program: got %q, want done", got)
	}
}
