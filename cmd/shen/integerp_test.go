// D3: (integer? X) must terminate for every X.
//
// The kernel defines integer? in sys.kl as
//
//	(defun integer? (V) (and (number? V)
//	                     (let W (shen.abs V)
//	                       (shen.integer-test? W (shen.magless W 1)))))
//
// and shen.magless doubles its accumulator until it exceeds |V|. Against
// infinity nothing ever exceeds it, so the call spins forever. A primitive
// that does not terminate is wrong under every reading of the spec, whatever
// one thinks overflow should do.
//
// The kernel's `defun` overwrites the native `integer?` binding at load time,
// so a guard installed in the primitive table alone would be dead code (that
// is exactly the state shen-lua's prims.lua:250 guard is in). These tests run
// the real binary after a full kernel load, which is the only place the answer
// counts.
package main

import (
	"context"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// buildShen compiles the CLI once for the calling test and returns its path.
func buildShen(t *testing.T) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	bin := filepath.Join(t.TempDir(), "shen")
	build := exec.CommandContext(ctx, "go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build shen: %v\n%s", err, out)
	}
	return bin
}

// evalBounded evaluates expr with the given binary under a hard deadline, so a
// non-terminating primitive fails the test instead of hanging the suite.
func evalBounded(t *testing.T, bin, expr string, timeout time.Duration) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "eval", "-e", expr)
	out, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("%s did not terminate within %s; output so far:\n%s", expr, timeout, out)
	}
	if err != nil {
		t.Fatalf("%s exited with error: %v\n%s", expr, err, out)
	}
	return strings.TrimSpace(string(out))
}

// Non-finite arguments are built arithmetically rather than written as
// literals, so the test does not depend on how the reader spells infinity.
const (
	klInf    = `(* (/ 1.0 1e-300) (/ 1.0 1e-300))`
	klNegInf = `(- 0 ` + klInf + `)`
	klNaN    = `(- ` + klInf + ` ` + klInf + `)`
)

func TestIntegerPTerminatesOnNonFinite(t *testing.T) {
	bin := buildShen(t)

	for _, tc := range []struct{ name, expr, want string }{
		{"+inf", `(integer? ` + klInf + `)`, "false"},
		{"-inf", `(integer? ` + klNegInf + `)`, "false"},
		{"nan", `(integer? ` + klNaN + `)`, "false"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := evalBounded(t, bin, tc.expr, 30*time.Second); got != tc.want {
				t.Fatalf("%s: got %q, want %q", tc.expr, got, tc.want)
			}
		})
	}
}

// The guard must not change the answer for anything finite: those still go
// through the kernel's own definition.
func TestIntegerPFiniteAnswersUnchanged(t *testing.T) {
	bin := buildShen(t)

	for _, tc := range []struct{ expr, want string }{
		{`(integer? 0)`, "true"},
		{`(integer? 5)`, "true"},
		{`(integer? -3)`, "true"},
		{`(integer? 5.5)`, "false"},
		{`(integer? -0.25)`, "false"},
		{`(integer? 1e18)`, "true"},
		{`(integer? 1e19)`, "true"},
		{`(integer? "abc")`, "false"},
		{`(integer? (cons 1 2))`, "false"},
		{`(integer? a)`, "false"},
	} {
		if got := evalBounded(t, bin, tc.expr, 30*time.Second); got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.expr, got, tc.want)
		}
	}
}
