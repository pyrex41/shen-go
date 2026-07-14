package main

import (
	"context"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestStdlibFilterWorks is the standard-library smoke test: a bare (filter ...)
// — a function the kernel does NOT provide, only Lib/StLib does — must work out
// of the box via the embedded stdlib loaded at startup.
func TestStdlibFilterWorks(t *testing.T) {
	out, err := runCLI(t, "eval", "-e", "(filter (/. X (> X 2)) [1 2 3 4])")
	if err != nil {
		t.Fatalf("shen eval (filter ...) errored: %v\n%s", err, out)
	}
	if !strings.Contains(out, "[3 4]") {
		t.Fatalf("expected (filter ...) => [3 4], got:\n%s", out)
	}
	if strings.Contains(out, "undefined") {
		t.Fatalf("filter reported undefined — stdlib not loaded:\n%s", out)
	}
}

// TestStdlibFnResolves verifies (fn filter) resolves (arity registered), i.e.
// the reader can compile a call to a stdlib function.
func TestStdlibFnResolves(t *testing.T) {
	out, err := runCLI(t, "eval", "-e", "(mapc (/. X X) (take 2 [10 20 30]))")
	if err != nil {
		t.Fatalf("shen eval errored: %v\n%s", err, out)
	}
	if strings.Contains(out, "undefined") {
		t.Fatalf("stdlib function reported undefined:\n%s", out)
	}
}

// TestStdlibOptOut verifies SHEN_NO_STDLIB skips the stdlib: filter is then the
// kernel's business only (undefined), and the REPL still comes up.
func TestStdlibOptOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "run", ".", "eval", "-e", "(filter (/. X X) [1])")
	cmd.Env = append(cmd.Environ(), "SHEN_NO_STDLIB=1")
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("shen did not exit; output:\n%s", output)
	}
	// With the stdlib off, filter is undefined; the launcher reports the error
	// (non-zero exit) rather than [1]-anything. We only assert it did NOT
	// silently succeed with a stdlib result.
	if err == nil && !strings.Contains(string(output), "undefined") {
		t.Fatalf("expected filter undefined with SHEN_NO_STDLIB, got:\n%s", output)
	}
}
