package kl

// Exact powers of ten for the kernel reader.
//
// kernel/sources/reader.shen computes a numeric literal's value itself:
//
//	(define expt
//	  _ 0       -> 1
//	  Base Expt -> (* Base (expt Base (- Expt 1)))  where (> Expt 0)
//	  Base Expt -> (/ (expt Base (+ Expt 1)) Base))
//
//	(define compute-E N Log10 -> (* N (expt 10 Log10)))
//
// so (expt 10 N) is N roundings of a float64 multiply (or divide). In a host
// with a bignum/rational tower -- shen-cl -- that is exact and the literal
// comes out right. In a float64-only host the error compounds, and the literal
// is simply a different number than it says:
//
//	1e25   ->  9.999999999999999e24
//	1e300  ->  1.0000000000000002e300
//	1e-6   ->  1.0000000000000002e-06   (so 0.000001 was wrong too)
//
// Repeated multiplication happens to be exact for 10^0..10^22 and repeated
// division for 10^-1..10^-5; the first divergences are 10^25 and 10^-6.
//
// InstallExactPow10 answers (expt 10 N) with the correctly rounded power
// instead. Nothing else about reading changes: compute-integer/compute-fraction
// keep the kernel's own summation, so the fix is confined to the one operation
// that was wrong. Over every numeric literal in kernel/ and src/ (3539 of them)
// the exact powers change no value at all; they change only the literals whose
// exponent reaches past 10^22 or 10^-5, which is exactly the defect.
//
// This is the interpreted half. kl/shenreader.go, the Go-native reader for the
// same grammar, calls pow10 directly, so the two readers stay bit-for-bit
// identical -- which cmd/shen's oracle parity tests check against the live
// interpreted kernel.

import (
	"math"
	"strconv"
)

const (
	pow10Min = -350 // below this every power underflows to 0
	pow10Max = 309  // above this every power overflows to +Inf
)

// pow10Table holds the correctly rounded 10^n for n in [pow10Min, pow10Max],
// built with strconv.ParseFloat (which is correctly rounded by contract).
// math.Pow is NOT usable here: it disagrees with the correctly rounded power
// for 252 of the 309 non-negative exponents.
var pow10Table = func() [pow10Max - pow10Min + 1]float64 {
	var t [pow10Max - pow10Min + 1]float64
	for n := pow10Min; n <= pow10Max; n++ {
		// ParseFloat returns +Inf/0 with ErrRange at the ends; that is the
		// value we want, so the error is deliberately ignored.
		v, _ := strconv.ParseFloat("1e"+strconv.Itoa(n), 64)
		t[n-pow10Min] = v
	}
	return t
}()

// pow10 returns 10**n as the correctly rounded float64, saturating to 0 and
// +Inf outside the representable exponent range (as ParseFloat itself does).
func pow10(n int) float64 {
	switch {
	case n < pow10Min:
		return 0
	case n > pow10Max:
		return math.Inf(1)
	}
	return pow10Table[n-pow10Min]
}

// InstallExactPow10 rebinds the kernel's shen.expt to a native that answers
// base-10 integer powers exactly and delegates everything else to the
// interpreted definition it replaces. It must run after ReaderMain (which
// defines shen.expt) and before any source is read.
//
// Only base 10 with an integral exponent is intercepted -- that is all the
// reader ever asks for. Any other call, including the non-integral exponent
// the kernel definition would loop forever on, is handed straight back to the
// interpreted expt so its behaviour is untouched.
func InstallExactPow10() {
	sym := MakeSymbol("shen.expt")
	// shen.expt lives in reader.kl, so a kernel that does not include the
	// reader never defines it. That is the normal shape of an eval-free
	// Yggdrasil shake: dropping eval drops the reader, and PrimFunc PANICS on
	// an unbound symbol (primitives.go). Every eval-free shaken artifact
	// therefore died at boot with a bare `panic: (kl.Obj) 0x...`, because the
	// generated main calls this helper outside its own error reporter.
	//
	// Guard the same way InstallIntegerGuard does: with nothing to delegate
	// to, there is nothing to optimise, so leave the binding alone. The
	// reader is what asks for base-10 expt, and a kernel without the reader
	// never makes that call.
	interpreted := mustSymbol(sym).function
	if interpreted == nil {
		return
	}
	native := MakeNative(func(e *ControlFlow) {
		base, exp := e.Get(1), e.Get(2)
		if *base == scmHeadNumber && *exp == scmHeadNumber {
			b, x := GetNumber(base), GetNumber(exp)
			if b == 10 && isPreciseInteger(x) && fitsInt(x) {
				e.Return(MakeNumber(pow10(int(x))))
				return
			}
		}
		e.TailApply(interpreted, base, exp)
	}, 2)
	BindSymbolFunc(sym, native)
}
