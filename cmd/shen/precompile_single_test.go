package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

// `make precompile` of a file with a single top-level form used to give a
// plugin that panicked on load ("partial apply unsupported"): compile-file
// wrapped the forms in (do ...), and (do X) is X, so the module thunk was the
// form itself. The answer must come from the plugin, so the source defines the
// function differently from the script.
func TestPrecompileSingleFormLoads(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a Go plugin through make precompile")
	}
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Skipf("go plugins are not supported on %s", runtime.GOOS)
	}
	if _, err := exec.LookPath("make"); err != nil {
		t.Skip("make not available")
	}
	bin := buildShen(t)
	dir := t.TempDir()
	src := filepath.Join(dir, "one.shen")
	so := filepath.Join(dir, "one.so")
	if err := os.WriteFile(src, []byte("(define only N -> (+ N 1))\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 280*time.Second)
	defer cancel()
	mk := exec.CommandContext(ctx, "make", "precompile", "FILE="+src, "OUT="+so)
	mk.Dir = filepath.Join("..", "..")
	if out, err := mk.CombinedOutput(); err != nil {
		t.Fatalf("make precompile: %v\n%s", err, out)
	}
	driver := filepath.Join(dir, "driver.shen")
	if err := os.WriteFile(driver, []byte("(output \"answer ~A~%\" (only 41))\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out, err := exec.CommandContext(ctx, bin, "-precompiled", so, "script", driver).CombinedOutput()
	if err != nil || !strings.Contains(string(out), "answer 42") || strings.Contains(string(out), "warning") {
		t.Fatalf("shen -precompiled one.so: %v\n%s", err, out)
	}
}
