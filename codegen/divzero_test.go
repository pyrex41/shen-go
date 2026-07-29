package codegen

import (
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

// Issue #10/#16: (/ x 0) must raise a catchable "division by zero" error rather
// than saturating to a bogus maxint. The guard lives in kl.PrimNumberDivide, so
// AOT-generated Go inherits it only for as long as codegen keeps routing `/`
// through that function. Inlining a raw Go `x / y` here -- the obvious
// "optimisation" for an arithmetic primitive -- would silently reintroduce the
// bug in compiled artifacts only, where no interpreter test would catch it.
func TestDivideIsEmittedThroughTheGuardedPrimitive(t *testing.T) {
	prim, ok := shenPrimitive["/"]
	if !ok {
		t.Fatal(`codegen no longer emits "/" as a primitive call; ` +
			"whatever replaced it must still guard a zero divisor")
	}
	if prim.Arity != 2 {
		t.Errorf(`"/" arity = %d, want 2`, prim.Arity)
	}
	if prim.Name != "PrimNumberDivide" {
		t.Errorf(`"/" emits %q, want "PrimNumberDivide" (the divide-by-zero guard)`,
			prim.Name)
	}
}

// And the primitive it names really does raise, so the mapping above is worth
// something. This is the same guard the interpreter and the bytecode VM use.
func TestPrimNumberDivideRaisesOnZeroDivisor(t *testing.T) {
	for _, tc := range []struct{ x, y float64 }{
		{1, 0}, {-1, 0}, {0, 0}, {1.5, 0}, {1, -0.0},
	} {
		func() {
			defer func() {
				r := recover()
				if r == nil {
					t.Errorf("(/ %v %v) did not raise", tc.x, tc.y)
					return
				}
				obj, isObj := r.(kl.Obj)
				if !isObj || !kl.IsError(obj) {
					t.Errorf("(/ %v %v) panicked with %v, want a kl error object", tc.x, tc.y, r)
				}
			}()
			kl.PrimNumberDivide(kl.MakeNumber(tc.x), kl.MakeNumber(tc.y))
		}()
	}

	// A non-zero divisor is untouched.
	if got := kl.GetNumber(kl.PrimNumberDivide(kl.MakeNumber(10), kl.MakeNumber(4))); got != 2.5 {
		t.Errorf("(/ 10 4) = %v, want 2.5", got)
	}
}
