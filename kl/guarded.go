package kl

import "unsafe"

// Guarded primitives for generated code. Each honours a rebinding exactly as
// the generated guard it replaces did (TypedIREnabled() &&
// HasCanonicalPrimitiveBinding(sym), else a dynamic call of the symbol's
// current function), but in one function rather than an expansion at each
// site, and with the checks generated code cannot need: a symbol here is known
// to be a symbol, and known to have a canonical binding. The check is the
// symbol's two words and the mode, and because it loads function each time a
// rebinding takes effect at the next call.
//
// typedIRMode is resolved at package init below, so a mode of 0 only appears
// after ResetTypedIRModeForTest; the slow paths resolve it.

func symCell(name string) *scmSymbol { return mustSymbol(MakeSymbol(name)) }

var (
	gsEq      = symCell("=")
	gsHd      = symCell("hd")
	gsTl      = symCell("tl")
	gsCons    = symCell("cons")
	gsConsP   = symCell("cons?")
	gsNot     = symCell("not")
	gsAdd     = symCell("+")
	gsSub     = symCell("-")
	gsMul     = symCell("*")
	gsDiv     = symCell("/")
	gsLt      = symCell("<")
	gsLe      = symCell("<=")
	gsGt      = symCell(">")
	gsGe      = symCell(">=")
	gsVecGet  = symCell("<-address")
	gsVecSet  = symCell("address->")
	gsVectorP = symCell("absvector?")
	gsValue   = symCell("value")
)

func init() { TypedIREnabled() }

func primIntact(s *scmSymbol) bool {
	return s.function == s.canonical && typedIRMode.Load() == 1
}

// guardedSlow1/2/3 is the path past a failed quick check: the full guard, then
// the symbol's current binding, with the error PrimFunc raises if it has none.
func guardedSlow1(e *ControlFlow, s *scmSymbol, prim func(Obj) Obj, a Obj) Obj {
	sym := Obj(&s.scmHead)
	if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym) {
		return prim(a)
	}
	return Call1(e, SymFn(sym), a)
}

func guardedSlow2(e *ControlFlow, s *scmSymbol, prim func(Obj, Obj) Obj, a, b Obj) Obj {
	sym := Obj(&s.scmHead)
	if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym) {
		return prim(a, b)
	}
	return Call2(e, SymFn(sym), a, b)
}

func guardedSlow3(e *ControlFlow, s *scmSymbol, prim func(Obj, Obj, Obj) Obj, a, b, c Obj) Obj {
	sym := Obj(&s.scmHead)
	if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym) {
		return prim(a, b, c)
	}
	return Call3(e, SymFn(sym), a, b, c)
}

// PrimGuard1/2/3 are the guarded call of any other primitive: sym is its
// symbol and prim its canonical implementation.
func PrimGuard1(e *ControlFlow, sym Obj, prim func(Obj) Obj, a Obj) Obj {
	s := (*scmSymbol)(unsafe.Pointer(sym))
	if s.canonical != nil && primIntact(s) {
		return prim(a)
	}
	return guardedSlow1(e, s, prim, a)
}

func PrimGuard2(e *ControlFlow, sym Obj, prim func(Obj, Obj) Obj, a, b Obj) Obj {
	s := (*scmSymbol)(unsafe.Pointer(sym))
	if s.canonical != nil && primIntact(s) {
		return prim(a, b)
	}
	return guardedSlow2(e, s, prim, a, b)
}

func PrimGuard3(e *ControlFlow, sym Obj, prim func(Obj, Obj, Obj) Obj, a, b, c Obj) Obj {
	s := (*scmSymbol)(unsafe.Pointer(sym))
	if s.canonical != nil && primIntact(s) {
		return prim(a, b, c)
	}
	return guardedSlow3(e, s, prim, a, b, c)
}

func GEq(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsEq) {
		if a == b {
			return True
		}
		return equal(a, b)
	}
	return guardedSlow2(e, gsEq, PrimEqual, a, b)
}

func GHd(e *ControlFlow, a Obj) Obj {
	if primIntact(gsHd) {
		return PrimHead(a)
	}
	return guardedSlow1(e, gsHd, PrimHead, a)
}

func GTl(e *ControlFlow, a Obj) Obj {
	if primIntact(gsTl) {
		return PrimTail(a)
	}
	return guardedSlow1(e, gsTl, PrimTail, a)
}

func GCons(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsCons) {
		return cons(a, b)
	}
	return guardedSlow2(e, gsCons, PrimCons, a, b)
}

func GConsP(e *ControlFlow, a Obj) Obj {
	if primIntact(gsConsP) {
		return PrimIsPair(a)
	}
	return guardedSlow1(e, gsConsP, PrimIsPair, a)
}

func GNot(e *ControlFlow, a Obj) Obj {
	if primIntact(gsNot) {
		return PrimNot(a)
	}
	return guardedSlow1(e, gsNot, PrimNot, a)
}

// Fixnum arithmetic is exact in both int and float64: a sum, difference or
// product of two values under 2^25 fits in 2^53, and MakeNumber of the float64
// result is MakeInteger of the int result (including a -0 product, which it
// makes fixnum 0). So the int path returns what the float64 path returns.

func GAdd(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsAdd) {
		if isFixnum(a) && isFixnum(b) {
			return MakeInteger(fixnum(a) + fixnum(b))
		}
		return PrimNumberAdd(a, b)
	}
	return guardedSlow2(e, gsAdd, PrimNumberAdd, a, b)
}

func GSub(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsSub) {
		if isFixnum(a) && isFixnum(b) {
			return MakeInteger(fixnum(a) - fixnum(b))
		}
		return PrimNumberSubtract(a, b)
	}
	return guardedSlow2(e, gsSub, PrimNumberSubtract, a, b)
}

func GMul(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsMul) {
		if isFixnum(a) && isFixnum(b) {
			return MakeInteger(fixnum(a) * fixnum(b))
		}
		return PrimNumberMultiply(a, b)
	}
	return guardedSlow2(e, gsMul, PrimNumberMultiply, a, b)
}

func GDiv(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsDiv) {
		return PrimNumberDivide(a, b)
	}
	return guardedSlow2(e, gsDiv, PrimNumberDivide, a, b)
}

func boolObj(b bool) Obj {
	if b {
		return True
	}
	return False
}

func GLt(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsLt) {
		if isFixnum(a) && isFixnum(b) {
			return boolObj(uintptr(unsafe.Pointer(a)) < uintptr(unsafe.Pointer(b)))
		}
		return PrimLessThan(a, b)
	}
	return guardedSlow2(e, gsLt, PrimLessThan, a, b)
}

func GLe(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsLe) {
		if isFixnum(a) && isFixnum(b) {
			return boolObj(uintptr(unsafe.Pointer(a)) <= uintptr(unsafe.Pointer(b)))
		}
		return PrimLessEqual(a, b)
	}
	return guardedSlow2(e, gsLe, PrimLessEqual, a, b)
}

func GGt(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsGt) {
		if isFixnum(a) && isFixnum(b) {
			return boolObj(uintptr(unsafe.Pointer(a)) > uintptr(unsafe.Pointer(b)))
		}
		return PrimGreatThan(a, b)
	}
	return guardedSlow2(e, gsGt, PrimGreatThan, a, b)
}

func GGe(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsGe) {
		if isFixnum(a) && isFixnum(b) {
			return boolObj(uintptr(unsafe.Pointer(a)) >= uintptr(unsafe.Pointer(b)))
		}
		return PrimGreatEqual(a, b)
	}
	return guardedSlow2(e, gsGe, PrimGreatEqual, a, b)
}

func GVecGet(e *ControlFlow, a, b Obj) Obj {
	if primIntact(gsVecGet) {
		return PrimVectorGet(a, b)
	}
	return guardedSlow2(e, gsVecGet, PrimVectorGet, a, b)
}

func GVecSet(e *ControlFlow, a, b, c Obj) Obj {
	if primIntact(gsVecSet) {
		return PrimVectorSet(a, b, c)
	}
	return guardedSlow3(e, gsVecSet, PrimVectorSet, a, b, c)
}

func GVectorP(e *ControlFlow, a Obj) Obj {
	if primIntact(gsVectorP) {
		return PrimIsVector(a)
	}
	return guardedSlow1(e, gsVectorP, PrimIsVector, a)
}

func GValue(e *ControlFlow, a Obj) Obj {
	if primIntact(gsValue) {
		return PrimValue(a)
	}
	return guardedSlow1(e, gsValue, PrimValue, a)
}
