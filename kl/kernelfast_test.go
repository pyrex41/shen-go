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
	if Call(&e, PrimFunc(MakeSymbol("fail"))) != MakeSymbol("...") {
		t.Fatalf("fail should return ...")
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

	func() {
		defer func() {
			if recover() == nil {
				t.Fatalf("<-vector of unwritten slot should raise")
			}
		}()
		Call(&e, PrimFunc(MakeSymbol("<-vector")), v, MakeInteger(2))
	}()

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
// rebinding is covered automatically. The names no other kl test bindDummy's
// (shen.abs, sum, union, string->symbol, …) are exactly the ones that were
// left unbound before, so this test fails under the old guard in any order.
func TestInstallKernelFastBindsEveryOverride(t *testing.T) {
	rows := parseInstallKernelFast(t)
	if len(rows) == 0 {
		t.Fatal("parseInstallKernelFast returned no rows")
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
