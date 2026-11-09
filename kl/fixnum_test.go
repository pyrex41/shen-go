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

// equal treats a fixnum and a boxed number as different, so = is correct only
// if every integral value in the fixnum range is a fixnum, however it was made.
func TestIntegralsInFixnumRangeAreFixnums(t *testing.T) {
	for _, v := range []int{fixnumMin, fixnumMin + 1, -(1 << 25) - 1, 1 << 25, fixnumMax - 1, 0, -1, 100000000 % fixnumMax} {
		if v < fixnumMin || v >= fixnumMax {
			continue // ±2^25 are inside the range only when it is wider than 2^26
		}
		f := float64(v)
		for name, o := range map[string]Obj{
			"MakeInteger": MakeInteger(v),
			"MakeNumber":  MakeNumber(f),
			"add":         PrimNumberAdd(MakeNumber(f/2), MakeNumber(f-f/2)),
			"multiply":    PrimNumberMultiply(MakeNumber(f), MakeNumber(1.0)),
			"divide":      PrimNumberDivide(MakeNumber(f*2), MakeNumber(2.0)),
			"slot":        slotFromNumber(f).objValue(),
		} {
			if !isFixnum(o) {
				t.Errorf("%s(%d) is not a fixnum", name, v)
			}
			if equal(o, MakeInteger(v)) != True {
				t.Errorf("%s(%d) is not = the fixnum", name, v)
			}
		}
	}
}
