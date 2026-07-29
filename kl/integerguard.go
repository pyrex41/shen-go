package kl

import "math"

// The kernel (sys.kl) defines
//
//	(defun integer? (V) (and (number? V)
//	                     (let W (shen.abs V)
//	                       (shen.integer-test? W (shen.magless W 1)))))
//	(defun shen.magless (V W) (let W2 (* W 2) (if (> W2 V) W (shen.magless V W2))))
//
// shen.magless doubles W until it exceeds |V|. For +-Inf nothing ever exceeds
// it and W just walks up to Inf itself; for NaN every comparison is false. In
// both cases the recursion never returns, so (integer? X) simply never
// answers. A primitive that does not terminate is wrong under every reading of
// the spec, independent of any question about what overflow ought to do.
//
// The native `integer?` in the primitive table (PrimIsInteger) already answers
// correctly, but it is not the binding that runs: the kernel's `defun`
// overwrites the symbol's function while sys.kl loads. A guard placed only in
// the primitive table is therefore dead code -- which is exactly the state
// shen-lua's equivalent guard (prims.lua:250) is in. The wrapper below has to
// be installed *after* the kernel modules have run.

// InstallIntegerGuard wraps the currently bound `integer?` with a native
// non-finite guard. +-Inf and NaN are answered directly with false; everything
// else tail-calls straight through to the binding that was in place, so the
// kernel's own definition remains the authority for every finite number.
//
// Must be called after the kernel modules have loaded (they define the
// interpreted integer? this wraps). Calling it twice is harmless but wasteful:
// the second call simply wraps the first wrapper.
func InstallIntegerGuard() {
	sym := MakeSymbol("integer?")
	inner := mustSymbol(sym).function
	if inner == nil {
		// Nothing loaded yet; the primitive-table binding (which handles
		// non-finite values correctly on its own) is still in place.
		return
	}
	BindSymbolFunc(sym, MakeNative(func(e *ControlFlow) {
		x := e.Get(1)
		if IsNumber(x) {
			// GetNumber covers fixnums too; those are always finite.
			if f := GetNumber(x); math.IsInf(f, 0) || math.IsNaN(f) {
				e.Return(False)
				return
			}
		}
		e.TailApply(inner, x)
	}, 1))
}
