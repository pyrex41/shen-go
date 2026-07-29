// D2: finite floats at or beyond 2^63 must not print as a saturated int64.
//
// 1e19 is an ordinary, exactly-representable double. Both number printers
// (scmHead.GoString and PrimStr) asked isPreciseInteger whether the value was
// an integer and, on "yes", narrowed with a bare int(f). Go leaves an
// out-of-range float64 -> int conversion implementation-defined; on arm64 it
// saturates, so 1e19, 1e300 and +Inf all came back as 9223372036854775807 and
// -1e19 as -9223372036854775808 -- a different number than the one held.
//
// isPreciseInteger was also lying about non-finite values: math.Ilogb answers
// MaxInt32 for +-Inf and NaN, which took its `exp >= 52` shortcut and called
// them integers.
package kl

import (
	"math"
	"strings"
	"testing"
)

// numberRenderings checks both printers at once. `str` (PrimStr) and the
// GoString used by print/error messages must never disagree about a number.
func numberRenderings(t *testing.T, f float64) (goString, str string) {
	t.Helper()
	o := MakeNumber(f)
	return ObjString(o), GetString(PrimStr(o))
}

func TestLargeFiniteFloatsDoNotSaturate(t *testing.T) {
	for _, tc := range []struct {
		f    float64
		want string
	}{
		// Beyond 2^63: no int can hold these, so they print as full
		// positional digits (D5) rather than as a saturated int64.
		{1e19, "10000000000000000000"},
		{-1e19, "-10000000000000000000"},
		{1e300, "1" + strings.Repeat("0", 300)},
		{-1e300, "-1" + strings.Repeat("0", 300)},
		// Exactly 2^63 -- one past math.MaxInt64, which is itself not
		// representable as a float64. The digits are the shortest round-trip
		// decimal laid out positionally, which is what shen-rust prints for
		// the same double (shen-cl's 9223372036854775808 is a bignum, not a
		// double).
		{9223372036854775808.0, "9223372036854776000"},
		// The largest exactly-representable integer strictly below 2^63
		// (2^63 - 1024) still renders as an exact integer.
		{9223372036854774784.0, "9223372036854774784"},
		// Ordinary values are untouched.
		{0, "0"},
		{42, "42"},
		{-42, "-42"},
		{3.5, "3.5"},
		{1e18, "1000000000000000000"},
		{-1e18, "-1000000000000000000"},
	} {
		gotGo, gotStr := numberRenderings(t, tc.f)
		if gotGo != tc.want {
			t.Errorf("ObjString(%v): got %q, want %q", tc.f, gotGo, tc.want)
		}
		if gotStr != tc.want {
			t.Errorf("str(%v): got %q, want %q", tc.f, gotStr, tc.want)
		}
	}
}

// Non-finite values must not print as a number they are not. shen-lua and
// shen-rust both render these as inf/-inf; this is a consequence of not
// narrowing, not a decision about whether overflow should signal at all.
func TestNonFiniteFloatsRender(t *testing.T) {
	for _, tc := range []struct {
		f    float64
		want string
	}{
		{math.Inf(1), "inf"},
		{math.Inf(-1), "-inf"},
		{math.NaN(), "nan"},
	} {
		gotGo, gotStr := numberRenderings(t, tc.f)
		if gotGo != tc.want {
			t.Errorf("ObjString(%v): got %q, want %q", tc.f, gotGo, tc.want)
		}
		if gotStr != tc.want {
			t.Errorf("str(%v): got %q, want %q", tc.f, gotStr, tc.want)
		}
	}
}

// isPreciseInteger must not claim +-Inf/NaN are integers: math.Ilogb returns
// MaxInt32 for them, which used to fall straight into the `exp >= 52` branch.
func TestIsPreciseIntegerRejectsNonFinite(t *testing.T) {
	for _, f := range []float64{math.Inf(1), math.Inf(-1), math.NaN()} {
		if isPreciseInteger(f) {
			t.Errorf("isPreciseInteger(%v) = true, want false", f)
		}
	}
	for _, f := range []float64{0, 1, -1, 42, 1e18, 1e19, 1e300} {
		if !isPreciseInteger(f) {
			t.Errorf("isPreciseInteger(%v) = false, want true", f)
		}
	}
	for _, f := range []float64{0.5, -3.25, 1e-300} {
		if isPreciseInteger(f) {
			t.Errorf("isPreciseInteger(%v) = true, want false", f)
		}
	}
}

// The same values reached through evaluation rather than constructed in Go.
// An overflow product is built as (* (/ 1.0 1e-300) (/ 1.0 1e-300)) so it does
// not depend on how the reader handles an `inf` literal.
func TestLargeFloatRenderingThroughEval(t *testing.T) {
	var ctx ControlFlow

	for _, tc := range []struct{ expr, want string }{
		{`1e19`, "10000000000000000000"},
		{`1e300`, "1" + strings.Repeat("0", 300)},
		{`(* (/ 1.0 1e-300) (/ 1.0 1e-300))`, "inf"},
		{`(str 1e19)`, `"10000000000000000000"`},
		{`(str 1e300)`, `"` + "1" + strings.Repeat("0", 300) + `"`},
		{`(str (* (/ 1.0 1e-300) (/ 1.0 1e-300)))`, `"inf"`},
		{`(str 2.5)`, `"2.5"`},
		{`(str 7)`, `"7"`},
	} {
		if got := ObjString(evalString(&ctx, tc.expr)); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.expr, got, tc.want)
		}
	}
}

// D5 (cross-port rendering convention): a FINITE INTEGRAL float renders as
// full positional decimal digits, with no exponent, whatever its magnitude.
//
// shen-cl prints 1e19 as 10000000000000000000 and shen-rust prints the same;
// shen-go was the only port emitting an exponent (1e+19). The digits are the
// shortest round-trip decimal rendered positionally -- exactly
// strconv.FormatFloat(f, 'f', -1, 64) -- which is byte-for-byte what shen-rust
// produces for the same float64 (2^70 -> 1180591620717411300000, not the
// double's exact value 1180591620717411303424; shen-cl prints the latter only
// because its numeric tower makes that value a bignum, not a double).
//
// Non-integral finite floats keep the shortest round-trip form (2.5, 1e-300),
// and +-inf/nan keep their existing spellings.
func TestIntegralFloatsRenderPositionally(t *testing.T) {
	for _, tc := range []struct {
		f    float64
		want string
	}{
		{1e19, "10000000000000000000"},
		{-1e19, "-10000000000000000000"},
		{1e22, "10000000000000000000000"},
		// 2^63 exactly: no int64 holds it, and it is integral, so all digits.
		// 2^63: shortest round-trip digits laid out positionally, which is
		// byte-for-byte what shen-rust prints for the same double.
		{9223372036854775808.0, "9223372036854776000"},
		// 2^70: shortest round-trip digits, zero padded -- what shen-rust prints.
		{1180591620717411303424.0, "1180591620717411300000"},
		// 1e300 positional is 301 characters.
		{1e300, "1" + strings.Repeat("0", 300)},
		{-1e300, "-1" + strings.Repeat("0", 300)},
		{math.MaxFloat64, "17976931348623157" + strings.Repeat("0", 292)},
		// Non-integral values are untouched.
		{2.5, "2.5"},
		{1e-300, "1e-300"},
		{0.1, "0.1"},
		{0, "0"},
		{42, "42"},
	} {
		gotGo, gotStr := numberRenderings(t, tc.f)
		if gotGo != tc.want {
			t.Errorf("ObjString(%v): got %q, want %q", tc.f, gotGo, tc.want)
		}
		if gotStr != tc.want {
			t.Errorf("str(%v): got %q, want %q", tc.f, gotStr, tc.want)
		}
	}
}

// The convention says nothing about a trailing ".0": shen-go has a single
// float64 number type and never emits one. (shen-rust's 10000000000.0 for
// (/ 1e300 1e290) and shen-cl's 1.0 for (* 2.0 0.5) come from their two-tier
// numeric towers, and the two ports do not agree with each other.)
func TestIntegralFloatsNeverGetATrailingPointZero(t *testing.T) {
	var ctx ControlFlow
	for _, tc := range []struct{ expr, want string }{
		{`(* 2.0 0.5)`, "1"},
		{`(/ 1e300 1e290)`, "10000000000"},
		{`(str (* 2.0 0.5))`, `"1"`},
	} {
		if got := ObjString(evalString(&ctx, tc.expr)); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.expr, got, tc.want)
		}
	}
}
