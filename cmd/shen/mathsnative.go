package main

import (
	"math"
	"os"

	"github.com/pyrex41/shen-go/kl"
)

// Native floor, ceiling, round and mod for the stdlib's Maths package.
//
// StLib computes floor, ceiling and round with maths.rounding-loop, which
// builds the answer one decimal digit at a time from 10^15 down, calling
// power at every step: about 100 µs a call. mod, div, modf, float->pair and
// random all go through floor, so any modular arithmetic written against the
// stdlib paid that on every operation.
//
// Each native is exact on a proven domain and hands every other input to the
// original Shen definition, captured before the rebinding, so behaviour
// outside the domain (error text for a non-number, division by zero, and the
// original's non-termination on NaN, +-Inf and some integers >= 2^53) is the
// original's by construction.
//
// The domain is a finite number of magnitude below 2^53. There every Guess
// rounding-loop forms is an exactly representable integer and every
// comparison is exact, so the loop returns exactly:
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
// cmd/shen/mathsnative_test.go runs thousands of cases through both sides and
// requires byte-identical output. SHEN_NO_MATHS_NATIVE=1 keeps the Shen
// definitions.

// mathsExactLimit bounds the domain: |x| < 2^53.
const mathsExactLimit = 1 << 53

// mathsNumber is x's value when x is a number in the natives' domain.
func mathsNumber(x kl.Obj) (float64, bool) {
	if !kl.IsNumber(x) {
		return 0, false
	}
	f := kl.GetNumber(x)
	return f, f > -mathsExactLimit && f < mathsExactLimit // false for NaN, +-Inf
}

// installMathsNatives rebinds floor, ceiling, round and mod once the stdlib
// is loaded. A name the stdlib did not bind (SHEN_NO_STDLIB, or a failed load)
// is left alone.
func installMathsNatives() {
	if os.Getenv("SHEN_NO_MATHS_NATIVE") != "" {
		return
	}
	unary := func(name string, f func(float64) float64) {
		sym := kl.MakeSymbol(name)
		orig := kl.SymbolFunction(sym)
		if orig == nil {
			return
		}
		kl.BindSymbolFunc(sym, kl.MakeNative(func(e *kl.ControlFlow) {
			x := e.Get(1)
			if v, ok := mathsNumber(x); ok {
				e.Return(kl.MakeNumber(f(v)))
				return
			}
			e.TailApply(orig, x)
		}, 1))
	}
	unary("floor", math.Floor)
	unary("ceiling", math.Ceil)
	unary("round", math.Round)

	modSym := kl.MakeSymbol("mod")
	origMod := kl.SymbolFunction(modSym)
	if origMod == nil {
		return
	}
	kl.BindSymbolFunc(modSym, kl.MakeNative(func(e *kl.ControlFlow) {
		x, y := e.Get(1), e.Get(2)
		if r, ok := mathsMod(x, y); ok {
			e.Return(r)
			return
		}
		e.TailApply(origMod, x, y)
	}, 2))
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
