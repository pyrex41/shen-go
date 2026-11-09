package certification

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"testing"
)

// deepTailCases are tail calls that must stay tail calls however the call is
// reached. The conformance suites never go deep enough to notice: a port passed
// all of them with mutual recursion through the last operand of and/or eating
// its stack. D is the depth, and allocates marks a case that allocates per
// iteration, because (F X Y) through a variable is curried, so its footprint
// rises to the collector's plateau.
var deepTailCases = []struct {
	name, defs, call, want string
	allocates              bool
}{
	{name: "and", want: "true", call: "(ev? D)", defs: `
(define ev? 0 -> true N -> (and (number? N) (od? (- N 1))))
(define od? 0 -> false N -> (and (number? N) (ev? (- N 1))))`},
	{name: "or", want: "true", call: "(ev? D)", defs: `
(define ev? 0 -> true N -> (or (not (number? N)) (od? (- N 1))))
(define od? 0 -> false N -> (or (not (number? N)) (ev? (- N 1))))`},
	{name: "if-else", want: "true", call: "(ev? D)", defs: `
(define ev? N -> (if (= N 0) true (od? (- N 1))))
(define od? N -> (if (= N 0) false (ev? (- N 1))))`},
	{name: "if-then", want: "true", call: "(ev? D)", defs: `
(define ev? N -> (if (not (= N 0)) (od? (- N 1)) true))
(define od? N -> (if (not (= N 0)) (ev? (- N 1)) false))`},
	{name: "cond-last", want: "true", call: "(ev? D)", defs: `
(define ev? N -> (cond ((= N 0) true) ((= N -1) false) (true (od? (- N 1)))))
(define od? N -> (cond ((= N 0) false) ((= N -1) false) (true (ev? (- N 1)))))`},
	{name: "cond-first", want: "true", call: "(ev? D)", defs: `
(define ev? N -> (cond ((> N 0) (od? (- N 1))) (true true)))
(define od? N -> (cond ((> N 0) (ev? (- N 1))) (true false)))`},
	{name: "let", want: "true", call: "(ev? D)", defs: `
(define ev? 0 -> true N -> (let M (- N 1) (od? M)))
(define od? 0 -> false N -> (let M (- N 1) (ev? M)))`},
	{name: "do", want: "true", call: "(ev? D)", defs: `
(define ev? 0 -> true N -> (do N (od? (- N 1))))
(define od? 0 -> false N -> (do N (ev? (- N 1))))`},
	{name: "nested", want: "true", call: "(ev? D)", defs: `
(define ev? 0 -> true N -> (and (number? N) (or false (let M (- N 1) (if (> M -1) (od? M) false)))))
(define od? 0 -> false N -> (and (number? N) (or false (let M (- N 1) (if (> M -1) (ev? M) false)))))`},
	{name: "var-self", want: "done", call: "(via (fn via) D)", allocates: true, defs: `
(define via _ 0 -> done F N -> (F F (- N 1)))`},
	{name: "var-mutual", want: "done", call: "(va (fn vb) (fn va) D)", allocates: true, defs: `
(define va _ _ 0 -> done F G N -> (F G F (- N 1)))
(define vb _ _ 0 -> done F G N -> (F G F (- N 1)))`},
	{name: "var-lambda", want: "done", call: "(lp (/. N (if (= N 0) done (lp2 N))) D)", allocates: true, defs: `
(define lp F N -> (F N))
(define lp2 N -> (lp (/. M (if (= M 0) done (lp2 M))) (- N 1)))`},
}

var resultRE = regexp.MustCompile(`RESULT (\S+)`)

// runDeepTail runs one case at depth and returns its result and peak RSS in MB.
func runDeepTail(t *testing.T, bin, dir string, env []string, defs, call string, depth int) (string, float64) {
	t.Helper()
	src := defs + "\n(pr (make-string \"RESULT ~A~%\" " + strings.ReplaceAll(call, "D", fmt.Sprint(depth)) + ") (stoutput))\n"
	file := filepath.Join(dir, "case.shen")
	if err := os.WriteFile(file, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(bin, "script", file)
	cmd.Dir = dir
	cmd.Env = env
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("depth %d: %v\n%s", depth, err, tail(string(out), 6))
	}
	m := resultRE.FindSubmatch(out)
	if m == nil {
		t.Fatalf("depth %d: no result in\n%s", depth, tail(string(out), 6))
	}
	rss := float64(cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss)
	if runtime.GOOS == "darwin" {
		rss /= 1024 // bytes on darwin, KiB elsewhere
	}
	return string(m[1]), rss / 1024
}

// TestDeepMutualTailCalls is the gate for unbounded tail calls: depth 5M in every
// tail position, with footprint flat. A broken tail call dies
// of stack overflow long before 5M (Go's limit is reached near 1.4M), and a slow
// leak shows as growth past the allowance.
func TestDeepMutualTailCalls(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping deep tail calls in -short mode")
	}
	if runtime.GOOS == "windows" {
		t.Skip("uses syscall.Rusage")
	}
	bin := filepath.Join(t.TempDir(), "shen")
	if out, err := exec.Command("go", "build", "-o", bin, "../cmd/shen").CombinedOutput(); err != nil {
		t.Fatalf("failed to build shen: %v\n%s", err, out)
	}
	{
		env := os.Environ()
		for _, tc := range deepTailCases {
			t.Run(tc.name, func(t *testing.T) {
				dir := t.TempDir()
				lo, hi, allowance := 50000, 5000000, 16.0
				if tc.allocates {
					// These rise to the collector's plateau (about 85 MB, set by
					// the ballast) only by about 5M. var-self goes on to 80M, past
					// the point where a fixnum read that pages in the sentinel
					// array shows: that grew it by 62 MB from 20M to 80M.
					lo, hi, allowance = 5000000, 20000000, 32.0
					if tc.name == "var-self" {
						lo, hi = 20000000, 80000000
					}
				}
				gotLo, rssLo := runDeepTail(t, bin, dir, env, tc.defs, tc.call, lo)
				gotHi, rssHi := runDeepTail(t, bin, dir, env, tc.defs, tc.call, hi)
				if gotLo != tc.want || gotHi != tc.want {
					t.Fatalf("got %s at %d and %s at %d, want %s", gotLo, lo, gotHi, hi, tc.want)
				}
				if rssHi > rssLo+allowance {
					t.Fatalf("footprint grew: %.0f MB at depth %d, %.0f MB at depth %d", rssLo, lo, rssHi, hi)
				}
				t.Logf("depth %d: %.0f MB, depth %d: %.0f MB", lo, rssLo, hi, rssHi)
			})
		}
	}
}
