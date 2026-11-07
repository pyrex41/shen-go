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

// A rebinding after compilation reaches a thaw that had been inlined.
func TestLowerKLRecompilesOnThawRebinding(t *testing.T) {
	thaw := mustSymbol(symThaw)
	savedStd, savedFn := stdThaw, thaw.function
	defer func() { stdThaw, thaw.function = savedStd, savedFn; lowerAssumptions = map[string]lowerAssumption{} }()
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("X")}, lowerRead(t, `(X)`)))
	stdThaw = thaw.function
	obj := CompileFunc("lower-test", nil, lowerRead(t, `(let X 1 (let W (freeze X) (if (= X 1) (thaw W) (lower-f W))))`))
	var e ControlFlow
	if got := ObjString(Call(&e, obj)); got != "1" {
		t.Fatalf("before rebinding: %s", got)
	}
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("X")}, MakeSymbol("rebound")))
	if got := ObjString(Call(&e, obj)); got != "rebound" {
		t.Fatalf("after rebinding: %s", got)
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

// Redefining a function replaces its entry instead of keeping every old body.
func TestLowerKLAssumptionsDoNotGrow(t *testing.T) {
	thaw := mustSymbol(symThaw)
	savedStd, savedFn := stdThaw, thaw.function
	defer func() { stdThaw, thaw.function = savedStd, savedFn; lowerAssumptions = map[string]lowerAssumption{} }()
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("X")}, lowerRead(t, `(X)`)))
	stdThaw = thaw.function
	lowerAssumptions = map[string]lowerAssumption{}
	for i := 0; i < 5; i++ {
		CompileFunc("lower-same", nil, lowerRead(t, `(let X 1 (let W (freeze X) (if (= X 1) (thaw W) (lower-f W))))`))
	}
	if len(lowerAssumptions) != 1 {
		t.Fatalf("%d entries for one name", len(lowerAssumptions))
	}
}

// A freeze passed to an inlined lambda becomes a let-bound freeze, which the
// compiler inlines; a rebound thaw must still be called there, whether the
// definition was compiled before the rebinding or after (review F1).
func TestLowerKLLambdaFreezeHonoursThawRebinding(t *testing.T) {
	thaw := mustSymbol(symThaw)
	savedStd, savedFn := stdThaw, thaw.function
	defer func() {
		stdThaw, thaw.function = savedStd, savedFn
		vmNoteRebind(thaw)
		lowerAssumptions = map[string]lowerAssumption{}
	}()
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("X")}, lowerRead(t, `(X)`)))
	stdThaw = thaw.function
	var e ControlFlow
	evalAll(t, &e, `(set lower-tc 0)`)
	before := CompileFunc("lower-before", []Obj{MakeSymbol("X")}, lowerRead(t, `(let G (lambda Y (thaw Y)) (G (freeze X)))`))
	direct := CompileFunc("lower-direct", []Obj{MakeSymbol("X")}, lowerRead(t, `((lambda Y (thaw Y)) (freeze X))`))
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("F")}, lowerRead(t, `(do (set lower-tc (+ 1 (value lower-tc))) (F))`)))
	after := CompileFunc("lower-after", []Obj{MakeSymbol("X")}, lowerRead(t, `(let G (lambda Y (thaw Y)) (G (freeze X)))`))
	for _, f := range []Obj{before, direct, after} {
		if got := ObjString(Call(&e, f, MakeInteger(5))); got != "5" {
			t.Fatalf("result %s", got)
		}
	}
	if got := evalAll(t, &e, `(value lower-tc)`); got != "3" {
		t.Fatalf("the rebound thaw ran %s times, want 3", got)
	}
}

// Every shape in which lowering could leave out a thaw call, compiled before
// and after thaw is rebound to a counting wrapper (review F1v). Master calls
// the rebound thaw once in each, except g, where (thaw (freeze E)) is written
// literally and the compiler leaves it out on master too.
func TestLowerKLThawShapesHonourRebinding(t *testing.T) {
	thaw := mustSymbol(symThaw)
	savedStd, savedFn := stdThaw, thaw.function
	defer func() {
		stdThaw, thaw.function = savedStd, savedFn
		vmNoteRebind(thaw)
		lowerAssumptions = map[string]lowerAssumption{}
	}()
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("X")}, lowerRead(t, `(X)`)))
	stdThaw = thaw.function
	shapes := []struct{ name, body string }{
		{"a", `(let F (freeze (cons X X)) (let G (lambda Y (thaw Y)) (G F)))`},
		{"b", `(let G (lambda A (lambda B (thaw B))) ((G 1) (freeze (cons X X))))`},
		{"c", `(let G (lambda A (lambda B (cons (thaw A) B))) ((G (freeze (cons X X))) 1))`},
		{"d", `(let G (lambda Y (lambda Z (thaw Y))) (let H (G (freeze (cons X X))) (H 1)))`},
		{"e", `(let F (freeze (cons X X)) ((lambda Y (thaw Y)) F))`},
		{"f", `(let G (lambda Y (if (= X 5) (thaw Y) (thaw Y))) (G (freeze (cons X X))))`},
		{"g", `(let G (lambda Y (thaw (freeze (cons Y Y)))) (G X))`},
		{"h", `(let G (lambda Y (freeze (cons Y Y))) (thaw (G X)))`},
		{"i", `(let F (freeze (cons X X)) (let H F (thaw H)))`},
		{"j", `(let G (lambda Y (cond ((= X 0) 0) (true (thaw Y)))) (G (freeze (cons X X))))`},
		{"k", `(let F (freeze (cons X X)) (let G (lambda Y (thaw Y)) (if (= X 5) (G F) (G F))))`},
	}
	compile := func(prefix string) []Obj {
		var fs []Obj
		for _, s := range shapes {
			fs = append(fs, CompileFunc("lower-"+prefix+s.name, []Obj{MakeSymbol("X")}, lowerRead(t, s.body)))
		}
		return fs
	}
	pre := compile("pre-")
	var e ControlFlow
	evalAll(t, &e, `(set lower-tc 0)`)
	BindSymbolFunc(symThaw, CompileFunc("thaw", []Obj{MakeSymbol("F")}, lowerRead(t, `(do (set lower-tc (+ 1 (value lower-tc))) (F))`)))
	post := compile("post-")
	for i, s := range shapes {
		want := "1"
		if s.name == "g" {
			want = "0"
		}
		for when, f := range map[string]Obj{"before": pre[i], "after": post[i]} {
			evalAll(t, &e, `(set lower-tc 0)`)
			wantv := "(5 . 5)"
			if s.name == "c" {
				wantv = "((5 . 5) . 1)"
			}
			if got := ObjString(Call(&e, f, MakeInteger(5))); got != wantv {
				t.Errorf("%s compiled %s: result %s", s.name, when, got)
			}
			if got := evalAll(t, &e, `(value lower-tc)`); got != want {
				t.Errorf("%s compiled %s the rebinding: the rebound thaw ran %s times, want %s", s.name, when, got, want)
			}
		}
	}
}
