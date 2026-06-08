// Package certification runs the CANONICAL Shen kernel test suite — the official
// ShenOSKernel acceptance tests shipped under kernel/tests (runme.shen ->
// harness.shen + kerneltests.shen). This is the external bar for "Certified" on
// shen-language.github.io; it is NOT part of our own test coverage.
//
// Keep this package isolated from our unit tests (which live in package kl and
// cmd/shen) so it is unambiguous that:
//
//   - this suite passes the canonical Shen language tests, 134/134, and
//   - our own Go unit tests are a separate, additional thing.
//
// Run just the certification:   make certify   (or: go test ./certification/)
// Run just our unit tests:       make test
// Plain `go test ./...` runs both.
package certification

import (
	"context"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

// TestKernelCertification builds the shen binary and runs the canonical kernel
// suite end-to-end, asserting a clean 100% pass rate (134/134, no failures).
//
// Loading in the suite is relative to the process working directory (the native
// reader ignores *home-directory*), so the binary runs with its CWD set to
// kernel/tests. It exits on stdin EOF (see cmd/shen repl), so no input draining
// or process kill is needed.
//
// This is an integration test (~10s: it builds and runs the binary). Skip it
// with `go test -short`.
func TestKernelCertification(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping canonical kernel certification suite in -short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// Build the binary into a temp dir.
	bin := filepath.Join(t.TempDir(), "shen")
	build := exec.CommandContext(ctx, "go", "build", "-o", bin, "../cmd/shen")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build shen: %v\n%s", err, out)
	}

	testsDir, err := filepath.Abs(filepath.Join("..", "kernel", "tests"))
	if err != nil {
		t.Fatalf("resolving kernel/tests dir: %v", err)
	}

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
	// regression against the canonical suite.
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
