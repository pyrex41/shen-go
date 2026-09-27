package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestStlibMathNativesMatchTheLibrary runs StLib's own floor, ceiling,
// round, mod and div next to the natives installStlibMath binds over them,
// on integers, halves, small fractions and random values of every magnitude
// rounding-loop terminates on, and requires identical results (the same
// float64, not just =). It also checks that a non-number still raises the
// library's own error.
func TestStlibMathNativesMatchTheLibrary(t *testing.T) {
	if testing.Short() {
		t.Skip("boots the kernel and StLib")
	}
	e := bootShen(t)
	loadStdlib(e)
	names := []string{"floor", "ceiling", "round", "mod", "div"}
	orig := map[string]kl.Obj{}
	for _, n := range names {
		orig[n] = kl.PrimFunc(kl.MakeSymbol(n))
	}
	installStlibMath()
	call := func(f kl.Obj, args ...kl.Obj) kl.Obj {
		r := kl.Try(e, kl.MakeNative(func(c *kl.ControlFlow) { c.TailApply(f, args...) }, 0))
		return r.Catch(kl.MakeNative(func(c *kl.ControlFlow) { c.Return(c.Get(1)) }, 1))
	}
	same := func(a, b kl.Obj) bool {
		if kl.IsError(a) || kl.IsError(b) {
			return kl.IsError(a) && kl.IsError(b) && kl.GetString(kl.PrimErrorToString(a)) == kl.GetString(kl.PrimErrorToString(b))
		}
		if kl.IsNumber(a) && kl.IsNumber(b) {
			return math.Float64bits(kl.GetNumber(a)) == math.Float64bits(kl.GetNumber(b))
		}
		return kl.PrimEqual(a, b) == kl.True
	}
	rng := rand.New(rand.NewSource(57))
	var xs []float64
	for _, x := range []float64{0, 1, -1, 2, 0.5, -0.5, 1.5, -1.5, 2.5, -2.5, 0.1, -0.1, 0.49999999999999994, 7.999, -7.999, 123456789.75, 1e10, 1e15 - 0.5, 999999999999999.9} {
		xs = append(xs, x)
	}
	for i := 0; i < 300; i++ {
		mag := math.Pow(10, float64(rng.Intn(15)))
		x := (rng.Float64()*2 - 1) * mag
		if i%3 == 0 {
			x = math.Round(x)
		}
		if i%5 == 0 {
			x = math.Floor(x) + 0.5
		}
		xs = append(xs, x)
	}
	for _, x := range xs {
		for _, n := range []string{"floor", "ceiling", "round"} {
			a := call(orig[n], kl.MakeNumber(x))
			b := call(kl.PrimFunc(kl.MakeSymbol(n)), kl.MakeNumber(x))
			if !same(a, b) {
				t.Errorf("(%s %v): library %s, native %s", n, x, kl.ObjString(a), kl.ObjString(b))
			}
		}
	}
	ys := []float64{3, -3, 7, 2, 2.5, -0.75, 1e6, 13, -1e3}
	for i, x := range xs {
		for _, y := range ys[i%3 : i%3+4] {
			if math.Abs(x/y) >= 1e15 {
				continue // floor of the quotient would not terminate in StLib
			}
			for _, n := range []string{"mod", "div"} {
				a := call(orig[n], kl.MakeNumber(x), kl.MakeNumber(y))
				b := call(kl.PrimFunc(kl.MakeSymbol(n)), kl.MakeNumber(x), kl.MakeNumber(y))
				if !same(a, b) {
					t.Errorf("(%s %v %v): library %s, native %s", n, x, y, kl.ObjString(a), kl.ObjString(b))
				}
			}
		}
	}
	for _, n := range []string{"floor", "round"} {
		a := call(orig[n], kl.MakeSymbol("x"))
		b := call(kl.PrimFunc(kl.MakeSymbol(n)), kl.MakeSymbol("x"))
		if !kl.IsError(a) || !same(a, b) {
			t.Errorf("(%s x): library %s, native %s", n, kl.ObjString(a), kl.ObjString(b))
		}
	}
	for _, n := range []string{"mod", "div"} {
		a := call(orig[n], kl.MakeNumber(5), kl.MakeNumber(0))
		b := call(kl.PrimFunc(kl.MakeSymbol(n)), kl.MakeNumber(5), kl.MakeNumber(0))
		if !same(a, b) {
			t.Errorf("(%s 5 0): library %s, native %s", n, kl.ObjString(a), kl.ObjString(b))
		}
	}
}
