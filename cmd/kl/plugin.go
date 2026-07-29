package main

// Build side of the compile-to-Go path. Given a raw .go file emitted by bc->go
// (which defines `var <ExportName> = MakeNative(...)` plus the symXXX decls),
// these natives wrap it into a self-contained plugin package and `go build
// -buildmode=plugin` it.
//
//   (go-build-plugin RawGoPath ExportName OutSoPath)  -- build only (AOT).
//   (go-build-and-load RawGoPath ExportName)          -- build + load (dev/REPL).
//
// The wrapper appends an exported Install hook plus the runtime-bound globals
// bc->go can't know about (ns2_1set / try-catch). Running Install executes the
// generated module thunk, which binds every function in the file via defun.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"time"

	"github.com/pyrex41/shen-go/kl"
)

// emitArities: (emit-arities KlFile OutFile) scans a KL file for top-level
// (defun NAME (params...) body) forms and writes one "NAME ARITY" line per defun.
// The .so loader replays these via shen.store-arity so plugin-compiled functions
// resolve at the Shen REPL (the Shen reader needs registered arity; binding the
// KL function alone is not enough).
var emitArities = kl.MakeNative(func(e *kl.ControlFlow) {
	klFile := kl.GetString(e.Get(1))
	outFile := kl.GetString(e.Get(2))

	in, err := os.Open(klFile)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer in.Close()
	out, err := os.Create(outFile)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer out.Close()
	w := bufio.NewWriter(out)

	r := kl.NewSexpReader(in, false)
	for {
		form, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			e.Return(kl.MakeError(err.Error()))
			return
		}
		// (defun NAME (p1 p2 ...) body)
		head := kl.Car(form)
		if !kl.IsSymbol(head) || kl.GetSymbol(head) != "defun" {
			continue
		}
		nameObj := kl.Cadr(form)
		if !kl.IsSymbol(nameObj) {
			continue
		}
		name := kl.GetSymbol(nameObj)
		params := kl.Car(kl.Cdr(kl.Cdr(form)))
		arity := 0
		for p := params; p != kl.Nil; p = kl.Cdr(p) {
			arity++
		}
		fmt.Fprintf(w, "%s %d\n", name, arity)
	}
	w.Flush()
	e.Return(kl.Nil)
}, 2)

// installHookTmpl is appended to bc->go's output. The dot-import in the generated
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

// buildPlugin wraps a bc->go raw file into a standalone plugin package and builds
// it to outSoPath. Returns the build wall-time in ms.
func buildPlugin(rawPath, exportName, outSoPath string) (int64, error) {
	raw, err := os.ReadFile(rawPath)
	if err != nil {
		return 0, err
	}
	// Scratch package dir inside the module (so the kl import resolves). A plugin
	// is one package main, so it must be the sole .go file in its dir.
	scratch, err := os.MkdirTemp(".", "plugintmp")
	if err != nil {
		return 0, err
	}
	src := append(raw, []byte(fmt.Sprintf(installHookTmpl, exportName))...)
	if err := os.WriteFile(filepath.Join(scratch, "plugin.go"), src, 0644); err != nil {
		return 0, err
	}
	abs, err := filepath.Abs(outSoPath)
	if err != nil {
		return 0, err
	}
	t0 := time.Now()
	build := exec.Command("go", "build", "-buildmode=plugin", "-o", abs, "./"+scratch)
	if out, err := build.CombinedOutput(); err != nil {
		return 0, fmt.Errorf("go build failed: %v\n%s", err, out)
	}
	return time.Since(t0).Milliseconds(), nil
}

// goBuildPlugin: (go-build-plugin RawGoPath ExportName OutSoPath) -> OutSoPath.
// Build only; the produced .so is loaded later (e.g. by `shen -precompiled`).
var goBuildPlugin = kl.MakeNative(func(e *kl.ControlFlow) {
	rawPath := kl.GetString(e.Get(1))
	exportName := kl.GetString(e.Get(2))
	outSoPath := kl.GetString(e.Get(3))

	buildMS, err := buildPlugin(rawPath, exportName, outSoPath)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	fmt.Fprintf(os.Stderr, "[go-build-plugin] build=%dms so=%s\n", buildMS, outSoPath)
	e.Return(kl.MakeString(outSoPath))
}, 3)

// goBuildAndLoad: (go-build-and-load RawGoPath ExportName) -- build + load into
// this process (dev/REPL convenience; rebinds the functions immediately).
var goBuildAndLoad = kl.MakeNative(func(e *kl.ControlFlow) {
	rawPath := kl.GetString(e.Get(1))
	exportName := kl.GetString(e.Get(2))

	scratchSo, err := os.MkdirTemp(".", "pluginso")
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	soPath := filepath.Join(scratchSo, "fn.so")
	buildMS, err := buildPlugin(rawPath, exportName, soPath)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}

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
