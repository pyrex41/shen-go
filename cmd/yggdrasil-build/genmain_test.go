package main

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// referenceInstalls is the boot-helper sequence cmd/shen/main.go's regist()
// runs, in order. The generated program must install the same set: it boots the
// same kernel, so a helper that matters there matters here.
var referenceInstalls = []string{
	"InstallKernelFast",
	"InstallIntegerGuard",
	"InstallExactPow10",
	"InstallShenX",
	"InstallPr",
}

func genMainForTest(needsEval bool) string {
	m := &manifest{kernel: "kernel.kl", init: "shen.initialise", users: []string{"prog.kl"}, needsEval: needsEval}
	ku := []unit{{name: "kernelChunk0", goFile: "kernel_00.go", src: []byte("(defun f () 1)")}}
	uu := []unit{{name: "userChunk0", goFile: "user_00.go", src: []byte("(defun g () 2)")}}
	return genMain(m, ku, uu, []fnArity{{name: "g", arity: 0}})
}

// TestGenMainInstallsMatchReference guards against the generated boot drifting
// from cmd/shen's. It already had: the template ran InstallExactPow10 and
// InstallShenX but omitted InstallKernelFast and InstallIntegerGuard, so shaken
// artifacts silently lost the native arity/fn fast paths.
func TestGenMainInstallsMatchReference(t *testing.T) {
	for _, needsEval := range []bool{false, true} {
		src := genMainForTest(needsEval)
		var got []string
		for _, name := range referenceInstalls {
			// Match the CALL, not the bare name. The emitted template also
			// discusses these helpers in comments, and a substring match on the
			// name was satisfied by the comment alone -- deleting the
			// InstallIntegerGuard call left this test passing. It was asymmetric
			// only by luck: InstallKernelFast's comment happens not to spell its
			// own name, and any future comment that did would silently kill that
			// half too.
			if strings.Contains(src, "runHelper(\""+name+"\"") {
				got = append(got, name)
			}
		}
		if len(got) != len(referenceInstalls) {
			t.Errorf("needsEval=%v: generated main installs %v, cmd/shen's regist() installs %v; "+
				"a helper present in the reference boot but missing here silently changes "+
				"shaken-artifact behaviour", needsEval, got, referenceInstalls)
		}
	}
}

// TestReferenceInstallsMatchCmdShen keeps the list above honest. If cmd/shen
// gains or drops a boot helper, this fails rather than letting the assertion
// above quietly test a stale set.
func TestReferenceInstallsMatchCmdShen(t *testing.T) {
	b, err := os.ReadFile("../shen/main.go")
	if err != nil {
		t.Skipf("cannot read cmd/shen/main.go: %v", err)
	}
	found := regexp.MustCompile(`kl\.(Install\w+)\(\)`).FindAllStringSubmatch(string(b), -1)
	seen := map[string]bool{}
	var names []string
	for _, m := range found {
		if !seen[m[1]] {
			seen[m[1]] = true
			names = append(names, m[1])
		}
	}
	if strings.Join(names, ",") != strings.Join(referenceInstalls, ",") {
		t.Errorf("cmd/shen/main.go installs %v but referenceInstalls says %v; "+
			"update referenceInstalls (and genMain) to match", names, referenceInstalls)
	}
}

// TestGenMainGuardsBootHelpers pins that boot helpers run under a recover that
// reports which helper failed. Called bare, a helper panic escaped as an
// unreadable raw kl.Obj pointer -- that is how the eval-free target stayed
// broken without naming the cause.
func TestGenMainGuardsBootHelpers(t *testing.T) {
	src := genMainForTest(false)
	if !strings.Contains(src, "func runHelper(") {
		t.Fatal("generated main has no runHelper: a panicking boot helper would report a raw pointer")
	}
	for _, name := range []string{"InstallExactPow10", "InstallShenX"} {
		if !strings.Contains(src, `runHelper("`+name+`"`) {
			t.Errorf("%s is called bare; a panic there is unattributable", name)
		}
	}
}

// TestGenMainInstallsNativesBeforeInit pins the boot order Yggdrasil's
// conformance gate checks: the natives go in after each kernel chunk, as
// kl.BootKernel does after each module, so the kernel's inline init, its init
// pass and any arity replay all run on them. The template used to install them
// after shen.initialise, which the gate records as
// native_overrides_installed_after and refuses.
func TestGenMainInstallsNativesBeforeInit(t *testing.T) {
	for _, needsEval := range []bool{false, true} {
		src := genMainForTest(needsEval)
		install := strings.Index(src, `runHelper("InstallKernelFast"`)
		afterChunks := strings.Index(src, `runHelper("InstallExactPow10"`)
		initCall := strings.Index(src, `PrimFunc(MakeSymbol("shen.initialise"))`)
		replay := strings.Index(src, `MakeSymbol("shen.store-arity")`)
		user := strings.Index(src, "range userChunks")
		kernel := strings.Index(src, "range kernelChunks")
		if install < 0 || initCall < 0 || user < 0 || kernel < 0 {
			t.Fatalf("needsEval=%v: generated main lacks an expected step:\n%s", needsEval, src)
		}
		// The install sits inside the kernel-chunk loop (before the first
		// post-loop helper), so each chunk's definitions are replaced as they
		// appear and the kernel's inline init runs on the natives.
		if !(kernel < install && install < afterChunks && afterChunks < initCall && initCall < user) {
			t.Errorf("needsEval=%v: want kernel chunk loop containing InstallKernelFast < post-loop helpers < init < user chunks; "+
				"offsets kernel=%d install=%d post-loop=%d init=%d user=%d", needsEval, kernel, install, afterChunks, initCall, user)
		}
		if needsEval && !(install < replay) {
			t.Errorf("InstallKernelFast (%d) must precede the arity replay (%d)", install, replay)
		}
	}
}
