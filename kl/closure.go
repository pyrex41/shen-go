package kl

// Closure-compiled executor: an alternative to the bytecode VM, selected with
// SHEN_GO_EXEC=closure. Each KL expression becomes a small Go closure
// (cnode); a function is a tree of them run over a frame of Obj slots.
//
// Frames live on the Go stack, in a fixed-size array chosen per function
// (clEnter4..64), reached through a pointer the escape analyser is told not
// to follow (clNoescape). Nothing stores a frame pointer in the heap: lambdas
// copy the values they capture. Slots at or above fp are arguments then lets;
// slots below fp are this activation's upvalues, so a slot number is known the
// moment a variable is resolved. A function too big for the largest array gets
// a heap frame and is compiled in heapFrame mode, whose stores have barriers.
//
// A node returns a value, or one of two marks:
//   nil        a tail call is pending in ctl (kind == ControlFlowApply);
//   clSelfTail this frame's arguments were rewritten, run the body again.
// Only nodes in tail position can return them. A non-tail call resolves nil
// with trampoline, so everything else only ever sees values. nil cannot be a
// Shen value: <-address maps an empty slot to `undefined` and slotFromObj
// already treats nil as "not an Obj".
//
// The VM and this executor interoperate: a closure-compiled function is still
// a scmHeadBytecodeFunc, with fn.Cl set and no Code, and vmExec/vmExecSlots
// enter it through clApply.

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

type cnode func(ctl *ControlFlow, fp unsafe.Pointer) Obj
type ctest func(ctl *ControlFlow, fp unsafe.Pointer) bool

// clCode is one compiled function.
type clCode struct {
	arity int
	// enter runs an activation with up to four arguments; enterN any number.
	enter  func(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj
	body   cnode
	fpOff  uintptr // byte offset of fp in the frame array: 8*nupvals
	nslots int     // arguments + lets + upvalues
	bucket int     // frame array size; 0 means heap frame
	// direct: no upvalues and a frame of at most 8 slots, so a call node can
	// hold the frame itself and call body, skipping enter's Go frame.
	direct bool
}

var clEnabled = strings.EqualFold(os.Getenv("SHEN_GO_EXEC"), "closure")

// ClosureExecEnabled reports whether new definitions are closure-compiled.
func ClosureExecEnabled() bool { return clEnabled }

// clDirty is the one-load primitive guard. It is false while every primitive
// the compiler inlines is still bound to its canonical implementation and
// typed IR is on. When false an inlined primitive needs no further check; when
// true each site falls to clGuardOK, the per-symbol check, so a rebinding of
// `hd` slows every inlined site down but changes the meaning of only `hd`.
var clDirty bool

// noteRebind is called by everything that writes a symbol's function. A
// primitive turns the guard dirty only if some compiled site has inlined it,
// so a boot-time rebinding of a primitive nothing inlined yet costs nothing:
// the compiler sees it is no longer canonical and emits an ordinary call.
func noteRebind(s *scmSymbol, f Obj) {
	if s.clInlined && f != s.canonical {
		clDirty = true
	}
}

// clMark records that a site inlines sym, which must then be guarded.
func clMark(sym Obj) { mustSymbol(sym).clInlined = true }

func clGuardOK(sym Obj) bool { return TypedIREnabled() && HasCanonicalPrimitiveBinding(sym) }

//go:nosplit
//go:nocheckptr
func clNoescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

// clSelfMark's address is the self-tail mark.
var clSelfMark scmHead = 99

func clSelfTail() Obj { return Obj(&clSelfMark) }

// clLoad reads a slot. Inlined everywhere.
func clLoad(fp unsafe.Pointer, off int) Obj { return *(*Obj)(unsafe.Add(fp, off)) }

// clStore writes a stack-frame slot with no write barrier, exactly what the
// compiler does for a store to a local. It must never be used on a heap frame.
func clStore(fp unsafe.Pointer, off int, v Obj) {
	*(*uintptr)(unsafe.Add(fp, off)) = uintptr(unsafe.Pointer(v))
}

func clStoreHeap(fp unsafe.Pointer, off int, v Obj) { *(*Obj)(unsafe.Add(fp, off)) = v }

func clEnterSetup(ctl *ControlFlow, bf *scmBytecodeFunc, fp unsafe.Pointer) {
	for i, u := range bf.upvals {
		clStore(fp, -8*(i+1), u)
	}
	if ctl.stepLimit != 0 {
		ctl.tick()
	}
}

func clEnter4(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj {
	var fr [4]Obj
	fp := unsafe.Add(clNoescape(unsafe.Pointer(&fr)), cl.fpOff)
	clStore(fp, 0, a0)
	clStore(fp, 8, a1)
	clStore(fp, 16, a2)
	clStore(fp, 24, a3)
	if len(bf.upvals) != 0 || ctl.stepLimit != 0 {
		clEnterSetup(ctl, bf, fp)
	}
	return cl.body(ctl, fp)
}

func clEnter8(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj {
	var fr [8]Obj
	fp := unsafe.Add(clNoescape(unsafe.Pointer(&fr)), cl.fpOff)
	clStore(fp, 0, a0)
	clStore(fp, 8, a1)
	clStore(fp, 16, a2)
	clStore(fp, 24, a3)
	if len(bf.upvals) != 0 || ctl.stepLimit != 0 {
		clEnterSetup(ctl, bf, fp)
	}
	return cl.body(ctl, fp)
}

func clEnter16(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj {
	var fr [16]Obj
	fp := unsafe.Add(clNoescape(unsafe.Pointer(&fr)), cl.fpOff)
	clStore(fp, 0, a0)
	clStore(fp, 8, a1)
	clStore(fp, 16, a2)
	clStore(fp, 24, a3)
	if len(bf.upvals) != 0 || ctl.stepLimit != 0 {
		clEnterSetup(ctl, bf, fp)
	}
	return cl.body(ctl, fp)
}

func clEnter32(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj {
	var fr [32]Obj
	fp := unsafe.Add(clNoescape(unsafe.Pointer(&fr)), cl.fpOff)
	clStore(fp, 0, a0)
	clStore(fp, 8, a1)
	clStore(fp, 16, a2)
	clStore(fp, 24, a3)
	if len(bf.upvals) != 0 || ctl.stepLimit != 0 {
		clEnterSetup(ctl, bf, fp)
	}
	return cl.body(ctl, fp)
}

func clEnter64(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj {
	var fr [64]Obj
	fp := unsafe.Add(clNoescape(unsafe.Pointer(&fr)), cl.fpOff)
	clStore(fp, 0, a0)
	clStore(fp, 8, a1)
	clStore(fp, 16, a2)
	clStore(fp, 24, a3)
	if len(bf.upvals) != 0 || ctl.stepLimit != 0 {
		clEnterSetup(ctl, bf, fp)
	}
	return cl.body(ctl, fp)
}

// clEnterHeap is the entry for a function whose frame does not fit clEnter64.
func clEnterHeap(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, a0, a1, a2, a3 Obj) Obj {
	return clEnterSlice(ctl, bf, cl, []Obj{a0, a1, a2, a3}[:min(cl.arity, 4)])
}

// clEnterSlice runs an activation of any arity from a slice of arguments, on a
// heap frame. It serves arity above 4 and functions compiled in heapFrame mode.
// A heap frame is correct for stack-mode code only if its stores are barriered,
// so stack-mode code with more than four parameters uses clEnterStackN instead.
func clEnterSlice(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, args []Obj) Obj {
	if cl.bucket != 0 {
		return clEnterStackN(ctl, bf, cl, args)
	}
	fr := make([]Obj, cl.nslots+1)
	fp := unsafe.Add(unsafe.Pointer(&fr[0]), cl.fpOff)
	for i, a := range args {
		clStoreHeap(fp, 8*i, a)
	}
	for i, u := range bf.upvals {
		clStoreHeap(fp, -8*(i+1), u)
	}
	if ctl.stepLimit != 0 {
		ctl.tick()
	}
	return cl.body(ctl, fp)
}

func clEnterStackN(ctl *ControlFlow, bf *scmBytecodeFunc, cl *clCode, args []Obj) Obj {
	var fr8 [8]Obj
	var fr64 [64]Obj
	var base unsafe.Pointer
	if cl.bucket <= 8 {
		base = clNoescape(unsafe.Pointer(&fr8))
	} else {
		base = clNoescape(unsafe.Pointer(&fr64))
	}
	fp := unsafe.Add(base, cl.fpOff)
	for i, a := range args {
		clStore(fp, 8*i, a)
	}
	clEnterSetup(ctl, bf, fp)
	return cl.body(ctl, fp)
}

// clApply runs bf for vmExec. It reports through ctl, as a VM activation does:
// a value by ctl.Return; a pending tail call is already in ctl.
func clApply(ctl *ControlFlow, bf *scmBytecodeFunc, args []Obj) {
	cl := bf.cl
	var r Obj
	switch len(args) {
	case 0:
		r = cl.enter(ctl, bf, cl, nil, nil, nil, nil)
	case 1:
		r = cl.enter(ctl, bf, cl, args[0], nil, nil, nil)
	case 2:
		r = cl.enter(ctl, bf, cl, args[0], args[1], nil, nil)
	case 3:
		r = cl.enter(ctl, bf, cl, args[0], args[1], args[2], nil)
	case 4:
		r = cl.enter(ctl, bf, cl, args[0], args[1], args[2], args[3])
	default:
		r = clEnterSlice(ctl, bf, cl, args)
	}
	if r != nil {
		ctl.Return(r)
	}
}

func clApplySlots(ctl *ControlFlow, bf *scmBytecodeFunc, args []vmSlot) {
	var inline [8]Obj
	objs := inline[:0]
	for _, a := range args {
		objs = append(objs, a.objValue())
	}
	clApply(ctl, bf, objs)
}

// ---- compiler ----

type clxKind uint8

const (
	clGeneral clxKind = iota
	clLocal
	clConst
)

// clx is a compiled expression: its node, plus what a parent can do without
// calling that node.
type clx struct {
	fn   cnode
	kind clxKind
	off  int
	k    Obj
	// shape is set on an if node that tests a slot in place, so an enclosing
	// clause can fuse with it and save a Go frame on the common two-clause
	// pattern match: (cond ((= [] V) A) ((cons? V) B) (true C)).
	shape *clIfShape
}

// clIfShape: if slot off is k (or a pair, when k is nil) then a else b.
type clIfShape struct {
	off  int
	k    Obj
	a, b cnode
}

// clOperand is an inlined operand: fn == nil means "load the slot at off", so
// a local costs a test and a load instead of an indirect call and a spill.
type clOperand struct {
	fn  cnode
	off int
}

func (x clx) op() clOperand {
	if x.kind == clLocal {
		return clOperand{nil, x.off}
	}
	return clOperand{x.fn, 0}
}

func clGet(a *clOperand, ctl *ControlFlow, fp unsafe.Pointer) Obj {
	if a.fn == nil {
		return *(*Obj)(unsafe.Add(fp, a.off))
	}
	return a.fn(ctl, fp)
}

type clThunk struct {
	body   Obj
	locals map[Obj]int
}

type clUpval struct {
	sym Obj
	off int // slot in the enclosing function
}

type clCompiler struct {
	name      string
	arity     int
	nlocals   int
	locals    map[Obj]int
	upvals    []clUpval
	outer     *clCompiler
	thunks    map[int]clThunk
	selfLoop  bool
	heapFrame bool
}

const clMaxBucket = 64

// CompileFuncClosure compiles a KL function to closures.
func clCompileFunc(name string, params []Obj, body Obj) Obj {
	if !TypedIREnabled() {
		clDirty = true
	}
	fn := clCompileBody(name, params, body, nil, false)
	return makeBytecodeObj(fn, nil)
}

func clCompileBody(name string, params []Obj, body Obj, outer *clCompiler, heap bool) *BytecodeFunc {
	c := &clCompiler{name: name, arity: len(params), nlocals: len(params), locals: make(map[Obj]int, len(params)+2), outer: outer, heapFrame: heap}
	for i, p := range params {
		c.locals[p] = i
	}
	b := c.expr(body, true).fn
	if !heap && len(c.upvals)+max(c.nlocals, 4) > clMaxBucket {
		// Keep the upvalue resolution order stable in the enclosing compiler:
		// outer upvalues found by this pass are found again by the next.
		return clCompileBody(name, params, body, outer, true)
	}
	if c.selfLoop {
		inner := b
		b = func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			for {
				r := inner(ctl, fp)
				if r != Obj(&clSelfMark) {
					return r
				}
				if ctl.stepLimit != 0 {
					ctl.tick()
				}
			}
		}
	}
	code := &clCode{arity: len(params), body: b, fpOff: uintptr(8 * len(c.upvals))}
	code.nslots = len(c.upvals) + max(c.nlocals, 4)
	switch {
	case heap:
		code.enter = clEnterHeap
	case code.nslots <= 4:
		code.enter, code.bucket = clEnter4, 4
	case code.nslots <= 8:
		code.enter, code.bucket = clEnter8, 8
	case code.nslots <= 16:
		code.enter, code.bucket = clEnter16, 16
	case code.nslots <= 32:
		code.enter, code.bucket = clEnter32, 32
	default:
		code.enter, code.bucket = clEnter64, 64
	}
	code.direct = !heap && len(c.upvals) == 0 && code.nslots <= 8
	return &BytecodeFunc{Name: name, Arity: len(params), Nlocals: c.nlocals, Cl: code, clUpvalOffs: c.upvalOffs()}
}

func (c *clCompiler) upvalOffs() []int {
	if len(c.upvals) == 0 {
		return nil
	}
	o := make([]int, len(c.upvals))
	for i, u := range c.upvals {
		o[i] = u.off
	}
	return o
}

// resolve finds sym as a slot offset. Same search order as klCompiler.resolveVar.
func (c *clCompiler) resolve(sym Obj) (int, bool) {
	if idx, ok := c.locals[sym]; ok {
		return 8 * idx, true
	}
	for i, u := range c.upvals {
		if u.sym == sym {
			return -8 * (i + 1), true
		}
	}
	if c.outer != nil {
		off, ok := c.outer.resolve(sym)
		if !ok {
			return 0, false
		}
		c.upvals = append(c.upvals, clUpval{sym, off})
		return -8 * len(c.upvals), true
	}
	return 0, false
}

func (c *clCompiler) newLocal(sym Obj) int {
	slot := c.nlocals
	c.nlocals++
	c.locals[sym] = slot
	return slot
}

func clConstX(k Obj) clx {
	return clx{fn: func(*ControlFlow, unsafe.Pointer) Obj { return k }, kind: clConst, k: k}
}

func clLocalX(off int) clx {
	return clx{fn: func(_ *ControlFlow, fp unsafe.Pointer) Obj { return *(*Obj)(unsafe.Add(fp, off)) }, kind: clLocal, off: off}
}

func (c *clCompiler) expr(form Obj, tail bool) clx {
	switch *form {
	case scmHeadSymbol:
		if off, ok := c.resolve(form); ok {
			return clLocalX(off)
		}
		return clConstX(form)
	case scmHeadPair:
		return c.form(form, tail)
	}
	return clConstX(form)
}

func clGen(fn cnode) clx { return clx{fn: fn} }

func (c *clCompiler) form(form Obj, tail bool) clx {
	pair := mustPair(form)
	if !IsSymbol(pair.car) {
		return c.call(pair.car, pair.cdr, tail)
	}
	head, rest := pair.car, pair.cdr
	switch head {
	case symDefun:
		return c.defun(car(rest), cadr(rest), caddr(rest))
	case symLambda:
		return c.lambda("lambda", []Obj{car(rest)}, cadr(rest))
	case symFreeze:
		return c.lambda("lambda", nil, car(rest))
	case symLet:
		return c.let(car(rest), cadr(rest), caddr(rest), tail)
	case symIf:
		if listLength(rest) == 3 {
			return c.ifNode(car(rest), cadr(rest), caddr(rest), tail)
		}
		return c.call(head, rest, tail)
	case symAnd:
		return c.ifNode(car(rest), cadr(rest), False, tail)
	case symOr:
		return c.ifNode(car(rest), True, cadr(rest), tail)
	case symCond:
		return c.cond(rest, tail)
	case symDo:
		a := c.expr(car(rest), false)
		b := c.expr(cadr(rest), tail)
		af, bf := a.fn, b.fn
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			af(ctl, fp)
			return bf(ctl, fp)
		})
	case symType:
		return c.expr(car(rest), tail)
	case symTrapError:
		return c.trapError(car(rest), cadr(rest))
	}
	return c.call(head, rest, tail)
}

func (c *clCompiler) defun(name, params, body Obj) clx {
	in := c.closureNode(mustSymbol(name).str, ListToSlice(params), body)
	return c.callGlobal(MakeSymbol("defun"), []clx{clConstX(name), in}, false)
}

func (c *clCompiler) lambda(name string, params []Obj, body Obj) clx {
	return c.closureNode(name, params, body)
}

func (c *clCompiler) closureNode(name string, params []Obj, body Obj) clx {
	fn := clCompileBody(name, params, body, c, c.heapFrame && false)
	offs := fn.clUpvalOffs
	switch len(offs) {
	case 0:
		return clGen(func(*ControlFlow, unsafe.Pointer) Obj {
			o, _ := makeClosure(fn, 0)
			return o
		})
	case 1:
		o0 := offs[0]
		return clGen(func(_ *ControlFlow, fp unsafe.Pointer) Obj {
			o, up := makeClosure(fn, 1)
			up[0] = clLoad(fp, o0)
			return o
		})
	}
	return clGen(func(_ *ControlFlow, fp unsafe.Pointer) Obj {
		o, up := makeClosure(fn, len(offs))
		for i := range up {
			up[i] = clLoad(fp, offs[i])
		}
		return o
	})
}

func (c *clCompiler) let(x, val, body Obj, tail bool) clx {
	if fb, ok := isFreezeForm(val); ok {
		thaws, others := walkThawUses(body, x)
		if others == 0 && thaws > 0 {
			captured := make(map[Obj]int, len(c.locals))
			for sym, slot := range c.locals {
				captured[sym] = slot
			}
			oldIdx, hadOld := c.locals[x]
			slot := c.newLocal(x)
			if c.thunks == nil {
				c.thunks = make(map[int]clThunk)
			}
			c.thunks[slot] = clThunk{body: fb, locals: captured}
			r := c.expr(body, tail)
			delete(c.thunks, slot)
			if hadOld {
				c.locals[x] = oldIdx
			} else {
				delete(c.locals, x)
			}
			return r
		}
	}
	v := c.expr(val, false)
	oldIdx, hadOld := c.locals[x]
	off := 8 * c.newLocal(x)
	b := c.expr(body, tail)
	if hadOld {
		c.locals[x] = oldIdx
	} else {
		delete(c.locals, x)
	}
	vf, bf := v.fn, b.fn
	if c.heapFrame {
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			clStoreHeap(fp, off, vf(ctl, fp))
			return bf(ctl, fp)
		})
	}
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		clStore(fp, off, vf(ctl, fp))
		return bf(ctl, fp)
	})
}

func (c *clCompiler) ifNode(cond, thenE, elseE Obj, tail bool) clx {
	// Shapes that test in place. The compile order matches compileIf; a test
	// compiled here has no effect on what the branches see.
	if ok, p := isPair(cond); ok && IsSymbol(p.car) {
		args := ListToSlice(p.cdr)
		if p.car == symNumEq && len(args) == 2 && HasCanonicalPrimitiveBinding(p.car) {
			if !c.bothConst(args[0], args[1]) {
				clMark(p.car)
				x, y := c.expr(args[0], false), c.expr(args[1], false)
				if y.kind == clLocal && x.kind == clConst {
					x, y = y, x
				}
				a, be := c.expr(thenE, tail).fn, c.expr(elseE, tail)
				b := be.fn
				if x.kind == clLocal && y.kind == clConst && clIdentityConst(y.k) {
					off, k, sym := x.off, y.k, p.car
					slow := func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
						v := *(*Obj)(unsafe.Add(fp, off))
						if clDirty {
							if clSlowEq(ctl, sym, v, k) {
								return a(ctl, fp)
							}
							return b(ctl, fp)
						}
						if v == k {
							return a(ctl, fp)
						}
						return b(ctl, fp)
					}
					r := clx{fn: slow, shape: &clIfShape{off, k, a, b}}
					if be.shape != nil {
						r.fn = clFuse2(off, k, a, be.shape, slow)
					}
					return r
				}
				t := clTestEq(p.car, x, y)
				return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
					if t(ctl, fp) {
						return a(ctl, fp)
					}
					return b(ctl, fp)
				})
			}
		}
		if len(args) == 1 && IsSymbol(p.car) && GetSymbol(p.car) == "cons?" && guardedPrimitiveArity(p.car) == 1 {
			if off, ok := c.resolve(args[0]); ok && IsSymbol(args[0]) {
				clMark(p.car)
				a, b := c.expr(thenE, tail).fn, c.expr(elseE, tail).fn
				sym := p.car
				return clx{fn: func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
					v := *(*Obj)(unsafe.Add(fp, off))
					if clDirty {
						if clTruth(clGP(ctl, sym, gpConsp, v)) {
							return a(ctl, fp)
						}
						return b(ctl, fp)
					}
					if !clIsFix(v) && *v == scmHeadPair {
						return a(ctl, fp)
					}
					return b(ctl, fp)
				}, shape: &clIfShape{off, nil, a, b}}
			}
		}
	}
	t := c.test(cond)
	a, b := c.expr(thenE, tail).fn, c.expr(elseE, tail).fn
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		if t(ctl, fp) {
			return a(ctl, fp)
		}
		return b(ctl, fp)
	})
}

// clFuse2 is (if (= slot k) a <shape>) as one node. Any rebinding takes slow,
// the unfused tree with its per-site guards.
func clFuse2(off int, k Obj, a cnode, sh *clIfShape, slow cnode) cnode {
	off2, k2, a2, b2 := sh.off, sh.k, sh.a, sh.b
	if k2 == nil {
		return func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			if clDirty {
				return slow(ctl, fp)
			}
			if *(*Obj)(unsafe.Add(fp, off)) == k {
				return a(ctl, fp)
			}
			v := *(*Obj)(unsafe.Add(fp, off2))
			if !clIsFix(v) && *v == scmHeadPair {
				return a2(ctl, fp)
			}
			return b2(ctl, fp)
		}
	}
	return func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		if clDirty {
			return slow(ctl, fp)
		}
		if *(*Obj)(unsafe.Add(fp, off)) == k {
			return a(ctl, fp)
		}
		if *(*Obj)(unsafe.Add(fp, off2)) == k2 {
			return a2(ctl, fp)
		}
		return b2(ctl, fp)
	}
}

func (c *clCompiler) cond(clauses Obj, tail bool) clx {
	if *clauses == scmHeadNull {
		return clConstX(Nil)
	}
	clause := car(clauses)
	if car(clause) == True {
		return c.expr(cadr(clause), tail)
	}
	return c.condClause(car(clause), cadr(clause), cdr(clauses), tail)
}

// condClause is ifNode whose else branch is the remaining clauses.
func (c *clCompiler) condClause(test, action, rest Obj, tail bool) clx {
	// (cond (a b) rest...) is (if a b (cond rest...)): build the form rather
	// than repeat ifNode. The compiled meaning of cond with no clause left is
	// the Nil constant; a quoted rest cannot be confused with a call because
	// `cond` is a special form here too.
	return c.ifNode(test, action, cons(symCond, rest), tail)
}

func (c *clCompiler) trapError(body, handler Obj) clx {
	if def, ok := trapHandlerDefault(handler); ok {
		if args, ok := callForm(body, "value", 1); ok {
			a := c.expr(args[0], false)
			d := c.lambda("lambda", nil, def)
			return c.callGlobal(MakeSymbol("_kl.value/or"), []clx{a, d}, false)
		}
		if args, ok := callForm(body, "<-vector", 2); ok {
			a, b := c.expr(args[0], false), c.expr(args[1], false)
			d := c.lambda("lambda", nil, def)
			return c.callGlobal(MakeSymbol("_kl.<-vector/or"), []clx{a, b, d}, false)
		}
	}
	b := c.lambda("lambda", nil, body)
	h := c.expr(handler, false)
	return c.callGlobal(MakeSymbol("try-catch"), []clx{b, h}, false)
}

func (c *clCompiler) constValue(o Obj) (Obj, bool) {
	if o == nil {
		return nil, false
	}
	switch *o {
	case scmHeadNumber, scmHeadString, scmHeadBoolean, scmHeadNull:
		return o, true
	case scmHeadSymbol:
		if _, local := c.resolveNoCapture(o); !local {
			return o, true
		}
	}
	return nil, false
}

// resolveNoCapture answers "is sym a variable" like resolve. klCompiler.constValue
// calls resolveVar, which records a capture as a side effect, so this must too:
// the slot numbering of upvalues is otherwise the same either way.
func (c *clCompiler) resolveNoCapture(o Obj) (int, bool) { return c.resolve(o) }

func (c *clCompiler) bothConst(x, y Obj) bool {
	if _, ok := c.constValue(x); !ok {
		return false
	}
	_, ok := c.constValue(y)
	return ok
}

func (c *clCompiler) args(l []Obj) []clx {
	xs := make([]clx, len(l))
	for i, a := range l {
		xs[i] = c.expr(a, false)
	}
	return xs
}

func (c *clCompiler) call(fn Obj, args Obj, tail bool) clx {
	argList := make([]Obj, 0, 4)
	for l := args; *l == scmHeadPair; l = cdr(l) {
		argList = append(argList, car(l))
	}
	nArgs := len(argList)

	if IsSymbol(fn) {
		if n := guardedPrimitiveArity(fn); n != 0 && n == nArgs {
			clMark(fn)
			return clGuarded(fn, c.args(argList))
		}
		if nArgs == 1 && GetSymbol(fn) == "thaw" {
			if fb, ok := isFreezeForm(argList[0]); ok {
				return c.expr(fb, tail)
			}
			if slot, local := c.locals[argList[0]]; local && IsSymbol(argList[0]) {
				if thunk, ok := c.thunks[slot]; ok {
					locals := c.locals
					c.locals = thunk.locals
					r := c.expr(thunk.body, tail)
					c.locals = locals
					return r
				}
			}
		}
		if nArgs == 2 && fn == symNumEq && HasCanonicalPrimitiveBinding(fn) {
			if a, ok := c.constValue(argList[0]); ok {
				if b, ok := c.constValue(argList[1]); ok {
					return clConstX(equal(a, b))
				}
			}
		}
		if nArgs == 1 && GetSymbol(fn) == "intern" && HasCanonicalPrimitiveBinding(fn) && IsString(argList[0]) {
			return clConstX(PrimIntern(argList[0]))
		}
		if nArgs == 1 && fn == symNot && HasCanonicalPrimitiveBinding(fn) {
			clMark(fn)
			return clNot(fn, c.expr(argList[0], false))
		}
		if nArgs == 2 && HasCanonicalPrimitiveBinding(fn) && intrinsicOp2(fn) != 0 {
			clMark(fn)
			if folded, ok := foldFixnum2(fn, argList[0], argList[1]); ok {
				return clFolded(fn, folded, argList[0], argList[1])
			}
			return clArith(fn, c.expr(argList[0], false), c.expr(argList[1], false))
		}
		if tail && mustSymbol(fn).str == c.name && nArgs == c.arity {
			if _, isVar := c.resolve(fn); !isVar {
				return c.selfTail(c.args(argList))
			}
		}
		sym := mustSymbol(fn)
		if sym.function == nil {
			if off, ok := c.resolve(fn); ok {
				return c.callDynamic(clLocalX(off), c.args(argList), tail)
			}
		}
		return c.callGlobal(fn, c.args(argList), tail)
	}
	h := c.expr(fn, false)
	return c.callDynamic(h, c.args(argList), tail)
}

func (c *clCompiler) selfTail(args []clx) clx {
	c.selfLoop = true
	// Rewrite only the arguments that change: (fold F (F A X) Xs) keeps F.
	type move struct {
		off int
		op  clOperand
	}
	var mv []move
	for i, a := range args {
		if a.kind == clLocal && a.off == 8*i {
			continue
		}
		mv = append(mv, move{8 * i, a.op()})
	}
	if c.heapFrame || len(mv) > 4 {
		heap := c.heapFrame
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			var inline [8]Obj
			tmp := inline[:0]
			for i := range mv {
				tmp = append(tmp, clGet(&mv[i].op, ctl, fp))
			}
			for i := range mv {
				if heap {
					clStoreHeap(fp, mv[i].off, tmp[i])
				} else {
					clStore(fp, mv[i].off, tmp[i])
				}
			}
			return Obj(&clSelfMark)
		})
	}
	switch len(mv) {
	case 0:
		return clGen(func(*ControlFlow, unsafe.Pointer) Obj { return Obj(&clSelfMark) })
	case 1:
		o0, a0 := mv[0].off, mv[0].op
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			clStore(fp, o0, clGet(&a0, ctl, fp))
			return Obj(&clSelfMark)
		})
	case 2:
		o0, a0, o1, a1 := mv[0].off, mv[0].op, mv[1].off, mv[1].op
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			x := clGet(&a0, ctl, fp)
			y := clGet(&a1, ctl, fp)
			clStore(fp, o0, x)
			clStore(fp, o1, y)
			return Obj(&clSelfMark)
		})
	case 3:
		o0, a0, o1, a1, o2, a2 := mv[0].off, mv[0].op, mv[1].off, mv[1].op, mv[2].off, mv[2].op
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			x := clGet(&a0, ctl, fp)
			y := clGet(&a1, ctl, fp)
			z := clGet(&a2, ctl, fp)
			clStore(fp, o0, x)
			clStore(fp, o1, y)
			clStore(fp, o2, z)
			return Obj(&clSelfMark)
		})
	}
	o0, a0, o1, a1, o2, a2, o3, a3 := mv[0].off, mv[0].op, mv[1].off, mv[1].op, mv[2].off, mv[2].op, mv[3].off, mv[3].op
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		x := clGet(&a0, ctl, fp)
		y := clGet(&a1, ctl, fp)
		z := clGet(&a2, ctl, fp)
		w := clGet(&a3, ctl, fp)
		clStore(fp, o0, x)
		clStore(fp, o1, y)
		clStore(fp, o2, z)
		clStore(fp, o3, w)
		return Obj(&clSelfMark)
	})
}

func clUnbound(s *scmSymbol) {
	panic(MakeError(fmt.Sprintf("function %s not bound", s.str)))
}

// clTailN sets up a tail call. A native with no captured arguments runs here
// rather than one trampoline hop later: it returns before anything else runs,
// so the Go stack is no deeper when it tail-calls in turn.
func clTailN(ctl *ControlFlow, f Obj, args []Obj) Obj {
	ctl.tailApplySlice(f, args)
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); len(nf.captured) == 0 && nf.require == len(args) {
			if ctl.stepLimit != 0 {
				ctl.tick()
			}
			nf.fn(ctl)
			if ctl.kind == ControlFlowReturn {
				ret := ctl.data[ctl.pos]
				ctl.data = ctl.data[:ctl.pos]
				return ret
			}
		}
	}
	return nil
}

func clTail0(ctl *ControlFlow, f Obj) Obj { return clTailN(ctl, f, nil) }
func clTail1(ctl *ControlFlow, f, a Obj) Obj {
	t := [1]Obj{a}
	return clTailN(ctl, f, t[:])
}
func clTail2(ctl *ControlFlow, f, a, b Obj) Obj {
	t := [2]Obj{a, b}
	return clTailN(ctl, f, t[:])
}
func clTail3(ctl *ControlFlow, f, a, b, d Obj) Obj {
	t := [3]Obj{a, b, d}
	return clTailN(ctl, f, t[:])
}
func clTail4(ctl *ControlFlow, f, a, b, d, e Obj) Obj {
	t := [4]Obj{a, b, d, e}
	return clTailN(ctl, f, t[:])
}

// clCallCheck resolves a non-tail call's result.
func clFinish(ctl *ControlFlow, r Obj) Obj {
	if r == nil {
		return trampoline(ctl)
	}
	return r
}

func (c *clCompiler) callGlobal(fnSym Obj, args []clx, tail bool) clx {
	s := mustSymbol(fnSym)
	switch len(args) {
	case 0:
		if tail {
			return clGen(func(ctl *ControlFlow, _ unsafe.Pointer) Obj {
				f := s.function
				if f == nil {
					clUnbound(s)
				}
				return clTail0(ctl, f)
			})
		}
		return clGen(func(ctl *ControlFlow, _ unsafe.Pointer) Obj {
			f := s.function
			if f == nil {
				clUnbound(s)
			}
			return clDo0(ctl, f)
		})
	case 1:
		a := args[0].op()
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				f := s.function
				if f == nil {
					clUnbound(s)
				}
				return clTail1(ctl, f, clGet(&a, ctl, fp))
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			f := s.function
			if f == nil {
				clUnbound(s)
			}
			x := clGet(&a, ctl, fp)
			if *f == scmHeadBytecodeFunc {
				bf := mustBytecodeFunc(f)
				if cl := bf.cl; cl != nil && cl.arity == 1 {
					if cl.direct && ctl.stepLimit == 0 {
						var fr [8]Obj
						fr[0] = x
						r := cl.body(ctl, clNoescape(unsafe.Pointer(&fr)))
						if r == nil {
							return trampoline(ctl)
						}
						return r
					}
					r := cl.enter(ctl, bf, cl, x, nil, nil, nil)
					if r == nil {
						return trampoline(ctl)
					}
					return r
				}
			}
			return Call(ctl, f, x)
		})
	case 2:
		a, b := args[0].op(), args[1].op()
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				f := s.function
				if f == nil {
					clUnbound(s)
				}
				x := clGet(&a, ctl, fp)
				y := clGet(&b, ctl, fp)
				return clTail2(ctl, f, x, y)
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			f := s.function
			if f == nil {
				clUnbound(s)
			}
			x := clGet(&a, ctl, fp)
			y := clGet(&b, ctl, fp)
			if *f == scmHeadBytecodeFunc {
				bf := mustBytecodeFunc(f)
				if cl := bf.cl; cl != nil && cl.arity == 2 {
					if cl.direct && ctl.stepLimit == 0 {
						var fr [8]Obj
						fr[0] = x
						fr[1] = y
						r := cl.body(ctl, clNoescape(unsafe.Pointer(&fr)))
						if r == nil {
							return trampoline(ctl)
						}
						return r
					}
					r := cl.enter(ctl, bf, cl, x, y, nil, nil)
					if r == nil {
						return trampoline(ctl)
					}
					return r
				}
			}
			return Call(ctl, f, x, y)
		})
	case 3:
		a, b, d := args[0].op(), args[1].op(), args[2].op()
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				f := s.function
				if f == nil {
					clUnbound(s)
				}
				x := clGet(&a, ctl, fp)
				y := clGet(&b, ctl, fp)
				z := clGet(&d, ctl, fp)
				return clTail3(ctl, f, x, y, z)
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			f := s.function
			if f == nil {
				clUnbound(s)
			}
			x := clGet(&a, ctl, fp)
			y := clGet(&b, ctl, fp)
			z := clGet(&d, ctl, fp)
			if *f == scmHeadBytecodeFunc {
				bf := mustBytecodeFunc(f)
				if cl := bf.cl; cl != nil && cl.arity == 3 {
					if cl.direct && ctl.stepLimit == 0 {
						var fr [8]Obj
						fr[0] = x
						fr[1] = y
						fr[2] = z
						r := cl.body(ctl, clNoescape(unsafe.Pointer(&fr)))
						if r == nil {
							return trampoline(ctl)
						}
						return r
					}
					r := cl.enter(ctl, bf, cl, x, y, z, nil)
					if r == nil {
						return trampoline(ctl)
					}
					return r
				}
			}
			return Call(ctl, f, x, y, z)
		})
	case 4:
		a, b, d, e := args[0].op(), args[1].op(), args[2].op(), args[3].op()
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				f := s.function
				if f == nil {
					clUnbound(s)
				}
				x := clGet(&a, ctl, fp)
				y := clGet(&b, ctl, fp)
				z := clGet(&d, ctl, fp)
				w := clGet(&e, ctl, fp)
				return clTail4(ctl, f, x, y, z, w)
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			f := s.function
			if f == nil {
				clUnbound(s)
			}
			x := clGet(&a, ctl, fp)
			y := clGet(&b, ctl, fp)
			z := clGet(&d, ctl, fp)
			w := clGet(&e, ctl, fp)
			if *f == scmHeadBytecodeFunc {
				bf := mustBytecodeFunc(f)
				if cl := bf.cl; cl != nil && cl.arity == 4 {
					if cl.direct && ctl.stepLimit == 0 {
						var fr [8]Obj
						fr[0] = x
						fr[1] = y
						fr[2] = z
						fr[3] = w
						r := cl.body(ctl, clNoescape(unsafe.Pointer(&fr)))
						if r == nil {
							return trampoline(ctl)
						}
						return r
					}
					r := cl.enter(ctl, bf, cl, x, y, z, w)
					if r == nil {
						return trampoline(ctl)
					}
					return r
				}
			}
			return Call(ctl, f, x, y, z, w)
		})
	}
	ops := clOps(args)
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		f := s.function
		if f == nil {
			clUnbound(s)
		}
		return clCallMany(ctl, f, ops, fp, tail)
	})
}

func clOps(args []clx) []clOperand {
	ops := make([]clOperand, len(args))
	for i, a := range args {
		ops[i] = a.op()
	}
	return ops
}

// clCallMany is any call with five or more arguments, or a rare shape.
func clCallMany(ctl *ControlFlow, f Obj, ops []clOperand, fp unsafe.Pointer, tail bool) Obj {
	var inline [8]Obj
	args := inline[:0]
	for i := range ops {
		args = append(args, clGet(&ops[i], ctl, fp))
	}
	if tail {
		return clTailN(ctl, f, args)
	}
	return Call(ctl, f, args...)
}

func clDo0(ctl *ControlFlow, f Obj) Obj {
	if *f == scmHeadBytecodeFunc {
		bf := mustBytecodeFunc(f)
		if cl := bf.cl; cl != nil && cl.arity == 0 {
			return clFinish(ctl, cl.enter(ctl, bf, cl, nil, nil, nil, nil))
		}
	}
	return Call(ctl, f)
}

func clDo1(ctl *ControlFlow, f, x Obj) Obj {
	if *f == scmHeadBytecodeFunc {
		bf := mustBytecodeFunc(f)
		if cl := bf.cl; cl != nil && cl.arity == 1 {
			if cl.direct && ctl.stepLimit == 0 {
				var fr [8]Obj
				fr[0] = x
				return clFinish(ctl, cl.body(ctl, clNoescape(unsafe.Pointer(&fr))))
			}
			return clFinish(ctl, cl.enter(ctl, bf, cl, x, nil, nil, nil))
		}
	}
	return Call(ctl, f, x)
}

func clDo2(ctl *ControlFlow, f, x, y Obj) Obj {
	if *f == scmHeadBytecodeFunc {
		bf := mustBytecodeFunc(f)
		if cl := bf.cl; cl != nil && cl.arity == 2 {
			if cl.direct && ctl.stepLimit == 0 {
				var fr [8]Obj
				fr[0] = x
				fr[1] = y
				return clFinish(ctl, cl.body(ctl, clNoescape(unsafe.Pointer(&fr))))
			}
			return clFinish(ctl, cl.enter(ctl, bf, cl, x, y, nil, nil))
		}
	}
	return Call(ctl, f, x, y)
}

func (c *clCompiler) callDynamic(h clx, args []clx, tail bool) clx {
	hop := h.op()
	switch len(args) {
	case 0:
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj { return clTail0(ctl, clGet(&hop, ctl, fp)) })
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj { return clDo0(ctl, clGet(&hop, ctl, fp)) })
	case 1:
		a := args[0].op()
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				f := clGet(&hop, ctl, fp)
				return clTail1(ctl, f, clGet(&a, ctl, fp))
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			f := clGet(&hop, ctl, fp)
			return clDo1(ctl, f, clGet(&a, ctl, fp))
		})
	case 2:
		a, b := args[0].op(), args[1].op()
		if tail {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				f := clGet(&hop, ctl, fp)
				x := clGet(&a, ctl, fp)
				return clTail2(ctl, f, x, clGet(&b, ctl, fp))
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			f := clGet(&hop, ctl, fp)
			x := clGet(&a, ctl, fp)
			return clDo2(ctl, f, x, clGet(&b, ctl, fp))
		})
	}
	ops := clOps(args)
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		f := clGet(&hop, ctl, fp)
		return clCallMany(ctl, f, ops, fp, tail)
	})
}
