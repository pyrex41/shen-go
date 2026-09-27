package kl

import (
	"strings"
	"testing"
	"time"
)

func TestDynamicSelfTailReusesFrame(t *testing.T) {
	var ctx ControlFlow
	if got := evalString(&ctx, `(defun dynamic-self-55 (F N) (if (= N 0) 42 (F F (- N 1))))`); IsError(got) {
		t.Fatalf("definition failed: %s", ObjString(got))
	}
	f := mustSymbol(MakeSymbol("dynamic-self-55")).function
	if f == nil || *f != scmHeadBytecodeFunc {
		t.Fatal("expected a compiled function")
	}
	code := mustBytecodeFunc(f).fn.Code
	var dynamic bool
	for _, in := range code {
		if in.Op == OP_TAIL_CALL {
			dynamic = true
		}
		if in.Op == OP_SELF_TAIL_CALL {
			t.Fatal("test must exercise dynamic tail calls")
		}
	}
	if !dynamic {
		t.Fatal("test function has no dynamic tail call")
	}
	if got := Call(&ctx, f, f, MakeInteger(100000)); ObjString(got) != "42" {
		t.Fatalf("dynamic recursion returned %s, want 42", ObjString(got))
	}
}

func TestDynamicSelfTailHonorsStepLimit(t *testing.T) {
	var panicked any
	if !runWithin(5*time.Second, func() {
		defer func() { panicked = recover() }()
		ctx := ControlFlow{stepLimit: 1000}
		if def := evalString(&ctx, `(defun dynamic-spin-55 (F) (F F))`); IsError(def) {
			t.Fatalf("definition failed: %s", ObjString(def))
		}
		f := mustSymbol(MakeSymbol("dynamic-spin-55")).function
		Call(&ctx, f, f)
	}) {
		t.Fatal("dynamic self tail call exceeded timeout under step limit")
	}
	got, ok := panicked.(Obj)
	if !ok || got == nil || !IsError(got) || !strings.Contains(ObjString(got), "step limit exceeded") {
		t.Fatalf("want step-limit panic, got %v", panicked)
	}
}

func BenchmarkVMDynamicSelfTail(b *testing.B) {
	var ctx ControlFlow
	if got := evalString(&ctx, `(defun dynamic-bench-55 (F N) (if (= N 0) 42 (F F (- N 1))))`); IsError(got) {
		b.Fatalf("definition failed: %s", ObjString(got))
	}
	f := mustSymbol(MakeSymbol("dynamic-bench-55")).function
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if got := Call(&ctx, f, f, MakeInteger(10000)); IsError(got) {
			b.Fatalf("call failed: %s", ObjString(got))
		}
	}
}
