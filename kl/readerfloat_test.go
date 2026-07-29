// D6: numeric literals must parse to the correctly rounded double.
//
// The kernel reader (kernel/sources/reader.shen) builds a literal's value with
// its own arithmetic: (compute-E N Log10) is (* N (expt 10 Log10)) and expt is
// repeated multiplication/division. In a bignum host (shen-cl) that is exact.
// In a float64-only host every step rounds, and the error compounds: 10^25
// came out as 9.999999999999999e24 and 10^300 as 1.0000000000000002e300, so
// the literal 1e300 was not the double 1e300 at all. 10^-6 was already wrong
// (1.0000000000000002e-06), which made 0.000001 a different number from what
// every other reader produces.
//
// Repeated multiplication is exact only for 10^0..10^22 and 10^-1..10^-5; the
// first divergences are 10^25 and 10^-6. Answering (expt 10 N) with the
// correctly rounded power fixes every literal without touching the kernel's
// summation, so nothing else about reading changes.
package kl

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

// readerLiterals are literals whose value the naive repeated-multiplication
// expt got wrong, plus neighbours that it already got right.
var readerLiterals = []string{
	"1e300", "1e150", "1e25", "1e23", "1e22", "1e19", "1e0",
	"1e-6", "1e-30", "1e-300", "0.000001", "0.001",
	"-1e300", "1e308", "-1e-300", "3.14159", "0.5", "42",
	"1234567890123456",
}

// Two roundings the exact powers do NOT remove, because they are in the
// kernel's own algorithm rather than in expt:
//
//   - (compute-E N Log10) is (* N (expt 10 Log10)), so a mantissa that is not
//     a power of ten is rounded once into N and once again by the multiply:
//     2.5e-30 lands one ulp off the correctly rounded double.
//   - compute-integer and compute-fraction sum digit*10^pos, so a long digit
//     string accumulates: 123456789012345678 (past 2^53) reads as
//     123456789012345664, and 0.8660254037844 lands one ulp high. Exactly two
//     of the 3539 numeric literals in kernel/ and src/ are affected, and both
//     read the same before and after this change.
//
// Both are inherent to running the kernel's reader arithmetic in a float64-only
// numeric tower; shen-lua shows the same values, and shen-cl and shen-rust
// avoid them only by having a bignum (resp. i64) integer type. Removing them
// would mean replacing the kernel's arithmetic rather than correcting one
// operation of it. Pinned here so the limitation is visible and deliberate.
func TestKernelReaderRoundingsThatRemain(t *testing.T) {
	for _, tc := range []struct {
		lit  string
		want float64
	}{
		{"2.5e-30", 2.5e-30 * (1 + 1e-16)}, // one ulp above 2.5e-30
		{"123456789012345678", 123456789012345664},
		{"0.8660254037844", 0.8660254037844001},
	} {
		forms, err := ShenReadSExprs([]byte(tc.lit))
		if err != nil {
			t.Fatalf("ShenReadSExprs(%q): %v", tc.lit, err)
		}
		if got := GetNumber(Car(forms)); got != tc.want {
			t.Errorf("ShenReadSExprs(%q) = %v, want %v", tc.lit, got, tc.want)
		}
	}
}

// pow10 must be the correctly rounded power of ten for every exponent a reader
// can produce, including the clamped ends.
func TestPow10IsCorrectlyRounded(t *testing.T) {
	for n := -350; n <= 320; n++ {
		want, err := strconv.ParseFloat("1e"+strconv.Itoa(n), 64)
		if err != nil && !strings.Contains(err.Error(), "out of range") {
			t.Fatalf("ParseFloat 1e%d: %v", n, err)
		}
		if got := pow10(n); got != want {
			t.Errorf("pow10(%d) = %v, want %v", n, got, want)
		}
	}
	if got := pow10(100000); !math.IsInf(got, 1) {
		t.Errorf("pow10(100000) = %v, want +Inf", got)
	}
	if got := pow10(-100000); got != 0 {
		t.Errorf("pow10(-100000) = %v, want 0", got)
	}
}

// The Go-native Shen reader must agree with a correctly rounded parse. It is
// the file-load path, so a wrong value here is baked into every loaded source.
func TestShenReaderParsesLiteralsCorrectlyRounded(t *testing.T) {
	for _, lit := range readerLiterals {
		want, _ := strconv.ParseFloat(lit, 64)
		forms, err := ShenReadSExprs([]byte(lit))
		if err != nil {
			t.Fatalf("ShenReadSExprs(%q): %v", lit, err)
		}
		got := GetNumber(Car(forms))
		if got != want {
			t.Errorf("ShenReadSExprs(%q) = %v, want %v", lit, got, want)
		}
	}
}

// kl's own s-expression reader was always exact (it uses strconv.ParseFloat);
// pin that the two readers now agree, since only one of them was wrong.
func TestSexpAndShenReadersAgreeOnLiterals(t *testing.T) {
	for _, lit := range readerLiterals {
		o, err := NewSexpReader(strings.NewReader(lit+" "), false).Read()
		if err != nil {
			t.Fatalf("NewSexpReader(%q): %v", lit, err)
		}
		forms, err := ShenReadSExprs([]byte(lit))
		if err != nil {
			t.Fatalf("ShenReadSExprs(%q): %v", lit, err)
		}
		if a, b := GetNumber(o), GetNumber(Car(forms)); a != b {
			t.Errorf("%q: NewSexpReader = %v, ShenReadSExprs = %v", lit, a, b)
		}
	}
}

// The reported symptom: 1e300 printed as a number it was not.
func TestBigLiteralRendersAsItsOwnDigits(t *testing.T) {
	var ctx ControlFlow
	want := "1" + strings.Repeat("0", 300)
	forms, err := ShenReadSExprs([]byte("1e300"))
	if err != nil {
		t.Fatal(err)
	}
	if got := ObjString(Car(forms)); got != want {
		t.Errorf("ShenReadSExprs(1e300) prints %s, want %s", got, want)
	}
	if got := ObjString(evalString(&ctx, "1e300")); got != want {
		t.Errorf("eval 1e300 prints %s, want %s", got, want)
	}
}
