package main

import (
	"context"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

// TestKernelCertification runs the official Shen kernel test suite (the bar for
// "Certified" on shen-language.github.io) end-to-end and asserts a clean 100%
// pass rate. The suite lives in-repo at kernel/tests; loading is relative to the
// process working directory (the native reader ignores *home-directory*), so we
// run the binary with its CWD set to that directory.
//
// This is an integration test (~10s, builds and runs the binary). Skip it with
// `go test -short` for a sub-second unit-test cycle.
func TestKernelCertification(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping kernel certification suite in -short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// Build the binary into a temp dir (CWD here is the cmd/shen package dir).
	bin := filepath.Join(t.TempDir(), "shen")
	build := exec.CommandContext(ctx, "go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build shen: %v\n%s", err, out)
	}

	testsDir, err := filepath.Abs(filepath.Join("..", "..", "kernel", "tests"))
	if err != nil {
		t.Fatalf("resolving kernel/tests dir: %v", err)
	}

	// Feed the load command then let stdin EOF, so the REPL exits cleanly
	// (relies on the stdin-EOF fix; see main.go repl).
	cmd := exec.CommandContext(ctx, bin)
	cmd.Dir = testsDir
	cmd.Stdin = strings.NewReader(`(load "runme.shen")` + "\n")

	output, err := cmd.CombinedOutput()
	out := string(output)
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("kernel suite did not finish before timeout; output tail:\n%s", tail(out, 40))
	}
	if err != nil {
		t.Fatalf("shen exited with error: %v\noutput tail:\n%s", err, tail(out, 40))
	}

	// Every report section ends with a cumulative "failed ... N". Any N>=1 is a
	// regression.
	if m := regexp.MustCompile(`failed \.\.\. [1-9]`).FindString(out); m != "" {
		t.Fatalf("kernel suite reported failures (%q); output tail:\n%s", m, tail(out, 40))
	}
	if !strings.Contains(out, "pass rate ... 100%") {
		t.Fatalf("did not find a 100%% pass rate; output tail:\n%s", tail(out, 40))
	}
	if !strings.Contains(out, "failed ... 0") {
		t.Fatalf("did not find a clean 'failed ... 0' report; output tail:\n%s", tail(out, 40))
	}
}

// tail returns the last n lines of s, for compact failure messages.
func tail(s string, n int) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return strings.Join(lines, "\n")
}
