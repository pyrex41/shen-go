package kl

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

// withExec runs f with closure compilation on or off for new definitions.
func withExec(closure bool, f func()) {
	old, oldDirty := clEnabled, clDirty
	clEnabled = closure
	defer func() { clEnabled, clDirty = old, oldDirty }()
	f()
}

// clEvalProg evaluates each form in order and describes the last result, an
// error object printing as its message, so VM and closure runs can be compared.
func clEvalProg(forms []string) string {
	var ctx ControlFlow
	var last Obj
	for _, f := range forms {
		last = evalString(&ctx, f)
	}
	if IsError(last) {
		return "ERROR " + mustError(last).err
	}
	return ObjString(last)
}

// The executors must agree on every program here, including its error text.
func TestClosureMatchesVM(t *testing.T) {
	progs := map[string][]string{
		"fib":          {`(defun cl-fib (N) (if (< N 2) N (+ (cl-fib (- N 1)) (cl-fib (- N 2)))))`, `(cl-fib 20)`},
		"arg-order":    {`(defun cl-o (A B C D E F) (cons A (cons B (cons C (cons D (cons E (cons F ())))))))`, `(cl-o 1 2 3 4 5 6)`},
		"arg-swap":     {`(defun cl-sw (N A B) (if (= N 0) (cons A B) (cl-sw (- N 1) B A)))`, `(cl-sw 3 1 2)`},
		"self-5":       {`(defun cl-s5 (N A B C D) (if (= N 0) (cons A (cons B (cons C D))) (cl-s5 (- N 1) B C D A)))`, `(cl-s5 7 1 2 3 4)`},
		"self-6":       {`(defun cl-s6 (N A B C D E) (if (= N 0) (cons A (cons B (cons C (cons D E)))) (cl-s6 (- N 1) B C D E A)))`, `(cl-s6 7 1 2 3 4 5)`},
		"let-shadow":   {`(defun cl-l (X) (let X (+ X 1) (let Y (let X (* X 2) (+ X 1)) (cons X Y))))`, `(cl-l 3)`},
		"lambda-cap":   {`(defun cl-c (X) (let Y (+ X 1) (lambda Z (lambda W (cons X (cons Y (cons Z W)))))))`, `((((cl-c 1) 2)) 3)`},
		"upval-let":    {`(defun cl-u (X) (lambda Y (let Z (+ X Y) (let W (* Z 2) (cons Z (cons W X))))))`, `((cl-u 1) 2)`},
		"loop-closes":  {`(defun cl-lc (N A) (if (= N 0) A (let M (* N 10) (cl-lc (- N 1) (cons (freeze M) A)))))`, `(thaw (hd (tl (cl-lc 3 ()))))`},
		"partial":      {`(defun cl-p (A B C) (cons A (cons B C)))`, `(((cl-p 1) 2) 3)`},
		"partial2":     {`(defun cl-p2 (A B C) (cons A (cons B C)))`, `((cl-p2 1 2) 3)`},
		"over":         {`(defun cl-ov (A) (lambda B (cons A B)))`, `(cl-ov 1 2)`},
		"non-fn":       {`(defun cl-nf (F) (F 1))`, `(cl-nf 5)`},
		"sym-fn":       {`(defun cl-nf2 (F) (F 1))`, `(cl-nf2 hd)`},
		"unbound":      {`(defun cl-ub (X) (cl-no-such-fn X))`, `(cl-ub 1)`},
		"unbound-1st":  {`(set cl-seen 0)`, `(defun cl-ub2 () (cl-no-such-fn2 (set cl-seen 1)))`, `(trap-error (cl-ub2) (lambda E (value cl-seen)))`},
		"if-nonbool":   {`(defun cl-ib (X) (if X 1 2))`, `(cl-ib 5)`},
		"and-nonbool":  {`(defun cl-ab (X Y) (and X Y))`, `(cl-ab true 5)`},
		"or-nonbool":   {`(defun cl-ob (X Y) (or X Y))`, `(cl-ob false 5)`},
		"and-if":       {`(defun cl-ai (X Y) (if (and X Y) 1 2))`, `(cl-ai true 5)`},
		"or-if":        {`(defun cl-oi (X Y) (if (or X Y) 1 2))`, `(cl-oi 5 true)`},
		"not-nonbool":  {`(defun cl-nb (X) (not X))`, `(cl-nb 5)`},
		"not-if":       {`(defun cl-ni (X) (if (not X) 1 2))`, `(cl-ni 5)`},
		"cond-none":    {`(defun cl-cn (X) (cond ((= X 1) 2)))`, `(cl-cn 5)`},
		"hd-num":       {`(defun cl-h (X) (hd X))`, `(cl-h 5)`},
		"hd-nil":       {`(defun cl-hn (X) (tl X))`, `(cl-hn ())`},
		"add-str":      {`(defun cl-as (X) (+ X 1))`, `(cl-as "a")`},
		"lt-str":       {`(defun cl-ls (X) (if (< X 1) 1 2))`, `(cl-ls "a")`},
		"mul-str":      {`(defun cl-ms (X) (* X 2))`, `(cl-ms "a")`},
		"div-zero":     {`(defun cl-dz (X) (/ X 0))`, `(cl-dz 1)`},
		"div":          {`(defun cl-d (X Y) (/ X Y))`, `(cl-d 7 2)`},
		"big-add":      {fmt.Sprintf(`(defun cl-ba (X) (+ X %d))`, fixnumMax-1), fmt.Sprintf(`(cl-ba %d)`, fixnumMax-1)},
		"neg-edge":     {`(defun cl-ne (X) (- X 1))`, fmt.Sprintf(`(cl-ne %d)`, fixnumMin)},
		"top-edge":     {`(defun cl-te (X) (+ X 1))`, fmt.Sprintf(`(cl-te %d)`, fixnumMax-1)},
		"add-ll-edge":  {`(defun cl-ae (X Y) (+ X Y))`, fmt.Sprintf(`(cl-ae %d 1)`, fixnumMax-1)},
		"sub-ll-edge":  {`(defun cl-se (X Y) (- X Y))`, fmt.Sprintf(`(cl-se %d 1)`, fixnumMin)},
		"sub-ll":       {`(defun cl-sl (X Y) (- X Y))`, `(cl-sl 3 10)`},
		"float":        {`(defun cl-f (X) (* (+ X 0.5) 2))`, `(cl-f 1.25)`},
		"big-mul":      {`(defun cl-bm (X) (* X X))`, fmt.Sprintf(`(cl-bm %d)`, fixnumMax-1)},
		"huge":         {`(defun cl-hg (X) (* X X))`, `(cl-hg 1000000000000)`},
		"cmp-float":    {`(defun cl-cf (X) (cons (< X 1.5) (cons (<= X 1.5) (cons (> X 1.5) (>= X 1.5)))))`, `(cl-cf 1.5)`},
		"cmp-fix":      {`(defun cl-cx (X Y) (cons (< X Y) (cons (<= X Y) (cons (> X Y) (>= X Y)))))`, `(cl-cx -1 1)`},
		"cmp-big":      {`(defun cl-cb (X Y) (cons (< X Y) (cons (> X Y) (= X Y))))`, `(cl-cb 100000000 100000000)`},
		"eq-str":       {`(defun cl-es (X) (if (= X "a") 1 2))`, `(cl-es (cn "" "a"))`},
		"eq-big":       {`(defun cl-eb (X) (if (= X 100000000) 1 2))`, `(cl-eb (* 10000 10000))`},
		"eq-list":      {`(defun cl-el (X Y) (= X Y))`, `(cl-el (cons 1 ()) (cons 1 ()))`},
		"eq-nil":       {`(defun cl-en (X) (cons (= X ()) (= () X)))`, `(cl-en ())`},
		"eq-num-sym":   {`(defun cl-ens (X) (= X a))`, `(cl-ens 1)`},
		"vec":          {`(defun cl-v (N) (let V (absvector N) (let W (address-> V 0 7) (<-address W 0))))`, `(cl-v 2)`},
		"vec-empty":    {`(defun cl-ve (N) (<-address (absvector N) 0))`, `(cl-ve 2)`},
		"vec-oob":      {`(defun cl-vo (N) (<-address (absvector N) N))`, `(cl-vo 2)`},
		"vec-oob-set":  {`(defun cl-vs (N) (address-> (absvector N) N 1))`, `(cl-vs 2)`},
		"vec-frac":     {`(defun cl-vf (N) (<-address (absvector N) 0.5))`, `(cl-vf 2)`},
		"vec-notvec":   {`(defun cl-vn (N) (<-address N 0))`, `(cl-vn 2)`},
		"str-ops":      {`(defun cl-str (S) (cons (pos S 1) (cons (tlstr S) (cons (string->n S) (n->string 65)))))`, `(cl-str "héllo")`},
		"preds":        {`(defun cl-pr (X) (cons (number? X) (cons (string? X) (cons (symbol? X) (cons (cons? X) (cons (absvector? X) (boolean? X)))))))`, `(cl-pr a)`},
		"freeze":       {`(defun cl-fz (X) (let G (freeze (+ X 1)) (if (= X 0) (thaw G) (+ 1 (thaw G)))))`, `(cl-fz 1)`},
		"thaw-esc":     {`(defun cl-te2 (X) (let G (freeze (+ X 1)) (cons G (thaw G))))`, `(tl (cl-te2 1))`},
		"thaw-shadow":  {`(defun cl-ts (X) (let G (freeze X) (let X 9 (thaw G))))`, `(cl-ts 1)`},
		"trap":         {`(defun cl-t (X) (trap-error (hd X) (lambda E (error-to-string E))))`, `(cl-t 5)`},
		"trap-fn":      {`(defun cl-tf (X) (trap-error (/ X 0) (lambda E (cons X (error-to-string E)))))`, `(cl-tf 5)`},
		"trap-val":     {`(defun cl-tv () (trap-error (value cl-unset-var) (lambda E 7)))`, `(cl-tv)`},
		"trap-oob":     {`(defun cl-tb () (trap-error (<-vector (absvector 1) 0) (lambda E 8)))`, `(cl-tb)`},
		"nested-defun": {`(defun cl-nd (X) (defun cl-ndi (Y) (+ X Y)))`, `(cl-nd 5)`, `(cl-ndi 2)`},
		"fold":         {`(defun cl-fo () (cons (+ 1 2) (cons (- 5 7) (* 3 4))))`, `(cl-fo)`},
		"eq-fold":      {`(defun cl-qf () (cons (= 1 1) (= a b)))`, `(cl-qf)`},
		"intern":       {`(defun cl-in () (intern "abc"))`, `(cl-in)`},
		"do":           {`(set cl-z 0)`, `(defun cl-do () (do (set cl-z 5) (value cl-z)))`, `(cl-do)`},
		"local-head":   {`(defun cl-lh (cl-unbound-name X) (cl-unbound-name X))`, `(cl-lh (lambda Y (+ Y 1)) 1)`},
		"param-cons":   {`(defun cl-pc (cons) (cons cons cons))`, `(cl-pc 1)`},
		"expr-head":    {`(defun cl-eh (X) ((lambda Y (+ Y X)) 1))`, `(cl-eh 1)`},
		"freeze-0":     {`(defun cl-f0 () (thaw (freeze 3)))`, `(cl-f0)`},
		"call-0":       {`(defun cl-z0 () 4)`, `(defun cl-z1 () (+ 1 (cl-z0)))`, `(cl-z1)`},
		"tail-native":  {`(defun cl-tn (X) (hd-fn X))`, `(defun hd-fn (X) (intern X))`, `(cl-tn "q")`},
		"mutual":       {`(defun cl-ev (N) (if (= N 0) true (cl-od (- N 1))))`, `(defun cl-od (N) (if (= N 0) false (cl-ev (- N 1))))`, `(cl-ev 100001)`},
		"redefine":     {`(defun cl-r1 () 1)`, `(defun cl-r2 () (cl-r1))`, `(defun cl-r1 () 2)`, `(cl-r2)`},
		"tail-in-args": {`(defun cl-t1 (N) (if (= N 0) 7 (cl-t2 (- N 1))))`, `(defun cl-t2 (N) (cl-t1 N))`, `(cons (cl-t1 3) (cons (cl-t2 3) ()))`},
		"self-redef":   {`(defun cl-sr (N) (if (= N 0) 0 (do (defun cl-sr (N) 42) (cl-sr (- N 1)))))`, `(cl-sr 3)`},
	}
	for name, forms := range progs {
		t.Run(name, func(t *testing.T) {
			var vm, cl string
			withExec(false, func() { vm = clEvalProg(forms) })
			withExec(true, func() { cl = clEvalProg(forms) })
			if vm != cl {
				t.Fatalf("VM %q, closure %q", vm, cl)
			}
		})
	}
}

// A function past the largest frame array has a heap frame; lets in it must
// reach the collector.
func TestClosureHugeFrame(t *testing.T) {
	body := "(cons X0 (cons X1 (cons X69 (hd (cons (freeze X69) ())))))"
	for i := 69; i >= 0; i-- {
		prev := "N"
		if i > 0 {
			prev = fmt.Sprintf("X%d", i-1)
		}
		body = fmt.Sprintf("(let X%d (cons %s %s) %s)", i, prev, prev, body)
	}
	form := fmt.Sprintf(`(defun cl-huge (N) (if (= N 0) 0 %s))`, body)
	withExec(true, func() {
		defer debug.SetGCPercent(debug.SetGCPercent(1))
		var ctx ControlFlow
		evalString(&ctx, form)
		f := mustBytecodeFunc(mustSymbol(MakeSymbol("cl-huge")).function)
		if f.cl == nil || f.cl.bucket != 0 {
			t.Fatalf("cl-huge should have a heap frame, code %+v", f.cl)
		}
		for i := 0; i < 2000; i++ {
			if r := evalString(&ctx, `(hd (hd (cl-huge 1)))`); IsError(r) {
				t.Fatal(ObjString(r))
			}
		}
	})
}

// Rebinding an inlined primitive must reach every compiled site, including
// sites compiled before the rebinding, and restoring it must restore meaning.
func TestClosureRebindPrimitive(t *testing.T) {
	cases := []struct{ prim, def, call, want string }{
		{"+", `(defun cl-rb (X) (+ X 1))`, `(cl-rb 1)`, "100"},
		{"+", `(defun cl-rb (X Y) (+ X Y))`, `(cl-rb 1 1)`, "100"},
		{"-", `(defun cl-rb (X) (- X 1))`, `(cl-rb 1)`, "100"},
		{"-", `(defun cl-rb (X Y) (- X Y))`, `(cl-rb 1 1)`, "100"},
		{"*", `(defun cl-rb (X Y) (* X Y))`, `(cl-rb 1 1)`, "100"},
		{"/", `(defun cl-rb (X Y) (/ X Y))`, `(cl-rb 1 1)`, "100"},
		{"hd", `(defun cl-rb (X) (hd X))`, `(cl-rb (cons 1 2))`, "100"},
		{"tl", `(defun cl-rb (X) (tl X))`, `(cl-rb (cons 1 2))`, "100"},
		{"cons", `(defun cl-rb (X) (cons X X))`, `(cl-rb 1)`, "100"},
		{"not", `(defun cl-rb (X) (not X))`, `(cl-rb true)`, "100"},
		{"cons?", `(defun cl-rb (X) (cons? X))`, `(cl-rb 1)`, "100"},
		{"number?", `(defun cl-rb (X) (number? X))`, `(cl-rb 1)`, "100"},
		{"=", `(defun cl-rb (X Y) (= X Y))`, `(cl-rb 1 1)`, "100"},
		{"<", `(defun cl-rb (X Y) (< X Y))`, `(cl-rb 1 1)`, "100"},
		{"+", `(defun cl-rb () (+ 1 2))`, `(cl-rb)`, "100"},
		{"=", `(defun cl-rb (X) (if (= X 1) 100 (+ X 0)))`, `(cl-rb 2)`, "Error(if requires a boolean)"},
		{"<", `(defun cl-rb (X) (if (< X 1) 100 0))`, `(cl-rb 2)`, "Error(if requires a boolean)"},
	}
	for _, tc := range cases {
		t.Run(tc.prim+" "+tc.def, func(t *testing.T) {
			withExec(true, func() {
				var ctx ControlFlow
				sym := MakeSymbol(tc.prim)
				saved := mustSymbol(sym).function
				defer BindSymbolFunc(sym, saved)
				before := ObjString(evalString(&ctx, strings.Join([]string{`(do `, tc.def, ` `, tc.call, `)`}, "")))
				nargs := 1
				if strings.Contains("+ - * / = < cons", tc.prim) && tc.prim != "cons?" {
					nargs = 2
				}
				BindSymbolFunc(sym, MakeNative(func(c *ControlFlow) { c.Return(MakeInteger(100)) }, nargs))
				if !clDirty {
					t.Fatal("rebinding an inlined primitive left the guard clean")
				}
				if got := ObjString(evalString(&ctx, tc.call)); got != tc.want {
					t.Fatalf("after rebinding got %s, want %s", got, tc.want)
				}
				BindSymbolFunc(sym, saved)
				if got := ObjString(evalString(&ctx, tc.call)); got != before {
					t.Fatalf("after restoring got %s, want %s", got, before)
				}
			})
		})
	}
}

// A rebinding of one primitive leaves the others meaning what they meant, on
// the slow path the guard falls to.
func TestClosureDirtyOtherPrimitives(t *testing.T) {
	withExec(true, func() {
		var ctx ControlFlow
		evalString(&ctx, `(defun cl-do2 (X Y) (cons (+ X Y) (cons (- X Y) (cons (< X Y) (cons (hd (cons X Y)) (if (= X Y) 1 (if (cons? X) 2 3)))))))`)
		want := ObjString(evalString(&ctx, `(cl-do2 3 4)`))
		clDirty = true
		if got := ObjString(evalString(&ctx, `(cl-do2 3 4)`)); got != want {
			t.Fatalf("dirty guard changed the result: %s, want %s", got, want)
		}
	})
}

func TestClosureUnboundedTailCalls(t *testing.T) {
	const n = "3000000"
	progs := map[string][]string{
		"self":    {`(defun cl-ts1 (N) (if (= N 0) ok (cl-ts1 (- N 1))))`, `(cl-ts1 ` + n + `)`},
		"mutual":  {`(defun cl-m1 (N) (if (= N 0) ok (cl-m2 (- N 1))))`, `(defun cl-m2 (N) (if (= N 0) ok (cl-m1 (- N 1))))`, `(cl-m1 ` + n + `)`},
		"dynamic": {`(defun cl-dy (F N) (if (= N 0) ok (F F (- N 1))))`, `(defun cl-dy-go (N) (cl-dy (cl-dy-get) N))`, `(defun cl-dy-get () (lambda F (lambda N (cl-dy-step F N))))`, `(defun cl-dy-step (F N) (if (= N 0) ok ((F F) (- N 1))))`, `(cl-dy-go ` + n + `)`},
		"let":     {`(defun cl-tl (N) (let M (- N 1) (if (< M 0) ok (cl-tl M))))`, `(cl-tl ` + n + `)`},
		"native":  {`(defun cl-tn1 (N) (if (= N 0) ok (cl-tn2 N)))`, `(defun cl-tn2 (N) (cl-tn1 (- N 1)))`, `(cl-tn1 ` + n + `)`},
	}
	for name, forms := range progs {
		t.Run(name, func(t *testing.T) {
			withExec(true, func() {
				defer debug.SetMaxStack(debug.SetMaxStack(32 << 20))
				if got := clEvalProg(forms); got != "ok" {
					t.Fatalf("got %s", got)
				}
			})
		})
	}
}

// Deep non-tail recursion uses the Go stack: it has to survive stack growth and
// collections while frames hold the only reference to live data.
func TestClosureDeepRecursionUnderGC(t *testing.T) {
	withExec(true, func() {
		defer debug.SetGCPercent(debug.SetGCPercent(5))
		var ctx ControlFlow
		evalString(&ctx, `(defun cl-mk (N) (if (= N 0) () (let C (cons N N) (cons (hd C) (cl-mk (- N 1))))))`)
		evalString(&ctx, `(defun cl-len (L A) (if (cons? L) (cl-len (tl L) (+ A (hd L))) A))`)
		for i := 0; i < 3; i++ {
			got := ObjString(evalString(&ctx, `(cl-len (cl-mk 200000) 0)`))
			if got != "20000100000" {
				t.Fatalf("got %s", got)
			}
		}
	})
}

func TestClosureStepLimit(t *testing.T) {
	withExec(true, func() {
		progs := []string{
			`(do (defun cl-sl1 (N) (cl-sl1 N)) (cl-sl1 0))`,
			`(do (defun cl-sl2 (N) (cl-sl3 N)) (do (defun cl-sl3 (N) (cl-sl2 N)) (cl-sl2 0)))`,
			`(do (defun cl-sl4 (F) ((F F))) (cl-sl4 (lambda F (freeze (cl-sl4 F)))))`,
		}
		for _, p := range progs {
			ctx := ControlFlow{stepLimit: 100000}
			r := evalString(&ctx, p)
			if !IsError(r) || mustError(r).err != "eval step limit exceeded" {
				t.Fatalf("%s: got %s", p, ObjString(r))
			}
		}
	})
}

// VM code and closure code call each other, in every direction.
func TestClosureInteropWithVM(t *testing.T) {
	var ctx ControlFlow
	withExec(false, func() {
		evalString(&ctx, `(defun cl-vm1 (N) (if (= N 0) vm (cl-c1 (- N 1))))`)
		evalString(&ctx, `(defun cl-vm2 (N) (+ 1 (cl-c2 N)))`)
		evalString(&ctx, `(defun cl-vm3 (N A B C D E) (cons N (cons A (cons B (cons C (cons D E))))))`)
	})
	withExec(true, func() {
		evalString(&ctx, `(defun cl-c1 (N) (if (= N 0) cl (cl-vm1 (- N 1))))`)
		evalString(&ctx, `(defun cl-c2 (N) (if (= N 0) 0 (+ 1 (cl-vm2 (- N 1)))))`)
		evalString(&ctx, `(defun cl-c3 (N) (cl-vm3 N 2 3 4 5 6))`)
		if got := ObjString(evalString(&ctx, `(cl-vm1 1000001)`)); got != "cl" {
			t.Fatalf("mutual tail: %s", got)
		}
		if got := ObjString(evalString(&ctx, `(cl-vm2 1000)`)); got != "2001" {
			t.Fatalf("nested: %s", got)
		}
		if got := ObjString(evalString(&ctx, `(cl-c3 1)`)); !strings.HasPrefix(got, "(1 2") {
			t.Fatalf("arity 6: %s", got)
		}
	})
}

// BenchmarkDefine measures what one define costs each executor, and what it
// keeps alive, on a function shaped like compiled Shen pattern matching.
func benchDefine(b *testing.B, closure bool) {
	body := "(shen.f-error f)"
	for i := 0; i < 8; i++ {
		body = fmt.Sprintf("(if (cons? V0) (let H (hd V0) (let T (tl V0) (if (= H %d) (cons H (f T (+ V1 %d))) %s))) V1)", i, i, body)
	}
	src := "(defun f (V0 V1) " + body + ")"
	form, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		b.Fatal(err)
	}
	rest := mustPair(form).cdr
	params, fbody := ListToSlice(cadr(rest)), caddr(rest)
	withExec(closure, func() {
		var before, after runtime.MemStats
		keep := make([]Obj, 200)
		runtime.GC()
		runtime.ReadMemStats(&before)
		for i := range keep {
			keep[i] = CompileFunc("f", params, fbody)
		}
		runtime.GC()
		runtime.ReadMemStats(&after)
		retained := float64(after.HeapAlloc-before.HeapAlloc) / float64(len(keep))
		runtime.KeepAlive(keep)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CompileFunc("f", params, fbody)
		}
		b.ReportMetric(retained, "retained-bytes/fn")
	})
}

func BenchmarkDefineVM(b *testing.B)      { benchDefine(b, false) }
func BenchmarkDefineClosure(b *testing.B) { benchDefine(b, true) }
