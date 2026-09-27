package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
	// Build the interpreter once through the shared bounded helper. Using
	// `go run` here makes the test's runtime deadline include a cold rebuild;
	// under the Nix check hook that rebuild can exceed 30 seconds before the
	// child has produced any output.
	bin := buildShen(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "eval", "-e", expr)
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

// TestCompileFileLowersLetToBind pins the let lowering: a let becomes a
// (bind X V) register binding, not ((lambda [X] Body) V), which allocated a
// closure per let and entered the body through TailApply. Binders are renamed
// apart so shadowing lets and a let reusing a parameter's name never declare
// one Go variable twice, and an inner lambda rebinding the name keeps its own.
func TestCompileFileLowersLetToBind(t *testing.T) {
	compiled := compileSource(t, "(defun f (X) (let X (+ X 1) (let X (* X 2) (let _ (g X) (lambda X X)))))")
	if strings.Contains(compiled, "tailapply tmp") || strings.Contains(compiled, "call tmp") {
		t.Fatalf("let still applied through a closure: %s", compiled)
	}
	binds := regexp.MustCompile(`\(bind (\S+) `).FindAllStringSubmatch(compiled, -1)
	if len(binds) != 3 {
		t.Fatalf("want 3 bind forms, got %d: %s", len(binds), compiled)
	}
	seen := map[string]bool{"X": true}
	for _, b := range binds {
		if seen[b[1]] {
			t.Fatalf("bind target %s is not fresh: %s", b[1], compiled)
		}
		seen[b[1]] = true
	}
	// The second let's value reads the first binder, (g X) reads the second.
	if !strings.Contains(compiled, "($global *) "+binds[0][1]+" ") || !strings.Contains(compiled, "($global g) "+binds[1][1]+")") {
		t.Fatalf("let values do not read the enclosing binders: %s", compiled)
	}
	if !strings.Contains(compiled, "(lambda (X) (return X))") {
		t.Fatalf("inner lambda X was renamed: %s", compiled)
	}
}
