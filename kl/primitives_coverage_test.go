package kl

import (
	"bytes"
	"testing"
)

// TestPrimitivesViaEval drives the bulk of the primitive surface through the
// evaluator, covering arithmetic, list ops, type predicates, string ops, logic
// and global value/set that the existing suite left at 0% unit coverage.
// (Bare unbound symbols self-evaluate, so e.g. (symbol? hello) sees the symbol.)
func TestPrimitivesViaEval(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		// arithmetic
		{"subtract int", "(- 5 3)", "2"},
		// Issue #11: floats render in shortest round-trip form (3.5, not 3.500000).
		{"subtract float", "(- 5.5 2.0)", "3.5"},
		{"multiply int", "(* 6 7)", "42"},
		{"multiply float", "(* 1.5 3.0)", "4.5"},
		{"divide whole", "(/ 20 4)", "5"},
		{"divide fractional", "(/ 7 2)", "3.5"},

		// list ops
		{"hd", "(hd (cons 1 (cons 2 ())))", "1"},
		{"tl", "(tl (cons 1 (cons 2 ())))", "(2)"},
		{"cons builds pair", "(cons 1 (cons 2 ()))", "(1 2)"},

		// type predicates
		{"number? yes", "(number? 42)", "true"},
		{"number? no", "(number? foo)", "false"},
		{"string? yes", `(string? "hi")`, "true"},
		{"string? no", "(string? 1)", "false"},
		{"symbol? yes", "(symbol? hello)", "true"},
		{"symbol? no", "(symbol? 1)", "false"},
		{"variable? upper", "(variable? X)", "true"},
		{"variable? lower", "(variable? x)", "false"},
		{"integer? yes", "(integer? 42)", "true"},
		{"integer? no", "(integer? 4.5)", "false"},
		{"cons? yes", "(cons? (cons 1 ()))", "true"},
		{"cons? no", "(cons? 1)", "false"},
		{"absvector? yes", "(absvector? (absvector 3))", "true"},
		{"absvector? no", "(absvector? 1)", "false"},

		// string ops
		{"string->n", `(string->n "A")`, "65"},
		{"n->string", "(n->string 65)", `"A"`},
		{"cn", `(cn "foo" "bar")`, `"foobar"`},
		{"tlstr", `(tlstr "hello")`, `"ello"`},
		{"pos", `(pos "hello" 1)`, `"e"`},

		// logic
		{"not true", "(not true)", "false"},
		{"not false", "(not false)", "true"},

		// str renders each atom type
		{"str number", "(str 42)", `"42"`},
		{"str float", "(str 4.5)", `"4.5"`},
		{"str symbol", "(str foo)", `"foo"`},
		{"str bool", "(str true)", `"true"`},
		{"str nil", "(str ())", `"()"`},

		// global value/set
		{"set then value", "(do (set foo 99) (value foo))", "99"},

		// vector set/get round trip (address-> returns the vector)
		{"vector round trip", "(<-address (address-> (absvector 3) 1 7) 1)", "7"},
	}

	var ctx ControlFlow
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ObjString(evalString(&ctx, c.input))
			if got != c.want {
				t.Errorf("%s: got %q, want %q", c.input, got, c.want)
			}
		})
	}
}

// TestFloatShortestFormIssue11 pins the issue #11 fix: non-integral floats
// render in Go's shortest round-trip form (2.5, not 2.500000) via both the
// `str` primitive (PrimStr) and ObjString (GoString) — matching shen-cl and
// the other ports. Integral-valued floats keep printing without a decimal,
// as the kernel expects.
func TestFloatShortestFormIssue11(t *testing.T) {
	var ctx ControlFlow

	// `str` primitive surface.
	strCases := []struct{ input, want string }{
		{"(str 2.5)", `"2.5"`},
		{"(str 4.5)", `"4.5"`},
		{"(str (/ 7 2))", `"3.5"`},
		{"(str (/ 1 4))", `"0.25"`},
		{"(str 0.1)", `"0.1"`},
		// integral-valued floats and integers print without a decimal point
		{"(str (* 2.0 3.0))", `"6"`},
		{"(str 42)", `"42"`},
	}
	for _, c := range strCases {
		if got := ObjString(evalString(&ctx, c.input)); got != c.want {
			t.Errorf("%s: got %q, want %q", c.input, got, c.want)
		}
	}

	// ObjString / GoString surface (the REPL/test rendering path).
	objCases := []struct{ input, want string }{
		{"(/ 7 2)", "3.5"},
		{"(* 1.5 2.5)", "3.75"},
		{"(/ 1 4)", "0.25"},
		{"(* 2.0 3.0)", "6"},
	}
	for _, c := range objCases {
		if got := ObjString(evalString(&ctx, c.input)); got != c.want {
			t.Errorf("ObjString %s: got %q, want %q", c.input, got, c.want)
		}
	}
}

// TestVectorUninitializedSlot exercises the PrimVectorGet nil-slot path, which
// returns the special `undefined` raw object.
func TestVectorUninitializedSlot(t *testing.T) {
	var ctx ControlFlow
	got := evalString(&ctx, "(<-address (absvector 3) 0)")
	if got != undefined {
		t.Errorf("uninitialized slot: got %q, want the undefined object", ObjString(got))
	}
}

// TestValueUnboundPanics covers the unbound branch of PrimValue.
func TestValueUnboundPanics(t *testing.T) {
	var ctx ControlFlow
	got := ObjString(evalString(&ctx,
		`(trap-error (value never-bound-xyz) (lambda E (error-to-string E)))`))
	want := `"variable never-bound-xyz not bound"`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestListAndTypeHelpers exercises the exported list helpers and type accessors
// directly (Car/Cdr/Cons/Cadr/cadddr and the Is*/Get* family).
func TestListAndTypeHelpers(t *testing.T) {
	l := Cons(MakeInteger(1),
		Cons(MakeInteger(2),
			Cons(MakeInteger(3),
				Cons(MakeInteger(4), Nil))))

	if got := ObjString(Car(l)); got != "1" {
		t.Errorf("Car: got %q", got)
	}
	if got := ObjString(Cadr(l)); got != "2" {
		t.Errorf("Cadr: got %q", got)
	}
	if got := ObjString(cadddr(l)); got != "4" {
		t.Errorf("cadddr: got %q", got)
	}
	if got := ObjString(Cdr(l)); got != "(2 3 4)" {
		t.Errorf("Cdr: got %q", got)
	}
	if got := ObjString(Cons(MakeInteger(5), Nil)); got != "(5)" {
		t.Errorf("Cons: got %q", got)
	}

	if ok, _ := isPair(l); !ok {
		t.Error("isPair(list) should be true")
	}
	if ok, _ := isPair(MakeInteger(1)); ok {
		t.Error("isPair(int) should be false")
	}

	if !IsNumber(MakeNumber(3.5)) {
		t.Error("IsNumber(3.5) should be true")
	}
	if IsNumber(MakeString("x")) {
		t.Error("IsNumber(string) should be false")
	}
	if !IsString(MakeString("hi")) {
		t.Error("IsString should be true")
	}
	if got := GetString(MakeString("hi")); got != "hi" {
		t.Errorf("GetString: got %q", got)
	}
	if got := GetSymbol(MakeSymbol("sym")); got != "sym" {
		t.Errorf("GetSymbol: got %q", got)
	}
	if got := GetInteger(MakeInteger(42)); got != 42 {
		t.Errorf("GetInteger(fixnum): got %d", got)
	}
	if got := GetInteger(MakeNumber(3.9)); got != 3 {
		t.Errorf("GetInteger(float): got %d", got)
	}
}

// TestStinputEOFFlag covers the ResetStinputEOF/StinputEOF accessors used by the
// CLI to detect stdin EOF and exit the REPL cleanly.
func TestStinputEOFFlag(t *testing.T) {
	ResetStinputEOF()
	if StinputEOF() {
		t.Error("StinputEOF should be false after reset")
	}
}

// TestCharStreamPredicates covers shen.char-stinput?/shen.char-stoutput?, which
// always report false in this implementation.
func TestCharStreamPredicates(t *testing.T) {
	var ctl ControlFlow
	for _, name := range []string{"shen.char-stinput?", "shen.char-stoutput?"} {
		fn := ctl.Global(MakeSymbol(name))
		if got := Call(&ctl, fn, MakeInteger(0)); got != False {
			t.Errorf("%s: got %q, want false", name, ObjString(got))
		}
	}
}

// TestIfPrimitiveDirectCall covers PrimIf invoked as an ordinary function (as
// opposed to the if special form handled by the evaluator/compiler).
func TestIfPrimitiveDirectCall(t *testing.T) {
	var ctl ControlFlow
	fn := ctl.Global(MakeSymbol("if"))
	if got := Call(&ctl, fn, True, MakeInteger(1), MakeInteger(2)); ObjString(got) != "1" {
		t.Errorf("if true: got %q, want 1", ObjString(got))
	}
	if got := Call(&ctl, fn, False, MakeInteger(1), MakeInteger(2)); ObjString(got) != "2" {
		t.Errorf("if false: got %q, want 2", ObjString(got))
	}
}

// TestObjStringRendering covers the GoString branches for the non-atom object
// types (vector, procedure, error, stream, primitive, compiled fn, improper
// pair) that ObjString must render.
func TestObjStringRendering(t *testing.T) {
	var ctl ControlFlow

	checks := []struct {
		name string
		obj  Obj
		want string
	}{
		{"vector", MakeVector(2), "#vector"},
		{"error", MakeError("boom"), "Error(boom)"},
		{"stream", MakeStream(new(bytes.Buffer)), "#stream"},
		{"improper pair", Cons(MakeInteger(1), MakeInteger(2)), "(1 . 2)"},
		{"primitive", ctl.Global(MakeSymbol("hd")), "#primitive(hd)"},
	}
	for _, c := range checks {
		if got := ObjString(c.obj); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}

	// procedure (from a lambda) and a compiled function (from a defun)
	if got := ObjString(evalString(&ctl, "(lambda x x)")); got != "#procedure" {
		t.Errorf("procedure: got %q, want #procedure", got)
	}
	evalString(&ctl, "(defun rendered-fn (x) x)")
	if got := ObjString(ctl.Global(MakeSymbol("rendered-fn"))); got != "#compiled(rendered-fn)" {
		t.Errorf("compiled fn: got %q, want #compiled(rendered-fn)", got)
	}
}

// TestStrRendering covers the PrimStr branches for non-atom types.
func TestStrRendering(t *testing.T) {
	var ctl ControlFlow
	str := ctl.Global(MakeSymbol("str"))

	if got := GetString(Call(&ctl, str, MakeError("boom"))); got != "#<err boom>" {
		t.Errorf("str error: got %q", got)
	}
	if got := GetString(Call(&ctl, str, MakeStream(new(bytes.Buffer)))); got != "#<stream >" {
		t.Errorf("str stream: got %q", got)
	}
	if got := GetString(Call(&ctl, str, ctl.Global(MakeSymbol("hd")))); got != "#<primitive hd>" {
		t.Errorf("str primitive: got %q", got)
	}
	if got := GetString(Call(&ctl, str, evalString(&ctl, "(lambda x x)"))); got != "#<procedure >" {
		t.Errorf("str procedure: got %q", got)
	}
}

// TestPrimHash covers the native hash: deterministic, equal keys hash equal,
// result stays in the kernel's [1, K-1] bucket-index contract, and limit<=1 is 1.
func TestPrimHash(t *testing.T) {
	const k = 256
	limit := MakeInteger(k)

	// Determinism + equal keys -> equal hash.
	a := GetInteger(PrimHash(MakeString("session-token-42"), limit))
	b := GetInteger(PrimHash(MakeString("session-token-42"), limit))
	if a != b {
		t.Errorf("hash not deterministic: %d vs %d", a, b)
	}

	// Range [1, K-1] for many distinct keys (offset 3+hash must fit a 3+K vector).
	for i := 0; i < 500; i++ {
		h := GetInteger(PrimHash(MakeString("k"+ObjString(MakeInteger(i))), limit))
		if h < 1 || h > k-1 {
			t.Fatalf("hash %d out of [1,%d] for key %d", h, k-1, i)
		}
	}

	// limit <= 1 collapses to 1.
	if GetInteger(PrimHash(MakeString("x"), MakeInteger(1))) != 1 {
		t.Error("hash with limit 1 should be 1")
	}

	// Reasonable distribution: distinct keys shouldn't all collide.
	seen := map[int]bool{}
	for i := 0; i < 100; i++ {
		seen[GetInteger(PrimHash(MakeString("distinct-"+ObjString(MakeInteger(i))), limit))] = true
	}
	if len(seen) < 50 {
		t.Errorf("poor hash distribution: only %d buckets for 100 keys", len(seen))
	}
}
