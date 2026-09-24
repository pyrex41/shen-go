package kl

// Booting the KLambda kernel from source, out of kernel/klambda.
// cmd/yggdrasil-build and the equivalence harness (equiv.go and
// cmd/kl equiv-check) boot this way. cmd/shen boots its precompiled modules in
// the same order; see regist in cmd/shen/main.go.

import (
	"fmt"
	"os"
	"path/filepath"
)

// KernelLoadOrder is the canonical load order of the KLambda kernel under
// kernel/klambda, mirrored from compiled/precompile.kl and cmd/shen's regist().
// The Tarver S42 kernel has no shen.initialise: each module runs its own
// top-level init forms as it loads, so this order must not change.
var KernelLoadOrder = []string{
	"sys.kl", "writer.kl", "core.kl", "reader.kl", "declarations.kl",
	"toplevel.kl", "macros.kl", "load.kl", "prolog.kl", "sequent.kl",
	"track.kl", "t-star.kl", "yacc.kl", "types.kl",
	"extension-launcher.kl",
}

// BootKernel loads the KLambda kernel from dir, a kernel/klambda directory,
// into the process the same way cmd/shen boots its compiled kernel:
//
//   - sys.kl defines the interpreted `hash`. The native FNV-1a hash is swapped
//     in right after, before declarations.kl builds the first property vector,
//     so that every property vector is written and read with one hash.
//   - InstallKernelFast runs after every module. In cmd/shen it runs once,
//     after all the modules, because the compiled modules never execute the
//     interpreted put/get. Here the modules are interpreted, and since 5edf47e
//     the interpreted `put` escapes the interpreter's recover during
//     declarations.kl, because of a tail call inside trap-error (issue #46).
//     So the natives have to be in place as soon as sys.kl has defined the
//     names they replace. InstallKernelFast is unconditional (issue #49); it
//     is called per module because each module's own (defun ...) of a rebound
//     name (fn in reader.kl, arity in declarations.kl, shen.pvar? in
//     prolog.kl, remove in track.kl, ...) overwrites the native until the
//     next call restores it, and the interpreted put/get must not be live
//     during declarations.kl (issue #46).
//   - InstallIntegerGuard runs after the natives, as in cmd/shen. With the
//     canonical integer? restored by InstallKernelFast it is a no-op. It is
//     called so that the two boots stay step for step the same.
//   - InstallPr binds shen.native-pr, which is the first thing cmd/shen's
//     fixPrHush does.
//
// cmd/shen does more than this. It also installs the exact base-10 power and
// the shen.x extensions (InstallExactPow10 and InstallShenX), redefines `pr`
// on top of shen.native-pr, wraps read-file, and registers load-native's
// arity. Those are reader-precision, REPL and plugin concerns, not part of
// booting the kernel.
func BootKernel(e *ControlFlow, dir string) error {
	for i, f := range KernelLoadOrder {
		p := filepath.Join(dir, f)
		if _, err := os.Stat(p); err != nil {
			return fmt.Errorf("boot kernel: %v", err)
		}
		form := Cons(MakeSymbol("load-file"), Cons(MakeString(p), Nil))
		if res := Eval(e, form); IsError(res) {
			return fmt.Errorf("boot kernel: load-file %s: %s", p, GetString(PrimErrorToString(res)))
		}
		if i == 0 {
			BindSymbolFunc(MakeSymbol("hash"), MakePrimitive("hash", 2, PrimHash))
		}
		InstallKernelFast()
	}
	InstallIntegerGuard()
	InstallPr()
	return nil
}

// FindKernelDir walks up from start looking for kernel/klambda/sys.kl and
// returns that kernel/klambda directory.
func FindKernelDir(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		cand := filepath.Join(dir, "kernel", "klambda")
		if _, err := os.Stat(filepath.Join(cand, "sys.kl")); err == nil {
			return cand, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("kernel/klambda not found above %s", start)
		}
		dir = parent
	}
}
