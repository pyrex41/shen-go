package main

import (
	"bytes"
	"os"
	"path/filepath"
	"plugin"
	"runtime"
	"testing"

	"github.com/pyrex41/shen-go/codegen"
	"github.com/pyrex41/shen-go/kl"
)

func TestGeneratedPluginInstallContract(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a Go plugin")
	}
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Skipf("Go plugins unsupported on %s", runtime.GOOS)
	}
	dir := t.TempDir()
	rawPath := filepath.Join(dir, "generated.go")
	soPath := filepath.Join(dir, "generated.so")
	bc := kl.Cons(kl.MakeSymbol("return"), kl.Cons(kl.Cons(kl.MakeSymbol("$const"), kl.Cons(kl.MakeNumber(9), kl.Nil)), kl.Nil))
	var raw bytes.Buffer
	if err := codegen.New().HandleBodyObj(bc, "Generated", &raw); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(rawPath, raw.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := buildPlugin(rawPath, "Generated", soPath); err != nil {
		t.Fatalf("buildPlugin: %v", err)
	}
	p, err := plugin.Open(soPath)
	if err != nil {
		t.Fatalf("plugin.Open: %v", err)
	}
	sym, err := p.Lookup("Install")
	if err != nil {
		t.Fatalf("Install lookup: %v", err)
	}
	install, ok := sym.(func(*kl.ControlFlow))
	if !ok {
		t.Fatalf("Install has type %T, want func(*kl.ControlFlow)", sym)
	}
	var ctl kl.ControlFlow
	install(&ctl)
}
