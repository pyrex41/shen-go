package kl

// Inlined primitives for the closure executor (see closure.go). Every node has
// the shape: evaluate operands, test clDirty, fast path, else the exact
// semantics of the VM instruction it replaces, which ends in the ordinary
// dynamic call when the primitive is no longer canonical.

import "unsafe"

func clU(o Obj) uintptr {
	return uintptr(unsafe.Pointer(o)) - uintptr(unsafe.Pointer(&addrForFixnum))
}

func clIsFix(o Obj) bool { return clU(o) < fixnumCount }

func clFromU(u uintptr) Obj { return Obj(unsafe.Add(unsafe.Pointer(&addrForFixnum), u)) }

func clTruth(o Obj) bool {
	if o == True {
		return true
	}
	if o != False {
		panic(MakeError("if requires a boolean"))
	}
	return false
}

// clIdentityConst reports whether (= X k) is exactly X == k for every X: equal
// compares fixnums, symbols, booleans and () by identity.
func clIdentityConst(k Obj) bool {
	if isFixnum(k) {
		return true
	}
	switch *k {
	case scmHeadSymbol, scmHeadBoolean, scmHeadNull:
		return true
	}
	return false
}

// clDyn is the dynamic call of the current binding, as vmDynamicCall.
func clDyn(ctl *ControlFlow, sym Obj, args ...Obj) Obj {
	s := mustSymbol(sym)
	if s.function == nil {
		clUnbound(s)
	}
	return Call(ctl, s.function, args...)
}

// clGP is OP_GUARDED_PRIM: the typed fast path while canonical, and the
// ordinary call when it declines, so errors are the primitive's own.
func clGP(ctl *ControlFlow, sym Obj, id int32, args ...Obj) Obj {
	var inline [4]vmSlot
	slots := inline[:0]
	for _, a := range args {
		slots = append(slots, slotFromObj(a))
	}
	if r, ok := vmGuardedPrimitive(sym, id+1, slots); ok {
		return r.objValue()
	}
	return clDyn(ctl, sym, args...)
}

func clSlowEq(ctl *ControlFlow, sym, x, y Obj) bool {
	if clGuardOK(sym) {
		return x == y || equal(x, y) == True
	}
	return clTruth(clDyn(ctl, sym, x, y))
}

func clEqObj(x, y Obj) bool { return x == y || equal(x, y) == True }

func clTestEq(sym Obj, x, y clx) ctest {
	if x.kind == clLocal && y.kind == clConst {
		off, k := x.off, y.k
		if clIdentityConst(k) {
			return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
				v := *(*Obj)(unsafe.Add(fp, off))
				if clDirty {
					return clSlowEq(ctl, sym, v, k)
				}
				return v == k
			}
		}
		return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
			v := *(*Obj)(unsafe.Add(fp, off))
			if clDirty {
				return clSlowEq(ctl, sym, v, k)
			}
			return clEqObj(v, k)
		}
	}
	a, b := x.op(), y.op()
	return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
		v := clGet(&a, ctl, fp)
		w := clGet(&b, ctl, fp)
		if clDirty {
			return clSlowEq(ctl, sym, v, w)
		}
		return clEqObj(v, w)
	}
}

func clCmpSem(op uint8, x, y Obj) bool {
	if u, t := clU(x), clU(y); u|t < fixnumCount {
		switch op {
		case OP_LT:
			return u < t
		case OP_LE:
			return u <= t
		case OP_GT:
			return u > t
		}
		return u >= t
	}
	switch op {
	case OP_LT:
		return numCmpLT(x, y) == True
	case OP_LE:
		return numCmpLE(x, y) == True
	case OP_GT:
		return numCmpLT(y, x) == True
	}
	return numCmpLE(y, x) == True
}

func clSlowCmp(ctl *ControlFlow, op uint8, sym, x, y Obj) bool {
	if clGuardOK(sym) {
		return clCmpSem(op, x, y)
	}
	return clTruth(clDyn(ctl, sym, x, y))
}

func clTestCmp(op uint8, sym Obj, x, y clx) ctest {
	a, b := x.op(), y.op()
	switch op {
	case OP_LT:
		return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			u, t := clU(v), clU(w)
			if u|t < fixnumCount && !clDirty {
				return u < t
			}
			return clSlowCmp(ctl, op, sym, v, w)
		}
	case OP_LE:
		return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			u, t := clU(v), clU(w)
			if u|t < fixnumCount && !clDirty {
				return u <= t
			}
			return clSlowCmp(ctl, op, sym, v, w)
		}
	case OP_GT:
		return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			u, t := clU(v), clU(w)
			if u|t < fixnumCount && !clDirty {
				return u > t
			}
			return clSlowCmp(ctl, op, sym, v, w)
		}
	}
	return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
		v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
		u, t := clU(v), clU(w)
		if u|t < fixnumCount && !clDirty {
			return u >= t
		}
		return clSlowCmp(ctl, op, sym, v, w)
	}
}

// test compiles form as a condition, with the VM's outcome for a non-boolean.
func (c *clCompiler) test(form Obj) ctest {
	if ok, p := isPair(form); ok && IsSymbol(p.car) {
		args := ListToSlice(p.cdr)
		switch {
		case p.car == symAnd:
			a, b := c.test(car(p.cdr)), c.test(cadr(p.cdr))
			return func(ctl *ControlFlow, fp unsafe.Pointer) bool { return a(ctl, fp) && b(ctl, fp) }
		case p.car == symOr:
			a, b := c.test(car(p.cdr)), c.test(cadr(p.cdr))
			return func(ctl *ControlFlow, fp unsafe.Pointer) bool { return a(ctl, fp) || b(ctl, fp) }
		case len(args) == 2 && HasCanonicalPrimitiveBinding(p.car):
			if p.car == symNumEq && !c.bothConst(args[0], args[1]) {
				clMark(p.car)
				return clTestEq(p.car, c.expr(args[0], false), c.expr(args[1], false))
			}
			if op := intrinsicOp2(p.car); op >= OP_LT && op <= OP_GE {
				clMark(p.car)
				return clTestCmp(op, p.car, c.expr(args[0], false), c.expr(args[1], false))
			}
		case len(args) == 1 && GetSymbol(p.car) == "cons?" && guardedPrimitiveArity(p.car) == 1:
			clMark(p.car)
			a, sym := c.expr(args[0], false).op(), p.car
			return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
				v := clGet(&a, ctl, fp)
				if clDirty {
					return clTruth(clGP(ctl, sym, gpConsp, v))
				}
				return !clIsFix(v) && *v == scmHeadPair
			}
		}
	}
	x := c.expr(form, false).fn
	return func(ctl *ControlFlow, fp unsafe.Pointer) bool {
		v := x(ctl, fp)
		if v == True {
			return true
		}
		if v != False {
			panic(MakeError("if requires a boolean"))
		}
		return false
	}
}

func clNot(sym Obj, x clx) clx {
	a := x.op()
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		v := clGet(&a, ctl, fp)
		if clDirty && !clGuardOK(sym) {
			return clDyn(ctl, sym, v)
		}
		if v == True {
			return False
		}
		if v == False {
			return True
		}
		panic(MakeError("not: expected boolean"))
	})
}

func clFolded(sym, k, x, y Obj) clx {
	return clGen(func(ctl *ControlFlow, _ unsafe.Pointer) Obj {
		if clDirty && !clGuardOK(sym) {
			return clDyn(ctl, sym, x, y)
		}
		return k
	})
}

func clAddSem(ctl *ControlFlow, sym, x, y Obj) Obj {
	if clDirty && !clGuardOK(sym) {
		return clDyn(ctl, sym, x, y)
	}
	return numAdd(x, y)
}

func clSubSem(ctl *ControlFlow, sym, x, y Obj) Obj {
	if clDirty && !clGuardOK(sym) {
		return clDyn(ctl, sym, x, y)
	}
	return numSub(x, y)
}

func clMulSem(ctl *ControlFlow, sym, x, y Obj) Obj {
	if clDirty && !clGuardOK(sym) {
		return clDyn(ctl, sym, x, y)
	}
	return numMul(x, y)
}

func clDivSem(ctl *ControlFlow, sym, x, y Obj) Obj {
	if clDirty && !clGuardOK(sym) {
		return clDyn(ctl, sym, x, y)
	}
	if IsNumber(x) && IsNumber(y) {
		return MakeNumber(TypedDivideValue(GetNumber(x), GetNumber(y)))
	}
	return PrimNumberDivide(x, y)
}

// clArith compiles a 2-argument intrinsic.
func clArith(sym Obj, x, y clx) clx {
	op := intrinsicOp2(sym)
	switch op {
	case OP_ADD:
		if y.kind == clConst && isFixnum(y.k) {
			return clAddK(sym, x, y.k, fixnum(y.k))
		}
		if x.kind == clConst && isFixnum(x.k) {
			return clAddK(sym, y, x.k, fixnum(x.k))
		}
		a, b := x.op(), y.op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			u, t := clU(v), clU(w)
			if u|t < fixnumCount && !clDirty {
				if r := u + t + uintptr(fixnumOffset); r < fixnumCount {
					return clFromU(r)
				}
			}
			return clAddSem(ctl, sym, v, w)
		})
	case OP_SUB:
		if y.kind == clConst && isFixnum(y.k) {
			return clSubK(sym, x, y.k, fixnum(y.k))
		}
		a, b := x.op(), y.op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			u, t := clU(v), clU(w)
			if u|t < fixnumCount && !clDirty {
				if r := u - t - uintptr(fixnumOffset); r < fixnumCount {
					return clFromU(r)
				}
			}
			return clSubSem(ctl, sym, v, w)
		})
	case OP_MUL:
		a, b := x.op(), y.op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			return clMulSem(ctl, sym, v, w)
		})
	case OP_DIV:
		a, b := x.op(), y.op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			return clDivSem(ctl, sym, v, w)
		})
	}
	// A comparison used as a value returns whatever a rebinding returns; only
	// a comparison used as a condition demands a boolean.
	a, b := x.op(), y.op()
	if op == OP_EQ {
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
			if clDirty && !clGuardOK(sym) {
				return clDyn(ctl, sym, v, w)
			}
			if clEqObj(v, w) {
				return True
			}
			return False
		})
	}
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		v, w := clGet(&a, ctl, fp), clGet(&b, ctl, fp)
		if clDirty && !clGuardOK(sym) {
			return clDyn(ctl, sym, v, w)
		}
		if clCmpSem(op, v, w) {
			return True
		}
		return False
	})
}

// fixnumOffset turns the sum of two encoded fixnums back into an encoding:
// (a-min)+(b-min)+min = (a+b)-min; subtraction uses its negation.
const fixnumOffset = uint64(1<<64 + fixnumMin)

func clAddK(sym Obj, x clx, k Obj, d int) clx {
	delta := uintptr(d)
	if x.kind == clLocal {
		off := x.off
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v := *(*Obj)(unsafe.Add(fp, off))
			if u := clU(v); u < fixnumCount && !clDirty {
				if r := u + delta; r < fixnumCount {
					return clFromU(r)
				}
			}
			return clAddSem(ctl, sym, v, k)
		})
	}
	f := x.fn
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		v := f(ctl, fp)
		if u := clU(v); u < fixnumCount && !clDirty {
			if r := u + delta; r < fixnumCount {
				return clFromU(r)
			}
		}
		return clAddSem(ctl, sym, v, k)
	})
}

func clSubK(sym Obj, x clx, k Obj, d int) clx {
	delta := uintptr(-d)
	if x.kind == clLocal {
		off := x.off
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v := *(*Obj)(unsafe.Add(fp, off))
			if u := clU(v); u < fixnumCount && !clDirty {
				if r := u + delta; r < fixnumCount {
					return clFromU(r)
				}
			}
			return clSubSem(ctl, sym, v, k)
		})
	}
	f := x.fn
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		v := f(ctl, fp)
		if u := clU(v); u < fixnumCount && !clDirty {
			if r := u + delta; r < fixnumCount {
				return clFromU(r)
			}
		}
		return clSubSem(ctl, sym, v, k)
	})
}

// clGuarded compiles a primitive of OP_GUARDED_PRIM.
func clGuarded(sym Obj, args []clx) clx {
	id := guardedPrimID(sym)
	switch id {
	case gpHd, gpTl:
		isHd := id == gpHd
		if args[0].kind == clLocal {
			off := args[0].off
			if isHd {
				return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
					v := *(*Obj)(unsafe.Add(fp, off))
					if !clIsFix(v) && *v == scmHeadPair && !clDirty {
						return (*scmPair)(unsafe.Pointer(v)).car
					}
					return clGP(ctl, sym, id, v)
				})
			}
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				v := *(*Obj)(unsafe.Add(fp, off))
				if !clIsFix(v) && *v == scmHeadPair && !clDirty {
					return (*scmPair)(unsafe.Pointer(v)).cdr
				}
				return clGP(ctl, sym, id, v)
			})
		}
		f := args[0].fn
		if isHd {
			return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
				v := f(ctl, fp)
				if !clIsFix(v) && *v == scmHeadPair && !clDirty {
					return (*scmPair)(unsafe.Pointer(v)).car
				}
				return clGP(ctl, sym, id, v)
			})
		}
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v := f(ctl, fp)
			if !clIsFix(v) && *v == scmHeadPair && !clDirty {
				return (*scmPair)(unsafe.Pointer(v)).cdr
			}
			return clGP(ctl, sym, id, v)
		})
	case gpCons:
		a, b := args[0].op(), args[1].op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			x := clGet(&a, ctl, fp)
			y := clGet(&b, ctl, fp)
			if clDirty {
				return clGP(ctl, sym, id, x, y)
			}
			return cons(x, y)
		})
	case gpConsp:
		a := args[0].op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v := clGet(&a, ctl, fp)
			if clDirty {
				return clGP(ctl, sym, id, v)
			}
			if !clIsFix(v) && *v == scmHeadPair {
				return True
			}
			return False
		})
	case gpVectorGet:
		a, b := args[0].op(), args[1].op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v := clGet(&a, ctl, fp)
			i := clGet(&b, ctl, fp)
			if !clDirty && !clIsFix(v) && *v == scmHeadVector && clIsFix(i) {
				vec := (*scmVector)(unsafe.Pointer(v)).vector
				n := int(clU(i)) + fixnumMin
				if uint(n) < uint(len(vec)) {
					if o := vec[n]; o != nil {
						return o
					}
				}
			}
			return clGP(ctl, sym, id, v, i)
		})
	case gpVectorSet:
		a, b, d := args[0].op(), args[1].op(), args[2].op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			v := clGet(&a, ctl, fp)
			i := clGet(&b, ctl, fp)
			x := clGet(&d, ctl, fp)
			if !clDirty && !clIsFix(v) && *v == scmHeadVector && clIsFix(i) {
				vec := (*scmVector)(unsafe.Pointer(v)).vector
				n := int(clU(i)) + fixnumMin
				if uint(n) < uint(len(vec)) {
					vec[n] = x
					return v
				}
			}
			return clGP(ctl, sym, id, v, i, x)
		})
	}
	switch len(args) {
	case 1:
		a := args[0].op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			return clGP(ctl, sym, id, clGet(&a, ctl, fp))
		})
	case 2:
		a, b := args[0].op(), args[1].op()
		return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
			x := clGet(&a, ctl, fp)
			return clGP(ctl, sym, id, x, clGet(&b, ctl, fp))
		})
	}
	a, b, d := args[0].op(), args[1].op(), args[2].op()
	return clGen(func(ctl *ControlFlow, fp unsafe.Pointer) Obj {
		x := clGet(&a, ctl, fp)
		y := clGet(&b, ctl, fp)
		return clGP(ctl, sym, id, x, y, clGet(&d, ctl, fp))
	})
}
