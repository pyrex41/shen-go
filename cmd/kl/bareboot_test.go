package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestBareEvalLoopBootsKernel is issue #46's minimal reproduction: feed the
// kernel's load-file forms to the bare `kl` interpreter, which installs no
// natives, and require the kernel to boot. Until nativeVectorRefOr accepted the
// KL kernel's own fail filler, declarations.kl died with "implementation error
// in shen.change-pointer-value". Run as a subprocess so the interpreter's
// state is exactly what a user of `kl` gets.
func TestBareEvalLoopBootsKernel(t *testing.T) {
	kernelDir, err := kl.FindKernelDir(".")
	if err != nil {
		t.Fatal(err)
	}
	bin := filepath.Join(t.TempDir(), "kl")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	var in strings.Builder
	for _, f := range kl.KernelLoadOrder {
		in.WriteString("(load-file \"" + filepath.Join(kernelDir, f) + "\")\n")
	}
	// Probes: the arity table declarations.kl built with the interpreted put
	// must read back, through the interpreted get and through arity.
	in.WriteString("(get map arity (value *property-vector*))\n")
	in.WriteString("(arity map)\n")
	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader(in.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("kl: %v\n%s", err, out.String())
	}
	got := out.String()
	for _, bad := range []string{"Panic:", "Recovered in", "implementation error"} {
		if strings.Contains(got, bad) {
			t.Fatalf("bare kl loop failed to boot the kernel (%q in output):\n%s", bad, got)
		}
	}
	// Each prompt line ends with the value printed for that form; the two
	// probes are the last two.
	lines := strings.Split(strings.TrimSpace(got), "\n")
	// The loop prints a prompt for the form it never gets (EOF); drop it.
	if last := lines[len(lines)-1]; strings.HasSuffix(last, "#>") {
		lines = lines[:len(lines)-1]
	}
	if len(lines) < 2 {
		t.Fatalf("unexpected output:\n%s", got)
	}
	for _, l := range lines[len(lines)-2:] {
		if !strings.HasSuffix(l, " 2") {
			t.Errorf("probe line %q should end with the arity 2\n%s", l, got)
		}
	}
}
