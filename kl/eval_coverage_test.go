package kl

import "testing"

// TestTreeWalkerSpecialForms covers the AST-interpreter paths for and/or/cond/if.
// Top-level forms (not inside a defun) are evaluated by the tree-walker, so these
// exercise evalAnd/evalOr/evalCond/evalIf rather than the bytecode VM. The
// short-circuit cases put a would-error expression in the dead branch to prove
// it is never evaluated.
func TestTreeWalkerSpecialForms(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"and both true", "(and true true)", "true"},
		{"and first false short-circuits", "(and false (undefined-fn 1))", "false"},
		{"and second false", "(and true false)", "false"},

		{"or first true short-circuits", "(or true (undefined-fn 1))", "true"},
		{"or both false", "(or false false)", "false"},
		{"or second true", "(or false true)", "true"},

		{"cond picks matching clause", "(cond (false 1) (true 2))", "2"},
		{"cond first match wins", "(cond ((= 1 1) 10) ((= 2 2) 20))", "10"},
		{"cond no match returns nil", "(cond (false 1) (false 2))", "()"},
	}

	var ctx ControlFlow
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ObjString(evalString(&ctx, c.input))
			if got != c.want {
				t.Errorf("%s: got %q, want %q", c.input, got, c.want)
			}
		})
	}
}

// TestTreeWalkerIfNonBoolean covers the evalIf error branch: a non-boolean
// condition must raise a catchable error rather than be treated as truthy,
// matching the compiled VM path.
func TestTreeWalkerIfNonBoolean(t *testing.T) {
	var ctx ControlFlow
	got := ObjString(evalString(&ctx,
		`(trap-error (if 42 1 2) (lambda E (error-to-string E)))`))
	if got != `"if requires a boolean"` {
		t.Errorf("got %q, want %q", got, `"if requires a boolean"`)
	}
}

// TestCurriedLambdaOverApplication exercises the apply() over-application branch
// (len(args) > proc.arity): a single-param lambda returning a lambda, applied to
// two args at once.
func TestCurriedLambdaOverApplication(t *testing.T) {
	var ctx ControlFlow
	got := ObjString(evalString(&ctx, "((lambda x (lambda y (+ x y))) 1 2)"))
	if got != "3" {
		t.Errorf("got %q, want 3", got)
	}
}

// TestVMNumericEdges covers the VM numeric fast paths that the existing suite
// missed: float <= (numCmpLE) and partial application of a compiled function
// (vmApply with fewer args than arity), plus the compiler intrinsic dispatch for
// every comparison/arith op through a compiled defun.
func TestVMNumericEdges(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		// numCmpLE float path through a compiled defun
		{"compiled <= floats true", "(do (defun le (X Y) (<= X Y)) (le 1.5 2.5))", "true"},
		{"compiled <= floats false", "(do (defun le (X Y) (<= X Y)) (le 3.5 2.5))", "false"},
		{"compiled <= floats equal", "(do (defun le (X Y) (<= X Y)) (le 2.5 2.5))", "true"},

		// partial application of a compiled (bytecode) fn -> vmPartialApply via vmApply
		{"partial then apply", "(do (defun add (X Y) (+ X Y)) ((add 5) 10))", "15"},

		// intrinsic op dispatch through compiled defuns
		{"intrinsic +", "(do (defun f (X Y) (+ X Y)) (f 2 3))", "5"},
		{"intrinsic -", "(do (defun f (X Y) (- X Y)) (f 9 4))", "5"},
		{"intrinsic *", "(do (defun f (X Y) (* X Y)) (f 6 7))", "42"},
		{"intrinsic <", "(do (defun f (X Y) (< X Y)) (f 1 2))", "true"},
		{"intrinsic >", "(do (defun f (X Y) (> X Y)) (f 1 2))", "false"},
		{"intrinsic >=", "(do (defun f (X Y) (>= X Y)) (f 2 2))", "true"},
	}

	var ctx ControlFlow
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ObjString(evalString(&ctx, c.input))
			if got != c.want {
				t.Errorf("%s: got %q, want %q", c.input, got, c.want)
			}
		})
	}
}
