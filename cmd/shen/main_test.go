// This file holds OUR OWN tests for the shen CLI — distinct from the canonical
// Shen kernel certification suite, which lives in package ./certification.
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestPipedStdinEOFExitsRepl verifies the CLI exits cleanly when piped stdin
// reaches EOF, rather than looping forever on "error: empty stream".
func TestPipedStdinEOFExitsRepl(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", ".")
	cmd.Stdin = strings.NewReader("(version)\n")

	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("shen did not exit after stdin EOF; output:\n%s", output)
	}
	if err != nil {
		t.Fatalf("shen exited with error: %v\n%s", err, output)
	}

	out := string(output)
	if !strings.Contains(out, "41.2") {
		t.Fatalf("expected version output, got:\n%s", out)
	}
	if strings.Contains(out, "error: empty stream") {
		t.Fatalf("unexpected empty stream error after stdin EOF:\n%s", out)
	}
}

// runCLI runs the shen CLI with the given arguments and returns combined output.
func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", append([]string{"run", "."}, args...)...)
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("shen %v did not exit; output:\n%s", args, output)
	}
	return string(output), err
}

// TestLauncherEval verifies the kernel launcher CLI protocol
// (extension-launcher.kl): `shen eval -e EXPR` evaluates and prints.
func TestLauncherEval(t *testing.T) {
	out, err := runCLI(t, "eval", "-e", "(+ 1 2)")
	if err != nil {
		t.Fatalf("shen eval exited with error: %v\n%s", err, out)
	}
	if !strings.Contains(out, "3") {
		t.Fatalf("expected eval result 3, got:\n%s", out)
	}
}

// TestLauncherQuietFileWrite verifies quiet mode (-q sets *hush*) does not
// suppress pr writes to file streams — the failure mode that produced 0-byte
// output files when hosting the ratatoskr stage-1 shaker. See fixPrHush.
func TestLauncherQuietFileWrite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "out.txt")
	expr := `(let S (open "` + path + `" out) (do (pr "payload" S) (close S)))`
	out, err := runCLI(t, "eval", "-q", "-e", expr)
	if err != nil {
		t.Fatalf("shen eval -q exited with error: %v\n%s", err, out)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("output file not written: %v", err)
	}
	if string(data) != "payload" {
		t.Fatalf("expected file content %q, got %q", "payload", data)
	}
}

// TestLauncherVersionAndBadArgs verifies --version bypasses the Go flag parser
// and unknown commands fail with a nonzero exit.
func TestLauncherVersionAndBadArgs(t *testing.T) {
	out, err := runCLI(t, "--version")
	if err != nil {
		t.Fatalf("shen --version exited with error: %v\n%s", err, out)
	}
	if !strings.Contains(out, "41.2") {
		t.Fatalf("expected version string, got:\n%s", out)
	}

	out, err = runCLI(t, "no-such-command", "x")
	if err == nil {
		t.Fatalf("expected nonzero exit for unknown command, got success:\n%s", out)
	}
	if !strings.Contains(out, "Invalid argument") {
		t.Fatalf("expected invalid-argument error, got:\n%s", out)
	}
}

// runCLIStdin runs the shen CLI with the given stdin (one form per line),
// returns combined output and the exit error.
func runCLIStdin(t *testing.T, stdin string, args ...string) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", append([]string{"run", "."}, args...)...)
	cmd.Stdin = strings.NewReader(stdin)
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("shen %v did not exit; output:\n%s", args, output)
	}
	return string(output), err
}

// TestReplSurvivesApplyNonFunction is the freeze of the exact user-reported
// repro: piping `(overflow->str)` then `(+ 40 2)` through the REPL. Before
// the fix this killed the process (SIGSEGV from IsError(nil)) before the
// second form ever ran. After the fix the REPL prints the error, returns
// to the prompt, and evaluates the next form.
func TestReplSurvivesApplyNonFunction(t *testing.T) {
	out, err := runCLIStdin(t, "(overflow->str)\n(+ 40 2)\n")
	if err != nil {
		t.Fatalf("repl exited nonzero: %v\n%s", err, out)
	}
	if !strings.Contains(out, "can't apply non function: overflow->str") {
		t.Fatalf("expected the catchable error message, got:\n%s", out)
	}
	if !strings.Contains(out, "42") {
		t.Fatalf("expected second form to evaluate to 42, got:\n%s", out)
	}
	// The hallmark of the uncatchable-panic regression is a Go goroutine
	// stack dump in the output. Refuse it.
	if strings.Contains(out, "goroutine ") && strings.Contains(out, "[running]") {
		t.Fatalf("repl leaked a Go goroutine dump (uncatchable panic regression):\n%s", out)
	}
}

// TestReplSurvivesAdversarialSession walks the REPL through a sequence of
// errors interleaved with valid forms, and asserts every prompt got an
// answer line, no goroutine dump leaked, and the final arithmetic value
// is correct. This is the integration counterpart to
// TestEvalSurvivesAdversarialSequence in package kl.
func TestReplSurvivesAdversarialSession(t *testing.T) {
	input := strings.Join([]string{
		"(overflow->str)",          // apply unbound symbol
		"(value never-bound-xyz)",  // unbound variable
		`(simple-error "boom")`,    // explicit error
		"(if 42 1 2)",              // type error on if
		"(+ 1 2)",                  // valid form, must work
		"(42 1)",                   // apply non-function literal
		"(* 6 7)",                  // valid form, final answer
	}, "\n") + "\n"

	out, err := runCLIStdin(t, input)
	if err != nil {
		t.Fatalf("repl exited nonzero: %v\n%s", err, out)
	}
	if strings.Contains(out, "goroutine ") && strings.Contains(out, "[running]") {
		t.Fatalf("repl leaked a Go goroutine dump:\n%s", out)
	}
	// Each prompt N must appear: (0-) ... (7-).
	for n := 0; n <= 7; n++ {
		marker := fmt.Sprintf("(%d-)", n)
		if !strings.Contains(out, marker) {
			t.Fatalf("missing prompt %s — repl quit early:\n%s", marker, out)
		}
	}
	// Specific spot checks on catchable errors.
	for _, msg := range []string{
		"can't apply non function: overflow->str",
		"variable never-bound-xyz not bound",
		"boom",
		"if requires a boolean",
	} {
		if !strings.Contains(out, msg) {
			t.Fatalf("expected error message %q in transcript, got:\n%s", msg, out)
		}
	}
	// Valid arithmetic still works after all those errors.
	if !strings.Contains(out, "42") {
		t.Fatalf("(*) 6 7 should print 42, got:\n%s", out)
	}
}

// TestLauncherEvalReportsError verifies `shen eval -e EXPR` on an
// adversarial input: it must exit nonzero, print a clean error line, and
// NOT leak a Go panic trace.
func TestLauncherEvalReportsError(t *testing.T) {
	out, err := runCLI(t, "eval", "-e", "(overflow->str)")
	if err == nil {
		t.Fatalf("expected nonzero exit on adversarial eval, got success:\n%s", out)
	}
	if !strings.Contains(out, "can't apply non function: overflow->str") {
		t.Fatalf("expected catchable error message in output, got:\n%s", out)
	}
	if strings.Contains(out, "goroutine ") && strings.Contains(out, "[running]") {
		t.Fatalf("launcher leaked a Go goroutine dump:\n%s", out)
	}
}
