package kl

// klCompiler compiles a KL form to bytecode.
type klCompiler struct {
	fn     *BytecodeFunc
	locals map[Obj]int // symbol pointer -> slot index
	upvals []upvalInfo // upvalues captured from outer scope
	outer  *klCompiler // enclosing compiler (for closures)
	hints  []TypeHint
	thunks map[int]inlineThunk // local slot -> nonescaping freeze
}

// Keep the definition's lexical environment, not the thaw site's bindings.
// Slots are never reused, so shadowing cannot change the captured values.
type inlineThunk struct {
	body   Obj
	locals map[Obj]int
}

type upvalInfo struct {
	sym      Obj
	outerRef varRef // where to find it in the outer scope
}

type varRef struct {
	kind  varKind
	index int
}

type varKind int

const (
	varLocal varKind = iota
	varUpval
	varGlobal
)

// CompileFunc compiles a KL top-level defun body to a BytecodeFunc object.
// params is the list of parameter symbols; body is the KL expression.
func CompileFunc(name string, params []Obj, body Obj) Obj {
	c := newCompiler(name, len(params), params, nil)
	c.compileExpr(body, true)
	c.fn.TypeHints = append([]TypeHint(nil), c.hints...)
	return makeBytecodeObj(c.fn, nil)
}

func newCompiler(name string, arity int, params []Obj, outer *klCompiler) *klCompiler {
	c := &klCompiler{
		fn: &BytecodeFunc{
			Name:    name,
			Arity:   arity,
			Nlocals: arity,
		},
		locals: make(map[Obj]int),
		outer:  outer,
	}
	for i, p := range params {
		c.locals[p] = i
	}
	return c
}

func (c *klCompiler) emit(op uint8, a, b int32) int {
	idx := len(c.fn.Code)
	c.fn.Code = append(c.fn.Code, Instr{Op: op, A: a, B: b})
	return idx
}

func (c *klCompiler) addConst(v Obj) int32 {
	for i, cv := range c.fn.Consts {
		if cv == v {
			return int32(i)
		}
	}
	c.fn.Consts = append(c.fn.Consts, v)
	return int32(len(c.fn.Consts) - 1)
}

func (c *klCompiler) newLocal(sym Obj) int32 {
	slot := int32(c.fn.Nlocals)
	c.fn.Nlocals++
	c.locals[sym] = int(slot)
	return slot
}

func (c *klCompiler) resolveVar(sym Obj) (varKind, int) {
	if idx, ok := c.locals[sym]; ok {
		return varLocal, idx
	}
	for i, uv := range c.upvals {
		if uv.sym == sym {
			return varUpval, i
		}
	}
	if c.outer != nil {
		outerKind, outerIdx := c.outer.resolveVar(sym)
		if outerKind == varGlobal {
			return varGlobal, 0
		}
		uvIdx := len(c.upvals)
		c.upvals = append(c.upvals, upvalInfo{
			sym:      sym,
			outerRef: varRef{kind: outerKind, index: outerIdx},
		})
		return varUpval, uvIdx
	}
	return varGlobal, 0
}

// patchJump patches a previously emitted jump instruction's A operand.
func (c *klCompiler) patchJump(instrIdx int) {
	target := len(c.fn.Code)
	c.fn.Code[instrIdx].A = int32(target - instrIdx - 1)
}

// compileExpr compiles a KL expression.  tail indicates whether this is
// in tail position (may emit OP_TAIL_CALL / OP_RETURN instead of OP_CALL).
func (c *klCompiler) compileExpr(form Obj, tail bool) {
	switch *form {
	case scmHeadNumber, scmHeadString, scmHeadBoolean, scmHeadNull:
		c.emit(OP_LOAD_CONST, c.addConst(form), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return

	case scmHeadSymbol:
		kind, idx := c.resolveVar(form)
		switch kind {
		case varLocal:
			c.emit(OP_LOAD_LOCAL, int32(idx), 0)
		case varUpval:
			c.emit(OP_LOAD_UPVAL, int32(idx), 0)
		case varGlobal:
			// In value position, a symbol not in scope evaluates to itself.
			c.emit(OP_LOAD_CONST, c.addConst(form), 0)
		}
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return

	case scmHeadPair:
		c.compileForm(form, tail)
		return

	default:
		// Vector, stream, procedure, native, etc. — treat as constants.
		c.emit(OP_LOAD_CONST, c.addConst(form), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}
}

// compileForm handles pair forms: special forms and function calls.
func (c *klCompiler) compileForm(form Obj, tail bool) {
	pair := mustPair(form)
	if !IsSymbol(pair.car) {
		// Complex function expression: (expr args...)
		c.compileCall(pair.car, pair.cdr, tail)
		return
	}
	head := pair.car
	rest := pair.cdr

	switch head {
	case symDefun:
		// (defun f (x...) body)
		name := car(rest)
		params := cadr(rest)
		body := caddr(rest)
		c.compileDefun(name, params, body)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}

	case symLambda:
		// (lambda x body)
		param := car(rest)
		body := cadr(rest)
		c.compileLambda([]Obj{param}, body)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}

	case symFreeze:
		// (freeze body)
		body := car(rest)
		c.compileLambda(nil, body)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}

	case symLet:
		// (let x val body)
		x := car(rest)
		val := cadr(rest)
		body := caddr(rest)
		c.compileLet(x, val, body, tail)

	case symIf:
		args := rest
		if listLength(args) == 3 {
			c.compileIf(car(args), cadr(args), caddr(args), tail)
		} else {
			// (if ...) with wrong arg count: fall through as a call
			c.compileCall(head, rest, tail)
		}

	case symAnd:
		// (and a b) => (if a b false)
		a := car(rest)
		b := cadr(rest)
		c.compileIf(a, b, False, tail)

	case symOr:
		// (or a b) => (if a true b)
		a := car(rest)
		b := cadr(rest)
		c.compileIf(a, True, b, tail)

	case symCond:
		c.compileCond(rest, tail)

	case symDo:
		// (do a b)
		a := car(rest)
		b := cadr(rest)
		c.compileExpr(a, false)
		c.emit(OP_POP, 0, 0)
		c.compileExpr(b, tail)

	case symType:
		// (type x T) is advisory metadata; compile x unchanged so malformed
		// annotations never introduce a new runtime error.
		if typeName := cadr(rest); IsSymbol(typeName) {
			name := GetSymbol(typeName)
			c.hints = append(c.hints, TypeHint{PC: len(c.fn.Code), Name: name, Kinds: KindSetForAnnotation(name), Source: TypeHintSourceAnnotation})
		}
		c.compileExpr(car(rest), tail)

	case symTrapError:
		// (trap-error body handler)
		// Lower to: (try-catch (freeze body) handler)
		body := car(rest)
		handler := cadr(rest)
		c.compileTrapError(body, handler, tail)

	default:
		// Regular function call with a symbol in head position.
		c.compileCall(head, rest, tail)
	}
}

func (c *klCompiler) compileDefun(name, params, body Obj) {
	paramSlice := ListToSlice(params)
	nameStr := mustSymbol(name).str
	// Thread c as outer so the body can close over the enclosing lexical scope,
	// matching the interpreter which does makeProcedure(..., env).
	inner := newCompiler(nameStr, len(paramSlice), paramSlice, c)
	inner.compileExpr(body, true)
	inner.fn.TypeHints = append([]TypeHint(nil), inner.hints...)

	defunSym := c.addConst(MakeSymbol("defun"))
	nameConst := c.addConst(name)
	c.emit(OP_LOAD_GLOBAL, defunSym, 0)
	c.emit(OP_LOAD_CONST, nameConst, 0)

	// If the body referenced outer variables, emit their values as upvalues
	// and create a closure; otherwise emit the bytecode func as a plain constant.
	for _, uv := range inner.upvals {
		switch uv.outerRef.kind {
		case varLocal:
			c.emit(OP_LOAD_LOCAL, int32(uv.outerRef.index), 0)
		case varUpval:
			c.emit(OP_LOAD_UPVAL, int32(uv.outerRef.index), 0)
		}
	}
	innerObj := makeBytecodeObj(inner.fn, nil)
	innerConst := c.addConst(innerObj)
	c.emit(OP_MAKE_CLOSURE, innerConst, int32(len(inner.upvals)))
	c.emit(OP_CALL, 2, 0)
}

func (c *klCompiler) compileLambda(params []Obj, body Obj) {
	inner := newCompiler("lambda", len(params), params, c)
	inner.compileExpr(body, true)
	inner.fn.TypeHints = append([]TypeHint(nil), inner.hints...)

	// The inner compiler may have discovered upvalues; emit loads for them.
	for _, uv := range inner.upvals {
		switch uv.outerRef.kind {
		case varLocal:
			c.emit(OP_LOAD_LOCAL, int32(uv.outerRef.index), 0)
		case varUpval:
			c.emit(OP_LOAD_UPVAL, int32(uv.outerRef.index), 0)
		}
	}

	innerObj := makeBytecodeObj(inner.fn, nil) // upvals populated at runtime
	innerConst := c.addConst(innerObj)
	c.emit(OP_MAKE_CLOSURE, innerConst, int32(len(inner.upvals)))
}

func isFreezeForm(o Obj) (Obj, bool) {
	ok, p := isPair(o)
	if !ok || p.car != symFreeze {
		return nil, false
	}
	return car(p.cdr), true
}

func walkThawUses(form, sym Obj) (thaws, others int) {
	if form == nil || form == Nil {
		return 0, 0
	}
	if form == sym {
		return 0, 1
	}
	ok, p := isPair(form)
	if !ok {
		return 0, 0
	}
	head, rest := p.car, p.cdr
	if head == MakeSymbol("thaw") {
		args := ListToSlice(rest)
		if len(args) == 1 && args[0] == sym {
			return 1, 0
		}
	}
	if head == symLambda {
		param := car(rest)
		if param == sym {
			return 0, 0
		}
		t, o := walkThawUses(cadr(rest), sym)
		return 0, t + o // A reference captured by another closure may escape.
	}
	if head == symFreeze {
		t, o := walkThawUses(car(rest), sym)
		return 0, t + o
	}
	if head == symDefun {
		for _, param := range ListToSlice(cadr(rest)) {
			if param == sym {
				return 0, 0
			}
		}
		t, o := walkThawUses(caddr(rest), sym)
		return 0, t + o
	}
	if head == symTrapError {
		// trap-error compiles its protected body as a zero-argument closure.
		t, o := walkThawUses(car(rest), sym)
		t2, o2 := walkThawUses(cadr(rest), sym)
		return t2, t + o + o2
	}
	if head == symCond {
		// Clauses are syntax, not calls (in particular, (thaw Go) is a
		// test named thaw and a value Go, not a thaw of Go).
		for _, clause := range ListToSlice(rest) {
			for _, expr := range ListToSlice(clause) {
				t, o := walkThawUses(expr, sym)
				thaws += t
				others += o
			}
		}
		return thaws, others
	}
	if head == symLet {
		x := car(rest)
		t, o := walkThawUses(cadr(rest), sym)
		if x == sym {
			return t, o
		}
		t2, o2 := walkThawUses(caddr(rest), sym)
		return t + t2, o + o2
	}
	t, o := walkThawUses(head, sym)
	for _, a := range ListToSlice(rest) {
		t1, o1 := walkThawUses(a, sym)
		t += t1
		o += o1
	}
	return t, o
}

func (c *klCompiler) compileLet(x, val, body Obj, tail bool) {
	// S37+ factoriser: (let Go (freeze Else) ... (thaw Go) ...).
	// If Go is only thawed locally, inline Else in its definition's scope.
	if fb, ok := isFreezeForm(val); ok {
		thaws, others := walkThawUses(body, x)
		if others == 0 && thaws > 0 {
			captured := make(map[Obj]int, len(c.locals))
			for sym, slot := range c.locals {
				captured[sym] = slot
			}
			oldIdx, hadOld := c.locals[x]
			slot := int(c.newLocal(x))
			if c.thunks == nil {
				c.thunks = make(map[int]inlineThunk)
			}
			c.thunks[slot] = inlineThunk{body: fb, locals: captured}
			c.compileExpr(body, tail)
			delete(c.thunks, slot)
			if hadOld {
				c.locals[x] = oldIdx
			} else {
				delete(c.locals, x)
			}
			return
		}
	}
	// Evaluate val.
	c.compileExpr(val, false)
	// Assign to a new local slot (or reuse if x is already a local).
	// We always allocate a new slot to handle shadowing correctly.
	oldIdx, hadOld := c.locals[x]
	slot := c.newLocal(x)
	c.emit(OP_STORE_LOCAL, slot, 0)
	// Compile body.
	c.compileExpr(body, tail)
	// Restore previous mapping (for shadowing).
	if hadOld {
		c.locals[x] = oldIdx
	} else {
		delete(c.locals, x)
	}
}

func (c *klCompiler) compileIf(cond, thenExpr, elseExpr Obj, tail bool) {
	c.compileExpr(cond, false)
	// Emit JUMP_FALSE with placeholder; patch after then-branch.
	jumpFalseIdx := c.emit(OP_JUMP_FALSE, 0, 0)
	c.compileExpr(thenExpr, tail)
	if !tail {
		// Emit JUMP over else-branch; patch after else.
		jumpOverIdx := c.emit(OP_JUMP, 0, 0)
		c.patchJump(jumpFalseIdx)
		c.compileExpr(elseExpr, tail)
		c.patchJump(jumpOverIdx)
	} else {
		// In tail position both branches already emit RETURN; just patch the false jump.
		c.patchJump(jumpFalseIdx)
		c.compileExpr(elseExpr, tail)
	}
}

func (c *klCompiler) compileCond(clauses Obj, tail bool) {
	if *clauses == scmHeadNull {
		// No matching clause: return Nil.
		c.emit(OP_LOAD_CONST, c.addConst(Nil), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}
	clause := car(clauses)
	rest := cdr(clauses)
	cond := car(clause)
	action := cadr(clause)

	if cond == True {
		// (cond (true action) ...) — unconditional
		c.compileExpr(action, tail)
		return
	}

	c.compileExpr(cond, false)
	jumpFalseIdx := c.emit(OP_JUMP_FALSE, 0, 0)
	c.compileExpr(action, tail)
	if !tail {
		jumpOverIdx := c.emit(OP_JUMP, 0, 0)
		c.patchJump(jumpFalseIdx)
		c.compileCond(rest, tail)
		c.patchJump(jumpOverIdx)
	} else {
		c.patchJump(jumpFalseIdx)
		c.compileCond(rest, tail)
	}
}

// compileTrapError lowers (trap-error body handler) to (try-catch (freeze body) handler).
// Stack layout for CALL/TAIL_CALL 2: [..., fn, arg1, arg2]
func (c *klCompiler) compileTrapError(body, handler Obj, tail bool) {
	// Recognised absence checks: trap-error is only recovering a missing
	// binding or vector slot, and the handler ignores the error object.
	// Port-performance.md: these are presence tests, not general exceptions.
	if def, ok := trapHandlerDefault(handler); ok {
		if c.tryCompileValueOr(body, def, tail) {
			return
		}
		if c.tryCompileVectorRefOr(body, def, tail) {
			return
		}
	}
	tryCatchSym := c.addConst(MakeSymbol("try-catch"))
	c.emit(OP_LOAD_GLOBAL, tryCatchSym, 0)
	c.compileLambda(nil, body)
	c.compileExpr(handler, false)
	// trap-error must NOT be a tail call: the panic/recover in try-catch
	// needs a real Go stack frame to catch from.
	c.emit(OP_CALL, 2, 0)
	if tail {
		c.emit(OP_RETURN, 0, 0)
	}
}

func trapHandlerDefault(handler Obj) (Obj, bool) {
	ok, p := isPair(handler)
	if !ok || p.car != MakeSymbol("lambda") {
		return nil, false
	}
	param := car(p.cdr)
	body := cadr(p.cdr)
	if formContainsSym(body, param) {
		return nil, false
	}
	return body, true
}

func formContainsSym(form, sym Obj) bool {
	if form == sym {
		return true
	}
	ok, p := isPair(form)
	if !ok {
		return false
	}
	return formContainsSym(p.car, sym) || formContainsSym(p.cdr, sym)
}

func callForm(form Obj, name string, arity int) ([]Obj, bool) {
	ok, p := isPair(form)
	if !ok || !IsSymbol(p.car) || GetSymbol(p.car) != name {
		return nil, false
	}
	var args []Obj
	for l := p.cdr; l != Nil && l != nil; {
		aok, ap := isPair(l)
		if !aok {
			return nil, false
		}
		args = append(args, ap.car)
		l = ap.cdr
	}
	if len(args) != arity {
		return nil, false
	}
	return args, true
}

func (c *klCompiler) tryCompileValueOr(body, def Obj, tail bool) bool {
	args, ok := callForm(body, "value", 1)
	if !ok {
		return false
	}
	c.emit(OP_LOAD_GLOBAL, c.addConst(MakeSymbol("_kl.value/or")), 0)
	c.compileExpr(args[0], false)
	c.compileLambda(nil, def)
	c.emit(OP_CALL, 2, 0)
	if tail {
		c.emit(OP_RETURN, 0, 0)
	}
	return true
}

func (c *klCompiler) tryCompileVectorRefOr(body, def Obj, tail bool) bool {
	args, ok := callForm(body, "<-vector", 2)
	if !ok {
		return false
	}
	c.emit(OP_LOAD_GLOBAL, c.addConst(MakeSymbol("_kl.<-vector/or")), 0)
	c.compileExpr(args[0], false)
	c.compileExpr(args[1], false)
	c.compileLambda(nil, def)
	c.emit(OP_CALL, 3, 0)
	if tail {
		c.emit(OP_RETURN, 0, 0)
	}
	return true
}

// intrinsicOp maps a 2-arg primitive symbol to a fast-path opcode.
// Returns 0 if no fast path exists.
func intrinsicOp2(sym Obj) uint8 {
	if IsSymbol(sym) && GetSymbol(sym) == "/" {
		return OP_DIV
	}
	switch sym {
	case symAdd:
		return OP_ADD
	case symSub:
		return OP_SUB
	case symMul:
		return OP_MUL
	case symLT:
		return OP_LT
	case symLE:
		return OP_LE
	case symGT:
		return OP_GT
	case symGE:
		return OP_GE
	case symNumEq:
		return OP_EQ
	}
	return 0
}

func guardedPrimitiveArity(sym Obj) int {
	if !IsSymbol(sym) || !HasCanonicalPrimitiveBinding(sym) {
		return 0
	}
	switch GetSymbol(sym) {
	case "number?", "integer?", "string?", "symbol?", "cons?", "absvector?", "variable?", "empty?", "boolean?", "hd", "tl", "tlstr", "string->n", "n->string", "absvector":
		return 1
	case "cn", "pos", "cons", "<-address":
		return 2
	case "address->":
		return 3
	}
	return 0
}

// foldFixnum2 folds only tagged-integer literals. No type information is
// inferred through locals or calls, so floats and other dynamic numeric
// boundaries remain on the existing runtime paths.
func (c *klCompiler) constValue(o Obj) (Obj, bool) {
	if o == nil {
		return nil, false
	}
	switch *o {
	case scmHeadNumber, scmHeadString, scmHeadBoolean, scmHeadNull:
		return o, true
	case scmHeadSymbol:
		kind, _ := c.resolveVar(o)
		if kind == varGlobal {
			return o, true
		}
	}
	return nil, false
}

func foldFixnum2(sym Obj, x, y Obj) (Obj, bool) {
	if !isFixnum(x) || !isFixnum(y) {
		return nil, false
	}
	a, b := fixnum(x), fixnum(y)
	switch sym {
	case symAdd:
		return MakeInteger(a + b), true
	case symSub:
		return MakeInteger(a - b), true
	case symMul:
		return MakeInteger(a * b), true
	}
	return nil, false
}

func (c *klCompiler) compileCall(fn Obj, args Obj, tail bool) {
	nArgs := 0
	argList := make([]Obj, 0, 4)
	for l := args; *l == scmHeadPair; l = cdr(l) {
		argList = append(argList, car(l))
		nArgs++
	}

	// Guarded primitives carry the canonical symbol in the instruction. The
	// VM executes a typed fast path only while that binding remains unchanged;
	// otherwise it performs the ordinary dynamic call.
	if IsSymbol(fn) {
		n := guardedPrimitiveArity(fn)
		if n != 0 && n == nArgs {
			for _, a := range argList {
				c.compileExpr(a, false)
			}
			c.emit(OP_GUARDED_PRIM, int32(n), c.addConst(fn))
			if tail {
				c.emit(OP_RETURN, 0, 0)
			}
			return
		}
	}

	// ---- (thaw (freeze E)) => E ----
	if IsSymbol(fn) && nArgs == 1 && GetSymbol(fn) == "thaw" {
		if fb, ok := isFreezeForm(argList[0]); ok {
			c.compileExpr(fb, tail)
			return
		}
		if slot, local := c.locals[argList[0]]; local && IsSymbol(argList[0]) {
			if thunk, ok := c.thunks[slot]; ok {
				locals := c.locals
				c.locals = thunk.locals
				c.compileExpr(thunk.body, tail)
				c.locals = locals
				return
			}
		}
	}

	// ---- fold (= literal literal) ----
	if nArgs == 2 && fn == symNumEq && HasCanonicalPrimitiveBinding(fn) {
		if a, ok := c.constValue(argList[0]); ok {
			if b, ok := c.constValue(argList[1]); ok {
				c.emit(OP_LOAD_CONST, c.addConst(equal(a, b)), 0)
				if tail {
					c.emit(OP_RETURN, 0, 0)
				}
				return
			}
		}
	}

	// ---- intern of a string literal: fold to the interned symbol/bool ----
	if IsSymbol(fn) && nArgs == 1 && GetSymbol(fn) == "intern" && HasCanonicalPrimitiveBinding(fn) && IsString(argList[0]) {
		c.emit(OP_LOAD_CONST, c.addConst(PrimIntern(argList[0])), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}

	// ---- 1-arg intrinsics ----
	if IsSymbol(fn) && nArgs == 1 && fn == symNot && HasCanonicalPrimitiveBinding(fn) {
		c.compileExpr(argList[0], false)
		c.emit(OP_NOT, 0, c.addConst(fn))
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}

	// ---- 2-arg arithmetic intrinsics ----
	if IsSymbol(fn) && nArgs == 2 && HasCanonicalPrimitiveBinding(fn) {
		if op := intrinsicOp2(fn); op != 0 {
			if folded, ok := foldFixnum2(fn, argList[0], argList[1]); ok {
				// Keep the original operands and symbol so rebinding after
				// compilation still takes the dynamic path.
				// Folded constants retain their canonical operation in B without
				// growing the constant pool (the original operands are in C/D).
				c.emit(OP_GUARDED_CONST, c.addConst(folded), -int32(op))
				idx := len(c.fn.Code) - 1
				c.fn.Code[idx].C = int32(fixnum(argList[0]))
				c.fn.Code[idx].D = int32(fixnum(argList[1]))
				if tail {
					c.emit(OP_RETURN, 0, 0)
				}
				return
			}
			c.compileExpr(argList[0], false)
			c.compileExpr(argList[1], false)
			// Keep the symbol in the instruction so the VM can detect a
			// redefinition after compilation and fall back to the dynamic call.
			c.emit(op, 0, c.addConst(fn))
			if tail {
				c.emit(OP_RETURN, 0, 0)
			}
			return
		}
	}

	// ---- Self-tail-call optimization ----
	if tail && IsSymbol(fn) && mustSymbol(fn).str == c.fn.Name {
		kind, _ := c.resolveVar(fn)
		if kind == varGlobal && nArgs == c.fn.Arity {
			// It's a self-call in tail position: emit args then OP_SELF_TAIL_CALL.
			for _, a := range argList {
				c.compileExpr(a, false)
			}
			c.emit(OP_SELF_TAIL_CALL, int32(nArgs), 0)
			return
		}
	}

	// ---- General call ----
	if IsSymbol(fn) {
		// Match interpreter precedence (evalFunction): if the symbol has a
		// global function binding, that takes priority over a same-named local.
		// This handles cases like a parameter named 'cons' shadowing the global.
		sym := mustSymbol(fn)
		if sym.function != nil {
			c.emit(OP_LOAD_GLOBAL, c.addConst(fn), 0)
		} else {
			kind, idx := c.resolveVar(fn)
			switch kind {
			case varLocal:
				c.emit(OP_LOAD_LOCAL, int32(idx), 0)
			case varUpval:
				c.emit(OP_LOAD_UPVAL, int32(idx), 0)
			case varGlobal:
				c.emit(OP_LOAD_GLOBAL, c.addConst(fn), 0)
			}
		}
	} else {
		c.compileExpr(fn, false)
	}

	for _, a := range argList {
		c.compileExpr(a, false)
	}

	if tail {
		c.emit(OP_TAIL_CALL, int32(nArgs), 0)
	} else {
		c.emit(OP_CALL, int32(nArgs), 0)
	}
}
