package main

import (
	"context"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestPipedStdinEOFExitsRepl(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
	if !strings.Contains(out, "41.1") {
		t.Fatalf("expected version output, got:\n%s", out)
	}
	if strings.Contains(out, "error: empty stream") {
		t.Fatalf("unexpected empty stream error after stdin EOF:\n%s", out)
	}
}
