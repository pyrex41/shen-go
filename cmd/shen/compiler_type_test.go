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

// compileSource exercises the same compile-file path used by the generated
// kernel, rather than testing a Go-side reimplementation of the Shen parser.
func compileSource(t *testing.T, source string) string {
	t.Helper()
	in := filepath.Join(t.TempDir(), "input.kl")
	out := filepath.Join(t.TempDir(), "output.bc")
	if err := os.WriteFile(in, []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	compiler, err := filepath.Abs("../../src/compiler.shen")
	if err != nil {
		t.Fatal(err)
	}
	expr := "(do (load " + shenString(compiler) + ") (compile-file " +
		shenString(in) + " " + shenString(out) + "))"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "run", ".", "eval", "-e", expr)
	result, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("compile-file failed: %v\n%s", err, result)
	}
	compiled, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	return string(compiled)
}

// Shen string literals use double quotes; test paths are temporary and do not
// contain quotes, so this escaping is sufficient and keeps the expression
// readable in failure output.
func shenString(path string) string {
	return `"` + strings.ReplaceAll(path, `"`, `\"`) + `"`
}

func TestCompileFilePreservesTypeWrappers(t *testing.T) {
	compiled := compileSource(t, "(do (type (+ 1 2) number) (defun typed-id (X) (type X number)))")
	if !strings.Contains(compiled, "ignore ($type number") {
		t.Fatalf("non-tail type wrapper missing from compiled IR: %s", compiled)
	}
	if !strings.Contains(compiled, "return ($type number") {
		t.Fatalf("tail type wrapper missing from compiled IR: %s", compiled)
	}
}

func TestCompileFileMalformedTypeRemainsLegacy(t *testing.T) {
	compiled := compileSource(t, "(type (+ 1 2))")
	if strings.Contains(compiled, "$type") {
		t.Fatalf("malformed type form unexpectedly serialized as $type: %s", compiled)
	}
	if !strings.Contains(compiled, "($global type)") {
		t.Fatalf("malformed type form did not retain legacy global call: %s", compiled)
	}
}
