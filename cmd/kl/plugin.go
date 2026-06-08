package main

// Spike: the runtime half of the compile-to-Go path. Given a raw .go file
// emitted by bc->go (which defines `var <ExportName> = MakeNative(...)` plus the
// symXXX decls), this native:
//   1. appends a self-contained Install hook (+ the runtime-bound globals bc->go
//      can't know about: ns2_1set / try_1catch),
//   2. shells out to `go build -buildmode=plugin`,
//   3. plugin.Open + Lookup("Install") + runs it (which executes the generated
//      module thunk, binding the function via defun).
//
// Bound as (go-build-and-load RawGoPath ExportName) in main.go. Returns () on
// success or an error. Prints phase timings (build / load) to stderr.

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"time"

	"github.com/tiancaiamao/shen-go/kl"
)

// installHook is appended to bc->go's output. The dot-import in the generated
// file means ControlFlow/PrimFunc/etc. need no kl. prefix. Unused package-level
// vars are legal in Go, so declaring ns2_1set/try_1catch unconditionally is fine.
const installHookTmpl = `

var ns2_1set Obj
var try_1catch Obj

func Install(e *ControlFlow) {
	ns2_1set = PrimFunc(MakeSymbol("defun"))
	try_1catch = PrimFunc(MakeSymbol("try-catch"))
	Call(e, %s)
}
`

var goBuildAndLoad = kl.MakeNative(func(e *kl.ControlFlow) {
	rawPath := kl.GetString(e.Get(1))
	exportName := kl.GetString(e.Get(2))

	raw, err := os.ReadFile(rawPath)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}

	// Scratch package dir inside the module (so the kl import resolves).
	scratch, err := os.MkdirTemp(".", "plugintmp")
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	src := append(raw, []byte(fmt.Sprintf(installHookTmpl, exportName))...)
	srcPath := filepath.Join(scratch, "plugin.go")
	if err := os.WriteFile(srcPath, src, 0644); err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}

	soPath := filepath.Join(scratch, "fn.so")
	t0 := time.Now()
	build := exec.Command("go", "build", "-buildmode=plugin", "-o", soPath, "./"+scratch)
	if out, err := build.CombinedOutput(); err != nil {
		e.Return(kl.MakeError(fmt.Sprintf("go build failed: %v\n%s", err, out)))
		return
	}
	buildMS := time.Since(t0).Milliseconds()

	t1 := time.Now()
	p, err := plugin.Open(soPath)
	if err != nil {
		e.Return(kl.MakeError("plugin.Open: " + err.Error()))
		return
	}
	sym, err := p.Lookup("Install")
	if err != nil {
		e.Return(kl.MakeError("Lookup(Install): " + err.Error()))
		return
	}
	install, ok := sym.(func(*kl.ControlFlow))
	if !ok {
		e.Return(kl.MakeError("Install has unexpected type"))
		return
	}
	install(e)
	loadMS := time.Since(t1).Milliseconds()

	fmt.Fprintf(os.Stderr, "[go-build-and-load] build=%dms load=%dms so=%s\n", buildMS, loadMS, soPath)
	e.Return(kl.Nil)
}, 2)
