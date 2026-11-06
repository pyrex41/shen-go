package kl

import (
	"strings"
	"testing"
)

// A direct-call caller must see the same behaviour as Call, whatever the callee
// is and whatever happens to the binding.

func directEval(t *testing.T, e *ControlFlow, src string) Obj {
	t.Helper()
	r := NewSexpReader(strings.NewReader(src), false)
	form, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}
	return Eval(e, form)
}

func TestDirectCallRebindingTakesEffectAtNextCall(t *testing.T) {
	e := &ControlFlow{}
	sym := MakeSymbol("direct-test-rebind")
	BindSymbolFunc(sym, MakeNative1(func(_ *ControlFlow, a Obj) Obj { return MakeInteger(1) }))
	if got := Call1(e, SymFn(sym), Nil); got != MakeInteger(1) {
		t.Fatalf("got %s", ObjString(got))
	}
	// To an old-convention native.
	BindSymbolFunc(sym, MakeNative(func(c *ControlFlow) { c.Return(MakeInteger(2)) }, 1))
	if got := Call1(e, SymFn(sym), Nil); got != MakeInteger(2) {
		t.Fatalf("got %s", ObjString(got))
	}
	// To bytecode.
	directEval(t, e, "(defun direct-test-rebind (X) 3)")
	if got := Call1(e, SymFn(sym), Nil); got != MakeInteger(3) {
		t.Fatalf("got %s", ObjString(got))
	}
}

func TestDirectUnboundMessage(t *testing.T) {
	sym := MakeSymbol("direct-test-unbound")
	defer func() {
		r := recover()
		o, ok := r.(Obj)
		if !ok || !IsError(o) || mustError(o).err != "variable direct-test-unbound not bound" {
			t.Fatalf("recovered %v", r)
		}
	}()
	SymFn(sym)
}

func TestDirectOldEntryAndTrampoline(t *testing.T) {
	e := &ControlFlow{}
	// Through Call, apply and the VM, a direct native behaves like any other,
	// including a tail call from its body.
	target := MakeNative2(func(_ *ControlFlow, a, b Obj) Obj { return PrimCons(a, b) })
	viaTail := MakeNative1(func(c *ControlFlow, a Obj) Obj { return TailCall2(c, target, a, Nil) })
	want := "[7]"
	if got := Call(e, viaTail, MakeInteger(7)); ObjString(got) != ObjString(Cons(MakeInteger(7), Nil)) {
		t.Fatalf("Call: got %s want %s", ObjString(got), want)
	}
	if got := Call1(e, viaTail, MakeInteger(7)); ObjString(got) != ObjString(Cons(MakeInteger(7), Nil)) {
		t.Fatalf("Call1: got %s", ObjString(got))
	}
	sym := MakeSymbol("direct-test-tail")
	BindSymbolFunc(sym, viaTail)
	if got := directEval(t, e, "(direct-test-tail 7)"); ObjString(got) != ObjString(Cons(MakeInteger(7), Nil)) {
		t.Fatalf("VM: got %s", ObjString(got))
	}
}

func TestDirectUnboundedTailCalls(t *testing.T) {
	e := &ControlFlow{}
	symA, symB := MakeSymbol("direct-test-a"), MakeSymbol("direct-test-b")
	count := func(next Obj) Obj {
		return MakeNative1(func(c *ControlFlow, n Obj) Obj {
			if n == MakeInteger(0) {
				return True
			}
			return TailCall1(c, SymFn(next), MakeInteger(fixnum(n)-1))
		})
	}
	BindSymbolFunc(symA, count(symB))
	BindSymbolFunc(symB, count(symA))
	if got := Call1(e, SymFn(symA), MakeInteger(5000000)); got != True {
		t.Fatalf("got %s", ObjString(got))
	}
}

func TestDirectStepLimitCharged(t *testing.T) {
	e := &ControlFlow{}
	e.SetStepLimit(100)
	loop := MakeSymbol("direct-test-forever")
	BindSymbolFunc(loop, MakeNative0(func(c *ControlFlow) Obj { return TailCall0(c, SymFn(loop)) }))
	got := func() (r interface{}) {
		defer func() { r = recover() }()
		Call0(e, SymFn(loop))
		return nil
	}()
	o, ok := got.(Obj)
	if !ok || !IsError(o) || !strings.Contains(mustError(o).err, "step limit") {
		t.Fatalf("got %v", got)
	}
	// A non-tail direct call charges one step, as Call does.
	e2, e3 := &ControlFlow{}, &ControlFlow{}
	e2.SetStepLimit(1000)
	e3.SetStepLimit(1000)
	id := MakeNative1(func(_ *ControlFlow, a Obj) Obj { return a })
	for i := 0; i < 10; i++ {
		Call1(e2, id, Nil)
		Call(e3, id, Nil)
	}
	if e2.steps != e3.steps {
		t.Fatalf("Call1 charged %d, Call charged %d", e2.steps, e3.steps)
	}
}

func TestDirectWrongArityMessageUnchanged(t *testing.T) {
	e := &ControlFlow{}
	f := MakeNative2(func(_ *ControlFlow, a, b Obj) Obj { return a })
	defer func() {
		o, ok := recover().(Obj)
		if !ok || !IsError(o) || mustError(o).err != "partial apply unsupported" {
			t.Fatalf("got %v", o)
		}
	}()
	Call1(e, f, Nil)
}

func TestGuardedHelpersHonourRebinding(t *testing.T) {
	e := &ControlFlow{}
	cases := []struct {
		name string
		call func() Obj
	}{
		{"=", func() Obj { return GEq(e, Nil, Nil) }},
		{"hd", func() Obj { return GHd(e, Cons(Nil, Nil)) }},
		{"tl", func() Obj { return GTl(e, Cons(Nil, Nil)) }},
		{"cons?", func() Obj { return GConsP(e, Nil) }},
		{"not", func() Obj { return GNot(e, True) }},
		{"+", func() Obj { return GAdd(e, MakeInteger(1), MakeInteger(2)) }},
		{"-", func() Obj { return GSub(e, MakeInteger(1), MakeInteger(2)) }},
		{"*", func() Obj { return GMul(e, MakeInteger(1), MakeInteger(2)) }},
		{"/", func() Obj { return GDiv(e, MakeInteger(1), MakeInteger(2)) }},
		{"<", func() Obj { return GLt(e, MakeInteger(1), MakeInteger(2)) }},
		{"<=", func() Obj { return GLe(e, MakeInteger(1), MakeInteger(2)) }},
		{">", func() Obj { return GGt(e, MakeInteger(1), MakeInteger(2)) }},
		{">=", func() Obj { return GGe(e, MakeInteger(1), MakeInteger(2)) }},
		{"value", func() Obj { return GValue(e, MakeSymbol("direct-test-value")) }},
		{"cn", func() Obj {
			return PrimGuard2(e, MakeSymbol("cn"), PrimStringConcat, MakeString("a"), MakeString("b"))
		}},
	}
	PrimSet(MakeSymbol("direct-test-value"), Nil)
	marker := MakeSymbol("direct-test-rebound")
	for _, c := range cases {
		sym := MakeSymbol(c.name)
		old := SymFn(sym)
		arity := MustNative(old).require
		if got := c.call(); got == marker {
			t.Fatalf("%s: rebound result before rebinding", c.name)
		}
		BindSymbolFunc(sym, MakeNative(func(ctl *ControlFlow) { ctl.Return(marker) }, arity))
		got := c.call()
		BindSymbolFunc(sym, old)
		if got != marker {
			t.Errorf("%s: rebinding ignored, got %s", c.name, ObjString(got))
		}
		if got := c.call(); got == marker {
			t.Errorf("%s: canonical binding not restored", c.name)
		}
	}
}

func TestGuardedHelpersTypedIROff(t *testing.T) {
	t.Setenv("SHEN_GO_TYPED_IR", "off")
	ResetTypedIRModeForTest()
	defer ResetTypedIRModeForTest()
	if TypedIREnabled() {
		t.Fatal("off-switch not read")
	}
	e := &ControlFlow{}
	// Off means the dynamic path, seen here through the step charge.
	e.SetStepLimit(1000)
	if GAdd(e, MakeInteger(1), MakeInteger(2)) != MakeInteger(3) {
		t.Fatal("off: wrong sum")
	}
	if e.steps == 0 {
		t.Fatal("off: fast path taken")
	}
}

func TestGuardedArithmeticMatchesPrimitives(t *testing.T) {
	e := &ControlFlow{}
	vals := []Obj{
		MakeInteger(0), MakeInteger(1), MakeInteger(-1), MakeInteger(fixnumMax - 1), MakeInteger(fixnumMin),
		MakeInteger(fixnumMax), MakeInteger(1 << 40), MakeNumber(1.5), MakeNumber(-0.25), MakeNumber(1e300),
	}
	ops := []struct {
		name string
		g    func(*ControlFlow, Obj, Obj) Obj
		p    func(Obj, Obj) Obj
	}{
		{"+", GAdd, PrimNumberAdd}, {"-", GSub, PrimNumberSubtract}, {"*", GMul, PrimNumberMultiply},
		{"<", GLt, PrimLessThan}, {"<=", GLe, PrimLessEqual}, {">", GGt, PrimGreatThan}, {">=", GGe, PrimGreatEqual},
	}
	for _, op := range ops {
		for _, a := range vals {
			for _, b := range vals {
				if got, want := op.g(e, a, b), op.p(a, b); !Equal(got, want) || ObjString(got) != ObjString(want) || isFixnum(got) != isFixnum(want) {
					t.Errorf("%s %s %s: got %s want %s", ObjString(a), op.name, ObjString(b), ObjString(got), ObjString(want))
				}
			}
		}
	}
}
