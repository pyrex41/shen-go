package kl

// Booting the KLambda kernel from source (kernel/klambda). cmd/yggdrasil-build
// and the equivalence harness (equiv.go, cmd/kl equiv-check) boot this way;
// cmd/shen boots its precompiled modules in the same order (cmd/shen/main.go
// regist).

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

// BootKernel loads the KLambda kernel from dir (a kernel/klambda directory)
// into the process the same way cmd/shen boots its compiled kernel:
//
//   - sys.kl defines the interpreted `hash`; the native FNV-1a hash is swapped
//     in right after, before declarations.kl builds the first property
//     dictionary, so every dictionary is written and read with one hash;
//   - InstallKernelFast runs after every module. In cmd/shen it runs once, after
//     all modules, because the compiled modules never execute the interpreted
//     put/get. Here the modules are interpreted, and since 5edf47e the
//     interpreted `put` (a tail call inside trap-error) escapes the
//     interpreter's recover during declarations.kl (issue #46), so the natives
//     have to be in place as soon as sys.kl has defined the names they replace.
//     InstallKernelFast only rebinds names the kernel has defined, so calling
//     it per module is the way to catch each module's definitions;
//   - InstallIntegerGuard runs after the natives, as in cmd/shen; with the
//     canonical integer? restored by InstallKernelFast it is a no-op, and is
//     called so the two boots stay step-for-step the same;
//   - InstallPr binds shen.native-pr, as cmd/shen's fixPrHush does.
//
// cmd/shen additionally wraps read-file and registers load-native's arity;
// those are REPL/plugin concerns and are not part of booting the kernel.
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
