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
