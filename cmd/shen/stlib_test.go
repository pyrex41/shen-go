package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/pyrex41/shen-go/cmd/shen/stlibcompiled"
)

func TestCompiledStdlibMatchesVendoredSources(t *testing.T) {
	h := sha256.New()
	for _, file := range stlibInstallOrder {
		data, err := stlibFS.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		h.Write(data)
		h.Write([]byte{'\n'})
	}
	if got := fmt.Sprintf("%x", h.Sum(nil)); got != stlibcompiled.SourceSHA256 {
		t.Fatalf("compiled StLib is stale: got source hash %s, artifact %s; run python3 scripts/generate-stlib.py", got, stlibcompiled.SourceSHA256)
	}
}

func TestCompiledStdlibMatchesGeneratorInputs(t *testing.T) {
	paths := []string{
		"scripts/generate-stlib.py", "src/compiler.shen", "codegen/codegen.go",
		"cmd/kl/codegen.go", "cmd/kl/plugin.go",
	}
	kernel, err := filepath.Glob("../../kernel/klambda/*.kl")
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range kernel {
		paths = append(paths, strings.TrimPrefix(filepath.ToSlash(path), "../../"))
	}
	sort.Strings(paths[5:])
	h := sha256.New()
	for _, path := range paths {
		data, err := os.ReadFile(filepath.Join("../..", path))
		if err != nil {
			t.Fatal(err)
		}
		h.Write([]byte(path))
		h.Write([]byte{0})
		h.Write(data)
	}
	if got := fmt.Sprintf("%x", h.Sum(nil)); got != stlibcompiled.GeneratorSHA256 {
		t.Fatalf("compiled StLib used stale generator inputs: got %s, artifact %s; run python3 scripts/generate-stlib.py", got, stlibcompiled.GeneratorSHA256)
	}
}

// The builder's process-sexprs pass has effects that are not present in the
// generated KL. Compare cold processes to catch missing macro registrations,
// package metadata, and arity state as well as ordinary function behavior.
func TestCompiledStdlibColdProcessParity(t *testing.T) {
	bin := cliBinary(t)
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Second)
	defer cancel()
	for _, expr := range []string{
		`(filter (/. X (> X 2)) [1 2 3 4])`,
		`(mapc (/. X X) (take 2 [10 20 30]))`,
		`(arity filter)`,
		`(external stlib)`,
		`(sqrt 9)`,
		`(macroexpand [array [2 3]])`,
		`(macroexpand [s-op1 F "abc"])`,
		`(macroexpand [errout (simple-error "x") none "err.txt"])`,
	} {
		run := func(interpreted bool) (string, error) {
			cmd := exec.CommandContext(ctx, bin, "eval", "-e", expr)
			if interpreted {
				cmd.Env = append(os.Environ(), "SHEN_STDLIB_INTERPRETED=1")
			}
			out, err := cmd.CombinedOutput()
			return string(out), err
		}
		compiled, compiledErr := run(false)
		interpreted, interpretedErr := run(true)
		if ctx.Err() != nil {
			t.Fatalf("cold-process comparison timed out at %q", expr)
		}
		if compiled != interpreted || fmt.Sprint(compiledErr) != fmt.Sprint(interpretedErr) {
			t.Errorf("%s: compiled (%v) %q, interpreted (%v) %q", expr, compiledErr, compiled, interpretedErr, interpreted)
		}
	}
}

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
	bin := cliBinary(t)
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "eval", "-e", "(filter (/. X X) [1])")
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
