package kl

// Micro-benchmarks for the bytecode VM hot path. These define a KL function via
// a compiled defun and then time repeated calls, so they isolate per-call
// runtime cost (activation setup, argument passing, arithmetic) from parsing and
// kernel boot. Run with -benchmem to see allocations/op, which is what the
// runtime-allocation-reduction work targets.

import (
	"strings"
	"testing"
)

func benchEval(b *testing.B, def, call string) {
	b.Helper()
	var ctx ControlFlow
	// Define the function(s) once.
	for _, d := range strings.Split(def, "\n;;\n") {
		if strings.TrimSpace(d) == "" {
			continue
		}
		if res := evalString(&ctx, d); IsError(res) {
			b.Fatalf("defun errored: %s", ObjString(res))
		}
	}
	callSexp := mustReadOne(b, call)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if res := Eval(&ctx, callSexp); IsError(res) {
			b.Fatalf("call errored: %s", ObjString(res))
		}
	}
}

func mustReadOne(b *testing.B, src string) Obj {
	b.Helper()
	o, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		b.Fatalf("read %q: %v", src, err)
	}
	return o
}

// BenchmarkVMFib is non-tail double recursion: lots of small activations and
// integer arithmetic.
func BenchmarkVMFib(b *testing.B) {
	benchEval(b, `(defun fib (N) (if (< N 2) N (+ (fib (- N 1)) (fib (- N 2)))))`,
		`(fib 24)`)
}

// BenchmarkVMTak is arithmetic-heavy triple recursion.
func BenchmarkVMTak(b *testing.B) {
	benchEval(b, `(defun tak (X Y Z) (if (not (< Y X)) Z (tak (tak (- X 1) Y Z) (tak (- Y 1) Z X) (tak (- Z 1) X Y))))`,
		`(tak 18 12 6)`)
}

// BenchmarkVMSum is a tight self-tail-recursive integer loop (exercises
// OP_SELF_TAIL_CALL + arithmetic fast paths).
func BenchmarkVMSum(b *testing.B) {
	benchEval(b, `(defun sum (Acc N) (if (= N 0) Acc (sum (+ Acc N) (- N 1))))`,
		`(sum 0 100000)`)
}

// BenchmarkVMAck is deep non-tail recursion with general (non-self) tail calls.
func BenchmarkVMAck(b *testing.B) {
	benchEval(b, `(defun ack (M N) (if (= M 0) (+ N 1) (if (= N 0) (ack (- M 1) 1) (ack (- M 1) (ack M (- N 1))))))`,
		`(ack 2 100)`)
}

// BenchmarkVMSumMid sums to ~32M, staying inside the widened signed fixnum range
// [-2^25, 2^25). Before widening (fixnum max 2^20) every accumulator value above
// ~1M boxed a scmNumber; after widening this allocates nothing — the direct
// demonstration of the fixnum-widening win for mid-size positive integers.
func BenchmarkVMSumMid(b *testing.B) {
	benchEval(b, `(defun sum (Acc N) (if (= N 0) Acc (sum (+ Acc N) (- N 1))))`,
		`(sum 0 8000)`)
}

// BenchmarkVMSumNeg accumulates negatively to ~-32M. Negative integers ALWAYS
// boxed under the old unsigned fixnum range; with the signed range they are
// unboxed, so this drops from ~8000 allocs to none.
func BenchmarkVMSumNeg(b *testing.B) {
	benchEval(b, `(defun nsum (Acc N) (if (= N 0) Acc (nsum (- Acc N) (- N 1))))`,
		`(nsum 0 8000)`)
}
