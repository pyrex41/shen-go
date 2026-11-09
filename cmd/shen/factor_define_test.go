package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// A definition is compiled factorised, but ps, gensym, a where guard's effects
// and the results are what they would be without LowerKL and factorDefine, the
// two steps SHEN_GO_LOWER=off turns off.
func TestFactoredDefineChangesNothingVisible(t *testing.T) {
	const script = `
(define f [a X] -> X [a X Y] -> Y [b X] -> X _ -> 0)
(print (ps f)) (print (gensym x))
(print (map (function f) [[a 1] [a 1 2] [b 3] [b 3 4] [a] [] a]))
(set n 0)
(define tick -> (do (set n (+ 1 (value n))) false))
(define g [a X] -> X where (tick) [a X] -> (+ X 1) where (tick) [b X] -> X _ -> 0)
(print (ps g)) (print (g [a 1])) (print (value n))
(define h [a X Y] -> X [a X] -> (hd X) _ -> 0)
(print (trap-error (h [a 1]) (/. E (error-to-string E))))
(print (trap-error (h 1 2) (/. E (error-to-string E))))
(define typed-tuple (@p a X) -> X (@p b X) -> X "" -> 0 (@s "a" S) -> S _ -> 1)
(print (map (function typed-tuple) [(@p a 1) (@p b 2) "" "abc" "b" 5]))
(print (gensym x))
(define thawer X -> (let W (freeze X) (if (= X 1) (thaw W) (cons W []))))
(print (thawer 1))
(print (do (defun thaw (X) rebound-thaw) (thawer 1)))
(defun shen.factor (X) (simple-error "ran inside the compiler"))
(define after-factor [a X] -> X [a X Y] -> Y [b X] -> X _ -> 0)
(print (after-factor [b 1]))
`
	dir := t.TempDir()
	file := filepath.Join(dir, "s.shen")
	if err := os.WriteFile(file, []byte(script), 0o644); err != nil {
		t.Fatal(err)
	}
	bin := buildShen(t)
	run := func(env ...string) string {
		cmd := exec.Command(bin, "script", file)
		cmd.Env = append(os.Environ(), env...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%v: %s", err, out)
		}
		return string(out)
	}
	on, off := run(), run("SHEN_GO_LOWER=off")
	if on != off {
		t.Fatalf("with the pass:\n%s\nwithout:\n%s", on, off)
	}
	if !strings.Contains(on, "[cond [[and") {
		t.Fatalf("ps no longer shows the unfactorised form:\n%s", on)
	}
}

// Mutual tail recursion stays a tail call through each form the lowering steps
// rewrite. A lost tail call does not pass quietly here: a Go stack past 1 GB is
// a fatal error at this depth, as a non-tail version of the same loop shows.
func TestLoweringKeepsTailCalls(t *testing.T) {
	const script = `
(define and-a 0 -> true N -> (and (number? N) (and-b (- N 1))))
(define and-b 0 -> false N -> (and (number? N) (and-a (- N 1))))
(define or-a 0 -> false N -> (or (not (number? N)) (or-b (- N 1))))
(define or-b 0 -> true N -> (or (not (number? N)) (or-a (- N 1))))
(define if-a N -> (if (and (number? N) (= N 0)) done (if-b (- N 1))))
(define if-b N -> (if (or (= N 0) (= N -1)) done (if-a (- N 1))))
(defun cond-a (N) (cond ((= N 0) done) ((and (number? N) (> N 0)) (cond-b (- N 1))) (true err)))
(defun cond-b (N) (cond ((= N 0) done) ((or (< N 0) (> N 0)) (cond-a (- N 1))) (true err)))
(define let-a N -> (if (= N 0) done (let M (- N 1) (let-b M))))
(define let-b N -> (if (= N 0) done (let M (- N 1) (let K M (let-a K)))))
(define var-a F N -> (if (= N 0) done (F (function var-b) (- N 1))))
(define var-b F N -> (if (= N 0) done (F (function var-a) (- N 1))))
(define var-go G N -> (G (function var-go) N))
(define fn-a N -> (if (= N 0) done (let F (/. X Y (fn-b (- X Y))) (F N 1))))
(define fn-b N -> (if (= N 0) done (let F (/. X Y (fn-a (- X Y))) (F N 1))))
(define frz-a N -> (if (= N 0) done (let W (freeze (frz-b (- N 1))) (if (number? N) (thaw W) (cons W [])))))
(define frz-b N -> (if (= N 0) done (let W (freeze (frz-a (- N 1))) (if (number? N) (thaw W) (cons W [])))))
(define jmp-a N -> (if (= N 0) done (let W (freeze (jmp-b (- N 1))) (if (= N -1) (thaw W) (if (= N -2) 1 (thaw W))))))
(define jmp-b N -> (if (= N 0) done (let W (freeze (jmp-a (- N 1))) (if (= N -1) (thaw W) (if (= N -2) 1 (thaw W))))))
(define pat-a [a X] N -> X [a X Y] N -> Y [b X] 0 -> done [b X] N -> (pat-b [c X] (- N 1)) [c X] N -> 0 _ _ -> 1)
(define pat-b [a X] N -> X [a X Y] N -> Y [c X] 0 -> done [c X] N -> (pat-a [b X] (- N 1)) [d X] N -> 0 _ _ -> 1)
(print [(and-a 5000000) (or-a 5000000) (if-a 5000000) (cond-a 5000000) (let-a 5000000) (var-go (function var-a) 5000000)
        (fn-a 5000000) (frz-a 5000000) (jmp-a 5000000) (pat-a [b 1] 5000000)])
`
	file := filepath.Join(t.TempDir(), "s.shen")
	if err := os.WriteFile(file, []byte(script), 0o644); err != nil {
		t.Fatal(err)
	}
	out, err := exec.Command(buildShen(t), "script", file).CombinedOutput()
	if err != nil {
		t.Fatalf("%v: %.2000s", err, out)
	}
	const want = "[true false done done done done done done done done]"
	if !strings.Contains(string(out), want) {
		t.Fatalf("got %.2000s, want %s", out, want)
	}
}
