package kl

import (
	"fmt"
	"strings"
	"testing"
)

func lowerRead(t *testing.T, src string) Obj {
	t.Helper()
	form, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	return form
}

func lowerRun(form Obj) (out string) {
	defer func() {
		if r := recover(); r != nil {
			if o, ok := r.(Obj); ok && IsError(o) {
				out = "raised: " + ObjString(o)
				return
			}
			out = fmt.Sprint("panic: ", r)
		}
	}()
	var e ControlFlow
	return ObjString(Call(&e, CompileFunc("lower-test", nil, form)))
}

func closuresIn(fn *BytecodeFunc) int {
	n := 0
	for _, in := range fn.Code {
		if in.Op == OP_MAKE_CLOSURE {
			n++
		}
	}
	return n
}

// Every form means the same with and without the pass, including which error
// it raises and the order of its effects.
func TestLowerKLPreservesMeaning(t *testing.T) {
	saved := stdThaw
	defer func() { stdThaw = saved; lowerOff = false }()
	stdThaw = mustSymbol(symThaw).function
	for _, src := range []string{
		`(let W (lambda A (lambda B (+ A B))) ((W 1) 2))`,
		`(let A (freeze 1) (let B (freeze (if false (thaw A) (thaw A))) (if false (thaw B) (if true (thaw B) 3))))`,
		`(let W (lambda A (+ A 1)) (W 1))`,
		`(let W (lambda A (lambda B (+ A B))) (W 1 2))`,
		`(let W (lambda A (lambda B (+ A B))) (let F (W 1) (F 2)))`,
		`((lambda X (+ X 1)) 2)`,
		`((lambda X (lambda Y (+ X Y))) 2)`,
		`(let X 1 (let W (lambda A (+ A X)) (let X 2 (W X))))`,
		`(let A 5 (let W (lambda A (lambda B (+ A B))) ((W 1) A)))`,
		`(let W (lambda A (lambda A (+ A 1))) ((W 1) 2))`,
		`(let W (freeze foo) (let foo 1 (thaw W)))`,
		`(let W (lambda A foo) (let foo 1 (W 2)))`,
		`(let W (lambda A (cons A foo)) (W foo))`,
		`(let W (lambda A (let foo 1 (cons A foo))) (W foo))`,
		`(let W (lambda A (lambda foo (cons A foo))) ((W foo) 2))`,
		`(do (set lower-o 0) (let W (lambda A (lambda B (cons A B))) ((W (set lower-o (+ 1 (value lower-o)))) (set lower-o (* 10 (value lower-o))))))`,
		`(let W (lambda A (hd A)) (W 1))`,
		`(let W (lambda A (lambda B (hd A))) ((W 1) (tl 2)))`,
		`(let W (lambda A (lambda B (hd A))) ((W (tl 1)) (tl 2)))`,
		`(let F (lambda X (+ X 1)) (let W (lambda G (G 1)) (W F)))`,
		`(let W (lambda cons (cons 1 ())) (W 5))`,
		`(let hd (lambda X 1) (hd (cons 2 ())))`,
		`(let cons (freeze 1) (thaw cons))`,
		`(let W (freeze (simple-error "not run")) (if true 1 (thaw W)))`,
		`(let W (freeze (simple-error "run")) (if false 1 (thaw W)))`,
		`(let W (freeze 9) (if false (thaw W) (thaw (hd (cons W ())))))`,
		`(do (set lower-c 0) (let W (freeze (set lower-c (+ 1 (value lower-c)))) (do (thaw W) (thaw W))))`,
		`(do (set lower-c 0) (let W (lambda A (set lower-c (+ A (value lower-c)))) (do (W 1) (W 2))))`,
		`(let W (lambda A A) (lambda W (W 1)))`,
		`((let W (lambda A A) (lambda W (W 1))) (lambda X (+ X 1)))`,
		`(thaw (let X 5 (let W (lambda A (+ A X)) (freeze (W 1)))))`,
		`(((let X 5 (let W (lambda A (+ A X)) (lambda X (W X)))) 1))`,
		`((let X 5 (let W (lambda A (+ A X)) (lambda X (W X)))) 1)`,
		`(let W (lambda A (+ A 1)) (let W (W 1) (+ W 1)))`,
		`(let W (lambda A (W A)) 5)`,
		`(let W (lambda A (lambda B (+ A B))) (+ ((W 1) 2) ((W 3) 4)))`,
		`(let W (lambda A (+ A 1)) (W (W 1)))`,
		`(let W (lambda A (if A 1 2)) (W 3))`,
		`(let W (freeze 1) (trap-error (+ (thaw W) (hd 1)) (lambda E 7)))`,
		`(let W (lambda A (+ A 1)) (cond ((= 2 (W 1)) (W 2)) (true 0)))`,
		`(let W (lambda A A) (and (W true) (W false)))`,
		`(let W (lambda A (type A number)) (type (W 1) number))`,
		`(let Z 1 (let W (lambda A (let Z (+ A Z) Z)) (W Z)))`,
		`(let Y 1 (let Z 2 (let W (lambda A (let Y 7 (cons A Y))) (W Y))))`,
		`(let Y 1 (let W (lambda A (lambda Y (cons A Y))) ((W Y) 2)))`,
	} {
		form := lowerRead(t, src)
		before := ObjString(form)
		lowerOff = true
		want := lowerRun(form)
		lowerOff = false
		got := lowerRun(form)
		if strings.Contains(got, "#<") || strings.Contains(want, "#<") {
			// A closure prints by address; only its presence is comparable.
			if strings.Contains(got, "#<") != strings.Contains(want, "#<") {
				t.Errorf("%s: got %s, want %s", src, got, want)
			}
		} else if got != want {
			t.Errorf("%s: got %s, want %s", src, got, want)
		}
		if ObjString(form) != before {
			t.Errorf("%s: input mutated to %s", src, ObjString(form))
		}
	}
}

func TestLowerKLRemovesClauseClosures(t *testing.T) {
	saved := stdThaw
	defer func() { stdThaw = saved }()
	stdThaw = mustSymbol(symThaw).function
	for _, tt := range []struct {
		src  string
		want int
	}{
		{`(let W (lambda A (lambda B (+ A B))) ((W 1) 2))`, 0},
		{`(let X (cons 1 2) (let W (lambda A (lambda B (cons A B))) (if (cons? X) ((W (hd X)) (tl X)) (lower-f (freeze ((W 1) 2))))))`, 1},
		{`(let X 1 (let W (freeze (lower-g X)) (if (= X 1) (thaw W) (lower-f W))))`, 1},
		{`((lambda X (+ X 1)) 2)`, 0},
	} {
		obj := CompileFunc("lower-test", nil, lowerRead(t, tt.src))
		if got := closuresIn(mustBytecodeFunc(obj).fn); got != tt.want {
			t.Errorf("%s: %d closures, want %d", tt.src, got, tt.want)
		}
	}
}

// A rebinding of thaw before compilation leaves the call alone.
func TestLowerKLHonoursRebuiltThaw(t *testing.T) {
	saved := stdThaw
	defer func() { stdThaw = saved }()
	stdThaw = nil
	obj := CompileFunc("lower-test", nil, lowerRead(t, `(let X 1 (let W (freeze X) (if (= X 1) (thaw W) (lower-f W))))`))
	if got := closuresIn(mustBytecodeFunc(obj).fn); got != 1 {
		t.Errorf("%d closures, want the freeze kept", got)
	}
}

// Clause nests like the Prolog compiler's copy each body at two uses per
// level. The budget and the closure limit keep that from being exponential.
func TestLowerKLGrowthIsBounded(t *testing.T) {
	src := `(+ Z 1)`
	for i := 0; i < 30; i++ {
		w := fmt.Sprintf("W%d", i)
		a := fmt.Sprintf("A%d", i)
		src = fmt.Sprintf(`(let %s (lambda %s %s) (if (cons? X) (%s (hd X)) (lower-f (freeze (%s 1)))))`, w, a, src, w, w)
	}
	form := lowerRead(t, src)
	in := formSize(form)
	out := formSize(LowerKL(form))
	if out > in+lowerBudget {
		t.Fatalf("grew from %d to %d nodes", in, out)
	}
}

// A test built from and and or raises what the value form raises, from the same
// part, and runs the same parts.
func TestCompileTestMatchesValueForm(t *testing.T) {
	BindSymbolFunc(MakeSymbol("lower-id"), CompileFunc("lower-id", []Obj{MakeSymbol("X")}, MakeSymbol("X")))
	vals := []string{"true", "false", "5", "(hd 1)", "(do (set lower-n (+ 1 (value lower-n))) true)", "(do (set lower-n (+ 10 (value lower-n))) false)"}
	shapes := []string{
		"(if (and %s %s) 1 2)", "(if (or %s %s) 1 2)",
		"(cond ((and %s %s) 1) (true 2))", "(cond ((or %s %s) 1) (true 2))",
		"(+ 1 (if (and %s %s) 1 2))", "(+ 1 (if (or %s %s) 1 2))",
		"(if (and (or %s true) %s) 1 2)", "(if (or (and %s false) %s) 1 2)",
		"(if (and %s (and %[1]s %s)) 1 2)", "(if (or %s (or %[1]s %s)) 1 2)",
		"(if (and (= 1 1) (= %s %s)) 1 2)",
	}
	for _, shape := range shapes {
		for _, a := range vals {
			for _, b := range vals {
				src := fmt.Sprintf("(do (set lower-n 0) (cons "+shape+" (value lower-n)))", a, b)
				// The value form the compiler used before: the condition as an
				// if of its own, whose result the outer if tests.
				valueForm := strings.NewReplacer("(if (", "(if (lower-id (", "(cond ((", "(cond ((lower-id (", ") 1", ")) 1").Replace(src)
				got := lowerRun(lowerRead(t, src))
				want := lowerRun(lowerRead(t, valueForm))
				if got != want {
					t.Errorf("%s: got %s, want %s", src, got, want)
				}
			}
		}
	}
}
