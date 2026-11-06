package kl

import "unsafe"

// Native Prolog runtime: the kernel's prolog.kl helpers that every compiled
// defprolog clause and the type checker call, as direct natives. Each is the
// same computation as its defun, through the same primitives, so the values and
// error messages are the kernel's; kl/equiv.json checks them against the
// kernel's own definitions.
//
// These natives call each other and the primitives without going through the
// symbols, so they stand for a whole group of definitions. Rebinding any
// function in that group (prologDeps) rebinds every native in it back to the
// kernel's compiled definition (restorePrologKernel), and from then on every
// call in the group goes through its symbol, as in the kernel. A native object
// a program already holds, from (fn shen.lazyderef) say, keeps its behaviour.

var (
	symPvarNull = MakeSymbol("shen.-null-")
	symInfs     = MakeSymbol("shen.*infs*")
	one         = MakeInteger(1)

	prologNatives = map[string]Obj{}
	prologInstall []*scmSymbol
	prologArmed   []*scmSymbol
)

// prologDeps are the functions the natives' bodies call in the kernel, beyond
// the natives themselves.
var prologDeps = []string{
	"hd", "tl", "cons", "cons?", "=", "not", "+", "-", "absvector", "absvector?",
	"<-address", "address->", "value", "set", "thaw", "shen.pvar?",
}

// overrideDirect binds a kernel function to a direct native. Unlike
// overridePrimitive, the object is new at each install.
func overrideDirect(name string, arity int, fn any) {
	BindSymbolFunc(MakeSymbol(name), makeDirect(arity, fn))
}

func makeDirect(arity int, fn any) Obj {
	switch f := fn.(type) {
	case func(*ControlFlow) Obj:
		return MakeNative0(f)
	case func(*ControlFlow, Obj) Obj:
		return MakeNative1(f)
	case func(*ControlFlow, Obj, Obj) Obj:
		return MakeNative2(f)
	case func(*ControlFlow, Obj, Obj, Obj) Obj:
		return MakeNative3(f)
	case func(*ControlFlow, Obj, Obj, Obj, Obj) Obj:
		return MakeNative4(f)
	case func(*ControlFlow, Obj, Obj, Obj, Obj, Obj) Obj:
		return MakeNative5(f)
	case func(*ControlFlow, Obj, Obj, Obj, Obj, Obj, Obj) Obj:
		return MakeNative6(f)
	}
	panic("makeDirect: " + MustNative(MakeNative(nil, arity)).name)
}

// overrideProlog binds a member of the native Prolog group, remembering the
// kernel's definition to go back to. InstallKernelFast can run more than once.
func overrideProlog(name string, arity int, fn any) {
	s := symCell(name)
	n, ok := prologNatives[name]
	if !ok {
		n = makeDirect(arity, fn)
		prologNatives[name] = n
	}
	if s.function != n {
		s.kernelDef = s.function
		s.function = n
		vmNoteRebind(s)
	}
	prologInstall = append(prologInstall, s)
}

// armPrologNatives makes a rebinding of any group member or dependency restore
// the kernel's definitions. It runs last in InstallKernelFast, because install
// itself rebinds those symbols.
func armPrologNatives() {
	for _, s := range prologInstall {
		s.prologDep = true
	}
	for _, name := range prologDeps {
		symCell(name).prologDep = true
	}
	prologArmed, prologInstall = prologInstall, nil
}

func restorePrologKernel() {
	for _, name := range prologDeps {
		symCell(name).prologDep = false
	}
	for _, s := range prologArmed {
		s.prologDep = false
		if s.kernelDef != nil && prologNatives[s.str] == s.function {
			s.function = s.kernelDef
			vmNoteRebind(s)
		}
	}
	prologArmed = nil
}

// setSymbolFunction is the one writer of a symbol's function after boot. The
// register VM's specialized code depends on vmNoteRebind following any write.
func setSymbolFunction(s *scmSymbol, f Obj) {
	if s.prologDep {
		restorePrologKernel()
	}
	s.function = f
	vmNoteRebind(s)
}

func pvarp(x Obj) bool {
	if x == nil || isFixnum(x) || *x != scmHeadVector {
		return false
	}
	v := (*scmVector)(unsafe.Pointer(x)).vector
	return len(v) > 0 && v[0] == symShenPvar
}

func lazyderef(e *ControlFlow, x, b Obj) Obj {
	for pvarp(x) {
		w := PrimVectorGet(b, PrimVectorGet(x, one))
		if w == symPvarNull {
			return x
		}
		x = w
		e.Tick()
	}
	return x
}

func nativeLazyderef(e *ControlFlow, x, b Obj) Obj { return lazyderef(e, x, b) }

func deref(e *ControlFlow, x, b Obj) Obj {
	for {
		if ok, p := isPair(x); ok {
			e.Tick()
			car := deref(e, p.car, b)
			return cons(car, deref(e, p.cdr, b))
		}
		if !pvarp(x) {
			return x
		}
		w := PrimVectorGet(b, PrimVectorGet(x, one))
		if w == symPvarNull {
			return x
		}
		x = w
		e.Tick()
	}
}

func nativeDeref(e *ControlFlow, x, b Obj) Obj { return deref(e, x, b) }

func nativeBindv(_ *ControlFlow, x, v, b Obj) Obj {
	return PrimVectorSet(b, PrimVectorGet(x, one), v)
}

func nativeUnwind(_ *ControlFlow, x, b, r Obj) Obj {
	PrimVectorSet(b, PrimVectorGet(x, one), symPvarNull)
	return r
}

func bindThaw(e *ControlFlow, x, v, b, k Obj) Obj {
	PrimVectorSet(b, PrimVectorGet(x, one), v)
	r := Call0(e, k)
	if r == False {
		PrimVectorSet(b, PrimVectorGet(x, one), symPvarNull)
	}
	return r
}

func nativeBindBang(e *ControlFlow, x, v, b, k Obj) Obj { return bindThaw(e, x, v, b, k) }

// pvarCell is a Prolog variable in one allocation: it is what (absvector 2)
// gives, but the vector's slots live in the object.
type pvarCell struct {
	scmVector
	slots [2]Obj
}

func makePvar(n Obj) Obj {
	c := &pvarCell{}
	c.scmHead = scmHeadVector
	c.slots[0] = symShenPvar
	c.slots[1] = n
	c.vector = c.slots[:]
	return &c.scmHead
}

func nativeMakePvar(_ *ControlFlow, n Obj) Obj { return makePvar(n) }

func nativeNextticket(_ *ControlFlow, b, n Obj) Obj {
	return PrimVectorSet(PrimVectorSet(b, n, symPvarNull), one, PrimNumberAdd(n, one))
}

func nativeNewpv(_ *ControlFlow, b Obj) Obj {
	n := PrimVectorGet(b, one)
	v := makePvar(n)
	PrimVectorSet(PrimVectorSet(b, n, symPvarNull), one, PrimNumberAdd(n, one))
	return v
}

func nativeGc(_ *ControlFlow, b, r Obj) Obj {
	if r == False {
		n := PrimVectorGet(b, one)
		PrimVectorSet(b, one, PrimNumberSubtract(n, one))
	}
	return r
}

func nativeTicketNumber(_ *ControlFlow, b Obj) Obj { return PrimVectorGet(b, one) }

func nativeDecrementTicket(_ *ControlFlow, n, b Obj) Obj {
	return PrimVectorSet(b, one, PrimNumberSubtract(n, one))
}

func nativeUnlockedp(_ *ControlFlow, l Obj) Obj { return PrimVectorGet(l, one) }

func nativeIncinfs(_ *ControlFlow) Obj {
	return PrimSet(symInfs, PrimNumberAdd(one, PrimValue(symInfs)))
}

func occursCheck(e *ControlFlow, x, t Obj) bool {
	for {
		if equal(x, t) == True {
			return true
		}
		ok, p := isPair(t)
		if !ok {
			return false
		}
		e.Tick()
		if occursCheck(e, x, p.car) {
			return true
		}
		t = p.cdr
	}
}

func nativeOccursCheck(e *ControlFlow, x, t Obj) Obj { return boolObj(occursCheck(e, x, t)) }

// unify is shen.lzy= (occurs false) and shen.lzy=! (occurs true). A pair of
// lists is one freeze per pair, as in the kernel.
func unify(e *ControlFlow, occurs bool, x, y, b, k Obj) Obj {
	for {
		if equal(x, y) == True {
			return TailCall0(e, k)
		}
		if pvarp(x) && (!occurs || !occursCheck(e, x, deref(e, y, b))) {
			return bindThaw(e, x, y, b, k)
		}
		if pvarp(y) && (!occurs || !occursCheck(e, y, deref(e, x, b))) {
			return bindThaw(e, y, x, b, k)
		}
		okx, px := isPair(x)
		oky, py := isPair(y)
		if !okx || !oky {
			return False
		}
		k0 := k
		k = MakeNative0(func(e *ControlFlow) Obj {
			return unify(e, occurs, lazyderef(e, px.cdr, b), lazyderef(e, py.cdr, b), b, k0)
		})
		x, y = lazyderef(e, px.car, b), lazyderef(e, py.car, b)
		e.Tick()
	}
}

func nativeLzyEq(e *ControlFlow, x, y, b, k Obj) Obj { return unify(e, false, x, y, b, k) }

func nativeLzyEqBang(e *ControlFlow, x, y, b, k Obj) Obj { return unify(e, true, x, y, b, k) }

func nativeIs(e *ControlFlow, x, y, b, _, _, k Obj) Obj {
	return unify(e, false, lazyderef(e, x, b), lazyderef(e, y, b), b, k)
}

func nativeIsBang(e *ControlFlow, x, y, b, _, _, k Obj) Obj {
	return unify(e, true, lazyderef(e, x, b), lazyderef(e, y, b), b, k)
}

func nativeThawDirect(e *ControlFlow, f Obj) Obj { return TailCall0(e, f) }

func nativeFailDirect(_ *ControlFlow) Obj { return symShenFailBang }

func nativeParseFailurep(_ *ControlFlow, x Obj) Obj { return equal(x, symShenFailBang) }
