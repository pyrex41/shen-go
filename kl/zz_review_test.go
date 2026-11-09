package kl

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

func zzEval(ctx *ControlFlow, exp string) string {
	r := NewSexpReader(strings.NewReader(exp), true)
	sexp, err := r.Read()
	if err != nil {
		panic(err)
	}
	return ObjString(Eval(ctx, sexp))
}

// Two goroutines, each its own ControlFlow, defining and running functions.
func TestZZConcurrentDefine(t *testing.T) {
	overrideDirect("thaw", 1, nativeThawDirect)
	stdThaw = mustSymbol(MakeSymbol("thaw")).function
	var wg sync.WaitGroup
	for g := 0; g < 4; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			var ctx ControlFlow
			for i := 0; i < 300; i++ {
				name := fmt.Sprintf("zzf%d_%d", g, i%7)
				zzEval(&ctx, fmt.Sprintf("(defun %s (X) (let F (freeze (+ X 1)) (if (= X 0) (thaw F) (do (set zzk F) (thaw F)))))", name))
				zzEval(&ctx, fmt.Sprintf("(defun %sb (X) (let G (lambda Y (lambda Z (cons Y Z))) (if (cons? X) ((G (hd X)) (tl X)) X)))", name))
				got := zzEval(&ctx, fmt.Sprintf("(%s 3)", name))
				if got != "4" {
					t.Errorf("got %s", got)
					return
				}
			}
		}(g)
	}
	wg.Wait()
}

func zzLimited(t *testing.T, defs []string, final string) {
	t.Helper()
	var got string
	done := make(chan struct{})
	go func() {
		defer func() { recover(); close(done) }()
		ctx := ControlFlow{stepLimit: 200_000}
		for _, d := range defs {
			zzEval(&ctx, d)
		}
		got = zzEval(&ctx, final)
	}()
	select {
	case <-done:
	case <-time.After(20 * time.Second):
		t.Fatalf("step limit did not halt %q", final)
	}
	t.Logf("%s -> %.80s", final, got)
}

func TestZZStepLimitShapes(t *testing.T) {
	zzLimited(t, []string{"(defun zs1 (X) (let G (lambda Y (zs1 Y)) (G X)))"}, "(zs1 0)")
	zzLimited(t, []string{"(defun zs2 (X) (let F (freeze (zs2 X)) (if (= X 0) (thaw F) (thaw F))))"}, "(zs2 0)")
	zzLimited(t, []string{"(defun zs3 (X) (let F (freeze (zs3 X)) (thaw F)))"}, "(zs3 0)")
	zzLimited(t, []string{"(defun zs4 (X) (and true (zs4b X)))", "(defun zs4b (X) (or false (zs4 X)))"}, "(zs4 0)")
	zzLimited(t, []string{"(defun zs5 (X) (trap-error (zs5 X) (lambda E e)))"}, "(zs5 0)")
	zzLimited(t, []string{"(defun zs6 (F) (F F))"}, "(zs6 zs6)")
	zzLimited(t, []string{"(defun zs7 (X) ((lambda Y (zs7 Y)) X))"}, "(zs7 0)")
	zzLimited(t, []string{"(defun zs8 (N) (zs8 (+ N 1)))"}, "(zs8 0)")
	zzLimited(t, []string{"(defun zs9 (X) (zs9b X))", "(defun zs9b (X) ((lambda Y (zs9 Y)) X))"}, "(zs9 0)")
	zzLimited(t, []string{"(defun zs10 (X) (let G (lambda A (lambda B (zs10 B))) ((G 1) X)))"}, "(zs10 0)")
	zzLimited(t, []string{}, "((lambda F (F F)) (lambda F (F F)))")
}
