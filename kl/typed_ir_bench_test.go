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
