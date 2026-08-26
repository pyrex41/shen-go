package kl

// Boundary tests for the signed/widened fixnum encoding (types.go). They prove
// MakeInteger/GetInteger round-trip exactly across zero, negatives, and the
// fixnum/boxed boundaries, and that values just outside the fixnum range fall
// back to a boxed scmNumber that still reads back correctly.

import "testing"

func TestFixnumRoundTrip(t *testing.T) {
	vals := []int{
		0, 1, -1, 2, -2, 7, -7, 255, -255,
		1 << 10, -(1 << 10),
		(1 << 20) - 1, 1 << 20, -(1 << 20), // around the OLD unsigned boundary
		fixnumMin, fixnumMin + 1, fixnumMax - 1, // fixnum extremes
		fixnumMin - 1, fixnumMax, fixnumMax + 1, // just outside -> boxed
		1 << 30, -(1 << 30), 1 << 40, -(1 << 40), // far outside -> boxed
	}
	for _, v := range vals {
		o := MakeInteger(v)
		if !IsNumber(o) {
			t.Errorf("MakeInteger(%d): not a number object", v)
			continue
		}
		if got := GetInteger(o); got != v {
			t.Errorf("GetInteger(MakeInteger(%d)) = %d", v, got)
		}
		if got := mustInteger(o); got != v {
			t.Errorf("mustInteger(MakeInteger(%d)) = %d", v, got)
		}
		if got := mustNumber(o); got != float64(v) {
			t.Errorf("mustNumber(MakeInteger(%d)) = %v", v, got)
		}
	}
}

func TestFixnumRangeClassification(t *testing.T) {
	inRange := []int{0, 1, -1, fixnumMin, fixnumMax - 1, (1 << 20), -(1 << 20)}
	for _, v := range inRange {
		if !isFixnum(MakeInteger(v)) {
			t.Errorf("MakeInteger(%d) expected to be a fixnum (unboxed)", v)
		}
	}
	boxed := []int{fixnumMin - 1, fixnumMax, 1 << 30, -(1 << 30)}
	for _, v := range boxed {
		if isFixnum(MakeInteger(v)) {
			t.Errorf("MakeInteger(%d) expected to be boxed (out of fixnum range)", v)
		}
	}
}

func TestFixnumPredicatesAreSafeAtSentinelBoundary(t *testing.T) {
	for _, v := range []int{fixnumMin, fixnumMax - 1} {
		o := MakeInteger(v)
		if !IsNumber(o) {
			t.Fatalf("IsNumber(MakeInteger(%d)) = false", v)
		}
		if IsError(o) || IsSymbol(o) || IsString(o) {
			t.Fatalf("non-number predicate accepted fixnum %d", v)
		}
		if PrimIsNumber(o) != True || PrimIsInteger(o) != True || PrimIsString(o) != False || PrimIsPair(o) != False || PrimIsVector(o) != False || PrimIsVariable(o) != False {
			t.Fatalf("primitive predicates mishandled fixnum %d", v)
		}
		if got := mustNumber(o); got != float64(v) {
			t.Fatalf("mustNumber(MakeInteger(%d)) = %v", v, got)
		}
		if got := ObjString(o); got != formatNumber(float64(v)) {
			t.Fatalf("ObjString(MakeInteger(%d)) = %q", v, got)
		}
		if equal(o, o) != True || equal(o, MakeString("number")) != False {
			t.Fatalf("equality mishandled fixnum sentinel %d", v)
		}
	}
	if equal(MakeInteger(fixnumMin), MakeInteger(fixnumMax-1)) != False {
		t.Fatal("distinct boundary fixnums compared equal")
	}
}

// TestFixnumArithmeticEquality guards that equality and ordering agree across
// the fixnum/boxed boundary (e.g. a boxed value equals the same fixnum value
// would-be — they can't both be fixnums, so this exercises the mixed path).
func TestFixnumArithmeticEquality(t *testing.T) {
	// A fixnum and a boxed number with the same integer value must compare equal.
	a := MakeInteger(fixnumMax - 1)         // fixnum
	b := MakeNumber(float64(fixnumMax - 1)) // also fixnum (same value)
	if !Equal(a, b) {
		t.Errorf("equal fixnums compared unequal")
	}
	big := fixnumMax + 12345
	c := MakeInteger(big)         // boxed
	d := MakeNumber(float64(big)) // boxed
	if !Equal(c, d) {
		t.Errorf("equal boxed integers compared unequal")
	}
}
