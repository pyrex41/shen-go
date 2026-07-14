package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/tiancaiamao/shen-go/kl"
)

// The Shen standard library, vendored from Tarver's S41.2 refresh Lib/StLib
// (see stlib/PROVENANCE.md). The kernel itself ships none of these — filter,
// mapc, take/drop, fold*, the maths/strings/vectors/tuples helpers, etc. all
// live here. We embed the sources and load them at startup so a bare
// `(filter ...)` works out of the box, keeping the binary self-contained.
//
//go:embed stlib
var stlibFS embed.FS

// stlibInstallOrder mirrors Lib/StLib/install.shen. Order is load-bearing:
// each Maths/Strings/Vectors macros file and every .dtype must precede the
// sources that use them, and package-stlib.shen (which unions the sub-package
// externals into the `stlib` package) comes last.
var stlibInstallOrder = []string{
	"stlib/Symbols/symbols1.shen",
	"stlib/Symbols/symbols2.shen",
	"stlib/Maths/macros.shen",
	"stlib/Maths/maths.shen",
	"stlib/Maths/rationals.dtype",
	"stlib/Maths/rationals.shen",
	"stlib/Maths/complex.dtype",
	"stlib/Maths/complex.shen",
	"stlib/Maths/numerals.dtype",
	"stlib/Maths/numerals.shen",
	"stlib/Lists/lists.shen",
	"stlib/Strings/macros.shen",
	"stlib/Strings/strings.shen",
	"stlib/Strings/smart.shen",
	"stlib/Vectors/macros.shen",
	"stlib/IO/prettyprint.shen",
	"stlib/IO/delete-file.shen",
	"stlib/IO/files.shen",
	"stlib/Tuples/tuples.shen",
	"stlib/package-stlib.shen",
}

// loadStdlib loads the embedded standard library into the running image. The
// sources are concatenated (their top-level forms just run in sequence — no
// cross-file loads) into one temp file and loaded with type-checking OFF: the
// library is known-good, and skipping the check keeps startup fast (~0.12s).
// Loader chatter (one "(fn X)" per defun, plus the "run time:" line) is
// silenced by rebinding `pr` to a no-op for the duration, then restoring the
// real pr via fixPrHush. Errors are non-fatal — a broken stdlib must not stop
// the REPL from coming up — but are reported to stderr.
func loadStdlib(e *kl.ControlFlow) {
	var b strings.Builder
	for _, f := range stlibInstallOrder {
		data, err := stlibFS.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "shen: stdlib embed read %s: %v\n", f, err)
			return
		}
		b.Write(data)
		b.WriteByte('\n')
	}

	tmp, err := os.CreateTemp("", "shen-stdlib-*.shen")
	if err != nil {
		fmt.Fprintf(os.Stderr, "shen: stdlib temp: %v\n", err)
		return
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.WriteString(b.String()); err != nil {
		tmp.Close()
		fmt.Fprintf(os.Stderr, "shen: stdlib write: %v\n", err)
		return
	}
	tmp.Close()

	// Silence loader output: pr -> identity (returns its first arg unwritten).
	kl.BindSymbolFunc(kl.MakeSymbol("pr"), kl.MakeNative(func(e *kl.ControlFlow) {
		e.Return(e.Get(1))
	}, 2))
	// Whatever happens, put the real pr back.
	defer fixPrHush(e)

	evalStr := func(src string) kl.Obj {
		exp, err := kl.NewSexpReader(strings.NewReader(src), false).Read()
		if err != nil {
			return kl.MakeError(err.Error())
		}
		return kl.Eval(e, exp)
	}
	fail := func(what string, res kl.Obj) bool {
		if kl.IsError(res) {
			fmt.Fprintf(os.Stderr, "shen: stdlib %s: %s\n", what, kl.GetString(kl.PrimErrorToString(res)))
			return true
		}
		return false
	}

	if fail("tc-off", evalStr("(tc -)")) {
		return
	}
	loadForm := kl.Cons(kl.MakeSymbol("load"), kl.Cons(kl.MakeString(tmp.Name()), kl.Nil))
	if fail("load", kl.Eval(e, loadForm)) {
		return
	}
	// Declare every stlib external as a system function (mirrors the tail of
	// install.shen), then leave type-checking off for the REPL — the kernel's
	// default startup state.
	fail("systemf", evalStr("(let External (external stlib) ExternalF (filter (/. X (> (arity X) -1)) External) Systemf (map (fn systemf) ExternalF) ok)"))
}
