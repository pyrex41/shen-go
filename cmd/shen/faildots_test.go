package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCompiledKernelHasNoFailDots guards the regenerated kernel against the
// serialization accident of issue #54. compile-file (src/compiler.shen) once
// printed the kernel IR through ~R, whose shen.arg->str renders anything equal
// to (fail) as `...`, so (defun fail () shen.fail!) landed in cmd/shen/sys.go
// as a return of the symbol `...`. The generated kernel must never declare
// that symbol: nothing in kernel/klambda uses `...` as data.
func TestCompiledKernelHasNoFailDots(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no Go files in cmd/shen")
	}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(data), "\n") {
			if strings.Contains(line, `MakeSymbol("...")`) {
				t.Errorf("%s:%d declares the symbol `...`: the kernel was regenerated through the ~R printer (issue #54)", f, i+1)
			}
		}
	}
}
