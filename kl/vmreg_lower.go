package kl

import (
	"fmt"
	"unsafe"
)

// Register bytecode. The compiler still emits stack bytecode (BytecodeFunc.Code);
// on a function's first call lower() rewrites it, with no operand stack: stack
// depth d becomes register Nlocals+d, loads of locals and constants become
// operands of the instruction that uses them, and an operation's result goes
// straight to the let variable or self-tail-call parameter that wants it.
//
// Operands A-D are byte offsets (register or constant index * 16, or a jump
// distance in bytes from the next instruction), so the executor adds them to a
// base pointer without scaling. A is the destination, or the jump distance on
// a conditional jump, whose operands are then B and C.
const (
	opRMOV       = uint8(iota) // A = B
	opRLOADK                   // A = k[B]
	opRGETUP                   // A = upvals[B] (B an index, not an offset)
	opRLOADG                   // A = k[B].function
	opRJMP                     // ip += A
	opRJF                      // if B == false: ip += A
	opRJT                      // if B == true: ip += A (= NOT then JF)
	opRNOT                     // A = not B
	opRJNEKI                   // if B != k[C]: ip += A; k has identity equality
	opRJNEKN                   // if B != k[C]: ip += A; k is a finite number
	opRJNE                     // if B != C: ip += A
	opRJNCONSP                 // if not (cons? B): ip += A
	opRHD                      // A = hd B
	opRTL                      // A = tl B
	opRCONSP                   // A = cons? B
	opRCONS                    // A = cons B C
	opRVGET                    // A = B[C]
	opRVSET                    // B[C] = D; A = B
	opRGCONST                  // A = k[C], a folded literal operation; B = its pc in Code
	opRSLOW                    // slow twin of a guarded instruction (see vmSlowOp)
	opRPRIM                    // A = prim C(A..A+B) with sym in the guard table
	opRDIV                     // A = B / C
	opRLT                      // A = B < C
	opRLE                      // A = B <= C
	opRGT                      // A = B > C
	opRGE                      // A = B >= C
	opREQ                      // A = B = C
	opRCLOSURE                 // A = closure of Consts[B] over regs A..A+C
	opRRET                     // return A
	opRRETK                    // return k[A]
	opRSELF                    // restart this function: ip += A, charge a step
	opRCALLG                   // A = k[C].function, then opRCALL
	opRCALL                    // A = (A)(A+1 .. A+B); A+1 is the callee's register 0
	opRTAILCALLG               // A = k[C].function, then opRTAILCALL
	opRTAILCALL                // tail-call A with args A+1 .. A+B
	opRADD                     // A = B + C
	opRADDK                    // A = B + k[C]
	opRSUB
	opRSUBK
	opRMUL
	opRMULK
	opRJNLT // if not (B < C): ip += A
	opRJNLTK
	opRJNLE
	opRJNLEK
	opRJNGT
	opRJNGTK
	opRJNGE
	opRJNGEK
	opRTRAP // pads the code, so ip stays inside its allocation
)

// vmCode is a function's lowered form.
type vmCode struct {
	code  []Instr
	slow  []Instr // lazily built twin of code; see slowCode
	guard []int32 // per instruction: index+1 in Consts of the symbol it specializes
	fast  unsafe.Pointer
	k     unsafe.Pointer // constSlots, by byte offset
	n     int            // len(code)
	size  int            // registers in a frame
}

// slowCode returns the twin of code in which every guarded instruction is
// opRSLOW, at the same index and with the same operands. It runs when
// vmGuardsOK is false; vmSlowOp finds the original through the index.
func (rc *vmCode) slowCode() unsafe.Pointer {
	if rc.slow == nil {
		slow := append([]Instr(nil), rc.code...)
		for i := range slow {
			if rc.guard[i] != 0 {
				slow[i].Op = opRSLOW
			}
		}
		rc.slow = slow
	}
	return unsafe.Pointer(&rc.slow[0])
}

const (
	vTemp  = iota // in its canonical register, Nlocals+depth
	vLocal        // in local register idx, whose value it still equals
	vConst        // constant idx
)

type ventry struct {
	kind, idx int
	def       int // index in out of the instruction that wrote this temp, or -1
}

type regLowerer struct {
	fn      *BytecodeFunc
	out     []Instr
	guard   []int32
	v       []ventry
	barrier int // out index of the last label or jump
	maxReg  int
	fixups  [][2]int // out index, stack pc of target
}

func off(reg int) int32 { return int32(reg) << 4 }

func (l *regLowerer) emit(in Instr, sym int32) int {
	l.out = append(l.out, in)
	l.guard = append(l.guard, sym)
	return len(l.out) - 1
}

func (l *regLowerer) pos(depth int) int {
	r := l.fn.Nlocals + depth
	if r > l.maxReg {
		l.maxReg = r
	}
	return r
}

// flush puts stack entry d in its canonical register.
func (l *regLowerer) flush(d int) {
	e := &l.v[d]
	switch e.kind {
	case vLocal:
		l.emit(Instr{Op: opRMOV, A: off(l.pos(d)), B: off(e.idx)}, 0)
	case vConst:
		l.emit(Instr{Op: opRLOADK, A: off(l.pos(d)), B: off(e.idx)}, 0)
	default:
		return
	}
	*e = ventry{kind: vTemp, def: -1}
}

func (l *regLowerer) flushFrom(d int) {
	for ; d < len(l.v); d++ {
		l.flush(d)
	}
}

func (l *regLowerer) pop() ventry {
	e := l.v[len(l.v)-1]
	l.v = l.v[:len(l.v)-1]
	return e
}

// regOf returns the register of an entry just popped from depth d,
// materializing a constant in the entry's own canonical register.
func (l *regLowerer) regOf(e ventry, d int) int32 {
	switch e.kind {
	case vLocal:
		return off(e.idx)
	case vConst:
		l.emit(Instr{Op: opRLOADK, A: off(l.pos(d)), B: off(e.idx)}, 0)
	}
	return off(l.pos(d))
}

// pushResult records that the instruction just emitted wrote the new top of stack.
func (l *regLowerer) pushResult(retargetable bool) {
	def := -1
	if retargetable {
		def = len(l.out) - 1
	}
	l.pos(len(l.v))
	l.v = append(l.v, ventry{kind: vTemp, def: def})
}

func (l *regLowerer) jumpTo(idx, target int) {
	l.fixups = append(l.fixups, [2]int{idx, target})
	l.barrier = len(l.out)
}

// instrTouches reports whether in reads or writes the register at byte offset o.
// Only used for local registers; a call's operands are all temporaries.
func instrTouches(in *Instr, o int32) bool {
	switch in.Op {
	case opRLOADK, opRGETUP, opRLOADG, opRGCONST:
		return in.A == o
	case opRMOV, opRNOT, opRHD, opRTL, opRCONSP, opRADDK, opRSUBK, opRMULK:
		return in.A == o || in.B == o
	case opRCONS, opRVGET, opRADD, opRSUB, opRMUL, opRDIV, opRLT, opRLE, opRGT, opRGE, opREQ:
		return in.A == o || in.B == o || in.C == o
	case opRVSET:
		return in.A == o || in.B == o || in.C == o || in.D == o
	case opRJF, opRJT, opRJNEKI, opRJNEKN, opRJNCONSP,
		opRJNLTK, opRJNLEK, opRJNGTK, opRJNGEK:
		return in.B == o
	case opRJNE, opRJNLT, opRJNLE, opRJNGT, opRJNGE:
		return in.B == o || in.C == o
	case opRJMP, opRCALL, opRCALLG, opRPRIM, opRCLOSURE:
		return false
	}
	return true
}

func canRetarget(op uint8) bool {
	switch op {
	case opRMOV, opRLOADK, opRGETUP, opRLOADG, opRGCONST, opRNOT, opRHD, opRTL, opRCONSP,
		opRCONS, opRVGET, opRVSET, opRADD, opRADDK, opRSUB, opRSUBK, opRMUL, opRMULK,
		opRDIV, opRLT, opRLE, opRGT, opRGE, opREQ:
		return true
	}
	return false
}

// fuseGlobal drops a LOADG of the callee at depth d when only moves separate it
// from the call: nothing observable can happen between them, so the call can
// read the binding itself. It returns the symbol's constant offset, or -1.
func (l *regLowerer) fuseGlobal(d int) int32 {
	e := l.v[d]
	if e.def < 0 || e.def < l.barrier || l.out[e.def].Op != opRLOADG {
		return -1
	}
	for _, in := range l.out[e.def+1:] {
		if in.Op != opRMOV && in.Op != opRLOADK {
			return -1
		}
	}
	k := l.out[e.def].B
	l.out = append(l.out[:e.def], l.out[e.def+1:]...)
	l.guard = append(l.guard[:e.def], l.guard[e.def+1:]...)
	return k
}

// lower builds and installs fn's register code. Bytecode from klCompiler is its
// only input; anything else is a compiler bug and panics.
func (fn *BytecodeFunc) lower() *vmCode {
	code := fn.Code
	ks := fn.constSlots()
	isTarget := make([]bool, len(code)+1)
	depthAt := make([]int, len(code)+1)
	newpc := make([]int, len(code)+1)
	for pc := range code {
		switch code[pc].Op {
		case OP_JUMP, OP_JUMP_FALSE, OP_EQ_JF:
			t := pc + 1 + int(code[pc].A)
			if t <= pc {
				panic("lower: backward jump")
			}
			isTarget[t] = true
		}
	}
	l := &regLowerer{fn: fn, maxReg: fn.Nlocals}
	reachable := true
	for pc := 0; pc < len(code); pc++ {
		if isTarget[pc] {
			if reachable {
				l.flushFrom(0)
			} else {
				l.v = l.v[:0]
				for i := 0; i < depthAt[pc]; i++ {
					l.v = append(l.v, ventry{kind: vTemp, def: -1})
				}
				reachable = true
			}
			for i := range l.v {
				l.v[i].def = -1
			}
			l.barrier = len(l.out)
			newpc[pc] = len(l.out)
		}
		if !reachable {
			continue
		}
		in := &code[pc]
		// A compare or predicate whose only consumer is the next instruction,
		// a conditional jump nothing else jumps into, branches itself.
		nextJF := pc+1 < len(code) && code[pc+1].Op == OP_JUMP_FALSE && !isTarget[pc+1]
		branch := func(op uint8, b, c, d, sym int32) {
			l.flushFrom(0)
			idx := l.emit(Instr{Op: op, B: b, C: c, D: d}, sym)
			t := pc + 2 + int(code[pc+1].A)
			depthAt[t] = len(l.v)
			l.jumpTo(idx, t)
			pc++
		}
		switch in.Op {
		case OP_LOAD_CONST:
			l.pos(len(l.v))
			l.v = append(l.v, ventry{kind: vConst, idx: int(in.A), def: -1})
		case OP_LOAD_LOCAL:
			l.pos(len(l.v))
			l.v = append(l.v, ventry{kind: vLocal, idx: int(in.A), def: -1})
		case OP_LOAD_UPVAL:
			l.emit(Instr{Op: opRGETUP, A: off(l.pos(len(l.v))), B: in.A}, 0)
			l.pushResult(true)
		case OP_LOAD_GLOBAL:
			l.emit(Instr{Op: opRLOADG, A: off(l.pos(len(l.v))), B: off(int(in.A))}, 0)
			l.pushResult(true)
		case OP_GUARDED_CONST:
			sym := in.B
			if sym < 0 {
				sym = int32(l.fn.addIntrinsicConst(uint8(-sym)))
			}
			l.emit(Instr{Op: opRGCONST, A: off(l.pos(len(l.v))), B: int32(pc), C: off(int(in.A))}, sym+1)
			l.pushResult(true)
		case OP_STORE_LOCAL:
			e := l.pop()
			for d := range l.v {
				if l.v[d].kind == vLocal && l.v[d].idx == int(in.A) {
					l.flush(d)
				}
			}
			switch {
			case e.kind == vTemp && e.def >= 0 && e.def == len(l.out)-1 && e.def >= l.barrier && canRetarget(l.out[e.def].Op):
				l.out[e.def].A = off(int(in.A))
			case e.kind == vConst:
				l.emit(Instr{Op: opRLOADK, A: off(int(in.A)), B: off(e.idx)}, 0)
			default:
				l.emit(Instr{Op: opRMOV, A: off(int(in.A)), B: l.regOf(e, len(l.v))}, 0)
			}
		case OP_POP:
			l.pop()
		case OP_NOT:
			e := l.pop()
			d := len(l.v)
			b := l.regOf(e, d)
			if nextJF {
				branch(opRJT, b, 0, 0, in.B+1)
				continue
			}
			l.emit(Instr{Op: opRNOT, A: off(l.pos(d)), B: b}, in.B+1)
			l.pushResult(true)
		case OP_ADD, OP_SUB, OP_MUL, OP_DIV, OP_LT, OP_LE, OP_GT, OP_GE, OP_EQ, OP_EQ_JF:
			y := l.pop()
			x := l.pop()
			d := len(l.v)
			sym := in.B + 1
			if in.Op == OP_EQ || in.Op == OP_EQ_JF {
				if in.Op == OP_EQ_JF || nextJF {
					l.lowerEqBranch(pc, x, y, d, sym, in, code, branch, depthAt)
					if in.Op == OP_EQ_JF {
						continue
					}
					continue
				}
			}
			b := l.regOf(x, d)
			if t, ok := map[uint8]uint8{OP_LT: opRJNLT, OP_LE: opRJNLE, OP_GT: opRJNGT, OP_GE: opRJNGE}[in.Op]; ok && nextJF {
				if y.kind == vConst {
					branch(t+1, b, off(y.idx), 0, sym)
				} else {
					branch(t, b, l.regOf(y, d+1), 0, sym)
				}
				continue
			}
			var op uint8
			switch in.Op {
			case OP_ADD:
				op = opRADD
			case OP_SUB:
				op = opRSUB
			case OP_MUL:
				op = opRMUL
			case OP_DIV:
				op = opRDIV
			case OP_LT:
				op = opRLT
			case OP_LE:
				op = opRLE
			case OP_GT:
				op = opRGT
			case OP_GE:
				op = opRGE
			case OP_EQ:
				op = opREQ
			}
			if y.kind == vConst && (op == opRADD || op == opRSUB || op == opRMUL) {
				l.emit(Instr{Op: op + 1, A: off(l.pos(d)), B: b, C: off(y.idx)}, sym)
			} else {
				c := l.regOf(y, d+1)
				l.emit(Instr{Op: op, A: off(l.pos(d)), B: b, C: c}, sym)
			}
			l.pushResult(true)
		case OP_GP1_LOCAL:
			// Same as LOAD_LOCAL; GUARDED_PRIM 1 from here on.
			l.pos(len(l.v))
			l.v = append(l.v, ventry{kind: vLocal, idx: int(in.A), def: -1})
			if l.lowerPrim(pc, in.B, in.C, 1, nextJF, branch) {
				continue
			}
		case OP_GUARDED_PRIM:
			if l.lowerPrim(pc, in.B, in.C, int(in.A), nextJF, branch) {
				continue
			}
		case OP_JUMP_FALSE:
			e := l.pop()
			d := len(l.v)
			b := l.regOf(e, d)
			l.flushFrom(0)
			idx := l.emit(Instr{Op: opRJF, B: b}, 0)
			t := pc + 1 + int(in.A)
			depthAt[t] = len(l.v)
			l.jumpTo(idx, t)
		case OP_JUMP:
			l.flushFrom(0)
			idx := l.emit(Instr{Op: opRJMP}, 0)
			t := pc + 1 + int(in.A)
			depthAt[t] = len(l.v)
			l.jumpTo(idx, t)
			reachable = false
		case OP_RETURN:
			e := l.pop()
			if e.kind == vConst {
				l.emit(Instr{Op: opRRETK, A: off(e.idx)}, 0)
			} else {
				l.emit(Instr{Op: opRRET, A: l.regOf(e, len(l.v))}, 0)
			}
			reachable = false
		case OP_MAKE_CLOSURE:
			n := int(in.B)
			base := len(l.v) - n
			l.flushFrom(base)
			l.v = l.v[:base]
			l.emit(Instr{Op: opRCLOSURE, A: off(l.pos(base)), B: in.A, C: in.B}, 0)
			l.pushResult(false)
		case OP_CALL, OP_TAIL_CALL:
			n := int(in.A)
			base := len(l.v) - n - 1
			for d := base + 1; d < len(l.v); d++ {
				l.flush(d)
			}
			k := int32(-1)
			if l.v[base].kind == vTemp {
				k = l.fuseGlobal(base)
			} else {
				l.flush(base)
			}
			op := opRCALL
			if in.Op == OP_TAIL_CALL {
				op = opRTAILCALL
			}
			if k >= 0 {
				l.emit(Instr{Op: op - 1, A: off(l.pos(base)), B: in.A, C: k}, 0)
			} else {
				l.emit(Instr{Op: op, A: off(l.pos(base)), B: in.A}, 0)
			}
			l.v = l.v[:base]
			if in.Op == OP_TAIL_CALL {
				reachable = false
			} else {
				l.pushResult(false)
			}
		case OP_SELF_TAIL_CALL:
			l.lowerSelfTail(int(in.A), uint32(in.B))
			idx := l.emit(Instr{Op: opRSELF}, 0)
			l.out[idx].A = int32(-(idx + 1) * int(instrSize))
			l.barrier = len(l.out)
			l.v = l.v[:0]
			reachable = false
		default:
			panic(fmt.Sprintf("lower: opcode %d", in.Op))
		}
	}
	if reachable {
		panic("lower: code runs off its end")
	}
	newpc[len(code)] = len(l.out)
	for _, f := range l.fixups {
		l.out[f[0]].A = int32((newpc[f[1]] - f[0] - 1) * int(instrSize))
	}
	l.emit(Instr{Op: opRTRAP}, 0)
	rc := &vmCode{code: l.out, guard: l.guard, n: len(l.out), size: l.maxReg + 1}
	rc.fast = unsafe.Pointer(&l.out[0])
	if len(ks) > 0 {
		rc.k = unsafe.Pointer(&ks[0])
	}
	fn.rc = rc
	return rc
}

// addIntrinsicConst returns the Consts index of an intrinsic's symbol, for a
// folded operation that carried only its opcode.
func (fn *BytecodeFunc) addIntrinsicConst(op uint8) int {
	sym := intrinsicSymbolForOp(op)
	for i, c := range fn.Consts {
		if c == sym {
			return i
		}
	}
	fn.Consts = append(fn.Consts, sym)
	fn.kslots = nil
	fn.constSlots()
	return len(fn.Consts) - 1
}

// lowerEqBranch lowers (= x y) followed by a conditional jump. An EQ_JF
// carries the jump itself; an EQ is followed by a JUMP_FALSE. D=1 records that
// the constant was the left operand, for a rebound = to see the original order.
func (l *regLowerer) lowerEqBranch(pc int, x, y ventry, d int, sym int32, in *Instr, code []Instr, branch func(uint8, int32, int32, int32, int32), depthAt []int) {
	ks := l.fn.constSlots()
	var op uint8
	var b, c, flip int32
	kind := func(e ventry) uint8 {
		if e.kind != vConst {
			return 0
		}
		s := ks[e.idx]
		if s.obj == nil {
			return opRJNEKN
		}
		switch *s.obj {
		case scmHeadSymbol, scmHeadBoolean, scmHeadNull:
			return opRJNEKI
		}
		return 0
	}
	switch {
	case kind(y) != 0:
		op, b, c = kind(y), l.regOf(x, d), off(y.idx)
	case kind(x) != 0:
		op, b, c, flip = kind(x), l.regOf(y, d+1), off(x.idx), 1
	default:
		op = opRJNE
		b = l.regOf(x, d)
		c = l.regOf(y, d+1)
	}
	if in.Op == OP_EQ_JF {
		l.flushFrom(0)
		idx := l.emit(Instr{Op: op, B: b, C: c, D: flip}, sym)
		t := pc + 1 + int(in.A)
		depthAt[t] = len(l.v)
		l.jumpTo(idx, t)
		return
	}
	branch(op, b, c, flip, sym)
}

// lowerPrim lowers a guarded primitive of n arguments on top of the virtual
// stack. It returns true if it consumed the next instruction.
func (l *regLowerer) lowerPrim(pc int, symB, idC int32, n int, nextJF bool, branch func(uint8, int32, int32, int32, int32)) bool {
	sym := symB + 1
	id := idC - 1
	if idC == 0 {
		id = guardedPrimID(l.fn.Consts[symB])
	}
	base := len(l.v) - n
	var op uint8
	switch id {
	case gpHd:
		op = opRHD
	case gpTl:
		op = opRTL
	case gpConsp:
		op = opRCONSP
	case gpCons:
		op = opRCONS
	case gpVectorGet:
		op = opRVGET
	case gpVectorSet:
		op = opRVSET
	default:
		l.flushFrom(base)
		l.v = l.v[:base]
		l.emit(Instr{Op: opRPRIM, A: off(l.pos(base)), B: int32(n), C: id + 1}, sym)
		l.pushResult(false)
		return false
	}
	var regs [3]int32
	for i := n - 1; i >= 0; i-- {
		regs[i] = 0
	}
	es := append([]ventry(nil), l.v[base:]...)
	l.v = l.v[:base]
	for i, e := range es {
		regs[i] = l.regOf(e, base+i)
	}
	if op == opRCONSP && nextJF {
		branch(opRJNCONSP, regs[0], 0, 0, sym)
		return true
	}
	l.emit(Instr{Op: op, A: off(l.pos(base)), B: regs[0], C: regs[1], D: regs[2]}, sym)
	l.pushResult(true)
	return false
}

// lowerSelfTail assigns the top n entries to parameters, in parallel: registers
// 0..n-1, or with a nonzero mask, the parameters whose bits are set, in order.
// When an argument was just computed into a temporary and nothing since has
// used its parameter, the computation writes the parameter directly.
func (l *regLowerer) lowerSelfTail(n int, mask uint32) {
	base := len(l.v) - n
	for d := 0; d < base; d++ {
		l.flush(d) // nothing lives under a tail call; keep it correct regardless
	}
	es := l.v[base:]
	param := make([]int, n)
	for i, p := 0, 0; i < n; i++ {
		if mask != 0 {
			for mask&(1<<p) == 0 {
				p++
			}
		} else {
			p = i
		}
		param[i] = p
		p++
	}
	readByOther := func(i int) bool {
		for j := range es {
			if j != i && es[j].kind == vLocal && es[j].idx == param[i] {
				return true
			}
		}
		return false
	}
	done := make([]bool, n)
	for i := range es {
		e := &es[i]
		if e.kind == vLocal && e.idx == param[i] {
			done[i] = true
			continue
		}
		if e.kind == vTemp && e.def >= 0 && canRetarget(l.out[e.def].Op) && !readByOther(i) {
			o := off(param[i])
			clean := true
			for j := e.def + 1; j < len(l.out); j++ {
				if instrTouches(&l.out[j], o) {
					clean = false
					break
				}
			}
			if clean {
				l.out[e.def].A = o
				done[i] = true
			}
		}
	}
	// A parameter another argument still reads must be saved before it is
	// overwritten.
	for i := range es {
		if done[i] || es[i].kind != vLocal {
			continue
		}
		for j := range es {
			if !done[j] && j != i && param[j] == es[i].idx {
				l.flush(base + i)
				break
			}
		}
	}
	for i := range es {
		if done[i] {
			continue
		}
		e := es[i]
		switch e.kind {
		case vConst:
			l.emit(Instr{Op: opRLOADK, A: off(param[i]), B: off(e.idx)}, 0)
		case vLocal:
			l.emit(Instr{Op: opRMOV, A: off(param[i]), B: off(e.idx)}, 0)
		default:
			l.emit(Instr{Op: opRMOV, A: off(param[i]), B: off(l.pos(base + i))}, 0)
		}
	}
}

// vmSlowOp runs a guarded instruction the executor's inline path did not take:
// a rebound primitive, or operands outside the inline case. It returns how far
// to jump, in bytes, nonzero only for a conditional jump that is taken. The
// semantics are those of the stack instructions: a type miss reaches the real
// primitive, for its error.
//
//go:noinline
func vmSlowOp(ctl *ControlFlow, cur *scmBytecodeFunc, regs, ks unsafe.Pointer, in *Instr) int {
	fn := cur.fn
	rc := fn.rc
	if in.Op == opRSLOW {
		i := (uintptr(unsafe.Pointer(in)) - uintptr(unsafe.Pointer(&rc.slow[0]))) / instrSize
		in = &rc.code[i]
	}
	// Arithmetic and compares on numbers, with the guards intact: a fixnum or
	// boxed operand, typically a value out of a list. No Shen code can run.
	if vmGuardsOK {
		if op, ok := vmSlowNumeric(regs, ks, in); ok {
			return op
		}
	}
	var sym Obj
	if g := rc.guard[(uintptr(unsafe.Pointer(in))-uintptr(rc.fast))/instrSize]; g != 0 {
		sym = fn.Consts[g-1]
	}
	y := func() vmSlot {
		switch in.Op {
		case opRADDK, opRSUBK, opRMULK, opRJNLTK, opRJNLEK, opRJNGTK, opRJNGEK, opRJNEKI, opRJNEKN:
			return *r(ks, in.C)
		}
		return *r(regs, in.C)
	}
	binary := func(op uint8) vmSlot {
		return vmBinarySlowSym(ctl, op, sym, *r(regs, in.B), y())
	}
	jump := func(v vmSlot) int {
		o := v.objValue()
		if o == False {
			return int(in.A)
		}
		if o != True {
			panic(vmBadIf())
		}
		return 0
	}
	prim := func(id int32, args ...vmSlot) vmSlot {
		return vmGuardedPrimSlow(ctl, sym, id+1, args)
	}
	switch in.Op {
	case opRADD, opRADDK:
		*r(regs, in.A) = binary(OP_ADD)
	case opRSUB, opRSUBK:
		*r(regs, in.A) = binary(OP_SUB)
	case opRMUL, opRMULK:
		*r(regs, in.A) = binary(OP_MUL)
	case opRDIV:
		*r(regs, in.A) = binary(OP_DIV)
	case opRLT:
		*r(regs, in.A) = binary(OP_LT)
	case opRLE:
		*r(regs, in.A) = binary(OP_LE)
	case opRGT:
		*r(regs, in.A) = binary(OP_GT)
	case opRGE:
		*r(regs, in.A) = binary(OP_GE)
	case opRJNLT, opRJNLTK:
		return jump(binary(OP_LT))
	case opRJNLE, opRJNLEK:
		return jump(binary(OP_LE))
	case opRJNGT, opRJNGTK:
		return jump(binary(OP_GT))
	case opRJNGE, opRJNGEK:
		return jump(binary(OP_GE))
	case opREQ, opRJNE, opRJNEKI, opRJNEKN:
		a, b := *r(regs, in.B), y()
		if in.D == 1 {
			a, b = b, a
		}
		// Identity first: interned symbols, booleans, [], (fail), and any
		// shared heap object skip structural equal().
		var v vmSlot
		if a.obj != nil && a.obj == b.obj {
			v = vmSlot{obj: True}
		} else {
			v = vmBinarySlowSym(ctl, OP_EQ, sym, a, b)
		}
		if in.Op == opREQ {
			*r(regs, in.A) = v
			return 0
		}
		return jump(v)
	case opRNOT, opRJT:
		x := *r(regs, in.B)
		var v vmSlot
		f, dyn := vmSlot{}, false
		if !vmGuardsOK {
			f, dyn = vmIntrinsicFallback(ctl, sym, x)
		}
		switch {
		case dyn:
			v = f
		case x.obj == True:
			v.obj = False
		case x.obj == False:
			v.obj = True
		default:
			panic(MakeError("not: expected boolean"))
		}
		if in.Op == opRNOT {
			*r(regs, in.A) = v
			return 0
		}
		return jump(v)
	case opRHD:
		*r(regs, in.A) = prim(gpHd, *r(regs, in.B))
	case opRTL:
		*r(regs, in.A) = prim(gpTl, *r(regs, in.B))
	case opRCONSP:
		*r(regs, in.A) = prim(gpConsp, *r(regs, in.B))
	case opRJNCONSP:
		return jump(prim(gpConsp, *r(regs, in.B)))
	case opRCONS:
		*r(regs, in.A) = prim(gpCons, *r(regs, in.B), *r(regs, in.C))
	case opRVGET:
		*r(regs, in.A) = prim(gpVectorGet, *r(regs, in.B), *r(regs, in.C))
	case opRVSET:
		*r(regs, in.A) = prim(gpVectorSet, *r(regs, in.B), *r(regs, in.C), *r(regs, in.D))
	case opRPRIM:
		b := r(regs, in.A)
		*b = vmGuardedPrimSlow(ctl, sym, in.C, unsafe.Slice(b, int(in.B)))
	case opRGCONST:
		*r(regs, in.A) = vmGuardedConstSlow(ctl, fn.Consts, &fn.Code[in.B])
	default:
		panic(fmt.Sprintf("vmSlowOp: opcode %d", in.Op))
	}
	return 0
}

// vmSlowNumeric handles arithmetic and compares whose operands are numbers in a
// form the inline path does not read, boxed floats. ok is false for anything
// else, which needs the full path.
func vmSlowNumeric(regs, ks unsafe.Pointer, in *Instr) (jump int, ok bool) {
	var x, y *vmSlot
	switch in.Op {
	case opRADD, opRSUB, opRMUL, opRJNLT, opRJNLE, opRJNGT, opRJNGE:
		x, y = r(regs, in.B), r(regs, in.C)
	case opRADDK, opRSUBK, opRMULK, opRJNLTK, opRJNLEK, opRJNGTK, opRJNGEK:
		x, y = r(regs, in.B), r(ks, in.C)
	default:
		return 0, false
	}
	a, oka := slotNumber(*x)
	b, okb := slotNumber(*y)
	if !oka || !okb {
		return 0, false
	}
	var cond bool
	switch in.Op {
	case opRADD, opRADDK:
		*r(regs, in.A) = slotFromNumber(a + b)
		return 0, true
	case opRSUB, opRSUBK:
		*r(regs, in.A) = slotFromNumber(a - b)
		return 0, true
	case opRMUL, opRMULK:
		*r(regs, in.A) = slotFromNumber(a * b)
		return 0, true
	case opRJNLT, opRJNLTK:
		cond = a < b
	case opRJNLE, opRJNLEK:
		cond = a <= b
	case opRJNGT, opRJNGTK:
		cond = a > b
	default:
		cond = a >= b
	}
	if cond {
		return 0, true
	}
	return int(in.A), true
}
