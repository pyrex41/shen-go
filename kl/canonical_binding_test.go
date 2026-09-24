package kl

import (
	"testing"
	"time"
)

// TestHasCanonicalPrimitiveBindingIsLockFree pins the property issues #51 and
// #55 depend on: the guard executed at every specialized primitive site must
// not take primitiveRegistry.mu. It holds the registry write lock and calls
// the guard from another goroutine; a guard that still does the locked map
// lookup blocks until the timeout.
func TestHasCanonicalPrimitiveBindingIsLockFree(t *testing.T) {
	sym := MakeSymbol("+")
	old := mustSymbol(sym).function
	defer func() { mustSymbol(sym).function = old }()
	primitiveRegistry.mu.RLock()
	canonical := primitiveRegistry.canonical["+"]
	primitiveRegistry.mu.RUnlock()
	BindSymbolFunc(sym, canonical)

	primitiveRegistry.mu.Lock()
	defer primitiveRegistry.mu.Unlock()
	done := make(chan bool, 1)
	go func() { done <- HasCanonicalPrimitiveBinding(sym) }()
	select {
	case got := <-done:
		if !got {
			t.Fatal("canonical + binding was not recognized")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("HasCanonicalPrimitiveBinding blocked on primitiveRegistry.mu")
	}
}

func registryCanonical(name string) (Obj, bool) {
	primitiveRegistry.mu.RLock()
	defer primitiveRegistry.mu.RUnlock()
	canonical, ok := primitiveRegistry.canonical[name]
	return canonical, ok
}

// TestCanonicalPrimitiveCachedOnSymbol covers the register-first order: the
// first registration under a name is cached on the interned symbol, a later
// same-name registration does not overwrite it (matching the registry map),
// and the guard follows the symbol's function binding through the cached
// field: false while unbound, true when bound to the canonical object, false
// when shadowed, true again when rebound to the canonical object.
func TestCanonicalPrimitiveCachedOnSymbol(t *testing.T) {
	const name = "canon-probe-55"
	first := MakePrimitive(name, 1, func(x Obj) Obj { return x })
	sym := MakeSymbol(name)
	if got := mustSymbol(sym).canonical; got != first {
		t.Fatalf("symbol canonical = %p, want first registration %p", got, first)
	}
	if HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("unbound symbol reported as canonical")
	}
	BindSymbolFunc(sym, first)
	if !HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("binding to the canonical primitive was not recognized")
	}

	second := MakePrimitive(name, 1, func(x Obj) Obj { return Nil })
	if second == first {
		t.Fatal("second registration returned the first object")
	}
	if got := mustSymbol(sym).canonical; got != first {
		t.Fatalf("second registration overwrote the cached canonical: %p, want %p", got, first)
	}
	if got, _ := registryCanonical(name); got != first {
		t.Fatalf("registry canonical = %p, want %p", got, first)
	}
	if !HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("registering a shadow must not change the guard while the canonical is still bound")
	}
	BindSymbolFunc(sym, second)
	if HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("shadowed binding reported as canonical")
	}
	BindSymbolFunc(sym, first)
	if !HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("rebinding back to the canonical primitive was not recognized")
	}
}

// TestCanonicalPrimitiveCachedOnPreInternedSymbol covers the intern-first
// order: when the symbol already exists in the trie before registration,
// register must find that node (not create another) and cache on it.
func TestCanonicalPrimitiveCachedOnPreInternedSymbol(t *testing.T) {
	const name = "canon-test-preinterned"
	sym := MakeSymbol(name)
	if mustSymbol(sym).canonical != nil {
		t.Fatal("fresh symbol already has a canonical primitive")
	}
	if HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("never-registered symbol reported as canonical")
	}
	prim := MakePrimitive(name, 2, func(x, y Obj) Obj { return x })
	if got := mustSymbol(sym).canonical; got != prim {
		t.Fatalf("pre-interned symbol canonical = %p, want %p", got, prim)
	}
	if MakeSymbol(name) != sym {
		t.Fatal("registration re-interned the symbol under a different node")
	}
	BindSymbolFunc(sym, prim)
	if !HasCanonicalPrimitiveBinding(sym) {
		t.Fatal("pre-interned symbol bound to its canonical was not recognized")
	}
}

// TestCanonicalBindingOfPlusCachedOnSymbol checks the cached field on a real
// kernel primitive, and that a name nobody registered stays a false guard.
func TestCanonicalBindingOfPlusCachedOnSymbol(t *testing.T) {
	want, ok := registryCanonical("+")
	if !ok || want == nil {
		t.Fatal("+ has no canonical registration")
	}
	if got := mustSymbol(MakeSymbol("+")).canonical; got != want {
		t.Fatalf("+ symbol canonical = %p, registry has %p", got, want)
	}
	if HasCanonicalPrimitiveBinding(MakeSymbol("never-registered-xyz")) {
		t.Fatal("never-registered symbol reported as canonical")
	}
	if mustSymbol(MakeSymbol("never-registered-xyz")).canonical != nil {
		t.Fatal("never-registered symbol has a cached canonical")
	}
}

// TestCanonicalRegistrySymbolInvariant is the exhaustive consistency check:
// for every registered name the interned symbol carries exactly the registry's
// canonical object, and the guard equals "function == canonical". Run after
// InstallKernelFast so the canonicalOrMake / restoreCanonicalPrimitive paths
// in kernelfast.go are covered as well as init-time registration.
func TestCanonicalRegistrySymbolInvariant(t *testing.T) {
	InstallKernelFast()
	primitiveRegistry.mu.RLock()
	names := make([]string, 0, len(primitiveRegistry.canonical))
	canonicals := make(map[string]Obj, len(primitiveRegistry.canonical))
	for name, prim := range primitiveRegistry.canonical {
		names = append(names, name)
		canonicals[name] = prim
	}
	primitiveRegistry.mu.RUnlock()
	if len(names) == 0 {
		t.Fatal("primitive registry is empty")
	}
	for _, name := range names {
		want := canonicals[name]
		sym := MakeSymbol(name)
		s := mustSymbol(sym)
		if want == nil {
			t.Errorf("%q: registry holds a nil canonical", name)
			continue
		}
		if s.canonical != want {
			t.Errorf("%q: symbol canonical %p != registry canonical %p", name, s.canonical, want)
		}
		if got, wantGuard := HasCanonicalPrimitiveBinding(sym), s.function == want; got != wantGuard {
			t.Errorf("%q: HasCanonicalPrimitiveBinding = %v, want %v (function %p, canonical %p)", name, got, wantGuard, s.function, want)
		}
	}
}
