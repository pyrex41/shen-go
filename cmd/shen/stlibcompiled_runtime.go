package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pyrex41/shen-go/cmd/shen/stlibcompiled"
	"github.com/pyrex41/shen-go/kl"
)

// loadStdlibCompiled installs the source-bound StLib artifact. The reader's
// preprocessing effects are replayed from the generated snapshot: arities
// (including lambdatable entries) and package external/internal lists. Type
// declarations and function/macro definitions are in StlibMain itself.
func loadStdlibCompiled(e *kl.ControlFlow) {
	kl.BindSymbolFunc(kl.MakeSymbol("pr"), kl.MakeNative(func(e *kl.ControlFlow) {
		e.Return(e.Get(1))
	}, 2))
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
			fmt.Fprintf(os.Stderr, "shen: compiled stdlib %s: %s\n", what, kl.GetString(kl.PrimErrorToString(res)))
			return true
		}
		return false
	}
	if fail("tc-off", evalStr("(tc -)")) {
		return
	}
	storeArity := kl.PrimFunc(kl.MakeSymbol("shen.store-arity"))
	for _, line := range strings.Split(stlibcompiled.Arities, "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "shen: compiled stdlib bad arity: %q\n", line)
			return
		}
		if fail("arity", kl.Call(e, storeArity, kl.MakeSymbol(fields[0]), kl.MakeInteger(n))) {
			return
		}
	}
	if fail("install", stlibcompiled.Install(e)) {
		return
	}
	pv := kl.PrimValue(kl.MakeSymbol("*property-vector*"))
	put := kl.PrimFunc(kl.MakeSymbol("put"))
	list := func(names []string) kl.Obj {
		out := kl.Nil
		for i := len(names) - 1; i >= 0; i-- {
			out = kl.Cons(kl.MakeSymbol(names[i]), out)
		}
		return out
	}
	for _, p := range stlibcompiled.Packages {
		name := kl.MakeSymbol(p.Name)
		for _, attr := range []struct {
			name   string
			values []string
		}{{"shen.external-symbols", p.External}, {"shen.internal-symbols", p.Internal}} {
			if fail("package", kl.Call(e, put, name, kl.MakeSymbol(attr.name), list(attr.values), pv)) {
				return
			}
		}
	}
	// Bootstrap uses defmacro while lowering StLib, so those registrations
	// are absent from the emitted KL. Replay the six small source forms after
	// their helper definitions and package metadata are installed.
	macroForms, err := kl.ShenReadSExprs([]byte(stlibcompiled.Macros))
	if err != nil {
		fmt.Fprintf(os.Stderr, "shen: compiled stdlib macros: %v\n", err)
		return
	}
	if fail("macros", kl.Call(e, kl.PrimFunc(kl.MakeSymbol("shen.process-sexprs")), macroForms)) {
		return
	}
	fail("systemf", evalStr("(let External (external stlib) ExternalF (filter (/. X (> (arity X) -1)) External) Systemf (map (fn systemf) ExternalF) ok)"))
}
