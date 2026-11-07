package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// A rebound helper of the kernel's factoriser that defines a function
// (review N1). It must not deadlock, as it did while a lock was held across
// shen.factor, and must not run inside the compiler at all: master has no
// factoriser, so the helper never runs there, inner is never defined, and the
// output is master's.
func TestFactoriserHelperRebindingIsNotReentered(t *testing.T) {
	bin := buildShen(t)
	dir := t.TempDir()
	src := `(set *depth* 0)
(define mk-inner -> (eval-kl [defun inner [v] [cond [[and [cons? v] [= a [hd v]]] 1] [[and [cons? v] [= b [hd v]]] 2] [true 3]]]))
(eval-kl [defun shen.factor-recognisors [v] [do [set *depth* [+ 1 [value *depth*]]] [do [if [= [value *depth*] 1] [mk-inner] skip] v]]])
(define outer [a X] -> X [b X] -> X [c X] -> X _ -> 0)
(put inner arity 1)
(print [(outer [b 7]) (trap-error (inner [b 8]) (/. E (error-to-string E))) (value *depth*)])
`
	file := filepath.Join(dir, "n1.shen")
	if err := os.WriteFile(file, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, bin, "script", file).CombinedOutput()
	if ctx.Err() != nil {
		t.Fatalf("hung: a define reached from inside the factoriser deadlocked\n%s", out)
	}
	if want := `[7 "can't apply non function: inner" 0]`; err != nil || !strings.Contains(string(out), want) {
		t.Fatalf("got %v\n%s\nwant %s", err, out, want)
	}
}
