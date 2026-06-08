// Stage 0 host: loads the hand-written tak plugin, binds it, and calls it
// through the trampoline. Proves a plugin-provided Obj is callable by the host.
//
// Build & run:
//   go build -buildmode=plugin -o /tmp/tak.so ./cmd/plugintest/plugin
//   go build -o /tmp/ptest-host ./cmd/plugintest/host
//   /tmp/ptest-host /tmp/tak.so
package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/tiancaiamao/shen-go/kl"
)

func main() {
	soPath := "/tmp/tak.so"
	if len(os.Args) > 1 {
		soPath = os.Args[1]
	}

	p, err := plugin.Open(soPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "plugin.Open failed:", err)
		os.Exit(1)
	}

	compiledSym, err := p.Lookup("Compiled")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Lookup(Compiled) failed:", err)
		os.Exit(1)
	}
	nameSym, err := p.Lookup("SymName")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Lookup(SymName) failed:", err)
		os.Exit(1)
	}

	obj := *compiledSym.(*kl.Obj)
	name := *nameSym.(*string)

	kl.BindSymbolFunc(kl.MakeSymbol(name), obj)

	var e kl.ControlFlow
	res := kl.Call(&e, kl.PrimFunc(kl.MakeSymbol(name)),
		kl.MakeInteger(18), kl.MakeInteger(12), kl.MakeInteger(6))

	fmt.Printf("(%s 18 12 6) = %s\n", name, kl.ObjString(res))
}
