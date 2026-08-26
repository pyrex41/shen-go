package kl

import (
	"os"
	"testing"
)

func TestPrimitiveSpecAndCanonicalBinding(t *testing.T) {
	if HasCanonicalPrimitiveBinding(MakeInteger(fixnumMax - 1)) {
		t.Fatal("fixnum reported as a canonical primitive binding")
	}
	spec, ok := LookupPrimitiveSpec("+")
	if !ok || spec.Arity != 2 || !spec.Scalar || !spec.Result.Contains(KindNumber) {
		t.Fatalf("unexpected + spec: %#v, %v", spec, ok)
	}
	if len(spec.Args) != 2 || !spec.Args[0].Contains(KindNumber) {
		t.Fatalf("unexpected + args: %#v", spec.Args)
	}
	sym := MakeSymbol("+")
	old := mustSymbol(sym).function
	defer func() { mustSymbol(sym).function = old }()
	primitiveRegistry.mu.RLock()
	canonical := primitiveRegistry.canonical["+"]
	primitiveRegistry.mu.RUnlock()
	BindSymbolFunc(sym, canonical)
	if !HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("canonical + binding was not recognized")
	}
	BindSymbolFunc(sym, MakePrimitive("+shadow", 1, func(x Obj) Obj { return x }))
	if HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("shadowed + binding reported as canonical")
	}
}

func TestTypedIRModeDefaultsOnAndSupportsOff(t *testing.T) {
	old, had := os.LookupEnv("SHEN_GO_TYPED_IR")
	defer func() {
		ResetTypedIRModeForTest()
		if had {
			_ = os.Setenv("SHEN_GO_TYPED_IR", old)
		} else {
			_ = os.Unsetenv("SHEN_GO_TYPED_IR")
		}
	}()
	_ = os.Unsetenv("SHEN_GO_TYPED_IR")
	ResetTypedIRModeForTest()
	if !TypedIRModeEnabled() {
		t.Fatal("typed IR should default on")
	}
	_ = os.Setenv("SHEN_GO_TYPED_IR", "off")
	ResetTypedIRModeForTest()
	if TypedIRModeEnabled() {
		t.Fatal("typed IR off switch ignored")
	}
	_ = os.Setenv("SHEN_GO_TYPED_IR", "unexpected")
	ResetTypedIRModeForTest()
	if !TypedIRModeEnabled() {
		t.Fatal("unknown mode should retain enabled default")
	}
}
