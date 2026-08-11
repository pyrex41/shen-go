package main

import (
	"strings"
	"testing"
)

// TestLoadNativeIsCallableFromShen pins the arity registration for load-native.
//
// load-native is bound as a KL function at startup, but the Shen reader resolves
// an application through the symbol's *registered arity* (shen.record-and-evaluate
// in core.kl), not through the KL binding. Without a shen.store-arity call the
// symbol is bound and yet reported "fn: load-native is undefined" the moment it
// is called from Shen source — so the documented `(load-native "hot.so")` form,
// and with it the only way to re-install a precompiled plugin *after* a
// (load "file.shen") has rebound those functions onto the VM, did not exist.
//
// This asserts the symbol resolves. It deliberately calls it with a path that
// does not exist: the point is that we get load-native's own error (a failed
// plugin.Open), NOT "undefined" from the reader. Building a real .so needs the
// full precompile toolchain and is out of scope for a unit test.
func TestLoadNativeIsCallableFromShen(t *testing.T) {
	out, _ := runCLI(t, "eval", "-e", `(load-native "/nonexistent/definitely-not-here.so")`)

	if strings.Contains(out, "undefined") {
		t.Fatalf("load-native did not resolve from Shen source (arity not registered):\n%s", out)
	}
	// It must actually have reached loadNative and failed to open the .so.
	if !strings.Contains(out, "plugin.Open") {
		t.Fatalf("expected a plugin.Open failure from load-native, got:\n%s", out)
	}
}
