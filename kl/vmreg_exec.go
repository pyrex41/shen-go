package kl

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// The register executor. See vmreg_lower.go for the instruction set and how it
// comes from the stack bytecode.
//
// One Go activation of vmRun runs any number of bytecode activations: a call
// from bytecode to bytecode pushes a vmFrame and continues the same loop, so it
// costs neither a Go call nor Go stack, and non-tail recursion depth is limited
// by the arena rather than the goroutine stack. Natives and everything else
// still go through ctl.data and the trampoline; re-entering bytecode from them
// nests a new vmRun above the current frame.
//
// Register windows overlap: a call's callee and arguments are the caller's top
// temporaries, and the callee's register 0 is the first argument, so arguments
// are not copied. No frame is cleared on entry or exit. A register is always
// written before it is read, so stale slots above the top are invisible to Shen
// code; vmAttend clears them once per GC cycle, so a dead value is retained for
// at most one more cycle.
//
// Registers, constants and code are reached by unsafe pointer arithmetic with
// byte offsets precomputed by the lowerer: the code is generated here, from a
// stack depth it tracked exactly, and frame size is checked at each activation.

type vmFrame struct {
	fn   *scmBytecodeFunc
	ip   unsafe.Pointer
	regs unsafe.Pointer
	base int
	// crossed: this call moved the arena to a later block; cur is the block
	// to go back to. frameUnwind clears abandoned records, so it is only ever
	// set by the call in progress.
	crossed bool
	cur     int
}

const instrSize = unsafe.Sizeof(Instr{})

func r(p unsafe.Pointer, off int32) *vmSlot { return (*vmSlot)(unsafe.Add(p, int(off))) }

// vmAttn is nonzero when vmAttend has something to do. It is tested at each
// activation and after each call out of bytecode, the only points at which a
// rebind or a finished GC cycle can matter to running code, so that ordinary
// operations carry no guard at all.
var vmAttn atomic.Uint32

// vmGCEpoch counts finished GC cycles.
var vmGCEpoch atomic.Uint32

type vmGCSentinel struct{ _ *int }

// vmGCFinished runs after each cycle that found the sentinel unreachable. It
// re-arms the same object rather than allocating, so it adds no allocations.
func vmGCFinished(s *vmGCSentinel) {
	vmGCEpoch.Add(1)
	vmAttn.Store(1)
	runtime.SetFinalizer(s, vmGCFinished)
}

func init() { runtime.SetFinalizer(&vmGCSentinel{}, vmGCFinished) }

// vmAttend deals with vmAttn for the frame about to continue at ip: it clears
// stale arena slots if a GC cycle has finished, and returns ip in the slow twin
// of the code if the specialization guards are broken.
//
//go:noinline
func vmAttend(ctl *ControlFlow, fn *BytecodeFunc, ip unsafe.Pointer) unsafe.Pointer {
	if e := vmGCEpoch.Load(); e != ctl.frameEpoch {
		ctl.frameEpoch = e
		ctl.frameClearStale()
	}
	if vmGuardsOK {
		vmAttn.Store(0)
		return ip
	}
	vmAttn.Store(1)
	rc := fn.rc
	slow := rc.slowCode()
	lo, hi := uintptr(rc.fast), uintptr(rc.fast)+uintptr(rc.n)*instrSize
	if p := uintptr(ip); p >= lo && p < hi {
		return unsafe.Add(slow, p-lo)
	}
	return ip
}

// vmExec runs one activation of a bytecode function from outside bytecode. It
// either calls ctl.Return with the result or sets up a tail call.
func vmExec(ctl *ControlFlow, bf *scmBytecodeFunc, args []Obj) {
	mark := frameMark{ctl.frameCur, ctl.frameTop, ctl.vmDepth}
	rc := bf.fn.rc
	if rc == nil {
		rc = bf.fn.lower()
	}
	base := ctl.frameTop
	if base+rc.size > ctl.frameCap {
		base = ctl.frameEnsure(base, rc.size, nil)
	}
	blk := ctl.frameBlocks[ctl.frameCur]
	ctl.frameTop = base + rc.size
	regs := unsafe.Pointer(&blk[base])
	for i, a := range args {
		s := (*vmSlot)(unsafe.Add(regs, uintptr(i)*16))
		s.obj = a
	}
	v, tail := vmRun(ctl, bf, regs, base)
	if ctl.frameCur != mark.cur {
		ctl.frameLeaveBlock(mark.cur)
	}
	ctl.frameTop = mark.top
	if !tail {
		ctl.Return(v.objValue())
	}
}

func vmRun(ctl *ControlFlow, cur *scmBytecodeFunc, regs unsafe.Pointer, base int) (result vmSlot, tail bool) {
	rc := cur.fn.rc
	ip, ks := rc.fast, rc.k
	baseDepth := ctl.vmDepth

	// The step budget is charged once per activation and once per
	// self-tail-call iteration, not per instruction: bytecode only jumps
	// forward, so every unbounded execution passes through one of those two
	// points. In every non-fuzz process limited is false.
	limited := ctl.stepLimited()
	if limited {
		ctl.tick()
	}
	if vmAttn.Load() != 0 {
		ip = vmAttend(ctl, cur.fn, ip)
	}

	for {
		in := (*Instr)(ip)
		ip = unsafe.Add(ip, instrSize)
		switch in.Op {

		case OP_R_MOV:
			*r(regs, in.A) = *r(regs, in.B)

		case OP_R_LOADK:
			*r(regs, in.A) = *r(ks, in.B)

		case OP_R_GETUP:
			r(regs, in.A).obj = cur.upvals[in.B]

		case OP_R_LOADG:
			sym := (*scmSymbol)(unsafe.Pointer(r(ks, in.B).obj))
			f := sym.function
			if f == nil {
				vmUnbound(sym)
			}
			r(regs, in.A).obj = f

		case OP_R_JMP:
			ip = unsafe.Add(ip, int(in.A))

		case OP_R_JF:
			o := r(regs, in.B).obj
			if o == False {
				ip = unsafe.Add(ip, int(in.A))
			} else if o != True {
				vmBadIf()
			}

		case OP_R_JT:
			o := r(regs, in.B).obj
			if o == True {
				ip = unsafe.Add(ip, int(in.A))
			} else if o != False {
				goto slow // not's error, not if's
			}

		case OP_R_NOT:
			o := r(regs, in.B).obj
			if o == True {
				r(regs, in.A).obj = False
			} else if o == False {
				r(regs, in.A).obj = True
			} else {
				goto slow
			}

		case OP_R_JNEKI:
			if r(regs, in.B).obj != r(ks, in.C).obj {
				ip = unsafe.Add(ip, int(in.A))
			}

		case OP_R_JNEKN:
			x := r(regs, in.B)
			if x.obj != nil {
				goto slow
			}
			if x.number != r(ks, in.C).number {
				ip = unsafe.Add(ip, int(in.A))
			}

		case OP_R_JNE:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if x.number != y.number {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			if x.obj != nil && x.obj == y.obj {
				continue
			}
			goto slow

		case OP_R_JNCONSP:
			o := r(regs, in.B).obj
			if o == nil || *o != scmHeadPair {
				ip = unsafe.Add(ip, int(in.A))
			}

		case OP_R_HD:
			o := r(regs, in.B).obj
			if o == nil || *o != scmHeadPair {
				goto slow
			}
			r(regs, in.A).obj = (*scmPair)(unsafe.Pointer(o)).car

		case OP_R_TL:
			o := r(regs, in.B).obj
			if o == nil || *o != scmHeadPair {
				goto slow
			}
			r(regs, in.A).obj = (*scmPair)(unsafe.Pointer(o)).cdr

		case OP_R_CONSP:
			o := r(regs, in.B).obj
			if o != nil && *o == scmHeadPair {
				r(regs, in.A).obj = True
			} else {
				r(regs, in.A).obj = False
			}

		case OP_R_CONS:
			c := cons(r(regs, in.B).objValue(), r(regs, in.C).objValue())
			r(regs, in.A).obj = c

		case OP_R_VGET:
			o, i := r(regs, in.B).obj, r(regs, in.C)
			if o == nil || *o != scmHeadVector || i.obj != nil {
				goto slow
			}
			vec := (*scmVector)(unsafe.Pointer(o)).vector
			n := int(i.number)
			if float64(n) != i.number || uint(n) >= uint(len(vec)) || vec[n] == nil {
				goto slow
			}
			r(regs, in.A).obj = vec[n]

		case OP_R_VSET:
			o, i := r(regs, in.B).obj, r(regs, in.C)
			if o == nil || *o != scmHeadVector || i.obj != nil {
				goto slow
			}
			vec := (*scmVector)(unsafe.Pointer(o)).vector
			n := int(i.number)
			if float64(n) != i.number || uint(n) >= uint(len(vec)) {
				goto slow
			}
			vec[n] = r(regs, in.D).objValue()
			r(regs, in.A).obj = o

		case OP_R_GCONST:
			*r(regs, in.A) = *r(ks, in.C)

		case OP_R_SLOW, OP_R_PRIM, OP_R_DIV, OP_R_LT, OP_R_LE, OP_R_GT, OP_R_GE, OP_R_EQ:
			goto slow

		case OP_R_CLOSURE:
			b := r(regs, in.A)
			*b = vmMakeClosure(cur.fn.Consts[in.B], unsafe.Slice(b, int(in.C)))

		case OP_R_RET:
			result = *r(regs, in.A)
			goto ret

		case OP_R_RETK:
			result = *r(ks, in.A)
			goto ret

		case OP_R_SELF:
			ip = unsafe.Add(ip, int(in.A))
			if limited {
				ctl.tick()
			}

		case OP_R_CALLG:
			sym := (*scmSymbol)(unsafe.Pointer(r(ks, in.C).obj))
			f := sym.function
			if f == nil {
				vmUnbound(sym)
			}
			r(regs, in.A).obj = f
			fallthrough

		case OP_R_CALL:
			cs := r(regs, in.A)
			callee := cs.obj
			if callee != nil && *callee == scmHeadBytecodeFunc {
				bf := (*scmBytecodeFunc)(unsafe.Pointer(callee))
				f2 := bf.fn
				if f2.Arity == int(in.B) {
					rc2 := f2.rc
					if rc2 == nil {
						rc2 = f2.lower()
					}
					regs2 := unsafe.Add(unsafe.Pointer(cs), 16)
					base2 := base + int(in.A)>>4 + 1
					// Save the caller before frameEnsure can change blocks.
					d := ctl.vmDepth
					if d >= len(ctl.vmFrames) {
						ctl.vmFramesGrow()
					}
					fr := &ctl.vmFrames[d]
					fr.fn, fr.ip, fr.regs, fr.base = cur, ip, regs, base
					ctl.vmDepth = d + 1
					if base2+rc2.size > ctl.frameCap {
						was := ctl.frameCur
						base2 = ctl.frameEnsure(base2, rc2.size, unsafe.Slice((*vmSlot)(regs2), int(in.B)))
						regs2 = unsafe.Pointer(&ctl.frameBlocks[ctl.frameCur][base2])
						if ctl.frameCur != was {
							fr.crossed, fr.cur = true, was
						}
					}
					ctl.frameTop = base2 + rc2.size
					cur, regs, base = bf, regs2, base2
					ip, ks = rc2.fast, rc2.k
					if limited {
						ctl.tick()
					}
					if vmAttn.Load() != 0 {
						ip = vmAttend(ctl, f2, ip)
					}
					continue
				}
			}
			{
				res := vmCallOther(ctl, cs.objValue(), unsafe.Slice((*vmSlot)(unsafe.Add(unsafe.Pointer(cs), 16)), int(in.B)))
				*cs = res
				if vmAttn.Load() != 0 {
					ip = vmAttend(ctl, cur.fn, ip)
				}
			}

		case OP_R_TAILCALLG:
			sym := (*scmSymbol)(unsafe.Pointer(r(ks, in.C).obj))
			f := sym.function
			if f == nil {
				vmUnbound(sym)
			}
			r(regs, in.A).obj = f
			fallthrough

		case OP_R_TAILCALL:
			cs := r(regs, in.A)
			callee := cs.obj
			n := int(in.B)
			if callee != nil && *callee == scmHeadBytecodeFunc {
				bf := (*scmBytecodeFunc)(unsafe.Pointer(callee))
				f2 := bf.fn
				if f2.Arity == n {
					// Reuse this frame: move the arguments down. A tail call to
					// the same function was already a loop; a tail call to
					// another used to bounce off the trampoline, and is charged
					// the same two steps as that bounce.
					rc2 := f2.rc
					if rc2 == nil {
						rc2 = f2.lower()
					}
					src := unsafe.Pointer(cs)
					for i := 0; i < n; i++ {
						src = unsafe.Add(src, 16)
						*(*vmSlot)(unsafe.Add(regs, uintptr(i)*16)) = *(*vmSlot)(src)
					}
					if base+rc2.size > ctl.frameCap {
						was := ctl.frameCur
						base = ctl.frameEnsure(base, rc2.size, unsafe.Slice((*vmSlot)(regs), n))
						regs = unsafe.Pointer(&ctl.frameBlocks[ctl.frameCur][base])
						if d := ctl.vmDepth; ctl.frameCur != was && d > baseDepth {
							if fr := &ctl.vmFrames[d-1]; !fr.crossed {
								fr.crossed, fr.cur = true, was
							}
						}
					}
					ctl.frameTop = base + rc2.size
					if limited {
						ctl.tick()
						if bf != cur {
							ctl.tick()
						}
					}
					cur = bf
					ip, ks = rc2.fast, rc2.k
					if vmAttn.Load() != 0 {
						ip = vmAttend(ctl, f2, ip)
					}
					continue
				}
			}
			ctl.tailApplySlots(cs.objValue(), unsafe.Slice((*vmSlot)(unsafe.Add(unsafe.Pointer(cs), 16)), n))
			if ctl.vmDepth == baseDepth {
				return vmSlot{}, true
			}
			// The trampoline runs after this frame is gone, so that a chain of
			// tail calls through natives cannot grow the arena.
			{
				d := ctl.vmDepth - 1
				fr := &ctl.vmFrames[d]
				cur, ip, regs, base = fr.fn, fr.ip, fr.regs, fr.base
				rc2 := cur.fn.rc
				ks = rc2.k
				ctl.vmDepth = d
				if fr.crossed {
					fr.crossed = false
					ctl.frameLeaveBlock(fr.cur)
				}
				ctl.frameTop = base + rc2.size
				res := vmSlot{obj: trampoline(ctl)}
				*r(regs, (*Instr)(unsafe.Add(ip, -int(instrSize))).A) = res
				if vmAttn.Load() != 0 {
					ip = vmAttend(ctl, cur.fn, ip)
				}
			}

		case OP_R_ADD:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if f := x.number + y.number; f-f == 0 {
					d := r(regs, in.A)
					d.number = f
					if d.obj != nil {
						d.obj = nil
					}
					continue
				}
			}
			goto slow

		case OP_R_ADDK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if f := x.number + y.number; f-f == 0 {
					d := r(regs, in.A)
					d.number = f
					if d.obj != nil {
						d.obj = nil
					}
					continue
				}
			}
			goto slow

		case OP_R_SUB:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if f := x.number - y.number; f-f == 0 {
					d := r(regs, in.A)
					d.number = f
					if d.obj != nil {
						d.obj = nil
					}
					continue
				}
			}
			goto slow

		case OP_R_SUBK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if f := x.number - y.number; f-f == 0 {
					d := r(regs, in.A)
					d.number = f
					if d.obj != nil {
						d.obj = nil
					}
					continue
				}
			}
			goto slow

		case OP_R_MUL:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if f := x.number * y.number; f-f == 0 {
					d := r(regs, in.A)
					d.number = f
					if d.obj != nil {
						d.obj = nil
					}
					continue
				}
			}
			goto slow

		case OP_R_MULK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if f := x.number * y.number; f-f == 0 {
					d := r(regs, in.A)
					d.number = f
					if d.obj != nil {
						d.obj = nil
					}
					continue
				}
			}
			goto slow

		case OP_R_JNLT:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number < y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNLTK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number < y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNLE:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number <= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNLEK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number <= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNGT:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number > y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNGTK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number > y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNGE:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number >= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		case OP_R_JNGEK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number >= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto slow

		default:
			panic("vmRun: unknown opcode")
		}
		continue

	ret:
		{
			d := ctl.vmDepth
			if d == baseDepth {
				return result, false
			}
			d--
			fr := &ctl.vmFrames[d]
			cur, ip, regs, base = fr.fn, fr.ip, fr.regs, fr.base
			rc2 := cur.fn.rc
			ks = rc2.k
			ctl.vmDepth = d
			if fr.crossed {
				fr.crossed = false
				ctl.frameLeaveBlock(fr.cur)
			}
			ctl.frameTop = base + rc2.size
			*r(regs, (*Instr)(unsafe.Add(ip, -int(instrSize))).A) = result
			if vmAttn.Load() != 0 {
				ip = vmAttend(ctl, cur.fn, ip)
			}
			continue
		}

	slow:
		if vmSlowOp(ctl, cur, regs, ks, in) {
			ip = unsafe.Add(ip, int(in.A))
		}
		if vmAttn.Load() != 0 {
			ip = vmAttend(ctl, cur.fn, ip)
		}
	}
}
