package kl

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// Goroutines with their own ControlFlow define and run functions at the same
// time: the lowering registry, the factoriser's flags and the Prolog-native
// tables are shared, and must be safe under -race (review F2; master passes).
// Some definitions are factorised (a cond of and-tests), so factorDefine runs
// concurrently too. (A rebinding concurrent with compiling races on the
// symbol's function field, as on master, so it is not part of this test.)
func TestConcurrentDefine(t *testing.T) {
	thaw := mustSymbol(symThaw)
	savedStd, savedFn := stdThaw, thaw.function
	defer func() { stdThaw, thaw.function = savedStd, savedFn; vmNoteRebind(thaw) }()
	overrideDirect("thaw", 1, nativeThawDirect)
	stdThaw = thaw.function
	eval := func(ctx *ControlFlow, exp string) string {
		r := NewSexpReader(strings.NewReader(exp), true)
		sexp, err := r.Read()
		if err != nil {
			panic(err)
		}
		return ObjString(Eval(ctx, sexp))
	}
	var wg sync.WaitGroup
	for g := 0; g < 4; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			var ctx ControlFlow
			for i := 0; i < 300; i++ {
				name := fmt.Sprintf("ccf%d_%d", g, i%7)
				eval(&ctx, fmt.Sprintf("(defun %s (X) (let F (freeze (+ X 1)) (if (= X 0) (thaw F) (do (set ccz%d F) (thaw F)))))", name, g))
				eval(&ctx, fmt.Sprintf("(defun %sb (X) (let G (lambda Y (lambda Z (cons Y Z))) (if (cons? X) ((G (hd X)) (tl X)) X)))", name))
				eval(&ctx, fmt.Sprintf("(defun %sc (X) (cond ((and (cons? X) (= (hd X) a)) 1) ((and (cons? X) (= (hd X) b)) 2) (true 3)))", name))
				if got := eval(&ctx, fmt.Sprintf("(%s 3)", name)); got != "4" {
					t.Errorf("got %s", got)
					return
				}
			}
		}(g)
	}
	wg.Wait()
}
