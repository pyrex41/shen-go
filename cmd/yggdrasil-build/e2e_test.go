package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestArtifactBootsWithInitOnNatives builds a real artifact from the vendored
// kernel plus a manifest init= function that puts into *property-vector*, then
// runs it. It pins what TestGenMainInstallsNativesBeforeInit only checks in
// the emitted text: the natives are installed as the kernel chunks load, the
// init pass runs on them, and the property vectors built by the kernel's
// inline init are read back by the native get and arity. Before the natives
// accepted the KL kernel's fail filler, this run died in shen.initialise with
// "implementation error in shen.change-pointer-value".
func TestArtifactBootsWithInitOnNatives(t *testing.T) {
	if testing.Short() {
		t.Skip("builds the builder and an artifact")
	}
	root, err := findShenGoRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	tmp := t.TempDir()
	shaken := filepath.Join(tmp, "shaken")
	out := filepath.Join(tmp, "artifact")
	if err := os.MkdirAll(shaken, 0o755); err != nil {
		t.Fatal(err)
	}

	var kernel strings.Builder
	for _, f := range kl.KernelLoadOrder {
		src, err := os.ReadFile(filepath.Join(root, "kernel", "klambda", f))
		if err != nil {
			t.Fatal(err)
		}
		kernel.Write(src)
		kernel.WriteString("\n")
	}
	kernel.WriteString("(defun shen.initialise () (do (put e2e.x arity 7 (value *property-vector*)) " +
		"(put e2e.x e2e.tag (cons a (cons b ())) (value *property-vector*))))\n")
	prog := `(pr (shen.app (get e2e.x arity (value *property-vector*)) " " shen.a) (value *stoutput*))
(pr (shen.app (get e2e.x e2e.tag (value *property-vector*)) " " shen.a) (value *stoutput*))
(pr (shen.app (arity e2e.x) " " shen.a) (value *stoutput*))
(pr (shen.app (reverse (cons 1 (cons 2 (cons 3 ())))) " " shen.a) (value *stoutput*))
`
	manifest := "manifest-version=1\nkernel=kernel.kl\ninit=shen.initialise\nuser=prog.kl\nneeds-eval=false\n"
	for name, body := range map[string]string{"kernel.kl": kernel.String(), "prog.kl": prog, "yggdrasil.manifest.txt": manifest} {
		if err := os.WriteFile(filepath.Join(shaken, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	builder := filepath.Join(tmp, "yggdrasil-build")
	if o, err := exec.Command("go", "build", "-o", builder, ".").CombinedOutput(); err != nil {
		t.Fatalf("go build builder: %v\n%s", err, o)
	}
	if o, err := exec.Command(builder, "-shen-go", root, shaken, out).CombinedOutput(); err != nil {
		t.Fatalf("yggdrasil-build: %v\n%s", err, o)
	}
	build := exec.Command("go", "build", "-o", "prog", ".")
	build.Dir = out
	if o, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build artifact: %v\n%s", err, o)
	}
	o, err := exec.Command(filepath.Join(out, "prog")).CombinedOutput()
	if err != nil {
		t.Fatalf("artifact: %v\n%s", err, o)
	}
	if got, want := string(o), "7 [a b] 7 [3 2 1] "; got != want {
		t.Fatalf("artifact printed %q, want %q", got, want)
	}
}
