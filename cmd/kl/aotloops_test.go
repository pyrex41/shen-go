package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/codegen"
	"github.com/pyrex41/shen-go/kl"
)

// TestAOTLetAndSelfTailSemantics compiles KL through src/compiler.shen and
// codegen exactly as compiled/script.kl does for the kernel, builds the Go,
// and runs it. It pins the two lowerings that do not go through the
// trampoline: a let is a Go local (renamed, so shadowed and sequential lets
// of one name coexist), and a defun's tail call of itself is a loop while
// the name is still bound to the running code. The loop re-declares the
// parameters each iteration, so closures made in different iterations keep
// their own values; a function that redefines itself mid-loop continues in
// the new definition; and a step budget still stops the loop.
func TestAOTLetAndSelfTailSemantics(t *testing.T) {
	if testing.Short() {
		t.Skip("boots the kernel and builds generated Go")
	}
	kernelDir, err := kl.FindKernelDir(".")
	if err != nil {
		t.Fatal(err)
	}
	compiler, err := filepath.Abs("../../src/compiler.shen")
	if err != nil {
		t.Fatal(err)
	}
	bin := buildKL(t)
	tmp := t.TempDir()
	src := filepath.Join(tmp, "prog.kl")
	ir := filepath.Join(tmp, "prog.tmp")
	prog := `
(defun count-up (N Acc) (if (= N 0) Acc (count-up (- N 1) (+ Acc 1))))
(defun keep-closures (N Acc) (if (= N 0) Acc (keep-closures (- N 1) (cons (freeze N) Acc))))
(defun thaw-all (L) (if (= L ()) () (cons ((hd L)) (thaw-all (tl L)))))
(defun rebind-midway (N) (if (= N 0) done (if (= N 5) (do (defun rebind-midway (M) (cons rebound M)) (rebind-midway (- N 1))) (rebind-midway (- N 1)))))
(defun shadow (X) (let X (+ X 1) (let X (* X 2) (let F (lambda Y (+ Y X)) (F X)))))
(defun sequential (X) (do (let Y (+ X 1) Y) (let Y (+ X 2) (let Unused (+ X 3) Y))))
(defun forever (N) (forever (+ N 1)))
`
	if err := os.WriteFile(src, []byte(prog), 0o644); err != nil {
		t.Fatal(err)
	}
	var in strings.Builder
	for _, f := range kl.KernelLoadOrder {
		in.WriteString("(load-file \"" + filepath.Join(kernelDir, f) + "\")\n")
	}
	in.WriteString("(load \"" + compiler + "\")\n")
	in.WriteString("(set *maximum-print-sequence-size* 100000)\n")
	in.WriteString(`(compile-file "` + src + `" "` + ir + `")` + "\n")
	stdout, stderr := runKL(t, bin, in.String())
	f, err := os.Open(ir)
	if err != nil {
		t.Fatalf("%v\nkl output:\n%s%s", err, stdout, stderr)
	}
	defer f.Close()
	var gen bytes.Buffer
	cg := codegen.New()
	if err := cg.HandleBody(f, "Generated", &gen); err != nil {
		t.Fatal(err)
	}
	cg.HandleSymbol(&gen)
	if n := strings.Count(gen.String(), "goto __selftop"); n < 4 {
		t.Fatalf("expected self-tail loops in count-up, keep-closures, rebind-midway and forever; found %d:\n%s", n, gen.String())
	}

	dir, err := os.MkdirTemp(".", "_aotloops-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	if err := os.WriteFile(filepath.Join(dir, "generated.go"), gen.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	harness := `package main

import (
	"strings"
	"testing"

	kl "github.com/pyrex41/shen-go/kl"
)

var ns2_1set = kl.PrimFunc(kl.MakeSymbol("defun"))
var try_1catch = kl.PrimFunc(kl.MakeSymbol("try-catch"))

func TestGenerated(t *testing.T) {
	var e kl.ControlFlow
	kl.Call(&e, Generated)
	for _, c := range [][2]string{
		{"(count-up 1000000 0)", "1000000"},
		{"(thaw-all (keep-closures 4 ()))", "(1 2 3 4)"},
		{"(rebind-midway 10)", "(rebound . 4)"},
		{"(shadow 3)", "16"},
		{"(sequential 10)", "12"},
	} {
		exp, err := kl.NewSexpReader(strings.NewReader(c[0]), false).Read()
		if err != nil {
			t.Fatal(err)
		}
		if got := kl.ObjString(kl.Eval(&e, exp)); got != c[1] {
			t.Errorf("%s = %s, want %s", c[0], got, c[1])
		}
	}
	var limited kl.ControlFlow
	limited.SetStepLimit(10000)
	exp, _ := kl.NewSexpReader(strings.NewReader("(forever 0)"), false).Read()
	if got := kl.Eval(&limited, exp); !kl.IsError(got) || !strings.Contains(kl.ObjString(got), "step limit") {
		t.Errorf("(forever 0) under a step limit = %s, want the step-limit error", kl.ObjString(got))
	}
}
`
	if err := os.WriteFile(filepath.Join(dir, "generated_test.go"), []byte(harness), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated Go failed: %v\n%s", err, out)
	}
}
