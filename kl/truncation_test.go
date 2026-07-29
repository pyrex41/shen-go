// D7: a primitive must never quietly narrow an argument into a value the
// caller did not supply.
//
// Two survivors of the D4 narrowing pass truncated instead of erroring:
//
//   - (write-byte 321 S) wrote one byte, 0x41, because PrimWriteByte did
//     byte(n) on an int that mustInteger had happily produced. -1 wrote 0xFF.
//     (write-byte 65.5 S) wrote 'A': mustInteger checks range but not
//     integrality.
//   - (n->string 65.5) returned "A" for the same reason.
//
// Tarver's reference primitives settle both. Primitives/write-byte.lsp is
// (DEFUN write-byte (Byte S) (WRITE-BYTE Byte S)), and CL's WRITE-BYTE signals
// a type error for anything that is not an (UNSIGNED-BYTE 8) -- shen-cl does
// exactly that. Primitives/n-to-string.lsp:5 is (CODE-CHAR N) inside a
// trap-error that raises "~A is not a natural number".
//
// The ports split 2-2 on both (shen-go and shen-rust truncated, shen-cl and
// shen-lua errored); the reference implementation outranks the split.
//
// The errors must stay inside the Shen error system: trap-error has to catch
// them, and the message must name the value the caller passed, not a narrowed
// one.
package kl

import "testing"

func TestWriteByteRejectsNonBytes(t *testing.T) {
	var ctx ControlFlow

	for _, tc := range []struct{ name, expr, want string }{
		{
			"above 255",
			`(write-byte 321 (value *stoutput*))`,
			`"321 is not a byte"`,
		},
		{
			"exactly 256",
			`(write-byte 256 (value *stoutput*))`,
			`"256 is not a byte"`,
		},
		{
			"negative",
			`(write-byte -1 (value *stoutput*))`,
			`"-1 is not a byte"`,
		},
		{
			"a code point past the Latin-1 range",
			`(write-byte 20013 (value *stoutput*))`,
			`"20013 is not a byte"`,
		},
		{
			"fractional",
			`(write-byte 65.5 (value *stoutput*))`,
			`"65.5 is not a valid integer"`,
		},
		{
			// D4 pinned this spelling; keep it.
			"+inf",
			`(write-byte ` + klInfExpr + ` (value *stoutput*))`,
			`"inf is not a valid integer"`,
		},
	} {
		if got := trapped(&ctx, tc.expr); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.name, got, tc.want)
		}
	}

	// The whole byte range still writes, and the evaluator survives.
	for _, expr := range []string{
		`(write-byte 0 (value *stoutput*))`,
		`(write-byte 255 (value *stoutput*))`,
	} {
		if got := trapped(&ctx, expr); got == `"0 is not a byte"` || got == `"255 is not a byte"` {
			t.Errorf("%s: rejected an in-range byte (%s)", expr, got)
		}
	}
	if got := ObjString(evalString(&ctx, "(+ 40 2)")); got != "42" {
		t.Fatalf("post-error eval: got %q, want 42", got)
	}
}

func TestNumberToStringRejectsFractions(t *testing.T) {
	var ctx ControlFlow

	for _, tc := range []struct{ name, expr, want string }{
		{
			"a fractional code",
			`(n->string 65.5)`,
			`"65.5 is not a natural number"`,
		},
		{
			"a fraction below 1",
			`(n->string 0.5)`,
			`"0.5 is not a natural number"`,
		},
		{
			"a negative fraction",
			`(n->string -0.5)`,
			`"-0.5 is not a natural number"`,
		},
		// D4 pinned these spellings; keep them.
		{"negative", `(n->string -1)`, `"-1 is not a natural number"`},
		{"+inf", `(n->string ` + klInfExpr + `)`, `"inf is not a natural number"`},
	} {
		if got := trapped(&ctx, tc.expr); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.name, got, tc.want)
		}
	}

	// Whole codes still convert.
	for _, tc := range []struct{ expr, want string }{
		{`(n->string 65)`, `"A"`},
		{`(n->string 0)`, "\"\x00\""},
		{`(n->string 20013)`, `"中"`},
	} {
		if got := ObjString(evalString(&ctx, tc.expr)); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.expr, got, tc.want)
		}
	}
}

// A lone surrogate is NOT the same defect class, and is deliberately left
// alone.
//
// The reported symptom is that (n->string 55296) answers U+FFFD instead of
// erroring. But the reference primitive is (CODE-CHAR N), and in SBCL
// CODE-CHAR accepts the surrogate range: shen-cl round-trips
// (string->n (n->string 55296)) back to 55296. So the authority that settles
// the fraction case says here that n->string should NOT error. shen-rust and
// shen-lua do error, shen-go and shen-cl do not -- another 2-2 split, but this
// time the reference sides with shen-go.
//
// Go could not error "correctly" anyway: a Go string is UTF-8 and cannot hold
// a lone surrogate at all, so string(rune(0xD800)) is U+FFFD by language
// definition, not by a narrowing bug. Matching shen-cl here would mean
// changing shen-go's string representation, which is nothing like a range
// check on write-byte.
//
// Pinned so the behaviour is a decision rather than an accident.
func TestNumberToStringMapsLoneSurrogates(t *testing.T) {
	var ctx ControlFlow
	for _, expr := range []string{
		`(string->n (n->string 55296))`, // U+D800, first lone surrogate
		`(string->n (n->string 57343))`, // U+DFFF, last
	} {
		if got := ObjString(evalString(&ctx, expr)); got != "65533" {
			t.Errorf("%s: got %s, want 65533 (U+FFFD)", expr, got)
		}
	}
	// The code points either side of the surrogate block round-trip exactly.
	for _, tc := range []struct{ expr, want string }{
		{`(string->n (n->string 55295))`, "55295"},
		{`(string->n (n->string 57344))`, "57344"},
	} {
		if got := ObjString(evalString(&ctx, tc.expr)); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.expr, got, tc.want)
		}
	}
}
