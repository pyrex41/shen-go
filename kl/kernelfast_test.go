package kl

import (
	"strings"
	"testing"
)

// bindDummy makes overridePrimitive / overrideNative consider the symbol
// kernel-defined, which is the sparse-vs-full distinction InstallKernelFast
// uses. Tests that want to exercise an override install a dummy first.
func bindDummy(name string, arity int) {
	BindSymbolFunc(MakeSymbol(name), MakeNative(func(e *ControlFlow) {
		e.Return(False)
	}, arity))
}

func TestKernelOverridePredicates(t *testing.T) {
	for _, name := range []string{
		"empty?", "boolean?", "vector?", "tuple?", "shen.pvar?",
		"length", "reverse", "hdstr",
	} {
		bindDummy(name, 1)
	}
	InstallKernelFast()

	var e ControlFlow
	call1 := func(name string, x Obj) Obj {
		return Call(&e, PrimFunc(MakeSymbol(name)), x)
	}

	if call1("empty?", Nil) != True {
		t.Fatalf("empty? ()")
	}
	if call1("empty?", cons(True, Nil)) != False {
		t.Fatalf("empty? (true)")
	}
	if call1("empty?", False) != False {
		t.Fatalf("empty? false")
	}

	if call1("boolean?", True) != True || call1("boolean?", False) != True {
		t.Fatalf("boolean? true/false")
	}
	if call1("boolean?", MakeSymbol("true")) != False {
		t.Fatalf("boolean? on a symbol")
	}

	if call1("length", Nil) != MakeInteger(0) && GetInteger(call1("length", Nil)) != 0 {
		t.Fatalf("length ()")
	}
	lst := cons(MakeInteger(1), cons(MakeInteger(2), cons(MakeInteger(3), Nil)))
	if GetInteger(call1("length", lst)) != 3 {
		t.Fatalf("length 3, got %s", ObjString(call1("length", lst)))
	}
	rev := call1("reverse", lst)
	if GetInteger(car(rev)) != 3 || GetInteger(car(cdr(cdr(rev)))) != 1 {
		t.Fatalf("reverse, got %s", ObjString(rev))
	}

	if mustString(call1("hdstr", MakeString("ABC"))) != "A" {
		t.Fatalf("hdstr")
	}
}

func TestKernelOverrideTupleAndVector(t *testing.T) {
	bindDummy("@p", 2)
	bindDummy("tuple?", 1)
	bindDummy("vector", 1)
	bindDummy("vector?", 1)
	bindDummy("<-vector", 2)
	bindDummy("vector->", 3)
	bindDummy("fst", 1)
	bindDummy("snd", 1)
	bindDummy("fail", 0)
	bindDummy("put", 4)
	bindDummy("get", 3)
	InstallKernelFast()

	var e ControlFlow
	// (fail) is the symbol shen.fail!, as sys.kl's (defun fail () shen.fail!)
	// says; `...` is only how the printer shows it (issue #54).
	if got := Call(&e, PrimFunc(MakeSymbol("fail"))); got != symShenFailBang {
		t.Fatalf("fail should return shen.fail!, got %s", ObjString(got))
	}
	tup := Call(&e, PrimFunc(MakeSymbol("@p")), MakeInteger(1), MakeInteger(2))
	if Call(&e, PrimFunc(MakeSymbol("tuple?")), tup) != True {
		t.Fatalf("tuple? of @p")
	}
	if Call(&e, PrimFunc(MakeSymbol("vector?")), tup) != False {
		t.Fatalf("vector? of tuple should be false (tag is shen.tuple, not a limit)")
	}
	if GetInteger(Call(&e, PrimFunc(MakeSymbol("fst")), tup)) != 1 {
		t.Fatalf("fst")
	}
	if GetInteger(Call(&e, PrimFunc(MakeSymbol("snd")), tup)) != 2 {
		t.Fatalf("snd")
	}

	v := Call(&e, PrimFunc(MakeSymbol("vector")), MakeInteger(2))
	if Call(&e, PrimFunc(MakeSymbol("vector?")), v) != True {
		t.Fatalf("vector?")
	}
	Call(&e, PrimFunc(MakeSymbol("vector->")), v, MakeInteger(1), MakeString("x"))
	if mustString(Call(&e, PrimFunc(MakeSymbol("<-vector")), v, MakeInteger(1))) != "x" {
		t.Fatalf("<-vector after vector->")
	}

	mustRaise := func(what, want string, f func()) {
		t.Helper()
		defer func() {
			r := recover()
			if r == nil {
				t.Fatalf("%s should raise", what)
			}
			obj, ok := r.(Obj)
			if !ok || !IsError(obj) {
				t.Fatalf("%s: raised %v, want a Shen error", what, r)
			}
			if got := GetString(PrimErrorToString(obj)); got != want {
				t.Fatalf("%s: raised %q, want %q", what, got, want)
			}
		}()
		f()
	}
	mustRaise("<-vector of unwritten slot", "vector element not found\n", func() {
		Call(&e, PrimFunc(MakeSymbol("<-vector")), v, MakeInteger(2))
	})
	// The unwritten slot holds the kernel's own fail filler, so a Shen
	// program sees (= (<-address (vector 2) 1) shen.fail!) as true, as it
	// does on shen-cl and shen-scheme.
	if slot := PrimVectorGet(v, MakeInteger(2)); slot != symShenFailBang {
		t.Fatalf("(<-address (vector 2) 2) = %s, want shen.fail!", ObjString(slot))
	}
	// Issue #54's mixed-read case: a slot that a program set to shen.fail! is
	// indistinguishable from an unwritten one, so <-vector raises for it too.
	w := Call(&e, PrimFunc(MakeSymbol("vector")), MakeInteger(3))
	Call(&e, PrimFunc(MakeSymbol("vector->")), w, MakeInteger(1), symShenFailBang)
	mustRaise("<-vector of a slot holding shen.fail!", "vector element not found\n", func() {
		Call(&e, PrimFunc(MakeSymbol("<-vector")), w, MakeInteger(1))
	})
	// The symbol `...` is an ordinary value, not a filler.
	dots := MakeSymbol("...")
	Call(&e, PrimFunc(MakeSymbol("vector->")), w, MakeInteger(2), dots)
	if got := Call(&e, PrimFunc(MakeSymbol("<-vector")), w, MakeInteger(2)); got != dots {
		t.Fatalf("<-vector of a slot holding the symbol ... = %s, want ...", ObjString(got))
	}

	d := Call(&e, PrimFunc(MakeSymbol("vector")), MakeInteger(8))
	Call(&e, PrimFunc(MakeSymbol("put")), MakeSymbol("a"), MakeSymbol("b"), MakeInteger(1), d)
	got := Call(&e, PrimFunc(MakeSymbol("get")), MakeSymbol("a"), MakeSymbol("b"), d)
	if GetInteger(got) != 1 {
		t.Fatalf("put/get roundtrip, got %s", ObjString(got))
	}
}

func TestKernelOverrideSymbolAndVariable(t *testing.T) {
	bindDummy("symbol?", 1)
	bindDummy("variable?", 1)
	InstallKernelFast()

	var e ControlFlow
	sym := func(name string, x Obj) Obj {
		return Call(&e, PrimFunc(MakeSymbol(name)), x)
	}

	if sym("symbol?", MakeSymbol("foo")) != True {
		t.Fatalf("symbol? foo")
	}
	if sym("symbol?", MakeSymbol("{")) != True {
		t.Fatalf("symbol? {")
	}
	if sym("symbol?", MakeSymbol(":")) != True {
		t.Fatalf("symbol? :")
	}
	if sym("symbol?", MakeSymbol("1foo")) != False {
		t.Fatalf("symbol? 1foo (leading digit)")
	}
	if sym("symbol?", True) != False || sym("symbol?", MakeInteger(1)) != False {
		t.Fatalf("symbol? true/1")
	}
	if sym("symbol?", MakeString("foo")) != False {
		t.Fatalf("symbol? string")
	}

	if sym("variable?", MakeSymbol("Foo")) != True {
		t.Fatalf("variable? Foo")
	}
	if sym("variable?", MakeSymbol("foo")) != False {
		t.Fatalf("variable? foo")
	}
	if sym("variable?", MakeSymbol("A-1")) != True {
		t.Fatalf("variable? A-1")
	}
	if sym("variable?", MakeSymbol("A b")) != False {
		t.Fatalf("variable? with space")
	}
}

func TestKernelOverrideAppendAssocElement(t *testing.T) {
	bindDummy("append", 2)
	bindDummy("assoc", 2)
	bindDummy("element?", 2)
	InstallKernelFast()

	var e ControlFlow
	a := cons(MakeInteger(1), cons(MakeInteger(2), Nil))
	b := cons(MakeInteger(3), Nil)
	got := Call(&e, PrimFunc(MakeSymbol("append")), a, b)
	if GetInteger(car(got)) != 1 || GetInteger(car(cdr(cdr(got)))) != 3 {
		t.Fatalf("append, got %s", ObjString(got))
	}

	entry := cons(MakeSymbol("k"), MakeInteger(9))
	al := cons(entry, Nil)
	found := Call(&e, PrimFunc(MakeSymbol("assoc")), MakeSymbol("k"), al)
	if found != entry {
		t.Fatalf("assoc hit")
	}
	if Call(&e, PrimFunc(MakeSymbol("assoc")), MakeSymbol("z"), al) != Nil {
		t.Fatalf("assoc miss")
	}

	if Call(&e, PrimFunc(MakeSymbol("element?")), MakeInteger(2), a) != True {
		t.Fatalf("element? hit")
	}
	if Call(&e, PrimFunc(MakeSymbol("element?")), MakeInteger(9), a) != False {
		t.Fatalf("element? miss")
	}
}

func TestRestoreCanonicalNotAndInteger(t *testing.T) {
	InstallKernelFast()
	if !HasCanonicalPrimitiveBinding(MakeSymbol("not")) {
		t.Fatalf("not should be canonical after InstallKernelFast")
	}
	if !HasCanonicalPrimitiveBinding(MakeSymbol("integer?")) {
		t.Fatalf("integer? should be canonical after InstallKernelFast")
	}
	InstallIntegerGuard()
	if !HasCanonicalPrimitiveBinding(MakeSymbol("integer?")) {
		t.Fatalf("InstallIntegerGuard must not wrap a canonical integer?")
	}
}

func TestKernelOverrideHeadTailNthBound(t *testing.T) {
	bindDummy("head", 1)
	bindDummy("tail", 1)
	bindDummy("nth", 2)
	bindDummy("bound?", 1)
	bindDummy("concat", 2)
	bindDummy("==", 2)
	bindDummy("sum", 1)
	bindDummy("adjoin", 2)
	bindDummy("remove", 2)
	bindDummy("union", 2)
	bindDummy("map", 2)
	InstallKernelFast()

	var e ControlFlow
	lst := cons(MakeInteger(1), cons(MakeInteger(2), cons(MakeInteger(3), Nil)))
	if GetInteger(Call(&e, PrimFunc(MakeSymbol("head")), lst)) != 1 {
		t.Fatalf("head")
	}
	tl := Call(&e, PrimFunc(MakeSymbol("tail")), lst)
	if GetInteger(car(tl)) != 2 {
		t.Fatalf("tail")
	}
	if GetInteger(Call(&e, PrimFunc(MakeSymbol("nth")), MakeInteger(2), lst)) != 2 {
		t.Fatalf("nth")
	}
	if Call(&e, PrimFunc(MakeSymbol("==")), MakeInteger(1), MakeInteger(1)) != True {
		t.Fatalf("==")
	}
	if GetInteger(Call(&e, PrimFunc(MakeSymbol("sum")), lst)) != 6 {
		t.Fatalf("sum")
	}
	adj := Call(&e, PrimFunc(MakeSymbol("adjoin")), MakeInteger(0), lst)
	if GetInteger(car(adj)) != 0 {
		t.Fatalf("adjoin new")
	}
	if Call(&e, PrimFunc(MakeSymbol("adjoin")), MakeInteger(1), lst) != lst {
		t.Fatalf("adjoin existing")
	}
	rm := Call(&e, PrimFunc(MakeSymbol("remove")), MakeInteger(2), lst)
	if GetInteger(car(rm)) != 1 || GetInteger(car(cdr(rm))) != 3 {
		t.Fatalf("remove, got %s", ObjString(rm))
	}
	u := Call(&e, PrimFunc(MakeSymbol("union")), cons(MakeInteger(1), Nil), cons(MakeInteger(2), Nil))
	if GetInteger(car(u)) != 1 || GetInteger(car(cdr(u))) != 2 {
		t.Fatalf("union order, got %s", ObjString(u))
	}
	add1 := MakePrimitive("add1", 1, func(x Obj) Obj { return MakeInteger(GetInteger(x) + 1) })
	mapped := Call(&e, PrimFunc(MakeSymbol("map")), add1, lst)
	if GetInteger(car(mapped)) != 2 || GetInteger(car(cdr(cdr(mapped)))) != 4 {
		t.Fatalf("map, got %s", ObjString(mapped))
	}
}

// nativeNth formats its error the way the kernel does: it counts N down while
// dropping heads, then hands what is left to shen.app. Without shen.app (a
// lowered kernel slice that keeps nth but drops writer.kl) it must still raise
// a sane message naming the same remaining N and L.
func TestNativeNthErrorWithAndWithoutShenApp(t *testing.T) {
	bindDummy("nth", 2)
	InstallKernelFast()
	if kernelBound("nth") == nil {
		t.Fatalf("nth not bound")
	}

	appSym := MakeSymbol("shen.app")
	saved := kernelBound("shen.app")
	defer BindSymbolFunc(appSym, saved)

	var e ControlFlow
	lst := cons(MakeSymbol("a"), cons(MakeSymbol("b"), Nil))
	nth := func(n Obj, l Obj) (msg string) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatalf("nth %s: expected an error", ObjString(n))
			}
			o, ok := r.(Obj)
			if !ok {
				t.Fatalf("nth %s: raised %v, not a Shen error", ObjString(n), r)
			}
			msg = mustString(PrimErrorToString(o))
		}()
		Call(&e, PrimFunc(MakeSymbol("nth")), n, l)
		return ""
	}

	// With a shen.app in place, nth calls it with the remaining N and L.
	var calls []string
	BindSymbolFunc(appSym, MakeNative(func(e *ControlFlow) {
		x, s := e.Get(1), e.Get(2)
		calls = append(calls, ObjString(x)+"|"+ObjString(e.Get(3)))
		e.Return(MakeString("<" + ObjString(x) + ">" + mustString(s)))
	}, 3))
	if got, want := nth(MakeInteger(3), lst), "nth applied to <1>, <()>\n"; got != want {
		t.Errorf("with shen.app: got %q, want %q", got, want)
	}
	if got, want := len(calls), 2; got != want {
		t.Fatalf("shen.app calls: got %d (%v), want %d", got, calls, want)
	}
	if calls[0] != "()|shen.a" || calls[1] != "1|shen.a" {
		t.Errorf("shen.app calls: got %v, want [()|shen.a 1|shen.a]", calls)
	}

	// Without shen.app, the fallback prints the same remaining N and L.
	BindSymbolFunc(appSym, nil)
	if got, want := nth(MakeInteger(3), lst), "nth applied to 1, ()\n"; got != want {
		t.Errorf("without shen.app: got %q, want %q", got, want)
	}
	if got, want := nth(MakeInteger(0), lst), "nth applied to -2, ()\n"; got != want {
		t.Errorf("without shen.app, index 0: got %q, want %q", got, want)
	}
	// The success path never touches shen.app.
	if got := Call(&e, PrimFunc(MakeSymbol("nth")), MakeInteger(2), lst); got != MakeSymbol("b") {
		t.Errorf("nth 2: got %s, want b", ObjString(got))
	}
}

func TestReaderCharPredicates(t *testing.T) {
	bindDummy("shen.digit?", 1)
	bindDummy("shen.uppercase?", 1)
	bindDummy("shen.alpha?", 1)
	bindDummy("shen.alphanums?", 1)
	bindDummy("shen.+string?", 1)
	bindDummy("shen.hds=?", 2)
	InstallKernelFast()
	var e ControlFlow
	if Call(&e, PrimFunc(MakeSymbol("shen.digit?")), MakeInteger(48)) != True {
		t.Fatalf("digit? 48")
	}
	if Call(&e, PrimFunc(MakeSymbol("shen.digit?")), MakeInteger(47)) != False {
		t.Fatalf("digit? 47")
	}
	if Call(&e, PrimFunc(MakeSymbol("shen.uppercase?")), MakeInteger(65)) != True {
		t.Fatalf("uppercase? A")
	}
	if Call(&e, PrimFunc(MakeSymbol("shen.alpha?")), MakeInteger(int('+'))) != True {
		t.Fatalf("alpha? +")
	}
	if Call(&e, PrimFunc(MakeSymbol("shen.alphanums?")), MakeString("A-1")) != True {
		t.Fatalf("alphanums? A-1")
	}
	if Call(&e, PrimFunc(MakeSymbol("shen.+string?")), MakeString("")) != False {
		t.Fatalf("+string? empty")
	}
	lst := cons(MakeInteger(1), Nil)
	if Call(&e, PrimFunc(MakeSymbol("shen.hds=?")), lst, MakeInteger(1)) != True {
		t.Fatalf("hds=?")
	}
}

func TestTrapErrorValueOr(t *testing.T) {
	form, err := NewSexpReader(strings.NewReader(`(trap-error (value missing-xyz) (lambda E 42))`), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	fn := mustBytecodeFunc(CompileFunc("t", nil, form)).fn
	hasValueOr := false
	for _, k := range fn.Consts {
		if IsSymbol(k) && GetSymbol(k) == "_kl.value/or" {
			hasValueOr = true
		}
	}
	if !hasValueOr {
		t.Fatalf("trap-error (value ...) should compile to _kl.value/or, consts=%v", fn.Consts)
	}
	var e ControlFlow
	got := Call(&e, makeBytecodeObj(fn, nil))
	if GetInteger(got) != 42 {
		t.Fatalf("unbound value/or = %s", ObjString(got))
	}
}

func TestThawFreezeInline(t *testing.T) {
	form, err := NewSexpReader(strings.NewReader(`(thaw (freeze 7))`), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	fn := mustBytecodeFunc(CompileFunc("t", nil, form)).fn
	if fn.Code[0].Op != OP_LOAD_CONST {
		t.Fatalf("(thaw (freeze 7)) should be LOAD_CONST, code=%v", fn.Code)
	}
	var e ControlFlow
	if GetInteger(Call(&e, makeBytecodeObj(fn, nil))) != 7 {
		t.Fatalf("thaw freeze")
	}
}

func TestLetFreezeThawInline(t *testing.T) {
	src := `(let Go (freeze 9) (if false 1 (thaw Go)))`
	form, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	fn := mustBytecodeFunc(CompileFunc("t", nil, form)).fn
	for _, instr := range fn.Code {
		if instr.Op == OP_MAKE_CLOSURE {
			t.Fatalf("let-freeze-thaw should not allocate a closure: %v", fn.Code)
		}
	}
	var e ControlFlow
	if GetInteger(Call(&e, makeBytecodeObj(fn, nil))) != 9 {
		t.Fatalf("let freeze thaw")
	}
}

func TestEqualLiteralFold(t *testing.T) {
	form, err := NewSexpReader(strings.NewReader(`(= 1 1)`), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	fn := mustBytecodeFunc(CompileFunc("t", nil, form)).fn
	if fn.Code[0].Op != OP_LOAD_CONST {
		t.Fatalf("(= 1 1) should fold, code=%v", fn.Code)
	}
	if fn.Consts[fn.Code[0].A] != True {
		t.Fatalf("folded (= 1 1) const = %s", ObjString(fn.Consts[fn.Code[0].A]))
	}
}

func TestInternConstantFold(t *testing.T) {
	form, err := NewSexpReader(strings.NewReader(`(intern "foo")`), true).Read()
	if err != nil {
		t.Fatal(err)
	}
	fn := mustBytecodeFunc(CompileFunc("t", nil, form)).fn
	if fn.Code[0].Op != OP_LOAD_CONST {
		t.Fatalf("intern \"foo\" should fold to LOAD_CONST, code=%v", fn.Code)
	}
	if GetSymbol(fn.Consts[fn.Code[0].A]) != "foo" {
		t.Fatalf("folded const = %s", ObjString(fn.Consts[fn.Code[0].A]))
	}
}

func TestAnalyseSymbolName(t *testing.T) {
	ok := []string{"foo", "A", "e-1", "shen.fail!", "=", "+", "foo_bar"}
	no := []string{"", "1foo", "foo bar", ":", "{", ",", ";"}
	for _, s := range ok {
		if !analyseSymbolName(s) {
			t.Errorf("analyseSymbolName(%q) = false, want true", s)
		}
	}
	for _, s := range no {
		if analyseSymbolName(s) {
			t.Errorf("analyseSymbolName(%q) = true, want false", s)
		}
	}
}

// TestInstallKernelFastBindsEveryOverride pins issue #49: InstallKernelFast
// installs every rebinding whether or not the kernel defined the name. A
// Yggdrasil `lower` slice deletes the KL (defun NAME …) of each declared
// native_override, so if a native were only installed when the name was
// already bound, the lowered artifact would die with "variable vector not
// bound". The rows come from the same parser TestEquivTable uses, so a new
// rebinding is covered automatically. Every row's function binding is
// cleared first: the process-global symbol table is shared by the whole
// package, and TestEquivTable's BootKernel (kl/equiv_test.go sorts before
// this file) would otherwise have defined all 58 names already, masking the
// old guard under a plain `go test ./kl`. With the clearing, this test fails
// under the old guard in any order and passes without it.
func TestInstallKernelFastBindsEveryOverride(t *testing.T) {
	rows := parseInstallKernelFast(t)
	if len(rows) == 0 {
		t.Fatal("parseInstallKernelFast returned no rows")
	}
	for _, b := range rows {
		mustSymbol(MakeSymbol(b.kernelFn)).function = nil
	}
	InstallKernelFast()
	for _, b := range rows {
		sym := MakeSymbol(b.kernelFn)
		fn := func() (f Obj) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s is not bound after InstallKernelFast (%v); the kernelBound guard is back", b.kernelFn, r)
					f = nil
				}
			}()
			return PrimFunc(sym)
		}()
		if fn == nil {
			continue
		}
		if !IsNativeBinding(fn) {
			t.Errorf("%s: bound to %s after InstallKernelFast, want the Go native %s", b.kernelFn, ObjString(fn), b.native)
		}
		canonical := HasCanonicalPrimitiveBinding(sym)
		switch b.helper {
		case "overridePrimitive", "restoreCanonicalPrimitive":
			// canonicalOrMake / the registry object is what makes the install
			// idempotent under the AOT HasCanonicalPrimitiveBinding guards.
			if !canonical {
				t.Errorf("%s (%s): not the canonical primitive after InstallKernelFast", b.kernelFn, b.helper)
			}
		case "overrideNative":
			// MakeNative objects must never satisfy the canonical guard: for
			// symbol?/variable? that would reinstall the weaker type-tag check.
			if canonical {
				t.Errorf("%s (%s): unexpectedly canonical after InstallKernelFast", b.kernelFn, b.helper)
			}
		case "BindSymbolFunc":
			// arity via canonicalOrMake, fn via MakeNative: either is fine.
		default:
			t.Errorf("%s: unknown helper %q", b.kernelFn, b.helper)
		}
	}
}

// TestNativeFnUndefinedWithoutShenApp pins nativeFn's error path on a sparse
// kernel. Since #49 `fn` is installed whether or not the kernel defined
// shen.app, so a lowered slice that dropped shen.app must still get the
// kernel's "fn: X is undefined" message rather than "variable shen.app not
// bound". With shen.app bound, the message still goes through shen.app, as
// the kernel's (simple-error (cn "fn: " (shen.app F " is undefined\n"
// shen.a))) does.
func TestNativeFnUndefinedWithoutShenApp(t *testing.T) {
	InstallKernelFast()
	fn := PrimFunc(MakeSymbol("fn"))
	if !IsNativeBinding(fn) {
		t.Fatalf("fn is bound to %s, want the Go native", ObjString(fn))
	}
	// Empty lambda table, no property vector: (arity F) is -1 and the assoc
	// misses, which is the undefined-function path.
	lt, pv := mustSymbol(symLambdaTable), mustSymbol(symPropertyVector)
	savedLT, savedPV := lt.value, pv.value
	defer func() { lt.value, pv.value = savedLT, savedPV }()
	PrimSet(symLambdaTable, Nil)
	PrimSet(symPropertyVector, Nil)
	callFn := func(name string) (msg string) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatalf("(fn %s) returned instead of raising", name)
			}
			o, ok := r.(Obj)
			if !ok || !IsError(o) {
				t.Fatalf("(fn %s) panicked with %v, want a KL error", name, r)
			}
			msg = mustError(o).err
		}()
		var e ControlFlow
		Call(&e, fn, MakeSymbol(name))
		return ""
	}

	app := mustSymbol(symShenApp)
	saved := app.function
	defer func() { app.function = saved }()

	app.function = nil
	if got, want := callFn("equiv.undefined-fn"), "fn: equiv.undefined-fn is undefined\n"; got != want {
		t.Errorf("shen.app unbound: (fn equiv.undefined-fn) raised %q, want %q", got, want)
	}

	// A stand-in shen.app proves the bound path still delegates to it.
	app.function = MakeNative(func(e *ControlFlow) {
		e.Return(MakeString("<app " + ObjString(e.Get(1)) + ">" + mustString(e.Get(2))))
	}, 3)
	if got, want := callFn("equiv.undefined-fn"), "fn: <app equiv.undefined-fn> is undefined\n"; got != want {
		t.Errorf("shen.app bound: (fn equiv.undefined-fn) raised %q, want %q", got, want)
	}
}
