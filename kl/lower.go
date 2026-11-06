package kl

import (
	"os"
	"strings"
)

// LowerKL rewrites a KL expression into an equivalent one that allocates fewer
// closures, before any executor sees it. It is a pure source-to-source step:
// the input tree is never mutated (the kernel's ps output may share it), and
// what it returns is only ever compiled, never printed.
//
// The shape it targets is what the kernel's Prolog compiler emits for a clause
// head, one let-bound function per head pattern and one let-bound continuation
// per clause:
//
//	(let W (lambda A (lambda B Body))
//	  (if (cons? X)
//	      (let H (hd X) (let T (tl X) ((W H) T)))     ; read mode
//	      (if (shen.pvar? X)
//	          (shen.bind! X ... (freeze ((W P) Q)))   ; write mode
//	          false)))
//
// Constructing a closure has no effect but the allocation, so a let-bound
// lambda or freeze can move to where it is used. The pass walks W's let down
// through exclusive branches and sequencing:
//
//   - A saturated use, ((W a) b) or (thaw W), becomes (let A a (let B b Body)),
//     or Body. a is evaluated, then b, then Body, the order the call had. Inside
//     a nested closure this happens only for small bodies (lowerClosureCopy).
//   - Any other use gets the closure just there, on the path that needs it.
//
// So read mode above allocates nothing and write mode only its freeze.
//
// Name capture: a copy of Body runs under the binders crossed on the way to
// its use, so the walk stops at a binder of any symbol occurring in the
// closure, and a use is not inlined when a later argument mentions an earlier
// parameter. There is no renaming. Symbols with a global function are left
// alone, because compileCall prefers the global over a same-named local.
//
// Growth: each copy is charged against a budget per definition. A use outside
// the budget keeps its closure, and closures placed beside an inlined copy
// hold the input form rather than the optimised one, so nesting grows
// quadratically and not exponentially.
//
// Inlining means fewer activations, so the step limit (tick) is charged less
// often. A rebinding of thaw before compilation turns thaw inlining off, and one
// after it recompiles what assumed it (noteThawAssumption).
func LowerKL(body Obj) Obj {
	out, _ := lowerDefine(body)
	return out
}

// lowerDefine is LowerKL, also saying whether the result inlined a thaw of a
// freeze that has other uses. compileLet treats a freeze with only thaws as
// syntax, but the program calls thaw on one that also escapes, so that call has
// to keep seeing a rebinding (see noteThawAssumption).
func lowerDefine(body Obj) (Obj, bool) {
	if lowerOff || !worthLowering(body) {
		return body, false
	}
	l := &lowerer{budget: lowerBudget}
	return l.lower(body), l.assumesThaw
}

// lowerBudget is the number of form nodes one definition may gain by copying.
const lowerBudget = 1500

// lowerClosureCopy is the largest body copied into a nested closure. A use
// outside any closure runs instead of a call and is worth more.
const lowerClosureCopy = 150

var stdThaw Obj // the native bound to thaw by installKernelFast

func thawInlinable() bool {
	return stdThaw != nil && mustSymbol(symThaw).function == stdThaw
}

var symThaw Obj

// lowerOff is the off-switch, SHEN_GO_LOWER=off, for comparing with and without.
var lowerOff = strings.EqualFold(strings.TrimSpace(os.Getenv("SHEN_GO_LOWER")), "off")

// worthLowering is the cheap test that keeps ordinary definitions out of the
// pass: a let of a closure, or a literal lambda in head position.
func worthLowering(f Obj) bool {
	ok, p := isPair(f)
	if !ok {
		return false
	}
	if p.car == symLet {
		if rest := listSlice(p.cdr); len(rest) == 3 {
			if _, _, _, ok := closureChain(rest[1]); ok {
				return true
			}
		}
	}
	if hp, isp := isPairP(p.car); isp && hp.car == symLambda {
		return true
	}
	for l := Obj(f); ; {
		lok, lp := isPair(l)
		if !lok {
			return false
		}
		if worthLowering(lp.car) {
			return true
		}
		l = lp.cdr
	}
}

func isPairP(o Obj) (*scmPair, bool) {
	ok, p := isPair(o)
	return p, ok
}

type lowerer struct {
	budget      int
	assumesThaw bool
}

// listSlice returns the elements of a proper list, or nil for anything else.
func listSlice(o Obj) []Obj {
	var out []Obj
	for {
		ok, p := isPair(o)
		if !ok {
			if o != Nil {
				return nil
			}
			return out
		}
		out = append(out, p.car)
		o = p.cdr
	}
}

func sliceList(xs []Obj) Obj {
	r := Nil
	for i := len(xs) - 1; i >= 0; i-- {
		r = cons(xs[i], r)
	}
	return r
}

func formSize(f Obj) int {
	n := 1
	for {
		ok, p := isPair(f)
		if !ok {
			return n
		}
		n += formSize(p.car)
		f = p.cdr
	}
}

func isLowerSpecial(sym Obj) bool {
	switch sym {
	case symDefun, symLambda, symFreeze, symLet, symIf, symAnd, symOr, symCond, symDo, symType, symTrapError:
		return true
	}
	return false
}

// closureChain recognises a closure value: (freeze B) has depth 0 and thaw as
// its saturated use, and (lambda A (lambda B Body)) has depth 2.
func closureChain(v Obj) (params []Obj, body Obj, isFreeze bool, ok bool) {
	vp, isp := isPairP(v)
	if !isp {
		return nil, nil, false, false
	}
	if vp.car == symFreeze {
		rest := listSlice(vp.cdr)
		if len(rest) != 1 {
			return nil, nil, false, false
		}
		return nil, rest[0], true, true
	}
	if vp.car != symLambda {
		return nil, nil, false, false
	}
	for {
		lp, isl := isPairP(v)
		if !isl || lp.car != symLambda {
			break
		}
		rest := listSlice(lp.cdr)
		if len(rest) != 2 || !IsSymbol(rest[0]) {
			break
		}
		params = append(params, rest[0])
		v = rest[1]
	}
	if len(params) == 0 {
		return nil, nil, false, false
	}
	return params, v, false, true
}

func (l *lowerer) lower(f Obj) Obj {
	p, isp := isPairP(f)
	if !isp {
		return f
	}
	args := listSlice(p.cdr)
	if p.cdr != Nil && args == nil {
		return f // improper: not code this pass understands
	}
	mapArgs := func() Obj {
		out := make([]Obj, len(args))
		for i, a := range args {
			out[i] = l.lower(a)
		}
		return cons(p.car, sliceList(out))
	}
	switch p.car {
	case symDefun:
		return f
	case symLet:
		if len(args) != 3 {
			return mapArgs()
		}
		v := l.lower(args[1])
		b := l.lower(args[2])
		return l.lowerLet(args[0], args[1], v, b)
	case symLambda:
		if len(args) != 2 {
			return f
		}
		return cons(p.car, cons(args[0], cons(l.lower(args[1]), Nil)))
	case symType:
		if len(args) != 2 {
			return f
		}
		return cons(p.car, cons(l.lower(args[0]), cons(args[1], Nil)))
	case symCond:
		out := make([]Obj, len(args))
		for i, clause := range args {
			c := listSlice(clause)
			if len(c) != 2 {
				return f
			}
			out[i] = cons(l.lower(c[0]), cons(l.lower(c[1]), Nil))
		}
		return cons(p.car, sliceList(out))
	}
	if hp, ish := isPairP(p.car); ish {
		// ((lambda X Body) A) is (let X A Body).
		if hp.car == symLambda && len(args) == 1 {
			if la := listSlice(hp.cdr); len(la) == 2 && IsSymbol(la[0]) && !isLowerSpecial(la[0]) {
				return l.bindParam(la[0], l.lower(args[0]), l.lower(la[1]))
			}
		}
		return cons(l.lower(p.car), cdr(mapArgs()))
	}
	return mapArgs()
}

// lowerLet decides what happens to (let x orig body), where v is orig lowered
// and body is already lowered.
func (l *lowerer) lowerLet(x, orig, v, body Obj) Obj {
	keep := func() Obj { return cons(symLet, cons(x, cons(v, cons(body, Nil)))) }
	params, chainBody, isFreeze, ok := closureChain(v)
	if !ok || !IsSymbol(x) || isLowerSpecial(x) || mustSymbol(x).function != nil {
		return keep()
	}
	if isFreeze && !thawInlinable() {
		// Moving a freeze to its thaw would hide it from compileLet, which
		// inlines a let-bound freeze whatever thaw is bound to, as on master.
		return keep()
	}
	if isFreeze {
		// Thawed and nothing else: compileLet makes that a jump where it can,
		// which a copy per use is not.
		if thaws, others := walkThawUses(body, x); others == 0 && thaws > 1 {
			return keep()
		}
	}
	syms := map[Obj]bool{}
	collectSyms(orig, syms)
	collectSyms(v, syms)
	if syms[x] {
		return keep()
	}
	s := &sinker{l: l, w: x, value: v, params: params, body: chainBody, isFreeze: isFreeze, syms: syms, thawOK: thawInlinable()}
	before := l.budget
	out := s.sink(body)
	switch {
	case s.placed == 0:
		return out
	case s.inlined == 0:
		// Nothing inlined: moving the closure is free on one path, and on two
		// it is a copy that saves nothing but the allocation on a third.
		if s.placed == 1 {
			return out
		}
		l.budget = before
		return keep()
	}
	// A closure beside an inlined copy holds the input form: optimising both
	// copies at every level of a nest is what makes it exponential. Those
	// copies come out of the budget too, or a nest would grow without limit
	// once the budget had gone.
	cost := s.placed * formSize(orig)
	if cost > l.budget {
		l.budget = before
		return keep()
	}
	// What the first run spent leaves room for the cost, so with the cost taken
	// first the same uses still fit and this run makes the same decisions.
	l.budget = before - cost
	s.value, s.inlined, s.placed = orig, 0, 0
	return s.sink(body)
}

func collectSyms(f Obj, into map[Obj]bool) {
	for {
		ok, p := isPair(f)
		if !ok {
			if IsSymbol(f) {
				into[f] = true
			}
			return
		}
		collectSyms(p.car, into)
		f = p.cdr
	}
}

func occursIn(sym, f Obj) bool {
	for {
		ok, p := isPair(f)
		if !ok {
			return f == sym
		}
		if occursIn(sym, p.car) {
			return true
		}
		f = p.cdr
	}
}

type sinker struct {
	l        *lowerer
	w        Obj
	value    Obj
	params   []Obj
	body     Obj
	isFreeze bool
	syms     map[Obj]bool
	thawOK   bool
	inlined  int
	placed   int
	under    int  // nested closures around the form being walked
	abort    bool // a use under a closure could not be inlined
}

// place puts the closure in front of r, the smallest form that needs it. A use
// that is the value itself is just the closure, with no let of W to itself.
func (s *sinker) place(r Obj) Obj {
	if s.under > 0 {
		s.abort = true
		return r
	}
	s.placed++
	if r == s.w {
		return s.value
	}
	return cons(symLet, cons(s.w, cons(s.value, cons(r, Nil))))
}

func (s *sinker) sink(r Obj) Obj {
	if !occursIn(s.w, r) {
		return r
	}
	p, isp := isPairP(r)
	if !isp {
		return s.place(r)
	}
	if out, ok := s.saturated(r); ok {
		return out
	}
	args := listSlice(p.cdr)
	if p.cdr != Nil && args == nil {
		return s.place(r)
	}
	switch p.car {
	case symIf:
		if len(args) != 3 || occursIn(s.w, args[0]) {
			return s.place(r)
		}
		return cons(symIf, cons(args[0], cons(s.sink(args[1]), cons(s.sink(args[2]), Nil))))
	case symLet:
		if len(args) != 3 || args[0] == s.w || s.syms[args[0]] {
			return s.place(r)
		}
		inV, inB := occursIn(s.w, args[1]), occursIn(s.w, args[2])
		switch {
		case inV && inB:
			return s.place(r)
		case inV:
			return cons(symLet, cons(args[0], cons(s.sink(args[1]), cons(args[2], Nil))))
		}
		return cons(symLet, cons(args[0], cons(args[1], cons(s.sink(args[2]), Nil))))
	case symDo, symAnd, symOr:
		if len(args) != 2 {
			return s.place(r)
		}
		inA, inB := occursIn(s.w, args[0]), occursIn(s.w, args[1])
		switch {
		case inA && inB:
			return s.place(r)
		case inA:
			return cons(p.car, cons(s.sink(args[0]), cons(args[1], Nil)))
		}
		return cons(p.car, cons(args[0], cons(s.sink(args[1]), Nil)))
	case symFreeze:
		if len(args) == 1 {
			return s.sinkClosure(r, Nil, args[0], func(b Obj) Obj { return cons(symFreeze, cons(b, Nil)) })
		}
		return s.place(r)
	case symLambda:
		if len(args) == 2 && args[0] != s.w && !s.syms[args[0]] {
			return s.sinkClosure(r, args[0], args[1], func(b Obj) Obj { return cons(symLambda, cons(args[0], cons(b, Nil))) })
		}
		return s.place(r)
	case symDefun, symCond, symType, symTrapError:
		return s.place(r)
	}
	// A call: with W in exactly one argument, go into that argument. The head
	// is W itself in (W a b), which is a use of the value.
	if p.car == s.w {
		return s.place(r)
	}
	at := -1
	headIn := occursIn(s.w, p.car)
	for i, a := range args {
		if occursIn(s.w, a) {
			if at >= 0 || headIn {
				return s.place(r)
			}
			at = i
		}
	}
	if headIn {
		return cons(s.sink(p.car), p.cdr)
	}
	out := append([]Obj(nil), args...)
	out[at] = s.sink(args[at])
	return cons(p.car, sliceList(out))
}

// sinkClosure goes into a nested closure only if every use in it can be
// inlined: W built inside a closure would be built on each of its runs, and W
// built beside it is what the program did anyway.
func (s *sinker) sinkClosure(r, param, body Obj, rebuild func(Obj) Obj) Obj {
	budget, inlined, placed := s.l.budget, s.inlined, s.placed
	s.under++
	out := s.sink(body)
	s.under--
	if s.under == 0 && s.abort {
		s.abort = false
		s.l.budget, s.inlined, s.placed = budget, inlined, placed
		return s.place(r)
	}
	return rebuild(out)
}

// saturated inlines r if it is a complete application of W.
func (s *sinker) saturated(r Obj) (Obj, bool) {
	var actuals []Obj
	if s.isFreeze {
		p, _ := isPairP(r)
		if p.car != symThaw || !s.thawOK {
			return nil, false
		}
		a := listSlice(p.cdr)
		if len(a) != 1 || a[0] != s.w {
			return nil, false
		}
	} else {
		// ((W a) b) reads from the outside in.
		actuals = make([]Obj, len(s.params))
		f := r
		for i := len(s.params) - 1; i >= 0; i-- {
			p, isp := isPairP(f)
			if !isp {
				return nil, false
			}
			a := listSlice(p.cdr)
			if len(a) != 1 {
				return nil, false
			}
			actuals[i] = a[0]
			f = p.car
		}
		if f != s.w {
			return nil, false
		}
		for i, a := range actuals {
			if occursIn(s.w, a) {
				return nil, false
			}
			for _, earlier := range s.params[:i] {
				if occursIn(earlier, a) {
					return nil, false
				}
			}
		}
	}
	size := formSize(s.body)
	if size > s.l.budget || (s.under > 0 && size > lowerClosureCopy) {
		return nil, false
	}
	s.l.budget -= size
	s.inlined++
	if s.isFreeze {
		s.l.assumesThaw = true
	}
	out := s.body
	for i := len(actuals) - 1; i >= 0; i-- {
		out = s.l.bindParam(s.params[i], actuals[i], out)
	}
	return out, true
}

// bindParam is (let z a b), sharing the variable when a is a variable. A let of
// one variable to another costs a load and a store and needs neither.
func (l *lowerer) bindParam(z, a, b Obj) Obj {
	if IsSymbol(a) && IsSymbol(z) && a != z && !isLowerSpecial(z) && !isLowerSpecial(a) && copyPropSafe(z, a, b) {
		return substVar(b, z, a)
	}
	return cons(symLet, cons(z, cons(a, cons(b, Nil))))
}

// copyPropSafe: nothing in b rebinds a, and z is never in head position, where
// the choice between a global function and a local depends on the name.
func copyPropSafe(z, a, b Obj) bool {
	if mustSymbol(a).function != nil || mustSymbol(z).function != nil {
		return false
	}
	p, isp := isPairP(b)
	if !isp {
		return true
	}
	if p.car == z {
		return false
	}
	switch p.car {
	case symLet:
		rest := listSlice(p.cdr)
		if len(rest) >= 1 && rest[0] == a {
			return false
		}
	case symLambda:
		rest := listSlice(p.cdr)
		if len(rest) >= 1 && rest[0] == a {
			return false
		}
	case symDefun:
		return !occursIn(z, b) && !occursIn(a, b)
	}
	for f := Obj(p.cdr); ; {
		fp, ok := isPairP(f)
		if !ok {
			break
		}
		if !copyPropSafe(z, a, fp.car) {
			return false
		}
		f = fp.cdr
	}
	if _, ish := isPairP(p.car); ish {
		return copyPropSafe(z, a, p.car)
	}
	return true
}

// substVar replaces free z with a in f; copyPropSafe has passed.
func substVar(f, z, a Obj) Obj {
	if f == z {
		return a
	}
	p, isp := isPairP(f)
	if !isp {
		return f
	}
	switch p.car {
	case symLet:
		rest := listSlice(p.cdr)
		if len(rest) == 3 {
			v := substVar(rest[1], z, a)
			if rest[0] == z {
				return cons(symLet, cons(z, cons(v, cons(rest[2], Nil))))
			}
			return cons(symLet, cons(rest[0], cons(v, cons(substVar(rest[2], z, a), Nil))))
		}
	case symLambda:
		rest := listSlice(p.cdr)
		if len(rest) == 2 {
			if rest[0] == z {
				return f
			}
			return cons(symLambda, cons(rest[0], cons(substVar(rest[1], z, a), Nil)))
		}
	case symType:
		rest := listSlice(p.cdr)
		if len(rest) == 2 {
			return cons(symType, cons(substVar(rest[0], z, a), cons(rest[1], Nil)))
		}
	}
	// The head of a call is a symbol copyPropSafe has checked, or a form.
	head := p.car
	if _, ish := isPairP(head); ish {
		head = substVar(head, z, a)
	}
	var args []Obj
	for l := Obj(p.cdr); ; {
		lp, ok := isPairP(l)
		if !ok {
			if l != Nil {
				return f
			}
			break
		}
		args = append(args, substVar(lp.car, z, a))
		l = lp.cdr
	}
	return cons(head, sliceList(args))
}

// compileTest compiles a form in test position and returns the jumps taken when
// it is false, for the caller to patch. An and or or here needs no boolean on
// the stack to be tested again: a false (and a b) has already failed one test,
// and each part still gets the same "if requires a boolean" check, in the same
// order, from the jump that tests it.
func (c *klCompiler) compileTest(form Obj) []int {
	if p, isp := isPairP(form); isp && (p.car == symAnd || p.car == symOr) {
		if args := listSlice(p.cdr); len(args) == 2 {
			if p.car == symAnd {
				return append(c.compileTest(args[0]), c.compileTest(args[1])...)
			}
			aFalse := c.compileTest(args[0])
			toThen := c.emit(OP_JUMP, 0, 0)
			c.patchJumps(aFalse)
			bFalse := c.compileTest(args[1])
			c.patchJump(toThen)
			return bFalse
		}
	}
	c.compileExpr(form, false)
	return []int{c.emitJumpFalse()}
}

func (c *klCompiler) patchJumps(jumps []int) {
	for _, j := range jumps {
		c.patchJump(j)
	}
}

// placeThunk puts the one copy of a thunk reached by tail jumps. The let's body
// is complete and in tail position, so nothing falls through to it and the
// operand stack is empty at every jump.
func (c *klCompiler) placeThunk(thunk *inlineThunk) {
	if len(thunk.jumps) == 0 {
		return
	}
	c.patchJumps(thunk.jumps)
	locals := c.locals
	c.locals = thunk.locals
	c.compileExpr(thunk.body, true)
	c.locals = locals
}

// factorDefine runs the kernel's own pattern factoriser, shen.factor, over what
// define produced, as (factorise +) would, without changing what ps prints or
// what gensym returns next: ps keeps the form it was given, since this changes
// only what is compiled, and the counter is put back.
//
// Only tests that are recognisers and selectors qualify. The factoriser merges
// equal tests, which a where guard with an effect would notice. shen.factor also
// needs a cond ending in a true clause, which define's output always has and a
// hand-written defun may not.
func factorDefine(name string, params []Obj, body Obj) (out Obj) {
	out = body
	if lowerOff || stdFactor == nil || factorBroken {
		return body
	}
	p, isp := isPairP(body)
	if !isp || p.car != symCond {
		return body
	}
	// A rebound shen.factor must not start running inside the compiler. On
	// master it is reached only through (factorise +).
	factor := mustSymbol(symFactor).function
	if factor != stdFactor {
		return body
	}
	clauses := listSlice(p.cdr)
	if len(clauses) < 3 || len(clauses) > maxFactorClauses {
		return body
	}
	ands := 0
	for i, clause := range clauses {
		c := listSlice(clause)
		if len(c) != 2 {
			return body
		}
		if i == len(clauses)-1 {
			if c[0] != True {
				return body
			}
			break
		}
		if !factorableTest(c[0]) {
			return body
		}
		if tp, ok := isPairP(c[0]); ok && tp.car == symAnd {
			ands++
		}
	}
	if ands < 2 {
		return body
	}
	gensym := mustSymbol(symGensym)
	saved := gensym.value
	defer func() {
		gensym.value = saved
		if r := recover(); r != nil {
			out = body
		}
	}()
	var ctl ControlFlow
	form := Call(&ctl, factor, cons(symDefun, cons(MakeSymbol(name), cons(sliceList(params), cons(body, Nil)))))
	if r := listSlice(form); len(r) == 4 && r[0] == symDefun {
		return r[3]
	}
	return body
}

// maxFactorClauses keeps the kernel's factoriser, which is not linear in the
// number of clauses, away from generated tables.
const maxFactorClauses = 200

var symFactor, symGensym, stdFactor Obj // stdFactor: bound by installKernelFast

func init() {
	symFactor = MakeSymbol("shen.factor")
	symGensym = MakeSymbol("shen.*gensym*")
}

// factorableTest accepts the forms define's pattern compiler produces for tests
// that have no effect. An application of anything else is someone's where guard.
func factorableTest(f Obj) bool {
	p, isp := isPairP(f)
	if !isp {
		return true
	}
	args := listSlice(p.cdr)
	if !IsSymbol(p.car) || (p.cdr != Nil && args == nil) {
		return false
	}
	switch mustSymbol(p.car).str {
	case "and", "=", "cons?", "hd", "tl", "tuple?", "fst", "snd", "shen.+string?", "hdstr", "tlstr", "shen.+vector?", "hdv", "tlv":
		for _, a := range args {
			if !factorableTest(a) {
				return false
			}
		}
		return true
	}
	return false
}

// A definition compiled on an assumption about a binding is kept here, to be
// compiled again when that binding changes: a thaw it inlined (thaw rebound), or
// tests the factoriser merged (one of factorSym rebound, since a rebound test
// with an effect must run as many times as the clauses ask for it).
//
// The new code replaces the old through its first instruction (vmCode.redirect),
// so a function object already held by anyone reaches it, as does the next
// iteration of a self-tail loop. An activation already running keeps the code
// it is in, whose registers and constants it depends on, and a closure already
// built keeps its body.
type lowerAssumption struct {
	fn        *BytecodeFunc
	params    []Obj
	raw, form Obj // the definition as given, and after factorDefine
	thaw      bool
}

var (
	// By name: a redefinition replaces the entry, so reloading a file in a
	// long session does not keep every old body alive. An object of the old
	// definition still held somewhere then keeps its code.
	lowerAssumptions = map[string]lowerAssumption{}
	// factorBroken: a test the factoriser may merge has been rebound since boot,
	// so nothing is factorised any more. factorArmed is set at the end of boot,
	// which binds those symbols itself.
	factorBroken, factorArmed bool
)

// compileDefine is CompileFunc's work, without the object.
func compileDefine(name string, params []Obj, raw, form Obj) *BytecodeFunc {
	lowered, thaw := lowerDefine(form)
	fn := compileFuncBody(name, params, lowered)
	if thaw || form != raw {
		lowerAssumptions[name] = lowerAssumption{fn, params, raw, form, thaw}
	}
	return fn
}

// factorSym says whether factorised code depends on sym: the tests the
// factoriser may merge (see factorableTest), and thaw, which runs the failure
// continuations it introduces. Without that, a rebound thaw would be called by
// factorised code where the definition as written never calls it.
func factorSym(s *scmSymbol) bool {
	switch s.str {
	case "thaw", "=", "cons?", "hd", "tl", "tuple?", "fst", "snd", "shen.+string?", "hdstr", "tlstr", "shen.+vector?", "hdv", "tlv":
		return true
	}
	return false
}

// noteRebinding is called whenever a symbol's function changes.
func noteRebinding(s *scmSymbol) {
	isThaw := Obj(&s.scmHead) == symThaw
	isFactor := factorArmed && factorSym(s)
	if !isThaw && !isFactor {
		return
	}
	factorBroken = factorBroken || isFactor
	pending := lowerAssumptions
	lowerAssumptions = map[string]lowerAssumption{}
	for name, a := range pending {
		form := a.form
		if isFactor {
			form = a.raw
		}
		if form == a.form && !(isThaw && a.thaw) {
			lowerAssumptions[name] = a
			continue
		}
		nf := compileDefine(a.fn.Name, a.params, a.raw, form)
		if a.fn.rc == nil {
			// Never run, so there is no code to redirect and no activation to
			// disturb. Assumptions compileDefine just noted name nf; move them.
			*a.fn = *nf
			if e, ok := lowerAssumptions[name]; ok && e.fn == nf {
				e.fn = a.fn
				lowerAssumptions[name] = e
			}
			continue
		}
		a.fn.rc.redirectTo(makeBytecodeObj(nf, nil))
	}
}
