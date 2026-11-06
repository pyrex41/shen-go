package kl

import (
	"runtime"
	"strconv"
	"testing"
	"weak"
)

func evalAll(t *testing.T, ctx *ControlFlow, srcs ...string) string {
	t.Helper()
	res := ""
	for _, src := range srcs {
		res = ObjString(evalString(ctx, src))
	}
	return res
}

// A self-tail call assigns its parameters in parallel, however the lowerer
// orders the moves and retargets.
func TestRegisterSelfTailParallelAssign(t *testing.T) {
	var ctx ControlFlow
	for _, tc := range []struct{ def, call, want string }{
		{`(defun sw (A B N) (if (= N 0) (cons A B) (sw B A (- N 1))))`, `(sw 1 2 3)`, `(2 . 1)`},
		{`(defun rot (A B C N) (if (= N 0) (cons A (cons B C)) (rot B C A (- N 1))))`, `(rot 1 2 3 4)`, `(2 3 . 1)`},
		{`(defun ab (A B N) (if (= N 0) (cons A B) (ab (+ A B) (- A B) (- N 1))))`, `(ab 5 3 2)`, `(10 . 6)`},
		{`(defun bu (A B N) (if (= N 0) (cons A B) (bu (+ A 1) (+ A B) (- N 1))))`, `(bu 1 0 3)`, `(4 . 6)`},
		{`(defun ke (A B N) (if (= N 0) (cons A B) (ke A 7 (- N 1))))`, `(ke 1 2 1)`, `(1 . 7)`},
		{`(defun le (A N) (if (= N 0) A (let X (+ A 1) (le (* X X) (- N 1)))))`, `(le 1 2)`, `25`},
		{`(defun co (A B N) (if (= N 0) (cons A B) (co (if (> A 0) (- A 1) A) (if (> A 0) (cons A B) B) (- N 1))))`, `(co 2 () 3)`, `(0 1 2)`},
	} {
		if got := evalAll(t, &ctx, tc.def, tc.call); got != tc.want {
			t.Errorf("%s = %s, want %s", tc.call, got, tc.want)
		}
	}
}

// Rebinding a specialized primitive from inside a running activation takes
// effect at the next operation of that activation, and restoring it restores
// the primitive's behaviour.
func TestRegisterRebindMidActivation(t *testing.T) {
	var ctx ControlFlow
	orig := map[string]Obj{}
	for _, n := range []string{"+", "hd", "cons?", "="} {
		orig[n] = PrimFunc(MakeSymbol(n))
	}
	defer func() {
		for n, f := range orig {
			BindSymbolFunc(MakeSymbol(n), f)
		}
		if !vmGuardsOK {
			t.Error("vmGuardsOK did not recover after restoring the primitives")
		}
	}()
	evalAll(&testing.T{}, &ctx,
		`(defun rb+ (X) (+ X 1))`,
		`(defun rbhd (X) (hd X))`,
		`(defun rbcons (X) (if (cons? X) yes no))`,
		`(defun rbeq (X) (if (= X 1) yes no))`,
	)
	if got := evalAll(t, &ctx, `(rb+ 1)`); got != "2" {
		t.Fatalf("(rb+ 1) = %s", got)
	}
	// The caller's add runs after the callee rebinds +.
	evalAll(t, &ctx, `(defun inner () (defun + (A B) plus))`, `(defun outer (X) (do (inner) (+ X 1)))`)
	if got := evalAll(t, &ctx, `(outer 1)`); got != "plus" {
		t.Fatalf("after a callee rebinds +: %s, want plus", got)
	}
	if got := evalAll(t, &ctx, `(rb+ 1)`); got != "plus" {
		t.Fatalf("(rb+ 1) with + rebound = %s", got)
	}
	// While one is rebound, the others keep their exact behaviour and errors.
	for src, want := range map[string]string{
		`(rbhd (cons 7 8))`: "7", `(rbcons (cons 7 8))`: "yes", `(rbcons 7)`: "no", `(rbeq 1)`: "yes", `(rbeq 2)`: "no",
		`(trap-error (rbhd 5) (lambda E (error-to-string E)))`: `"mustPair"`,
	} {
		if got := evalAll(t, &ctx, src); got != want {
			t.Errorf("%s = %s, want %s", src, got, want)
		}
	}
	evalAll(t, &ctx, `(defun hd (X) head)`, `(defun cons? (X) maybe)`)
	if got := evalAll(t, &ctx, `(rbhd (cons 7 8))`); got != "head" {
		t.Errorf("rebound hd: %s", got)
	}
	if got := evalAll(t, &ctx, `(trap-error (rbcons 1) (lambda E (error-to-string E)))`); got != `"if requires a boolean"` {
		t.Errorf("rebound cons? returning a non-boolean: %s", got)
	}
	// = with the literal on the left keeps the source's argument order.
	evalAll(t, &ctx, `(defun = (A B) (cons? A))`)
	BindSymbolFunc(MakeSymbol("cons?"), orig["cons?"])
	evalAll(t, &ctx, `(defun eql (X) (if (= (cons 1 2) X) yes no))`, `(defun eqk (X) (if (= 0 X) yes no))`)
	if got := evalAll(t, &ctx, `(eqk (cons 1 2))`); got != "no" {
		t.Errorf("(= 0 X) rebound, operands swapped: %s", got)
	}
}

// Tail calls between bytecode functions and through natives do not grow the
// arena or the Go stack; non-tail recursion is limited by neither at this depth.
func TestRegisterCallDepth(t *testing.T) {
	var ctx ControlFlow
	const big = "3000000"
	for _, tc := range []struct{ def, call, want string }{
		{`(defun even-p (N) (if (= N 0) true (odd-p (- N 1))))`, ``, ``},
		{`(defun odd-p (N) (if (= N 0) false (even-p (- N 1))))`, `(even-p ` + big + `)`, `true`},
		{`(defun via-apply (F N) (if (= N 0) done (F F (- N 1))))`, `(via-apply (lambda F (lambda N (via-apply F N))) ` + big + `)`, `done`},
		{`(defun count-up (N) (if (= N 0) 0 (+ 1 (count-up (- N 1)))))`, `(count-up ` + big + `)`, big},
		{`(defun clos (N) ((lambda X (if (= X 0) end (clos (- X 1)))) N))`, `(clos 1000000)`, `end`},
		{`(defun tailthaw (N) (if (= N 0) thawed (thaw (freeze (tailthaw (- N 1))))))`, `(tailthaw ` + big + `)`, `thawed`},
	} {
		if tc.call == "" {
			evalAll(t, &ctx, tc.def)
			continue
		}
		if got := evalAll(t, &ctx, tc.def, tc.call); got != tc.want {
			t.Errorf("%s = %s, want %s", tc.call, got, tc.want)
		}
		if ctx.frameCur != 0 || ctx.frameTop != 0 || ctx.vmDepth != 0 {
			t.Errorf("after %s: arena not unwound: cur=%d top=%d depth=%d", tc.call, ctx.frameCur, ctx.frameTop, ctx.vmDepth)
		}
	}
}

// An error thrown from deep in a bytecode recursion, across block boundaries and
// a nested vmRun, unwinds to the trap-error, and execution resumes correctly in
// the frame that caught it.
func TestRegisterUnwindAcrossBlocks(t *testing.T) {
	var ctx ControlFlow
	got := evalAll(t, &ctx,
		`(defun boom (N) (if (= N 0) (simple-error "boom") (+ 1 (boom (- N 1)))))`,
		`(defun viamap (N) (hd (map (lambda X (boom X)) (cons N ()))))`,
		`(defun catcher (N A) (+ A (trap-error (boom N) (lambda E 1))))`,
		`(defun again (K) (if (= K 0) 0 (+ (catcher 5000 K) (again (- K 1)))))`,
		`(again 20)`,
	)
	if got != "230" {
		t.Fatalf("(again 20) = %s, want 230", got)
	}
	if ctx.frameCur != 0 || ctx.frameTop != 0 || ctx.vmDepth != 0 {
		t.Fatalf("arena not unwound: cur=%d top=%d depth=%d", ctx.frameCur, ctx.frameTop, ctx.vmDepth)
	}
}

// Frames are not cleared when they return, so a value only a dead frame held
// must still become unreachable: the arena's stale slots go once per GC cycle.
func TestRegisterStaleSlotsDoNotRetain(t *testing.T) {
	var ctx ControlFlow
	var w weak.Pointer[scmHead]
	mk := MakeSymbol("retention-probe")
	BindSymbolFunc(mk, MakeNative(func(e *ControlFlow) {
		o := cons(MakeSymbol("x"), Nil)
		w = weak.Make((*scmHead)(o))
		e.Return(o)
	}, 0))
	defer BindSymbolFunc(mk, nil)
	evalAll(t, &ctx,
		`(defun holds (X) (do (no-op) (+ 1 (if (cons? X) 1 2))))`,
		`(defun no-op () 0)`,
		`(defun probe () (do (holds (retention-probe)) 0))`,
		`(probe)`,
	)
	for i := 0; i < 50 && w.Value() != nil; i++ {
		runtime.GC()
		runtime.Gosched()
		evalAll(t, &ctx, `(no-op)`)
	}
	if w.Value() != nil {
		t.Fatal("a value in a returned frame stayed reachable across GC cycles")
	}
}

// Step limits still bound every kind of unbounded execution.
func TestRegisterStepLimitBoundsTailCalls(t *testing.T) {
	for _, def := range []string{
		`(defun spin (N) (spin2 (+ N 1)))|(defun spin2 (N) (spin N))`,
		`(defun spin (N) (spin (+ N 1)))|(defun unused () 0)`,
	} {
		var ctx ControlFlow
		for i := 0; i < len(def); i++ {
			if def[i] == '|' {
				evalAll(t, &ctx, def[:i], def[i+1:])
			}
		}
		ctx.SetStepLimit(100000)
		if got := ObjString(evalString(&ctx, `(spin 0)`)); got != "Error(eval step limit exceeded)" {
			t.Errorf("%s: %s", def, got)
		}
	}
}

// Many arguments, a closure's upvalues and numbers across the bytecode boundary.
func TestRegisterValueForms(t *testing.T) {
	var ctx ControlFlow
	for src, want := range map[string]string{
		`(* 1000000 1000000)`:                                                       strconv.Itoa(1000000 * 1000000),
		`(let X 33554431 (+ X 1))`:                                                  "33554432",
		`(let X 33554432 (= X 33554432))`:                                           "true",
		`(let X (/ 1 2) (cons X X))`:                                                "(0.5 . 0.5)",
		`(let X 0.5 (integer? (+ X X)))`:                                            "true",
		`(let X 2 (= (* X 0.5) 1))`:                                                 "true",
		`(let F (lambda X (+ X 0.25)) (F (F 1)))`:                                   "1.5",
		`(let V (absvector 2) (address-> V 1 (+ 1 1)))`:                             "",
		`(let V (absvector 2) (<-address (address-> V 1 2.5) 1))`:                   "2.5",
		`(trap-error (<-address (absvector 1) 3) (lambda E (error-to-string E)))`:   `"index 3 out of range 1"`,
		`(trap-error (<-address (absvector 1) 0.5) (lambda E (error-to-string E)))`: "",
		`(trap-error (not 1) (lambda E (error-to-string E)))`:                       `"not: expected boolean"`,
		`(trap-error (if 1 2 3) (lambda E (error-to-string E)))`:                    `"if requires a boolean"`,
		`(trap-error (if (not 1) 2 3) (lambda E (error-to-string E)))`:              `"not: expected boolean"`,
		`(trap-error (unbound-fn 1) (lambda E (error-to-string E)))`:                `"function unbound-fn not bound"`,
		`(= "ab" (cn "a" "b"))`:                                                     "true",
		`(= (cons 1 2) (cons 1 2))`:                                                 "true",
		`(= 1 a)`:                                                                   "false",
		`(< 1 2)`:                                                                   "true",
		`(>= 1 2)`:                                                                  "false",
	} {
		if want == "" {
			continue
		}
		// Inside a defun, so it runs in the VM rather than the interpreter.
		if got := evalAll(t, &ctx, `(defun value-form () `+src+`)`, `(value-form)`); got != want {
			t.Errorf("%s = %s, want %s", src, got, want)
		}
	}
	// Compared with the literal forms inside a function, on boxed values too.
	if got := evalAll(t, &ctx, `(defun lit (X) (if (= X 33554432) big (if (= 0.5 X) half (if (= a X) sym (if (= "s" X) str other)))))`,
		`(cons (lit (+ 33554431 1)) (cons (lit (/ 1 2)) (cons (lit a) (cons (lit "s") (cons (lit ()) ())))))`); got != "(big half sym str other)" {
		t.Errorf("lit: %s", got)
	}
}
