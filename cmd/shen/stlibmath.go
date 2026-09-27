package main

import (
	"math"

	"github.com/pyrex41/shen-go/kl"
)

// StLib's floor/ceiling/round search for the answer one decimal digit at a
// time (maths.shen, rounding-loop): about 45µs a call here, and mod and div
// sit on top of floor, so any numeric Shen code using them pays that
// (issue #57). installStlibMath rebinds them, after the library has loaded,
// to natives that compute the same values:
//
//   - floor, ceiling and round are exact in float64 for every input on which
//     rounding-loop terminates (|N| < 1e16; above that its (+ Guess (power 10
//     Exponent)) stops changing Guess and it never returns). round rounds a
//     half away from zero, as rounding-loop does (math.Round).
//   - mod and div repeat StLib's formula with the same float64 operations in
//     the same order, so their results are bit-for-bit the library's, down to
//     the rounding of non-integer arguments.
//
// Anything else (a non-number, a non-finite number, a zero divisor) is
// handed to the library's own definition, so errors read exactly as before.
// The type signatures StLib declared are untouched; only the function cell
// is rebound. SHEN_NO_STLIB_NATIVE=1 skips this.
func installStlibMath() {
	unary := func(name string, f func(float64) float64) {
		sym := kl.MakeSymbol(name)
		orig := kl.PrimFunc(sym)
		if orig == nil {
			return
		}
		kl.BindSymbolFunc(sym, kl.MakeNative(func(e *kl.ControlFlow) {
			x := e.Get(1)
			if n, ok := finite(x); ok {
				e.Return(kl.MakeNumber(f(n) + 0))
				return
			}
			e.TailApply(orig, x)
		}, 1))
	}
	binary := func(name string, f func(x, y float64) float64) {
		sym := kl.MakeSymbol(name)
		orig := kl.PrimFunc(sym)
		if orig == nil {
			return
		}
		kl.BindSymbolFunc(sym, kl.MakeNative(func(e *kl.ControlFlow) {
			x, y := e.Get(1), e.Get(2)
			if a, ok := finite(x); ok {
				if b, ok := finite(y); ok && b != 0 {
					if r := f(a, b); !math.IsNaN(r) && !math.IsInf(r, 0) {
						e.Return(kl.MakeNumber(r + 0))
						return
					}
				}
			}
			e.TailApply(orig, x, y)
		}, 2))
	}
	unary("floor", math.Floor)
	unary("ceiling", math.Ceil)
	unary("round", math.Round)
	// (define mod X Y -> (let Div (/ X Y) FloorDiv (floor Div)
	//   (if (and (integer? X) (integer? Y)) (round (* Y (- Div FloorDiv)))
	//       (* Y (- Div FloorDiv)))))
	binary("mod", func(x, y float64) float64 {
		div := x / y
		r := y * (div - math.Floor(div))
		if isInteger(x) && isInteger(y) {
			return math.Round(r)
		}
		return r
	})
	// (define div N D -> (floor (/ N D)))
	binary("div", func(x, y float64) float64 { return math.Floor(x / y) })
}

func finite(o kl.Obj) (float64, bool) {
	if !kl.IsNumber(o) {
		return 0, false
	}
	f := kl.GetNumber(o)
	return f, !math.IsNaN(f) && !math.IsInf(f, 0)
}

// isInteger is the kernel's integer? on a finite float64.
func isInteger(f float64) bool { return f == math.Trunc(f) }
