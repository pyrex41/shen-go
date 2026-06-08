// Stage 0 plugin viability check: a HAND-WRITTEN tak as a Go plugin, mirroring
// what the codegen would emit (MakeNative + trampoline + boxed Obj + PrimFunc
// self-reference). Proves the plugin ABI works against the kl package before we
// wire up real codegen.
//
// Build: go build -buildmode=plugin -o /tmp/tak.so ./cmd/plugintest/plugin
package main

import . "github.com/tiancaiamao/shen-go/kl"

// symtak is the self-reference, resolved at call time (after Install binds it),
// exactly like generated kernel code uses PrimFunc(symXXX).
var symtak = MakeSymbol("tak")

// takFn implements:
//   (defun tak (X Y Z)
//     (if (not (< Y X)) Z
//         (tak (tak (- X 1) Y Z) (tak (- Y 1) Z X) (tak (- Z 1) X Y))))
var takFn = MakeNative(func(__e *ControlFlow) {
	X := __e.Get(1)
	Y := __e.Get(2)
	Z := __e.Get(3)

	if True == PrimNot(PrimLessThan(Y, X)) {
		__e.Return(Z)
		return
	}

	one := MakeInteger(1)
	a := Call(__e, PrimFunc(symtak), PrimNumberSubtract(X, one), Y, Z)
	b := Call(__e, PrimFunc(symtak), PrimNumberSubtract(Y, one), Z, X)
	c := Call(__e, PrimFunc(symtak), PrimNumberSubtract(Z, one), X, Y)
	__e.TailApply(PrimFunc(symtak), a, b, c)
}, 3)

// Exported plugin contract (looked up by the host):
var Compiled Obj = takFn
var SymName string = "tak"
