package kl

// Booting the KLambda kernel from source, out of kernel/klambda.
// cmd/yggdrasil-build and the equivalence harness (equiv.go and
// cmd/kl equiv-check) boot this way. cmd/shen boots its precompiled modules the
// same way; see regist in cmd/shen/main.go.

import (
	"fmt"
	"os"
	"path/filepath"
)

// KernelLoadOrder is the load order of the KLambda kernel under
// kernel/klambda, mirrored from compiled/script.kl, compiled/precompile.kl
// and cmd/shen's regist(). ShenOSKernel 42.2 (community lineage) modules are
// plain lists of (defun ...) forms, so the order is not significant; all
// top-level state is set up afterwards by KernelInitForm.
var KernelLoadOrder = []string{
	"sys.kl", "writer.kl", "core.kl", "reader.kl", "declarations.kl",
	"toplevel.kl", "macros.kl", "load.kl", "prolog.kl", "sequent.kl",
	"track.kl", "t-star.kl", "yacc.kl", "types.kl", "dict.kl", "init.kl",
	"extension-launcher.kl",
}

// KernelInitForm is the form that initialises a kernel loaded from
// KernelLoadOrder: init.kl's shen.initialise builds *property-vector*, the
// arity table, the lambda forms and the datatype declarations.
const KernelInitForm = "(shen.initialise)"

// BootKernel loads the KLambda kernel from dir, a kernel/klambda directory,
// into the process the same way cmd/shen boots its compiled kernel:
//
//   - Every module in KernelLoadOrder is loaded (defuns only).
//   - The native FNV-1a `hash` replaces sys.kl's interpreted one, before
//     shen.initialise builds *property-vector*, so that every dict is
//     written and read with one hash.
//   - InstallKernelFast rebinds the hot kernel functions (get/put/unput on
//     the property dict, arity, fn, ...) to natives. It runs before
//     shen.initialise so the arity table and lambda forms are built by the
//     same get/put that later read them, and so the interpreted put never runs
//     during initialisation (issue #46).
//   - shen.initialise (KernelInitForm) runs last.
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
	for _, f := range KernelLoadOrder {
		p := filepath.Join(dir, f)
		if _, err := os.Stat(p); err != nil {
			return fmt.Errorf("boot kernel: %v", err)
		}
		form := Cons(MakeSymbol("load-file"), Cons(MakeString(p), Nil))
		if res := Eval(e, form); IsError(res) {
			return fmt.Errorf("boot kernel: load-file %s: %s", p, GetString(PrimErrorToString(res)))
		}
	}
	BindSymbolFunc(MakeSymbol("hash"), MakePrimitive("hash", 2, PrimHash))
	InstallKernelFast()
	init, err := ReadForm(KernelInitForm)
	if err != nil {
		return err
	}
	if res := Eval(e, init); IsError(res) {
		return fmt.Errorf("boot kernel: %s: %s", KernelInitForm, GetString(PrimErrorToString(res)))
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
