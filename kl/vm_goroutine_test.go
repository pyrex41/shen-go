package kl

import (
	"runtime/debug"
	"testing"
)

// A rebinding made by another goroutine reaches bytecode that is already
// running, at its next call out of the VM. The two goroutines synchronise
// through a channel inside a native, which is the only way a Go embedder can
// share a binding without a data race on the symbol itself; cmd/shen never
// runs Shen code on more than one goroutine (the zmq receiver goroutine calls
// only the socket's Recv, and -pprof's server only net/http). Run under -race
// it also checks that the guard flags are race-free.
func TestRegisterRebindFromAnotherGoroutine(t *testing.T) {
	// No collection meanwhile: a finished GC cycle also raises vmAttn, which
	// would hide a rebinding that failed to.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	plus := MakeSymbol("+")
	orig := PrimFunc(plus)
	defer BindSymbolFunc(plus, orig)

	ready, rebound := make(chan struct{}), make(chan struct{})
	BindSymbolFunc(MakeSymbol("gr-wait"), MakeNative0(func(*ControlFlow) Obj {
		close(ready)
		<-rebound
		return True
	}))
	var ctx ControlFlow
	evalAll(t, &ctx,
		`(defun gr-sum (N A) (if (= N 0) A (gr-sum (- N 1) (+ A 1))))`,
		`(defun gr-run () (let A (gr-sum 1000 0) (let W (gr-wait) (cons A (gr-sum 3 0)))))`,
	)
	go func() {
		<-ready
		BindSymbolFunc(plus, MakeNative2(func(_ *ControlFlow, a, b Obj) Obj { return MakeInteger(100) }))
		close(rebound)
	}()
	if got := evalAll(t, &ctx, `(gr-run)`); got != "(1000 . 100)" {
		t.Fatalf("(gr-run) = %s, want (1000 . 100): the rebinding was missed", got)
	}
}
