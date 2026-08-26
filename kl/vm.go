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
	OP_ADD           = uint8(13) // pop y,x: push x+y
	OP_SUB           = uint8(14) // pop y,x: push x-y
	OP_MUL           = uint8(15) // pop y,x: push x*y
	OP_DIV           = uint8(24) // pop y,x: push x/y
	OP_LT            = uint8(16) // pop y,x: push x<y
	OP_LE            = uint8(17) // pop y,x: push x<=y
	OP_GT            = uint8(18) // pop y,x: push x>y
	OP_GE            = uint8(19) // pop y,x: push x>=y
	OP_EQ            = uint8(20) // pop y,x: push x=y  (structural equality, delegates to equal())
	OP_NOT           = uint8(21) // pop x: push (not x)
	OP_GUARDED_CONST = uint8(22) // A=result const, B=symbol const, C=x const, D=y const
	OP_GUARDED_PRIM  = uint8(23) // A=arity, B=symbol const; args are on stack
)

// Instr is a single VM instruction. A-D are signed operands; C/D are used by
// guarded constant specializations and remain zero for legacy instructions.
type Instr struct {
	Op uint8
	A  int32
	B  int32
	C  int32
	D  int32
}

// BytecodeFunc holds the compiled representation of a KL function.
type BytecodeFunc struct {
	Name    string
	Arity   int
	Nlocals int // number of local slots (includes arity slots for params)
	Code    []Instr
	Consts  []Obj // constant pool (numbers, strings, symbols, nested BytecodeFunc objs)
	// TypeHints preserves advisory (type ...) annotations for later typed-IR
	// passes. Hints never alter runtime behavior or introduce errors.
	TypeHints []TypeHint
}

type TypeHint struct {
	PC   int
	Name string
	// Kinds is the conservative shape implied by Name. Unknown annotation
	// names intentionally map to KindUnknown; annotations are advisory only.
	Kinds KindSet
	// Source identifies this as source-level metadata rather than a proven
	// runtime fact. Consumers must retain guards before specializing it.
	Source TypeHintSource
}

// TypeHintSource records how a type fact was obtained. Source annotations are
// never proof: they are useful for choosing a guarded region, but cannot alter
// Shen semantics when the runtime value disagrees.
type TypeHintSource uint8

const (
	TypeHintSourceUnknown TypeHintSource = iota
	TypeHintSourceAnnotation
)

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

// vmGuardedPrimitive executes the small, side-effect-free primitive subset
// that has a typed representation. It returns ok=false when the arguments
// are not suitable for specialization; callers must then use the ordinary
// dynamic call so that the primitive's exact error behavior is retained.
func vmGuardedPrimitive(sym Obj, args []vmSlot) (vmSlot, bool) {
	if !TypedIREnabled() || !HasCanonicalPrimitiveBinding(sym) {
		return vmSlot{}, false
	}
	name := GetSymbol(sym)
	obj := func(i int) Obj { return args[i].objValue() }
	boolResult := func(o Obj) (vmSlot, bool) {
		if o != True && o != False {
			return vmSlot{}, false
		}
		return vmSlot{obj: o}, true
	}
	switch name {
	case "number?", "integer?", "string?", "symbol?", "cons?", "absvector?", "variable?":
		var r Obj
		switch name {
		case "number?":
			r = PrimIsNumber(obj(0))
		case "integer?":
			r = PrimIsInteger(obj(0))
		case "string?":
			r = PrimIsString(obj(0))
		case "symbol?":
			r = PrimIsSymbol(obj(0))
		case "cons?":
			r = PrimIsPair(obj(0))
		case "absvector?":
			r = PrimIsVector(obj(0))
		default:
			r = PrimIsVariable(obj(0))
		}
		return boolResult(r)
	case "not":
		o := obj(0)
		if o != True && o != False {
			return vmSlot{}, false
		}
		return slotFromObj(PrimNot(o)), true
	case "cons":
		return slotFromObj(PrimCons(obj(0), obj(1))), true
	case "hd", "tl":
		if TypedObjectKind(obj(0)) != KindPair {
			return vmSlot{}, false
		}
		if name == "hd" {
			h, _ := TypedPairHead(obj(0))
			return slotFromObj(h), true
		}
		t, _ := TypedPairTail(obj(0))
		return slotFromObj(t), true
	case "cn":
		x, okx := TypedString(obj(0))
		y, oky := TypedString(obj(1))
		if !okx || !oky {
			return vmSlot{}, false
		}
		return slotFromObj(TypedMaterializeString(x + y)), true
	case "tlstr":
		x, ok := TypedString(obj(0))
		if !ok {
			return vmSlot{}, false
		}
		return slotFromObj(TypedMaterializeString(TypedStringTailValue(x))), true
	case "pos":
		x, ok := TypedString(obj(0))
		n, okn := slotNumber(args[1])
		if !ok || !okn || !isPreciseInteger(n) || !fitsInt(n) {
			return vmSlot{}, false
		}
		return slotFromObj(TypedMaterializeString(TypedStringIndexValue(x, int(n)))), true
	case "string->n":
		x, ok := TypedString(obj(0))
		if !ok || len([]rune(x)) == 0 {
			return vmSlot{}, false
		}
		return slotFromInteger(int([]rune(x)[0])), true
	case "n->string":
		n, ok := slotNumber(args[0])
		if !ok || !isPreciseInteger(n) || !fitsInt(n) || n < 0 || n > unicodeMaxRune {
			return vmSlot{}, false
		}
		return slotFromObj(TypedMaterializeString(string(rune(int(n))))), true
	case "/":
		x, okx := slotNumber(args[0])
		y, oky := slotNumber(args[1])
		if !okx || !oky {
			return vmSlot{}, false
		}
		return slotFromNumber(TypedDivideValue(x, y)), true
	case "<-address":
		if TypedObjectKind(obj(0)) != KindVector {
			return vmSlot{}, false
		}
		n, ok := slotNumber(args[1])
		if !ok || !isPreciseInteger(n) || !fitsInt(n) {
			return vmSlot{}, false
		}
		return slotFromObj(TypedVectorGet(obj(0), int(n))), true
	case "address->":
		if TypedObjectKind(obj(0)) != KindVector {
			return vmSlot{}, false
		}
		n, ok := slotNumber(args[1])
		if !ok || !isPreciseInteger(n) || !fitsInt(n) {
			return vmSlot{}, false
		}
		return slotFromObj(TypedVectorSet(obj(0), int(n), obj(2))), true
	case "absvector":
		n, ok := slotNumber(args[0])
		if !ok || !isPreciseInteger(n) || !fitsInt(n) {
			return vmSlot{}, false
		}
		return slotFromObj(PrimAbsvector(MakeInteger(int(n)))), true
	}
	return vmSlot{}, false
}

const unicodeMaxRune = 0x10ffff

// vmIntrinsicFallback preserves dynamic rebinding semantics for bytecode
// instructions specialized at compile time. It returns the current function's
// result when the binding is no longer canonical (or specialization is off).
func vmIntrinsicFallback(ctl *ControlFlow, sym Obj, args ...vmSlot) (vmSlot, bool) {
	// Older callers can construct arithmetic instructions without the symbol
	// operand (B defaults to zero). In that case retain the original intrinsic
	// behavior instead of attempting to type-assert a nil/non-symbol constant.
	if !IsSymbol(sym) {
		return vmSlot{}, false
	}
	if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym) {
		return vmSlot{}, false
	}
	return vmDynamicCall(ctl, sym, args...)
}

// vmDynamicCall invokes the current binding unconditionally. This is needed
// when a guarded primitive's runtime type check misses: canonical bindings
// still need their ordinary (catchable) error behavior for bad arguments.
func vmDynamicCall(ctl *ControlFlow, sym Obj, args ...vmSlot) (vmSlot, bool) {
	if !IsSymbol(sym) {
		return vmSlot{}, false
	}
	s := mustSymbol(sym)
	if s.function == nil {
		panic(MakeError(fmt.Sprintf("function %s not bound", s.str)))
	}
	objs := make([]Obj, len(args))
	for i, arg := range args {
		objs[i] = arg.objValue()
	}
	ctl.tailApplySlice(s.function, objs)
	return slotFromObj(trampoline(ctl)), true
}

// instrConstSym defensively reads the optional symbol operand carried by
// specialized instructions. Zero-value/legacy Instr values have no such
// operand and should continue down their built-in fast path.
func instrConstSym(consts []Obj, index int32) Obj {
	if index < 0 || int(index) >= len(consts) {
		return nil
	}
	return consts[index]
}

func intrinsicSymbolForOp(op uint8) Obj {
	for _, name := range []string{"+", "-", "*", "<", "<=", ">", ">=", "="} {
		sym := MakeSymbol(name)
		if intrinsicOp2(sym) == op {
			return sym
		}
	}
	return nil
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

		case OP_GUARDED_CONST:
			var sym Obj
			if instr.B < 0 {
				sym = intrinsicSymbolForOp(uint8(-instr.B))
			} else if int(instr.B) < len(consts) {
				sym = consts[instr.B]
			}
			if sym != nil && TypedIREnabled() && HasCanonicalPrimitiveBinding(sym) {
				stack = append(stack, slotFromObj(consts[instr.A]))
				continue
			}
			if r, ok := vmIntrinsicFallback(ctl, sym, slotFromInteger(int(instr.C)), slotFromInteger(int(instr.D))); ok {
				stack = append(stack, r)
				continue
			}
			panic("vmExec: guarded constant failed without fallback")

		case OP_GUARDED_PRIM:
			n := int(instr.A)
			base := len(stack) - n
			if base < 0 || instr.B < 0 || int(instr.B) >= len(consts) {
				panic("vmExec: malformed guarded primitive")
			}
			args := stack[base:]
			if r, ok := vmGuardedPrimitive(consts[instr.B], args); ok {
				stack = append(stack[:base], r)
				continue
			}
			if r, ok := vmDynamicCall(ctl, consts[instr.B], args...); ok {
				stack = append(stack[:base], r)
				continue
			}
			panic("vmExec: guarded primitive failed without fallback")

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
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotAdd(x, y))

		case OP_SUB:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotSub(x, y))

		case OP_MUL:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotMul(x, y))

		case OP_DIV:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			a, oka := slotNumber(x)
			b, okb := slotNumber(y)
			if !oka || !okb {
				stack = append(stack, slotFromObj(PrimNumberDivide(x.objValue(), y.objValue())))
				continue
			}
			stack = append(stack, slotFromNumber(TypedDivideValue(a, b)))

		case OP_LT:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotCmp(x, y, OP_LT))

		case OP_LE:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotCmp(x, y, OP_LE))

		case OP_GT:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotCmp(x, y, OP_GT))

		case OP_GE:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotCmp(x, y, OP_GE))

		case OP_EQ:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x, y); ok {
				stack = append(stack, r)
				continue
			}
			stack = append(stack, slotEqual(x, y))

		case OP_NOT:
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if r, ok := vmIntrinsicFallback(ctl, instrConstSym(consts, instr.B), x); ok {
				stack = append(stack, r)
				continue
			}
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
