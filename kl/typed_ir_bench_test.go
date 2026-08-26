package kl

import "testing"

func BenchmarkTypedNumericFixnumAdd(b *testing.B) {
	x, y := MakeInteger(123), MakeInteger(456)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !isFixnum(numAdd(x, y)) {
			b.Fatal("fixnum add unexpectedly boxed")
		}
	}
}

func BenchmarkTypedNumericFloatFallback(b *testing.B) {
	x, y := MakeNumber(1.25), MakeNumber(2.5)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if got := mustNumber(numAdd(x, y)); got != 3.75 {
			b.Fatalf("got %v", got)
		}
	}
}

func BenchmarkTypedVMCompiledSum(b *testing.B) {
	benchEval(b, `(defun sum (Acc N) (if (= N 0) Acc (sum (+ Acc N) (- N 1))))`, `(sum 0 8000)`)
}

func BenchmarkTypedVMCompiledFib(b *testing.B) {
	benchEval(b, `(defun fib (N) (if (< N 2) N (+ (fib (- N 1)) (fib (- N 2)))))`, `(fib 24)`)
}

func BenchmarkTypedVMDynamicPrimitive(b *testing.B) {
	benchEval(b, `(defun numeric-loop (N) (if (= N 0) true (do (number? N) (numeric-loop (- N 1)))))`, `(numeric-loop 8000)`)
}

// The following benchmarks exercise the typed-region boundaries.  They are
// intentionally written as KL programs so VM and AOT implementations can use
// the same workload definitions.
func BenchmarkTypedVMBoolBranch(b *testing.B) {
	benchEval(b, `(defun bool-loop (N Acc) (if (= N 0) Acc (if (number? N) (bool-loop (- N 1) true) (bool-loop (- N 1) Acc))))`, `(bool-loop 8000 false)`)
}

func BenchmarkTypedVMUnicodeString(b *testing.B) {
	benchEval(b, `(defun string-loop (N S) (if (= N 0) S (string-loop (- N 1) (cn S "λ"))))`, `(string-loop 200 "λ")`)
}

func BenchmarkTypedVMPairList(b *testing.B) {
	benchEval(b, `(defun list-loop (N L) (if (= N 0) L (list-loop (- N 1) (cons N L))))`, `(list-loop 2000 ())`)
}

func BenchmarkTypedVMVector(b *testing.B) {
	benchEval(b, `(defun vector-loop (N V) (if (= N 0) V (do (address-> V 0 N) (vector-loop (- N 1) V))))`, `(vector-loop 8000 (absvector 1))`)
}

func BenchmarkTypedVMDynamicApply(b *testing.B) {
	benchEval(b, `(defun apply-add (F N) (if (= N 0) 0 (do (F N) (apply-add F (- N 1)))))`, `(apply-add (lambda X (+ X 1)) 8000)`)
}

func BenchmarkTypedVMFallbackHeavy(b *testing.B) {
	benchEval(b, `(defun fallback-heavy (N X) (if (= N 0) X (fallback-heavy (- N 1) (if (number? X) (+ X 1) (cn X "x")))))`, `(fallback-heavy 2000 "x")`)
}
