package main

import (
	"bytes"
	"os"
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
	bin := buildKL(t)
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
	// A form that raises prints `N #> Error(...)`; a clean boot prints none.
	for _, bad := range []string{"Panic:", "Recovered in", "Error(", "implementation error"} {
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

// buildKL builds the bare kl interpreter into a temporary directory.
func buildKL(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "kl")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

// runKL feeds stdin to the bare interpreter with extra environment entries
// and returns stdout and stderr separately.
func runKL(t *testing.T, bin, stdin string, env ...string) (stdout, stderr string) {
	t.Helper()
	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader(stdin)
	cmd.Env = append(os.Environ(), env...)
	var out, errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		t.Fatalf("kl: %v\n%s%s", err, out.String(), errb.String())
	}
	return out.String(), errb.String()
}

// TestDebugRecoverTracesToStderr is issue #52's diagnostic switch. An
// uncaught KL error prints `N #> Error(MESSAGE)` and nothing else by default;
// with SHEN_DEBUG_RECOVER=1 the recover trace (expression, message, Go stack)
// appears on stderr and still never on stdout. debugRecover is read at
// package init, so this runs the interpreter as a subprocess.
func TestDebugRecoverTracesToStderr(t *testing.T) {
	bin := buildKL(t)
	const form = "(tl 5)\n"
	const trace = "kl: recovered in Eval"

	stdout, stderr := runKL(t, bin, form, "SHEN_DEBUG_RECOVER=")
	if stderr != "" {
		t.Errorf("stderr should be empty without SHEN_DEBUG_RECOVER, got:\n%s", stderr)
	}
	if !strings.Contains(stdout, "Error(") || strings.Contains(stdout, "goroutine") {
		t.Errorf("stdout should print the error object without a stack dump, got:\n%s", stdout)
	}

	stdout, stderr = runKL(t, bin, form, "SHEN_DEBUG_RECOVER=1")
	for _, want := range []string{trace, "(tl 5)", "goroutine "} {
		if !strings.Contains(stderr, want) {
			t.Errorf("SHEN_DEBUG_RECOVER=1: stderr lacks %q:\n%s", want, stderr)
		}
	}
	for _, bad := range []string{trace, "goroutine"} {
		if strings.Contains(stdout, bad) {
			t.Errorf("SHEN_DEBUG_RECOVER=1: %q leaked to stdout:\n%s", bad, stdout)
		}
	}
	if !strings.Contains(stdout, "Error(") {
		t.Errorf("SHEN_DEBUG_RECOVER=1: stdout should still print the error object, got:\n%s", stdout)
	}
}
