package kl

import (
	"fmt"
	"math"
	"unsafe"
)

// ---- Opcodes ----

const (
	OP_LOAD_CONST     = uint8(0)  // push consts[A]
	OP_LOAD_LOCAL     = uint8(1)  // push locals[A]
	OP_STORE_LOCAL    = uint8(2)  // locals[A] = pop()
	OP_LOAD_GLOBAL    = uint8(3)  // push mustSymbol(consts[A]).function
	OP_LOAD_UPVAL     = uint8(4)  // push upvals[A]
	OP_CALL           = uint8(5)  // non-tail call: fn at stack[top-A-1], A args above
	OP_TAIL_CALL      = uint8(6)  // tail call: same layout, signal trampoline
	OP_RETURN         = uint8(7)  // ctl.Return(pop())
	OP_JUMP           = uint8(8)  // pc += A (signed)
	OP_JUMP_FALSE     = uint8(9)  // pop(); if == False: pc += A
	OP_MAKE_CLOSURE   = uint8(10) // consts[A]=BytecodeFunc; pop B upvals; push closure
	OP_POP            = uint8(11) // discard top of stack
	OP_SELF_TAIL_CALL = uint8(12) // A args on stack → update locals[0..A-1], jump pc=0
	// Arithmetic fast-paths (avoid trampoline + float64 boxing for fixnums)
	OP_ADD = uint8(13) // pop y,x: push x+y
	OP_SUB = uint8(14) // pop y,x: push x-y
	OP_MUL = uint8(15) // pop y,x: push x*y
	OP_LT  = uint8(16) // pop y,x: push x<y
	OP_LE  = uint8(17) // pop y,x: push x<=y
	OP_GT  = uint8(18) // pop y,x: push x>y
	OP_GE  = uint8(19) // pop y,x: push x>=y
	OP_EQ  = uint8(20) // pop y,x: push x=y  (structural equality, delegates to equal())
	OP_NOT = uint8(21) // pop x: push (not x)
)

// Instr is a single VM instruction.  A and B are signed operands.
type Instr struct {
	Op uint8
	A  int32
	B  int32
}

// BytecodeFunc holds the compiled representation of a KL function.
type BytecodeFunc struct {
	Name    string
	Arity   int
	Nlocals int // number of local slots (includes arity slots for params)
	Code    []Instr
	Consts  []Obj // constant pool (numbers, strings, symbols, nested BytecodeFunc objs)
}

// scmBytecodeFunc wraps a BytecodeFunc as an Obj so it can live in the Shen heap.
type scmBytecodeFunc struct {
	scmHead
	fn     *BytecodeFunc
	upvals []Obj // captured values for closures
}

// vmSlot keeps finite numbers unboxed while retaining an Obj for all other
// Shen values. obj==nil denotes a numeric slot.
type vmSlot struct {
	obj    Obj
	number float64
}

func slotFromObj(o Obj) vmSlot {
	if o != nil && isFixnum(o) {
		return vmSlot{obj: o}
	}
	if o != nil && *o == scmHeadNumber {
		f := GetNumber(o)
		if !math.IsNaN(f) && !math.IsInf(f, 0) {
			return vmSlot{number: f}
		}
	}
	return vmSlot{obj: o}
}
func slotFromNumber(f float64) vmSlot {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return vmSlot{obj: MakeNumber(f)}
	}
	if isPreciseInteger(f) && f >= float64(fixnumMin) && f < float64(fixnumMax) {
		return vmSlot{obj: MakeInteger(int(f))}
	}
	return vmSlot{number: f}
}
func slotFromInteger(i int) vmSlot {
	if i >= fixnumMin && i < fixnumMax {
		return vmSlot{obj: MakeInteger(i)}
	}
	return vmSlot{number: float64(i)}
}
func slotFixnum(s vmSlot) (int, bool) {
	if s.obj != nil && isFixnum(s.obj) {
		return fixnum(s.obj), true
	}
	return 0, false
}
func slotNumber(s vmSlot) (float64, bool) {
	if s.obj == nil {
		return s.number, true
	}
	if isFixnum(s.obj) || *s.obj == scmHeadNumber {
		return GetNumber(s.obj), true
	}
	return 0, false
}
func (s vmSlot) objValue() Obj {
	if s.obj != nil {
		return s.obj
	}
	return MakeNumber(s.number)
}
func slotArithmetic(s vmSlot) Obj {
	if s.obj == nil {
		return MakeNumber(s.number)
	}
	return s.obj
}
func slotAdd(x, y vmSlot) vmSlot {
	if a, ok := slotFixnum(x); ok {
		if b, ok := slotFixnum(y); ok {
			return slotFromInteger(a + b)
		}
	}
	if a, ok := slotNumber(x); ok {
		if b, ok := slotNumber(y); ok {
			return slotFromNumber(a + b)
		}
	}
	return slotFromObj(numAdd(slotArithmetic(x), slotArithmetic(y)))
}
func slotSub(x, y vmSlot) vmSlot {
	if a, ok := slotFixnum(x); ok {
		if b, ok := slotFixnum(y); ok {
			return slotFromInteger(a - b)
		}
	}
	if a, ok := slotNumber(x); ok {
		if b, ok := slotNumber(y); ok {
			return slotFromNumber(a - b)
		}
	}
	return slotFromObj(numSub(slotArithmetic(x), slotArithmetic(y)))
}
func slotMul(x, y vmSlot) vmSlot {
	if a, ok := slotFixnum(x); ok {
		if b, ok := slotFixnum(y); ok {
			return slotFromInteger(a * b)
		}
	}
	if a, ok := slotNumber(x); ok {
		if b, ok := slotNumber(y); ok {
			return slotFromNumber(a * b)
		}
	}
	return slotFromObj(numMul(slotArithmetic(x), slotArithmetic(y)))
}
func slotBool(v bool) vmSlot {
	if v {
		return vmSlot{obj: True}
	}
	return vmSlot{obj: False}
}
func slotCmp(x, y vmSlot, op uint8) vmSlot {
	if a, ok := slotFixnum(x); ok {
		if b, ok := slotFixnum(y); ok {
			switch op {
			case OP_LT:
				return slotBool(a < b)
			case OP_LE:
				return slotBool(a <= b)
			case OP_GT:
				return slotBool(a > b)
			default:
				return slotBool(a >= b)
			}
		}
	}
	if a, ok := slotNumber(x); ok {
		if b, ok := slotNumber(y); ok {
			switch op {
			case OP_LT:
				return slotBool(a < b)
			case OP_LE:
				return slotBool(a <= b)
			case OP_GT:
				return slotBool(a > b)
			default:
				return slotBool(a >= b)
			}
		}
	}
	switch op {
	case OP_LT:
		return slotFromObj(numCmpLT(slotArithmetic(x), slotArithmetic(y)))
	case OP_LE:
		return slotFromObj(numCmpLE(slotArithmetic(x), slotArithmetic(y)))
	case OP_GT:
		return slotFromObj(numCmpLT(slotArithmetic(y), slotArithmetic(x)))
	default:
		return slotFromObj(numCmpLE(slotArithmetic(y), slotArithmetic(x)))
	}
}
func slotEqual(x, y vmSlot) vmSlot {
	if x.obj != nil && y.obj != nil {
		return vmSlot{obj: equal(x.obj, y.obj)}
	}
	if a, ok := slotNumber(x); ok {
		if b, ok := slotNumber(y); ok {
			return slotBool(a == b)
		}
	}
	return vmSlot{obj: equal(slotArithmetic(x), slotArithmetic(y))}
}

const scmHeadBytecodeFunc scmHead = 50

func makeBytecodeObj(fn *BytecodeFunc, upvals []Obj) Obj {
	tmp := &scmBytecodeFunc{
		scmHead: scmHeadBytecodeFunc,
		fn:      fn,
		upvals:  upvals,
	}
	return &tmp.scmHead
}

func mustBytecodeFunc(o Obj) *scmBytecodeFunc {
	return (*scmBytecodeFunc)(unsafe.Pointer(o))
}

// ---- Arithmetic helpers (fixnum fast paths) ----

func numAdd(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		return MakeInteger(fixnum(x) + fixnum(y))
	}
	return MakeNumber(mustNumber(x) + mustNumber(y))
}

func numSub(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		return MakeInteger(fixnum(x) - fixnum(y))
	}
	return MakeNumber(mustNumber(x) - mustNumber(y))
}

func numMul(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		return MakeInteger(fixnum(x) * fixnum(y))
	}
	return MakeNumber(mustNumber(x) * mustNumber(y))
}

func numCmpLT(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		if fixnum(x) < fixnum(y) {
			return True
		}
		return False
	}
	if mustNumber(x) < mustNumber(y) {
		return True
	}
	return False
}

func numCmpLE(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		if fixnum(x) <= fixnum(y) {
			return True
		}
		return False
	}
	if mustNumber(x) <= mustNumber(y) {
		return True
	}
	return False
}

// vmApply is the entry point called from apply() for bytecode functions.
// It handles arity checking and partial application, then calls vmExec.
func vmApply(ctl *ControlFlow, bfObj Obj, args []Obj) {
	bf := mustBytecodeFunc(bfObj)
	fn := bf.fn
	provided := len(args)

	switch {
	case provided < fn.Arity:
		// Partial application: build a closure that waits for the remaining args.
		ctl.Return(vmPartialApply(fn.Arity, args, bfObj))
	case provided == fn.Arity:
		vmExec(ctl, bf, args)
	case provided > fn.Arity:
		// Over-application: call with the right arity, then apply the result.
		res := Call(ctl, bfObj, args[:fn.Arity]...)
		ctl.TailApply(res, args[fn.Arity:]...)
	}
}

// vmPartialApply creates a closure that captures the supplied args and waits
// for the remaining (required - len(provided)) arguments.
func vmPartialApply(required int, providedArgs []Obj, proc Obj) Obj {
	symbols := makeTempSymbols(required)
	env1 := envExtend(Nil, symbols[:len(providedArgs)], providedArgs)

	args := Nil
	for i, count := len(symbols)-1, required-len(providedArgs); count > 0; count-- {
		args = cons(symbols[i], args)
		i--
	}

	body := Nil
	for i := len(symbols) - 1; i >= 0; i-- {
		body = cons(symbols[i], body)
	}
	body = cons(proc, body)

	return makeProcedure(args, body, env1)
}

// vmExec runs one activation of a bytecode function.
// It either calls ctl.Return with the result or sets up a tail call.
func vmExec(ctl *ControlFlow, bf *scmBytecodeFunc, args []Obj) {
	var inline [8]vmSlot
	slots := inline[:0]
	if len(args) <= len(inline) {
		slots = inline[:len(args)]
	} else {
		slots = make([]vmSlot, len(args))
	}
	for i, a := range args {
		slots[i] = slotFromObj(a)
	}
	vmExecSlots(ctl, bf, slots)
}

func vmExecSlots(ctl *ControlFlow, bf *scmBytecodeFunc, args []vmSlot) {
	fn := bf.fn
	upvals := bf.upvals

	// Locals and the operand stack share one slab (stack starts at len==0 just
	// past the locals; append only ever writes at indices >= Nlocals, so the
	// regions never overlap). Slabs are recycled via ctl.framePool — the
	// previous per-activation make() dominated alloc_space on SHA/prng. If a
	// frame needs more than 16 stack slots append reallocates the stack away
	// from the slab, which is rare and still correct (the original slab is
	// what we putFrame).
	slab := ctl.takeFrame(fn.Nlocals)
	locals := slab
	copy(locals, args)

	stack := slab[fn.Nlocals:]
	pc := 0
	code := fn.Code
	consts := fn.Consts

	// Hoisted out of the loop: see ControlFlow.stepLimited. In every
	// non-fuzz process this is false and the per-instruction cost of the
	// step counter collapses to one predictable register test.
	limited := ctl.stepLimited()

	for {
		if limited {
			ctl.tick()
		}
		instr := code[pc]
		pc++
		switch instr.Op {

		case OP_LOAD_CONST:
			stack = append(stack, slotFromObj(consts[instr.A]))

		case OP_LOAD_LOCAL:
			stack = append(stack, locals[instr.A])

		case OP_STORE_LOCAL:
			locals[instr.A] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		case OP_LOAD_GLOBAL:
			sym := mustSymbol(consts[instr.A])
			if sym.function == nil {
				panic(MakeError(fmt.Sprintf("function %s not bound", sym.str)))
			}
			stack = append(stack, slotFromObj(sym.function))

		case OP_LOAD_UPVAL:
			stack = append(stack, slotFromObj(upvals[instr.A]))

		case OP_CALL:
			n := int(instr.A)
			base := len(stack) - n - 1
			callee := stack[base].objValue()
			callArgs := stack[base+1:]
			var result Obj
			switch *callee {
			case scmHeadBytecodeFunc:
				// Exact-arity: run the callee directly. If it returns, skip the
				// trampoline hop; if it sets up a tail call, finish via trampoline.
				// Partial/over-application still goes through vmApply.
				bfc := mustBytecodeFunc(callee)
				if n == bfc.fn.Arity {
					vmExecSlots(ctl, bfc, callArgs)
					if ctl.kind == ControlFlowReturn {
						result = ctl.data[ctl.pos]
						ctl.data = ctl.data[:ctl.pos]
					} else {
						result = trampoline(ctl)
					}
				} else {
					ctl.tailApplySlots(callee, callArgs)
					result = trampoline(ctl)
				}
			case scmHeadNative:
				// Exact-arity natives without captured slots: set up ctl.data
				// and invoke directly, only entering the trampoline if the
				// native tail-applies (e.g. fn with arity 0).
				nf := MustNative(callee)
				if len(nf.captured) == 0 && n == nf.require {
					ctl.tailApplySlots(callee, callArgs)
					nf.fn(ctl)
					if ctl.kind == ControlFlowReturn {
						result = ctl.data[ctl.pos]
						ctl.data = ctl.data[:ctl.pos]
					} else {
						result = trampoline(ctl)
					}
				} else {
					ctl.tailApplySlots(callee, callArgs)
					result = trampoline(ctl)
				}
			default:
				ctl.tailApplySlots(callee, callArgs)
				result = trampoline(ctl)
			}
			stack = stack[:base]
			stack = append(stack, slotFromObj(result))

		case OP_TAIL_CALL:
			n := int(instr.A)
			base := len(stack) - n - 1
			callee := stack[base].objValue()
			ctl.tailApplySlots(callee, stack[base+1:])
			ctl.putFrame(slab)
			return

		case OP_RETURN:
			ctl.Return(stack[len(stack)-1].objValue())
			ctl.putFrame(slab)
			return

		case OP_JUMP:
			pc += int(instr.A)

		case OP_JUMP_FALSE:
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vo := v.objValue()
			if vo == False {
				pc += int(instr.A)
			} else if vo != True {
				panic(MakeError("if requires a boolean"))
			}

		case OP_MAKE_CLOSURE:
			nUpvals := int(instr.B)
			innerBFObj := consts[instr.A]
			innerBF := mustBytecodeFunc(innerBFObj)
			captured := make([]Obj, nUpvals)
			base := len(stack) - nUpvals
			for i := range captured {
				captured[i] = stack[base+i].objValue()
			}
			stack = stack[:base]
			closure := makeBytecodeObj(innerBF.fn, captured)
			stack = append(stack, slotFromObj(closure))

		case OP_POP:
			stack = stack[:len(stack)-1]

		case OP_SELF_TAIL_CALL:
			n := int(instr.A)
			// All n args are on top of stack; copy into locals[0..n-1] and loop.
			copy(locals[:n], stack[len(stack)-n:])
			stack = stack[:len(stack)-n]
			pc = 0

		case OP_ADD:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotAdd(x, y))

		case OP_SUB:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotSub(x, y))

		case OP_MUL:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotMul(x, y))

		case OP_LT:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotCmp(x, y, OP_LT))

		case OP_LE:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotCmp(x, y, OP_LE))

		case OP_GT:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotCmp(x, y, OP_GT))

		case OP_GE:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotCmp(x, y, OP_GE))

		case OP_EQ:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, slotEqual(x, y))

		case OP_NOT:
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			xo := x.objValue()
			if xo == True {
				stack = append(stack, slotFromObj(False))
			} else if xo == False {
				stack = append(stack, slotFromObj(True))
			} else {
				panic(MakeError("not: expected boolean"))
			}

		default:
			panic(fmt.Sprintf("vmExec: unknown opcode %d at pc %d", instr.Op, pc-1))
		}
	}
}
