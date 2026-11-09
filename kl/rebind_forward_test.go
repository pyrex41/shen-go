package kl

import "testing"

// A native of the Prolog group fetched before a rebinding forwards to the
// kernel's definition afterwards, by every entry: Call, CallN and the VM.
func TestForwardDirectReachesDefinition(t *testing.T) {
	n := makeDirect(2, func(e *ControlFlow, a, b Obj) Obj { return MakeString("native") })
	def := CompileFunc("fwd-def", []Obj{MakeSymbol("A"), MakeSymbol("B")}, cons(MakeSymbol("cons"), cons(MakeSymbol("A"), cons(MakeSymbol("B"), Nil))))
	var e ControlFlow
	if got := ObjString(Call2(&e, n, MakeInteger(1), MakeInteger(2))); got != `"native"` {
		t.Fatalf("before: %s", got)
	}
	forwardDirect(MustNative(n), def)
	if got := ObjString(Call2(&e, n, MakeInteger(1), MakeInteger(2))); got != "(1 . 2)" {
		t.Errorf("Call2 after: %s", got)
	}
	if got := ObjString(Call(&e, n, MakeInteger(1), MakeInteger(2))); got != "(1 . 2)" {
		t.Errorf("Call after: %s", got)
	}
	caller := CompileFunc("fwd-caller", []Obj{MakeSymbol("F")}, cons(MakeSymbol("F"), cons(MakeInteger(3), cons(MakeInteger(4), Nil))))
	if got := ObjString(Call(&e, caller, n)); got != "(3 . 4)" {
		t.Errorf("from bytecode after: %s", got)
	}
}
