package main

import (
	"encoding/json"
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

// TestArtifactBootsWithLoweredKernelDefuns pins issue #49 end to end. A
// Yggdrasil `lower --target go` slice deletes the top-level (defun NAME …) of
// every declared native_override, on the strength of kl/equiv.json saying the
// native replaces it. That only works if InstallKernelFast installs the native
// whether or not the kernel defined the name. Before the fix the artifact died
// while loading the kernel chunks with "variable vector not bound"
// (declarations.kl's (set *property-vector* (vector 20000)) runs before any
// user code; sys.kl's dropped (defun vector …) never triggered the install).
//
// kl/equiv.json is decoded with encoding/json: the file escapes < and > as
// < / >, so a textual match would silently keep <-vector, vector->,
// string->symbol, shen.byte->digit and shen.string->bytes.
func TestArtifactBootsWithLoweredKernelDefuns(t *testing.T) {
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

	equivSrc, err := os.ReadFile(filepath.Join(root, "kl", "equiv.json"))
	if err != nil {
		t.Fatal(err)
	}
	var equiv struct {
		Rows []struct {
			KernelFn string `json:"kernel_fn"`
		} `json:"rows"`
	}
	if err := json.Unmarshal(equivSrc, &equiv); err != nil {
		t.Fatalf("decode kl/equiv.json: %v", err)
	}
	lowered := map[string]int{}
	for _, r := range equiv.Rows {
		if r.KernelFn == "" {
			t.Fatal("kl/equiv.json row with empty kernel_fn")
		}
		if _, dup := lowered[r.KernelFn]; dup {
			t.Fatalf("kl/equiv.json lists %s twice", r.KernelFn)
		}
		lowered[r.KernelFn] = 0
	}

	var kernel strings.Builder
	dropped := 0
	for _, f := range kl.KernelLoadOrder {
		src, err := os.ReadFile(filepath.Join(root, "kernel", "klambda", f))
		if err != nil {
			t.Fatal(err)
		}
		forms, err := splitForms(src)
		if err != nil {
			t.Fatalf("%s: %v", f, err)
		}
		for _, form := range forms {
			if name, ok := defunName(form); ok {
				if _, drop := lowered[name]; drop {
					lowered[name]++
					dropped++
					continue
				}
			}
			kernel.Write(form)
			kernel.WriteString("\n")
		}
	}
	// Every row must correspond to exactly one dropped top-level defun, so a
	// row whose body is silently kept (or defined twice) is noticed here rather
	// than by the artifact passing for the wrong reason.
	for name, n := range lowered {
		if n != 1 {
			t.Errorf("kl/equiv.json row %s: dropped %d top-level defuns, want 1", name, n)
		}
	}
	if dropped != len(equiv.Rows) {
		t.Fatalf("dropped %d defuns, want one per kl/equiv.json row (%d)", dropped, len(equiv.Rows))
	}
	kernel.WriteString("(defun shen.initialise () (do (put e2e.x arity 7 (value *property-vector*)) " +
		"(put e2e.x e2e.tag (cons a (cons b ())) (value *property-vector*))))\n")
	// Each line calls at least one name whose KL body was dropped above:
	// vector/limit, get/put (via shen.initialise), arity, reverse, map, empty?,
	// <-vector on the fail filler, and fail itself.
	prog := `(pr (shen.app (limit (vector 3)) " " shen.a) (value *stoutput*))
(pr (shen.app (get e2e.x arity (value *property-vector*)) " " shen.a) (value *stoutput*))
(pr (shen.app (get e2e.x e2e.tag (value *property-vector*)) " " shen.a) (value *stoutput*))
(pr (shen.app (arity e2e.x) " " shen.a) (value *stoutput*))
(pr (shen.app (reverse (cons 1 (cons 2 (cons 3 ())))) " " shen.a) (value *stoutput*))
(pr (shen.app (map (lambda X (+ X 1)) (cons 1 (cons 2 ()))) " " shen.a) (value *stoutput*))
(pr (shen.app (empty? ()) " " shen.a) (value *stoutput*))
(pr (shen.app (trap-error (<-vector (vector 1) 1) (lambda E e2e.caught)) " " shen.a) (value *stoutput*))
(pr (shen.app (= (fail) (fail)) " " shen.a) (value *stoutput*))
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
		t.Fatalf("lowered artifact (%d kernel defuns dropped): %v\n%s", dropped, err, o)
	}
	if got, want := string(o), "3 7 [a b] 7 [3 2 1] [2 3] true e2e.caught true "; got != want {
		t.Fatalf("lowered artifact printed %q, want %q", got, want)
	}
}

// defunName returns NAME for a top-level `(defun NAME …)` form.
func defunName(form []byte) (string, bool) {
	rest, ok := strings.CutPrefix(string(form), "(defun ")
	if !ok {
		return "", false
	}
	rest = strings.TrimLeft(rest, " \t\r\n")
	end := strings.IndexAny(rest, " \t\r\n()")
	if end <= 0 {
		return "", false
	}
	return rest[:end], true
}
