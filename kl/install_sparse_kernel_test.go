package kl

import "testing"

// TestInstallHelpersToleratesSparseKernel pins the contract that every boot
// helper must survive a kernel that does not define the symbols it would like
// to optimise.
//
// This is not hypothetical. An eval-free Yggdrasil shake drops the reader, and
// with it shen.expt. InstallExactPow10 looked that symbol up with PrimFunc,
// which panics on an unbound symbol, so EVERY eval-free shaken artifact died at
// boot -- and because the generated main called the helper outside its own
// error reporter, the failure surfaced as a bare "panic: (kl.Obj) 0x...". The
// Go target's own reference fixture (fib) could not run at all, and nothing
// reported it because cmd/yggdrasil-build had no tests.
//
// A fresh symbol table has none of these bindings, so it is exactly the sparse
// kernel this must tolerate: any helper that panics here is broken for shaken
// artifacts.
func TestInstallHelpersToleratesSparseKernel(t *testing.T) {
	for _, tc := range []struct {
		name string
		fn   func()
	}{
		{"InstallExactPow10", InstallExactPow10},
		{"InstallIntegerGuard", InstallIntegerGuard},
		{"InstallKernelFast", InstallKernelFast},
		{"InstallShenX", InstallShenX},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("%s panicked on a kernel lacking its symbols: %v\n"+
						"Boot helpers must degrade to a no-op, as InstallIntegerGuard does; "+
						"a panic here breaks every eval-free shaken artifact at boot.", tc.name, r)
				}
			}()
			tc.fn()
		})
	}
}
