// Test fixture for the auto-engage-native path (cmd/shen/autonative.go).
//
// It is a hand-written stand-in for what `make precompile` emits: a plugin whose
// exported Install hook rebinds a Shen-level function. Hand-written on purpose —
// the real toolchain (Shen -> KL -> IR -> Go -> go build -buildmode=plugin) takes
// tens of seconds per file, which is too slow to sit inside a Go test, and the
// behaviour under test is the *loader's* (does a (load ...) of the plugin's
// source re-engage it?), not the code generator's.
//
// `probe` returns a marker string that says which implementation answered, so a
// test can tell whether the compiled version or the freshly-loaded VM version is
// in force. See cmd/shen/autonative_test.go.
//
// Build: go build -buildmode=plugin -o /tmp/probe.so ./cmd/plugintest/probeplugin
package main

import . "github.com/pyrex41/shen-go/kl"

var symprobe = MakeSymbol("probe")

// probeFn is the "compiled" implementation: (probe) -> "native".
var probeFn = MakeNative(func(e *ControlFlow) {
	e.Return(MakeString("native"))
}, 0)

// Install is the plugin contract cmd/shen's loader looks up: a
// func(*ControlFlow) that binds this file's functions, exactly as the generated
// installHookTmpl in cmd/kl/plugin.go does.
func Install(e *ControlFlow) {
	BindSymbolFunc(symprobe, probeFn)
}

// main is never called: this package exists only to be built with
// -buildmode=plugin. It is declared so a plain `go build ./...` links this main
// package instead of failing with "function main is undeclared".
func main() {}
