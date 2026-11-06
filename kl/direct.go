package kl

import (
	"fmt"
	"unsafe"
)

// Direct native calling convention.
//
// The original native convention passes arguments and the result through
// ctl.data: the caller appends callee and arguments (write barrier per word
// while the GC marks), the callee Gets each with a bounds check, Return appends
// the result, the caller reads it back. A direct native takes its arguments
// as Go parameters and returns its result:
//
//	func(e *ControlFlow, a, b Obj) Obj
//
// A nil result means "a tail call is pending in ctl.data": the body ended with
// `return TailCallN(...)`, and the nearest non-tail caller runs the trampoline.
// nil is not an Obj value any KL program can see, so it is free as the marker.
//
// Rebinding stays correct because nothing caches a callee: a generated call
// loads sym.function at each call (SymFn), so a redefinition takes effect at
// the next call, in the same load the old PrimFunc made.
//
// Every direct native also has the ordinary fn entry (directEntry[N]), so
// apply, Call, the VM and anything else that runs nf.fn(ctl) works unchanged.
// That entry is a shared static function, not a closure per native: it finds
// the native at ctl.data[ctl.pos], where every path into nf.fn has put it, so
// making a lambda costs no extra allocation.

// SymFn is PrimFunc for a sym already known to be a symbol, which generated
// code guarantees. It raises the same error and is small enough to inline.
func SymFn(sym Obj) Obj {
	f := (*scmSymbol)(unsafe.Pointer(sym)).function
	if f == nil {
		symUnbound(sym)
	}
	return f
}

func symUnbound(sym Obj) {
	panic(MakeError(fmt.Sprintf("variable %s not bound", mustSymbol(sym).str)))
}

// ReturnDirect finishes an old-convention native whose body uses the direct
// convention's result (generated code for arities above the direct maximum).
func (ctl *ControlFlow) ReturnDirect(r Obj) {
	if r != nil {
		ctl.Return(r)
	}
}

// directResult finishes a direct call: r, or the pending tail call's result.
func directResult(e *ControlFlow, r Obj) Obj {
	if r != nil {
		return r
	}
	return trampoline(e)
}

// exitBytecode finishes a vmExec run started by a direct call.
func exitBytecode(e *ControlFlow) Obj {
	if e.kind == ControlFlowReturn {
		ret := e.data[e.pos]
		e.data = e.data[:e.pos]
		return ret
	}
	return trampoline(e)
}

// bytecodeCallee returns f as a bytecode function of exactly this arity.
func bytecodeCallee(f Obj, n int) *scmBytecodeFunc {
	if f != nil && !isFixnum(f) && *f == scmHeadBytecodeFunc {
		if bf := mustBytecodeFunc(f); bf.fn.Arity == n {
			return bf
		}
	}
	return nil
}

// Direct0 is the direct convention at arity 0.
type Direct0 = func(*ControlFlow) Obj

// MakeNative0 makes a native with a direct entry of arity 0.
func MakeNative0(f Direct0) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry0,
		require: 0,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry0(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct0)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e))
}

// Call0 is Call at arity 0: same semantics, including the step-limit charge.
func Call0(e *ControlFlow, f Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 0 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct0)(unsafe.Pointer(&nf.direct)))(e)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call0Slow(e, f)
}

func call0Slow(e *ControlFlow, f Obj) Obj {
	if bf := bytecodeCallee(f, 0); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [0]Obj{}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f)
}

// TailCall0 records a tail call and returns the marker for a direct body to
// return.
func TailCall0(e *ControlFlow, f Obj) Obj {
	e.data = append(e.data[:e.pos], f)
	e.kind = ControlFlowApply
	return nil
}

// Direct1 is the direct convention at arity 1.
type Direct1 = func(*ControlFlow, Obj) Obj

// MakeNative1 makes a native with a direct entry of arity 1.
func MakeNative1(f Direct1) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry1,
		require: 1,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry1(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct1)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e, e.data[e.pos+1]))
}

// Call1 is Call at arity 1: same semantics, including the step-limit charge.
func Call1(e *ControlFlow, f, a Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 1 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct1)(unsafe.Pointer(&nf.direct)))(e, a)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call1Slow(e, f, a)
}

func call1Slow(e *ControlFlow, f, a Obj) Obj {
	if bf := bytecodeCallee(f, 1); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [1]Obj{a}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f, a)
}

// TailCall1 records a tail call and returns the marker for a direct body to
// return.
func TailCall1(e *ControlFlow, f, a Obj) Obj {
	e.data = append(e.data[:e.pos], f, a)
	e.kind = ControlFlowApply
	return nil
}

// Direct2 is the direct convention at arity 2.
type Direct2 = func(*ControlFlow, Obj, Obj) Obj

// MakeNative2 makes a native with a direct entry of arity 2.
func MakeNative2(f Direct2) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry2,
		require: 2,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry2(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct2)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e, e.data[e.pos+1], e.data[e.pos+2]))
}

// Call2 is Call at arity 2: same semantics, including the step-limit charge.
func Call2(e *ControlFlow, f, a, b Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 2 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct2)(unsafe.Pointer(&nf.direct)))(e, a, b)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call2Slow(e, f, a, b)
}

func call2Slow(e *ControlFlow, f, a, b Obj) Obj {
	if bf := bytecodeCallee(f, 2); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [2]Obj{a, b}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f, a, b)
}

// TailCall2 records a tail call and returns the marker for a direct body to
// return.
func TailCall2(e *ControlFlow, f, a, b Obj) Obj {
	e.data = append(e.data[:e.pos], f, a, b)
	e.kind = ControlFlowApply
	return nil
}

// Direct3 is the direct convention at arity 3.
type Direct3 = func(*ControlFlow, Obj, Obj, Obj) Obj

// MakeNative3 makes a native with a direct entry of arity 3.
func MakeNative3(f Direct3) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry3,
		require: 3,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry3(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct3)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e, e.data[e.pos+1], e.data[e.pos+2], e.data[e.pos+3]))
}

// Call3 is Call at arity 3: same semantics, including the step-limit charge.
func Call3(e *ControlFlow, f, a, b, c Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 3 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct3)(unsafe.Pointer(&nf.direct)))(e, a, b, c)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call3Slow(e, f, a, b, c)
}

func call3Slow(e *ControlFlow, f, a, b, c Obj) Obj {
	if bf := bytecodeCallee(f, 3); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [3]Obj{a, b, c}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f, a, b, c)
}

// TailCall3 records a tail call and returns the marker for a direct body to
// return.
func TailCall3(e *ControlFlow, f, a, b, c Obj) Obj {
	e.data = append(e.data[:e.pos], f, a, b, c)
	e.kind = ControlFlowApply
	return nil
}

// Direct4 is the direct convention at arity 4.
type Direct4 = func(*ControlFlow, Obj, Obj, Obj, Obj) Obj

// MakeNative4 makes a native with a direct entry of arity 4.
func MakeNative4(f Direct4) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry4,
		require: 4,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry4(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct4)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e, e.data[e.pos+1], e.data[e.pos+2], e.data[e.pos+3], e.data[e.pos+4]))
}

// Call4 is Call at arity 4: same semantics, including the step-limit charge.
func Call4(e *ControlFlow, f, a, b, c, d Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 4 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct4)(unsafe.Pointer(&nf.direct)))(e, a, b, c, d)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call4Slow(e, f, a, b, c, d)
}

func call4Slow(e *ControlFlow, f, a, b, c, d Obj) Obj {
	if bf := bytecodeCallee(f, 4); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [4]Obj{a, b, c, d}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f, a, b, c, d)
}

// TailCall4 records a tail call and returns the marker for a direct body to
// return.
func TailCall4(e *ControlFlow, f, a, b, c, d Obj) Obj {
	e.data = append(e.data[:e.pos], f, a, b, c, d)
	e.kind = ControlFlowApply
	return nil
}

// Direct5 is the direct convention at arity 5.
type Direct5 = func(*ControlFlow, Obj, Obj, Obj, Obj, Obj) Obj

// MakeNative5 makes a native with a direct entry of arity 5.
func MakeNative5(f Direct5) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry5,
		require: 5,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry5(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct5)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e, e.data[e.pos+1], e.data[e.pos+2], e.data[e.pos+3], e.data[e.pos+4], e.data[e.pos+5]))
}

// Call5 is Call at arity 5: same semantics, including the step-limit charge.
func Call5(e *ControlFlow, f, a, b, c, d, g Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 5 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct5)(unsafe.Pointer(&nf.direct)))(e, a, b, c, d, g)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call5Slow(e, f, a, b, c, d, g)
}

func call5Slow(e *ControlFlow, f, a, b, c, d, g Obj) Obj {
	if bf := bytecodeCallee(f, 5); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [5]Obj{a, b, c, d, g}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f, a, b, c, d, g)
}

// TailCall5 records a tail call and returns the marker for a direct body to
// return.
func TailCall5(e *ControlFlow, f, a, b, c, d, g Obj) Obj {
	e.data = append(e.data[:e.pos], f, a, b, c, d, g)
	e.kind = ControlFlowApply
	return nil
}

// Direct6 is the direct convention at arity 6.
type Direct6 = func(*ControlFlow, Obj, Obj, Obj, Obj, Obj, Obj) Obj

// MakeNative6 makes a native with a direct entry of arity 6.
func MakeNative6(f Direct6) Obj {
	tmp := scmNative{
		scmHead: scmHeadNative,
		fn:      directEntry6,
		require: 6,
		direct:  *(*unsafe.Pointer)(unsafe.Pointer(&f)),
	}
	return &tmp.scmHead
}

func directEntry6(e *ControlFlow) {
	nf := MustNative(e.data[e.pos])
	d := *(*Direct6)(unsafe.Pointer(&nf.direct))
	e.ReturnDirect(d(e, e.data[e.pos+1], e.data[e.pos+2], e.data[e.pos+3], e.data[e.pos+4], e.data[e.pos+5], e.data[e.pos+6]))
}

// Call6 is Call at arity 6: same semantics, including the step-limit charge.
func Call6(e *ControlFlow, f, a, b, c, d, g, h Obj) Obj {
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); nf.require == 6 && nf.direct != nil {
			if e.stepLimit != 0 {
				e.tick()
			}
			r := (*(*Direct6)(unsafe.Pointer(&nf.direct)))(e, a, b, c, d, g, h)
			if r != nil {
				return r
			}
			return trampoline(e)
		}
	}
	return call6Slow(e, f, a, b, c, d, g, h)
}

func call6Slow(e *ControlFlow, f, a, b, c, d, g, h Obj) Obj {
	if bf := bytecodeCallee(f, 6); bf != nil {
		if e.stepLimit != 0 {
			e.tick()
		}
		args := [6]Obj{a, b, c, d, g, h}
		vmExec(e, bf, args[:])
		return exitBytecode(e)
	}
	return Call(e, f, a, b, c, d, g, h)
}

// TailCall6 records a tail call and returns the marker for a direct body to
// return.
func TailCall6(e *ControlFlow, f, a, b, c, d, g, h Obj) Obj {
	e.data = append(e.data[:e.pos], f, a, b, c, d, g, h)
	e.kind = ControlFlowApply
	return nil
}

// callDirectSlots runs an exact-arity direct native from the VM's operand
// stack without going through ctl.data. nf.direct is non-nil and
// nf.require == len(args). It charges no step, like the VM's call of a native.
func callDirectSlots(ctl *ControlFlow, nf *scmNative, args []vmSlot) Obj {
	var r Obj
	switch len(args) {
	case 0:
		r = (*(*Direct0)(unsafe.Pointer(&nf.direct)))(ctl)
	case 1:
		r = (*(*Direct1)(unsafe.Pointer(&nf.direct)))(ctl, args[0].objValue())
	case 2:
		r = (*(*Direct2)(unsafe.Pointer(&nf.direct)))(ctl, args[0].objValue(), args[1].objValue())
	case 3:
		r = (*(*Direct3)(unsafe.Pointer(&nf.direct)))(ctl, args[0].objValue(), args[1].objValue(), args[2].objValue())
	case 4:
		r = (*(*Direct4)(unsafe.Pointer(&nf.direct)))(ctl, args[0].objValue(), args[1].objValue(), args[2].objValue(), args[3].objValue())
	case 5:
		r = (*(*Direct5)(unsafe.Pointer(&nf.direct)))(ctl, args[0].objValue(), args[1].objValue(), args[2].objValue(), args[3].objValue(), args[4].objValue())
	case 6:
		r = (*(*Direct6)(unsafe.Pointer(&nf.direct)))(ctl, args[0].objValue(), args[1].objValue(), args[2].objValue(), args[3].objValue(), args[4].objValue(), args[5].objValue())
	default:
		panic("callDirectSlots: arity")
	}
	if r != nil {
		return r
	}
	return trampoline(ctl)
}
