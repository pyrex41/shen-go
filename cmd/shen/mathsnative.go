package main

import (
	"math"
	"os"

	"github.com/pyrex41/shen-go/kl"
)

// Native floor, ceiling, round, mod, power, gcd, lcd and isqrt for the
// stdlib's Maths package.
//
// StLib's versions are loops written in Shen. floor, ceiling and round go
// through maths.rounding-loop, which builds the answer one decimal digit at a
// time from 10^15 down, calling power at every step: about 100 µs a call.
// mod, div, modf, float->pair and random all go through floor, so any modular
// arithmetic written against the stdlib paid that on every operation. power
// recurses once per unit of the exponent. gcd tries every divisor down from
// the smaller argument, lcd every odd one up to it, and isqrt counts up to
// the root.
//
// Each native is exact on a proven domain and hands every other input to the
// original Shen definition, captured before the rebinding. Behaviour outside
// the domain is therefore the original's by construction: error text for a
// non-number, division by zero, and non-termination or a fatal stack
// overflow where the original has one.
//
// floor, ceiling, round, mod. The domain is a finite number of magnitude
// below 2^53. There every Guess rounding-loop forms is an exactly
// representable integer and every comparison is exact, so the loop returns
// exactly:
//
//	floor:   math.Floor(N)
//	ceiling: math.Ceil(N)
//	round:   the nearer integer, ties up; negative N goes through
//	         (~ (round (~ N))), so ties are away from zero: math.Round(N).
//	         Its tie test compares Up = ceil-N with Down = N-floor, and both
//	         differences are exact in this range (Sterbenz), so it agrees
//	         with math.Round on every input, not just away from ties.
//
// mod is not integer modulus: it is the float formula
// (let Div (/ X Y) FloorDiv (floor Div) (* Y (- Div FloorDiv))), rounded when
// X and Y are both integers. The native evaluates that formula with the same
// float operations and the natives above, and defers to the original when Y
// is 0 or an intermediate leaves the domain.
//
// power N M is (* N (power N (- M 1))) down to (power _ 0) = 1: M
// multiplications by N, each rounded, which the native repeats in the same
// order. The domain is an integer 0 <= M <= 100000. A negative or fractional
// M never reaches 0 and the original recurses until the Go stack overflows
// (fatal, not catchable), and so does an M much past 100000; those still go
// to the original.
//
// gcd M N on integers is gcd(|M|, |N|), computed by Euclid, except that the
// original raises "division by zero" when either is 0 or |M| = |N| (its loop
// starts at a divisor of 0), and "gcd expects integer inputs" for anything
// else; both go to the original. Below 2^53 its (integer? (/ M D)) test is
// exact divisibility: a non-integer quotient M/D < 2^52 is at least 1/D from
// an integer, and the division's rounding error is smaller than that.
//
// lcd M N on integers is 2 when both are even. Otherwise it is the smallest
// odd D >= 3 with D <= min(M, N) dividing both, else 1. When both are
// positive, that D is the smallest prime factor of the odd part of
// gcd(M, N), found by trial division up to its square root. When either is
// <= 0 the loop never runs and the answer is 1. Non-integers go to the
// original.
//
// isqrt N counts S up from 0 to the first S with S*S >= N, returning S on
// equality and S-1 otherwise: floor(sqrt(N)) for N >= 0, and -1 for every
// negative N. Below 2^53 the native corrects math.Sqrt's estimate with exact
// products (S < 2^27).
//
// cmd/shen/mathsnative_test.go runs thousands of cases through both sides and
// requires byte-identical output. SHEN_NO_MATHS_NATIVE=1 keeps the Shen
// definitions.

// mathsExactLimit bounds the float domain: |x| < 2^53.
const mathsExactLimit = 1 << 53

// mathsPowerLimit is the largest exponent the native power takes. The
// original's recursion is verified to survive it; somewhere below 10^6 it
// overflows the Go stack.
const mathsPowerLimit = 100000

// mathsNumber is x's value when x is a number in the float domain.
func mathsNumber(x kl.Obj) (float64, bool) {
	if !kl.IsNumber(x) {
		return 0, false
	}
	f := kl.GetNumber(x)
	return f, f > -mathsExactLimit && f < mathsExactLimit // false for NaN, +-Inf
}

// mathsInt is x's value when x is an integer below 2^53 in magnitude.
func mathsInt(x kl.Obj) (int64, bool) {
	f, ok := mathsNumber(x)
	if !ok || kl.PrimIsInteger(x) != kl.True {
		return 0, false
	}
	return int64(f), true
}

// installMathsNatives rebinds the Maths natives once the stdlib is loaded. A
// name the stdlib did not bind (SHEN_NO_STDLIB, or a failed load) is left
// alone.
func installMathsNatives() {
	if os.Getenv("SHEN_NO_MATHS_NATIVE") != "" {
		return
	}
	unary := func(f func(float64) float64) func([]kl.Obj) (kl.Obj, bool) {
		return func(a []kl.Obj) (kl.Obj, bool) {
			if v, ok := mathsNumber(a[0]); ok {
				return kl.MakeNumber(f(v)), true
			}
			return nil, false
		}
	}
	bindMathsNative("floor", 1, unary(math.Floor))
	bindMathsNative("ceiling", 1, unary(math.Ceil))
	bindMathsNative("round", 1, unary(math.Round))
	bindMathsNative("mod", 2, func(a []kl.Obj) (kl.Obj, bool) { return mathsMod(a[0], a[1]) })
	bindMathsNative("power", 2, func(a []kl.Obj) (kl.Obj, bool) { return mathsPower(a[0], a[1]) })
	bindMathsNative("gcd", 2, func(a []kl.Obj) (kl.Obj, bool) { return mathsGcd(a[0], a[1]) })
	bindMathsNative("lcd", 2, func(a []kl.Obj) (kl.Obj, bool) { return mathsLcd(a[0], a[1]) })
	bindMathsNative("isqrt", 1, func(a []kl.Obj) (kl.Obj, bool) { return mathsIsqrt(a[0]) })
}

// bindMathsNative rebinds name to a native that returns fast(args) when that
// reports true, and otherwise tail-calls the definition it replaced.
func bindMathsNative(name string, arity int, fast func([]kl.Obj) (kl.Obj, bool)) {
	sym := kl.MakeSymbol(name)
	orig := kl.SymbolFunction(sym)
	if orig == nil {
		return
	}
	kl.BindSymbolFunc(sym, kl.MakeNative(func(e *kl.ControlFlow) {
		var buf [2]kl.Obj
		args := buf[:arity]
		for i := range args {
			args[i] = e.Get(i + 1)
		}
		if r, ok := fast(args); ok {
			e.Return(r)
			return
		}
		e.TailApply(orig, args...)
	}, arity))
}

// mathsMod evaluates StLib's mod formula, or reports false when the original
// must run instead.
func mathsMod(x, y kl.Obj) (kl.Obj, bool) {
	if !kl.IsNumber(x) || !kl.IsNumber(y) {
		return nil, false
	}
	xv, yv := kl.GetNumber(x), kl.GetNumber(y)
	if yv == 0 {
		return nil, false // (/ X 0) raises; let the original raise it
	}
	div := xv / yv
	if div <= -mathsExactLimit || div >= mathsExactLimit || div != div {
		return nil, false
	}
	r := yv * (div - math.Floor(div))
	if kl.PrimIsInteger(x) == kl.True && kl.PrimIsInteger(y) == kl.True {
		if r <= -mathsExactLimit || r >= mathsExactLimit || r != r {
			return nil, false
		}
		r = math.Round(r)
	}
	return kl.MakeNumber(r), true
}

// mathsPower multiplies by N, M times, in the original's order.
func mathsPower(n, m kl.Obj) (kl.Obj, bool) {
	if !kl.IsNumber(n) {
		return nil, false
	}
	k, ok := mathsInt(m)
	if !ok || k < 0 || k > mathsPowerLimit {
		return nil, false
	}
	nv, r := kl.GetNumber(n), 1.0
	for ; k > 0; k-- {
		r = nv * r
	}
	return kl.MakeNumber(r), true
}

func gcdInt(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// mathsGcd is gcd(|M|, |N|) where the original returns one.
func mathsGcd(x, y kl.Obj) (kl.Obj, bool) {
	m, ok1 := mathsInt(x)
	n, ok2 := mathsInt(y)
	if !ok1 || !ok2 {
		return nil, false
	}
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}
	if m == 0 || n == 0 || m == n {
		return nil, false // the original divides by zero
	}
	return kl.MakeInteger(int(gcdInt(m, n))), true
}

// mathsLcd is the lowest common divisor above 1 as StLib's lcd searches it.
func mathsLcd(x, y kl.Obj) (kl.Obj, bool) {
	m, ok1 := mathsInt(x)
	n, ok2 := mathsInt(y)
	if !ok1 || !ok2 {
		return nil, false
	}
	if m%2 == 0 && n%2 == 0 {
		return kl.MakeInteger(2), true
	}
	if m <= 0 || n <= 0 {
		return kl.MakeInteger(1), true
	}
	g := gcdInt(m, n)
	for g%2 == 0 {
		g /= 2
	}
	for d := int64(3); d*d <= g; d += 2 {
		if g%d == 0 {
			return kl.MakeInteger(int(d)), true
		}
	}
	if g > 1 {
		return kl.MakeInteger(int(g)), true // g is an odd prime
	}
	return kl.MakeInteger(1), true
}

// mathsIsqrt is floor(sqrt(N)) for N >= 0 and -1 for negative N.
func mathsIsqrt(x kl.Obj) (kl.Obj, bool) {
	v, ok := mathsNumber(x)
	if !ok {
		return nil, false
	}
	if v < 0 {
		return kl.MakeInteger(-1), true
	}
	s := math.Floor(math.Sqrt(v))
	for s*s > v {
		s--
	}
	for (s+1)*(s+1) <= v {
		s++
	}
	return kl.MakeNumber(s), true
}
