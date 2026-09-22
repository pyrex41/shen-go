package main

import (
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestBuilderBootsKernel pins issue #46: since 5edf47e the builder's boot
// panicked inside declarations.kl ("implementation error in
// shen.change-pointer-value") because it loaded the kernel on the interpreted
// put/get without ever calling kl.InstallKernelFast in its own process. The
// boot now goes through kl.BootKernel, which installs the natives as the
// modules load; this test boots the full compiler image the way main does
// and checks the pieces the build step needs are there.
func TestBuilderBootsKernel(t *testing.T) {
	root, err := findShenGoRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	var e kl.ControlFlow
	if err := bootCompilerImage(&e, root); err != nil {
		t.Fatalf("boot: %v", err)
	}
	for _, name := range []string{"put", "get", "arity", "fn", "reverse", "map", "shen.native-pr"} {
		f := kl.PrimFunc(kl.MakeSymbol(name))
		if !kl.IsNativeBinding(f) {
			t.Errorf("%s is not bound to a native after boot; InstallKernelFast did not run in the builder", name)
		}
	}
	for _, tc := range []struct{ src, want string }{
		{"(arity compile-file)", "2"},
		{"(get compile-file arity (value *property-vector*))", "2"},
		{"(reverse (cons 1 (cons 2 (cons 3 ()))))", "(3 2 1)"},
		{"(value *maximum-print-sequence-size*)", "100000"},
	} {
		v, err := evalKLValue(&e, tc.src)
		if err != nil {
			t.Fatalf("%s: %v", tc.src, err)
		}
		if got := kl.ObjString(v); got != tc.want {
			t.Errorf("%s = %s, want %s", tc.src, got, tc.want)
		}
	}
}
