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

// vmState is vmRun's register state, in memory during a call out of the loop.
// It is the point of the arrangement below: Go spills a value live across a
// call at a point dominating every use after a call, which for the loop's state
// is the dispatch, on every instruction, whichever path runs. So no state value
// in vmRun is live across a call. Every cold path saves it here, calls, and
// loads it back. A struct of five fields is not split into SSA values, so these
// really are loads and stores on those paths and nowhere else.
type vmState struct {
	cur      *scmBytecodeFunc
	ip       unsafe.Pointer
	regs, ks unsafe.Pointer
	base     int
}

func vmRun(ctl *ControlFlow, cur *scmBytecodeFunc, regs unsafe.Pointer, base int) (vmSlot, bool) {
	rc := cur.fn.rc
	ip, ks := rc.fast, rc.k
	baseDepth := ctl.vmDepth
	var sv vmState

	// The step budget is charged once per activation and once per
	// self-tail-call iteration, not per instruction: bytecode only jumps
	// forward, so every unbounded execution passes through one of those two
	// points. In every non-fuzz process limited is false.
	limited := ctl.stepLimited()
	if limited {
		sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
		ctl.tick()
		cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
	}
	if vmAttn.Load() != 0 {
		ip = vmAttend(ctl, cur.fn, ip)
	}

	for {
		in := (*Instr)(ip)
		ip = unsafe.Add(ip, instrSize)
		switch in.Op {

		case opRMOV:
			*r(regs, in.A) = *r(regs, in.B)

		case opRLOADK:
			*r(regs, in.A) = *r(ks, in.B)

		case opRGETUP:
			r(regs, in.A).obj = cur.upvals[in.B]

		case opRLOADG:
			sym := (*scmSymbol)(unsafe.Pointer(r(ks, in.B).obj))
			f := sym.function
			if f == nil {
				panic(vmUnbound(sym))
			}
			r(regs, in.A).obj = f

		case opRJMP:
			ip = unsafe.Add(ip, int(in.A))

		case opRJF:
			o := r(regs, in.B).obj
			if o == False {
				ip = unsafe.Add(ip, int(in.A))
			} else if o != True {
				panic(vmBadIf())
			}

		case opRJT:
			o := r(regs, in.B).obj
			if o == True {
				ip = unsafe.Add(ip, int(in.A))
			} else if o != False {
				goto slow // not's error, not if's
			}

		case opRNOT:
			o := r(regs, in.B).obj
			if o == True {
				r(regs, in.A).obj = False
			} else if o == False {
				r(regs, in.A).obj = True
			} else {
				goto slow
			}

		case opRJNEKI:
			if r(regs, in.B).obj != r(ks, in.C).obj {
				ip = unsafe.Add(ip, int(in.A))
			}

		case opRJNEKN:
			x := r(regs, in.B)
			if x.obj != nil {
				goto fixK
			}
			if x.number != r(ks, in.C).number {
				ip = unsafe.Add(ip, int(in.A))
			}

		case opRJNE:
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
			goto fixR

		case opRJNCONSP:
			o := r(regs, in.B).obj
			if o == nil || *o != scmHeadPair {
				ip = unsafe.Add(ip, int(in.A))
			}

		case opRHD:
			o := r(regs, in.B).obj
			if o == nil || *o != scmHeadPair {
				goto slow
			}
			r(regs, in.A).obj = (*scmPair)(unsafe.Pointer(o)).car

		case opRTL:
			o := r(regs, in.B).obj
			if o == nil || *o != scmHeadPair {
				goto slow
			}
			r(regs, in.A).obj = (*scmPair)(unsafe.Pointer(o)).cdr

		case opRCONSP:
			o := r(regs, in.B).obj
			if o != nil && *o == scmHeadPair {
				r(regs, in.A).obj = True
			} else {
				r(regs, in.A).obj = False
			}

		case opRCONS:
			sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
			vmConsOp(r(regs, in.A), r(regs, in.B), r(regs, in.C))
			cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base

		case opRVGET:
			o, i := r(regs, in.B).obj, r(regs, in.C)
			if o == nil || *o != scmHeadVector {
				goto slow
			}
			if i.obj != nil {
				goto fixR
			}
			vec := (*scmVector)(unsafe.Pointer(o)).vector
			n := int(i.number)
			if float64(n) != i.number || uint(n) >= uint(len(vec)) || vec[n] == nil {
				goto slow
			}
			r(regs, in.A).obj = vec[n]

		case opRVSET:
			o, i := r(regs, in.B).obj, r(regs, in.C)
			if o == nil || *o != scmHeadVector {
				goto slow
			}
			if i.obj != nil {
				goto fixR
			}
			vec := (*scmVector)(unsafe.Pointer(o)).vector
			n := int(i.number)
			if float64(n) != i.number || uint(n) >= uint(len(vec)) {
				goto slow
			}
			v := r(regs, in.D).obj
			if v != nil {
				vec[n] = v
				r(regs, in.A).obj = o
				continue
			}
			sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
			vmVsetNumber(&vec[n], r(regs, in.D).number)
			r(sv.regs, in.A).obj = o
			cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base

		case opRGCONST:
			*r(regs, in.A) = *r(ks, in.C)

		case opRSLOW, opRPRIM, opRDIV, opRLT, opRLE, opRGT, opRGE, opREQ:
			goto slow

		case opRCLOSURE:
			sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
			vmClosureOp(cur.fn, r(regs, in.A), in)
			cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base

		case opRRET:
			d := ctl.vmDepth
			if d == baseDepth {
				return *r(regs, in.A), false
			}
			v := r(regs, in.A)
			d--
			ctl.vmDepth = d
			fr := &ctl.vmFrames[d]
			cur, ip, base = fr.fn, fr.ip, fr.base
			rc2 := cur.fn.rc
			ks = rc2.k
			ctl.frameTop = base + rc2.size
			*r(fr.regs, (*Instr)(unsafe.Add(ip, -int(instrSize))).A) = *v
			regs = fr.regs
			if fr.crossed {
				goto crossback
			}
			if vmAttn.Load() != 0 {
				goto attend
			}

		case opRRETK:
			d := ctl.vmDepth
			if d == baseDepth {
				return *r(ks, in.A), false
			}
			v := r(ks, in.A)
			d--
			ctl.vmDepth = d
			fr := &ctl.vmFrames[d]
			cur, ip, base = fr.fn, fr.ip, fr.base
			rc2 := cur.fn.rc
			ks = rc2.k
			ctl.frameTop = base + rc2.size
			*r(fr.regs, (*Instr)(unsafe.Add(ip, -int(instrSize))).A) = *v
			regs = fr.regs
			if fr.crossed {
				goto crossback
			}
			if vmAttn.Load() != 0 {
				goto attend
			}

		case opRSELF:
			ip = unsafe.Add(ip, int(in.A))
			if limited {
				sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
				ctl.tick()
				cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
			}

		case opRCALLG:
			sym := (*scmSymbol)(unsafe.Pointer(r(ks, in.C).obj))
			f := sym.function
			if f == nil {
				panic(vmUnbound(sym))
			}
			r(regs, in.A).obj = f
			fallthrough

		case opRCALL:
			cs := r(regs, in.A)
			callee := cs.obj
			if callee != nil && *callee == scmHeadBytecodeFunc {
				bf := (*scmBytecodeFunc)(unsafe.Pointer(callee))
				f2 := bf.fn
				rc2 := f2.rc
				d := ctl.vmDepth
				base2 := base + int(in.A)>>4 + 1
				if f2.Arity == int(in.B) {
					if rc2 == nil || d >= len(ctl.vmFrames) || base2+rc2.size > ctl.frameCap {
						// Everything these paths need is recomputed afterwards.
						sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
						vmCallCold(ctl, &sv, in, bf)
						cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
						if limited {
							sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
							ctl.tick()
							cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
						}
						if vmAttn.Load() != 0 {
							goto attend
						}
						continue
					}
					fr := &ctl.vmFrames[d]
					fr.fn, fr.ip, fr.regs, fr.base = cur, ip, regs, base
					ctl.vmDepth = d + 1
					ctl.frameTop = base2 + rc2.size
					cur, regs, base = bf, unsafe.Add(unsafe.Pointer(cs), 16), base2
					ip, ks = rc2.fast, rc2.k
					if limited {
						sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
						ctl.tick()
						cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
					}
					if vmAttn.Load() != 0 {
						goto attend
					}
					continue
				}
			}
			sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
			vmCallOut(ctl, cs, int(in.B))
			cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
			if vmAttn.Load() != 0 {
				goto attend
			}

		case opRTAILCALLG:
			sym := (*scmSymbol)(unsafe.Pointer(r(ks, in.C).obj))
			f := sym.function
			if f == nil {
				panic(vmUnbound(sym))
			}
			r(regs, in.A).obj = f
			fallthrough

		case opRTAILCALL:
			cs := r(regs, in.A)
			callee := cs.obj
			n := int(in.B)
			if callee != nil && *callee == scmHeadBytecodeFunc {
				bf := (*scmBytecodeFunc)(unsafe.Pointer(callee))
				f2 := bf.fn
				rc2 := f2.rc
				if f2.Arity == n {
					if rc2 == nil || base+rc2.size > ctl.frameCap {
						sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
						rc2 = vmTailCold(ctl, &sv, baseDepth, f2)
						cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
					}
					// Reuse this frame: move the arguments down. A tail call to
					// the same function was already a loop; a tail call to
					// another used to bounce off the trampoline, and is charged
					// the same two steps as that bounce.
					from := unsafe.Pointer(r(regs, in.A))
					to := regs
					for i := 0; i < n; i++ {
						from = unsafe.Add(from, 16)
						*(*vmSlot)(to) = *(*vmSlot)(from)
						to = unsafe.Add(to, 16)
					}
					ctl.frameTop = base + rc2.size
					if limited {
						self := bf == cur
						sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
						ctl.tick()
						if !self {
							ctl.tick()
						}
						cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
					}
					cur = bf
					ip, ks = rc2.fast, rc2.k
					if vmAttn.Load() != 0 {
						goto attend
					}
					continue
				}
			}
			sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
			if vmTailOut(ctl, &sv, cs, n, baseDepth) {
				return vmSlot{}, true
			}
			cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
			if vmAttn.Load() != 0 {
				goto attend
			}
		case opRADD:
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
			goto fixR

		case opRADDK:
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
			goto fixK

		case opRSUB:
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
			goto fixR

		case opRSUBK:
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
			goto fixK

		case opRMUL:
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
			goto fixR

		case opRMULK:
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
			goto fixK

		case opRJNLT:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number < y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixR

		case opRJNLTK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number < y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixK

		case opRJNLE:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number <= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixR

		case opRJNLEK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number <= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixK

		case opRJNGT:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number > y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixR

		case opRJNGTK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number > y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixK

		case opRJNGE:
			x, y := r(regs, in.B), r(regs, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number >= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixR

		case opRJNGEK:
			x, y := r(regs, in.B), r(ks, in.C)
			if x.obj == nil && y.obj == nil {
				if !(x.number >= y.number) {
					ip = unsafe.Add(ip, int(in.A))
				}
				continue
			}
			goto fixK

		default:
			panic("vmRun: unknown opcode")
		}
		continue

	fixR:
		// A fixnum operand, typically fresh from a list or a caller outside
		// bytecode. Its register takes the float form, which the operation and
		// everything that later reads the register can use inline; the
		// instruction then runs again. Anything else is for vmSlowOp.
		{
			y := r(regs, in.C)
			if o := uintptr(unsafe.Pointer(y.obj)) - uintptr(unsafe.Pointer(&addrForFixnum)); o < fixnumCount {
				y.number = float64(int(o) + fixnumMin)
				y.obj = nil
				x := r(regs, in.B)
				if o := uintptr(unsafe.Pointer(x.obj)) - uintptr(unsafe.Pointer(&addrForFixnum)); o < fixnumCount {
					x.number = float64(int(o) + fixnumMin)
					x.obj = nil
				}
				ip = unsafe.Pointer(in)
				continue
			}
		}
	fixK:
		{
			x := r(regs, in.B)
			if o := uintptr(unsafe.Pointer(x.obj)) - uintptr(unsafe.Pointer(&addrForFixnum)); o < fixnumCount {
				x.number = float64(int(o) + fixnumMin)
				x.obj = nil
				ip = unsafe.Pointer(in)
				continue
			}
		}
		goto slow

	crossback:
		sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
		ctl.frameLeaveDepth()
		cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
		if vmAttn.Load() == 0 {
			continue
		}

	attend:
		sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
		sv.ip = vmAttend(ctl, sv.cur.fn, sv.ip)
		cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
		continue

	slow:
		sv.cur, sv.ip, sv.regs, sv.ks, sv.base = cur, ip, regs, ks, base
		jump := vmSlowOp(ctl, cur, regs, ks, in)
		cur, ip, regs, ks, base = sv.cur, sv.ip, sv.regs, sv.ks, sv.base
		ip = unsafe.Add(ip, jump)
		if vmAttn.Load() != 0 {
			goto attend
		}
	}
}

//go:noinline
func vmVsetNumber(p *Obj, f float64) { *p = MakeNumber(f) }

// vmConsOp is opRCONS, out of line since it allocates anyway.
//
//go:noinline
func vmConsOp(d, x, y *vmSlot) {
	d.obj = cons(x.objValue(), y.objValue())
}

//go:noinline
func vmClosureOp(fn *BytecodeFunc, b *vmSlot, in *Instr) {
	*b = vmMakeClosure(fn.Consts[in.B], unsafe.Slice(b, int(in.C)))
}

// vmCallOut is a non-tail call to anything but an exact-arity bytecode function.
//
//go:noinline
func vmCallOut(ctl *ControlFlow, cs *vmSlot, n int) {
	*cs = vmCallOther(ctl, cs.objValue(), unsafe.Slice((*vmSlot)(unsafe.Add(unsafe.Pointer(cs), 16)), n))
}

// vmCallCold is an exact-arity bytecode call that first has to lower the
// callee, grow the record stack or move to another arena block. It leaves st
// at the callee's first instruction.
//
//go:noinline
func vmCallCold(ctl *ControlFlow, st *vmState, in *Instr, bf *scmBytecodeFunc) {
	f2 := bf.fn
	rc2 := f2.rc
	if rc2 == nil {
		rc2 = f2.lower()
	}
	d := ctl.vmDepth
	if d >= len(ctl.vmFrames) {
		ctl.vmFramesGrow()
	}
	fr := &ctl.vmFrames[d]
	fr.fn, fr.ip, fr.regs, fr.base = st.cur, st.ip, st.regs, st.base
	ctl.vmDepth = d + 1
	cs := r(st.regs, in.A)
	regs2 := unsafe.Add(unsafe.Pointer(cs), 16)
	base2 := st.base + int(in.A)>>4 + 1
	if base2+rc2.size > ctl.frameCap {
		was := ctl.frameCur
		base2 = ctl.frameEnsure(base2, rc2.size, unsafe.Slice((*vmSlot)(regs2), int(in.B)))
		regs2 = unsafe.Pointer(&ctl.frameBlocks[ctl.frameCur][base2])
		if ctl.frameCur != was {
			fr.crossed, fr.cur = true, was
		}
	}
	ctl.frameTop = base2 + rc2.size
	st.cur, st.regs, st.base = bf, regs2, base2
	st.ip, st.ks = rc2.fast, rc2.k
}

// vmTailCold makes the current frame ready to become a frame of f2, which may
// move it to another arena block along with the pending call's window.
//
//go:noinline
func vmTailCold(ctl *ControlFlow, st *vmState, baseDepth int, f2 *BytecodeFunc) *vmCode {
	rc2 := f2.rc
	if rc2 == nil {
		rc2 = f2.lower()
	}
	if st.base+rc2.size > ctl.frameCap {
		was := ctl.frameCur
		// Keep the whole current frame: it holds the pending call's window.
		size := max(rc2.size, ctl.frameTop-st.base)
		base2 := ctl.frameEnsure(st.base, size, unsafe.Slice((*vmSlot)(st.regs), ctl.frameTop-st.base))
		st.regs, st.base = unsafe.Pointer(&ctl.frameBlocks[ctl.frameCur][base2]), base2
		if d := ctl.vmDepth; ctl.frameCur != was && d > baseDepth {
			if fr := &ctl.vmFrames[d-1]; !fr.crossed {
				fr.crossed, fr.cur = true, was
			}
		}
	}
	return rc2
}

// vmTailOut is a tail call to anything but an exact-arity bytecode function. It
// returns true if this was the base frame, leaving the call for the trampoline
// that entered vmRun. Otherwise it pops the frame before running the call, so
// that a chain of tail calls through natives cannot grow the arena, and leaves
// st in the caller with the result in place.
//
//go:noinline
func vmTailOut(ctl *ControlFlow, st *vmState, cs *vmSlot, n, baseDepth int) bool {
	ctl.tailApplySlots(cs.objValue(), unsafe.Slice((*vmSlot)(unsafe.Add(unsafe.Pointer(cs), 16)), n))
	d := ctl.vmDepth
	if d == baseDepth {
		return true
	}
	d--
	fr := &ctl.vmFrames[d]
	st.cur, st.ip, st.regs, st.base = fr.fn, fr.ip, fr.regs, fr.base
	rc2 := st.cur.fn.rc
	st.ks = rc2.k
	ctl.vmDepth = d
	if fr.crossed {
		fr.crossed = false
		ctl.frameLeaveBlock(fr.cur)
	}
	ctl.frameTop = st.base + rc2.size
	res := vmSlot{obj: trampoline(ctl)}
	*r(st.regs, (*Instr)(unsafe.Add(st.ip, -int(instrSize))).A) = res
	return false
}

// frameLeaveDepth finishes a return whose call had changed arena blocks.
//
//go:noinline
func (ctl *ControlFlow) frameLeaveDepth() {
	fr := &ctl.vmFrames[ctl.vmDepth]
	fr.crossed = false
	ctl.frameLeaveBlock(fr.cur)
}
