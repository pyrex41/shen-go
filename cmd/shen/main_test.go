// This file holds OUR OWN tests for the shen CLI — distinct from the canonical
// Shen kernel certification suite, which lives in package ./certification.
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
