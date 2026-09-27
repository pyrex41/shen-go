package kl

import (
	"math"
	"math/rand"
	"testing"
)

// isPreciseIntegerIlogb is the previous exponent/bit-shift implementation,
// kept as the reference the Trunc-based isPreciseInteger must agree with.
func isPreciseIntegerIlogb(f float64) bool {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return false
	}
	exp := math.Ilogb(f)
	if exp < 0 && exp != math.MinInt32 {
		return false
	}
	if exp >= 52 {
		return true
	}
	bits := math.Float64bits(f)
	return (bits << uint(12+exp)) == 0
}

func TestIsPreciseIntegerMatchesReference(t *testing.T) {
	edge := []float64{
		0, math.Copysign(0, -1), 1, -1, 0.5, -0.5, 1.5, 2.5, -2.5,
		math.Inf(1), math.Inf(-1), math.NaN(),
		math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64, math.MaxFloat64, -math.MaxFloat64,
		1 << 52, 1<<52 + 0.5, 1<<52 - 0.5, 1 << 53, 1<<53 + 2, -(1 << 53),
		4294967295, 4294967295.5, 1e300, 1e-300, 0.9999999999999999,
	}
	for _, f := range edge {
		if got, want := isPreciseInteger(f), isPreciseIntegerIlogb(f); got != want {
			t.Errorf("isPreciseInteger(%v) = %v, reference %v", f, got, want)
		}
	}
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 1_000_000; i++ {
		f := math.Float64frombits(r.Uint64())
		if i%2 == 0 { // bias toward integral and near-integral values
			f = math.Round(f) + float64(i%4)*0.25
		}
		if got, want := isPreciseInteger(f), isPreciseIntegerIlogb(f); got != want {
			t.Fatalf("isPreciseInteger(%v) = %v, reference %v", f, got, want)
		}
	}
}

// MakeNumber's fast path must give the same object kind and value as the
// general path for integral, fractional, negative-zero and boundary inputs.
func TestMakeNumberFastPathBoundaries(t *testing.T) {
	for _, f := range []float64{
		0, math.Copysign(0, -1), 1.5, -1.5,
		float64(fixnumMin), float64(fixnumMin) - 1, float64(fixnumMin) + 0.5,
		float64(fixnumMax), float64(fixnumMax) - 1, float64(fixnumMax) - 0.5,
		math.NaN(), math.Inf(1), 1e300,
	} {
		o := MakeNumber(f)
		wantFix := isPreciseIntegerIlogb(f) && f >= float64(fixnumMin) && f < float64(fixnumMax)
		if isFixnum(o) != wantFix {
			t.Errorf("MakeNumber(%v): fixnum=%v, want %v", f, isFixnum(o), wantFix)
		}
		if got := GetNumber(o); got != f && !(math.IsNaN(got) && math.IsNaN(f)) {
			t.Errorf("MakeNumber(%v) reads back %v", f, got)
		}
	}
}
