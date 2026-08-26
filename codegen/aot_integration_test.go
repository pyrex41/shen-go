package codegen

import (
	"bytes"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

func irList(xs ...kl.Obj) kl.Obj {
	out := kl.Nil
	for i := len(xs) - 1; i >= 0; i-- {
		out = kl.Cons(xs[i], out)
	}
	return out
}

// TestAOTGeneratesLegacyAndTypedIR exercises both the historical bytecode
// shape and the compiler's $type metadata wrapper. Metadata must be erased by
// code generation, preserving the same executable expression.
func TestAOTGeneratesLegacyAndTypedIR(t *testing.T) {
	sym := kl.MakeSymbol
	legacy := irList(sym("return"), irList(sym("$const"), kl.MakeNumber(1.5)))
	typed := irList(sym("return"), irList(sym("$type"), sym("number"), irList(sym("$const"), kl.MakeNumber(1.5))))
	for _, tc := range []struct {
		name string
		bc   kl.Obj
	}{
		{"legacy", legacy},
		{"typed", typed},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got bytes.Buffer
			cg := New()
			if err := cg.HandleBodyObj(tc.bc, "Generated", &got); err != nil {
				t.Fatalf("HandleBodyObj: %v", err)
			}
			if !strings.Contains(got.String(), "MakeNumber(1.5)") {
				t.Fatalf("generated source lost fractional constant:\n%s", got.String())
			}
			if _, err := parser.ParseFile(token.NewFileSet(), "generated.go", got.Bytes(), 0); err != nil {
				t.Fatalf("generated source is not valid Go: %v\n%s", err, got.String())
			}
		})
	}
}

// TestAOTGeneratedSourceCompiles writes a generated module into a temporary
// package and asks the Go toolchain to compile it against the current kl
// runtime. This catches import, symbol, and Go-version compatibility errors
// that parser-only tests cannot detect.
func TestAOTGeneratedSourceCompiles(t *testing.T) {
	bc := irList(kl.MakeSymbol("return"), irList(kl.MakeSymbol("$const"), kl.MakeNumber(42)))
	var src bytes.Buffer
	cg := New()
	if err := cg.HandleBodyObj(bc, "Generated", &src); err != nil {
		t.Fatal(err)
	}
	// Keep the temporary package below the repository so the module's go.mod
	// supplies the github.com/pyrex41/shen-go/kl import during `go test`.
	dir, err := os.MkdirTemp(".", "_aotcompile-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	if err := os.WriteFile(filepath.Join(dir, "generated.go"), src.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated Go failed to compile: %v\n%s\nsource:\n%s", err, out, src.String())
	}
}

// TestAOTTypedIRModeSwitchStillGenerates checks both supported deployment
// modes. The mode is consumed by the runtime/compiler; AOT generation must
// remain valid in either mode and must not depend on process-global state.
func TestAOTTypedIRModeSwitchStillGenerates(t *testing.T) {
	bc := irList(kl.MakeSymbol("return"), irList(kl.MakeSymbol("$const"), kl.MakeNumber(7)))
	old, had := os.LookupEnv("SHEN_GO_TYPED_IR")
	defer func() {
		if had {
			_ = os.Setenv("SHEN_GO_TYPED_IR", old)
		} else {
			_ = os.Unsetenv("SHEN_GO_TYPED_IR")
		}
	}()
	for _, mode := range []string{"on", "off"} {
		t.Run(mode, func(t *testing.T) {
			if err := os.Setenv("SHEN_GO_TYPED_IR", mode); err != nil {
				t.Fatal(err)
			}
			var src bytes.Buffer
			if err := New().HandleBodyObj(bc, "Generated", &src); err != nil {
				t.Fatalf("mode %s: %v", mode, err)
			}
			if !strings.Contains(src.String(), "var Generated = MakeNative") {
				t.Fatalf("mode %s emitted no module thunk", mode)
			}
		})
	}
}

func TestAOTGuardedTailApplyGeneratesValidGo(t *testing.T) {
	sym := kl.MakeSymbol
	for _, name := range []string{"+", "hd"} {
		args := []kl.Obj{irList(sym("$const"), kl.MakeNumber(1))}
		if name == "+" {
			args = append(args, irList(sym("$const"), kl.MakeNumber(2)))
		}
		body := irList(append([]kl.Obj{sym("tailapply"), irList(sym("$global"), sym(name))}, args...)...)
		var src bytes.Buffer
		if err := New().HandleBodyObj(body, "Generated", &src); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if _, err := parser.ParseFile(token.NewFileSet(), "generated.go", src.Bytes(), 0); err != nil {
			t.Fatalf("%s generated invalid tail call: %v\n%s", name, err, src.String())
		}
	}
}

// TestAOTGeneratedTypedExecution compiles a generated module together with a
// small Go test harness and executes the thunk.  Keeping the harness in the
// temporary package exercises the same ABI used by native plugins while also
// checking that typed and dynamic modes produce identical Shen values.
func TestAOTGeneratedTypedExecution(t *testing.T) {
	sym := kl.MakeSymbol
	global := func(name string) kl.Obj {
		return irList(sym("$global"), sym(name))
	}
	call := func(name string, args ...kl.Obj) kl.Obj {
		return irList(append([]kl.Obj{sym("call"), global(name)}, args...)...)
	}
	ret := func(x kl.Obj) kl.Obj { return irList(sym("return"), x) }

	cases := []struct {
		name  string
		body  kl.Obj
		check string
	}{
		{"arithmetic", ret(call("+", irList(sym("$const"), kl.MakeNumber(1.25)), irList(sym("$const"), kl.MakeNumber(2.5)))), `if n, ok := kl.TypedFloat64(got); !ok || n != 3.75 { t.Fatalf("got %s", kl.ObjString(got)) }`},
		{"flattened-chain", irList(sym("block"),
			irList(sym("<="), sym("x"), call("+", irList(sym("$const"), kl.MakeNumber(1)), irList(sym("$const"), kl.MakeNumber(2)))),
			irList(sym("<="), sym("y"), call("*", sym("x"), irList(sym("$const"), kl.MakeNumber(4)))),
			ret(sym("y"))), `if n, ok := kl.TypedFloat64(got); !ok || n != 12 { t.Fatalf("got %s", kl.ObjString(got)) }`},
		{"boolean", ret(call("not", irList(sym("$const"), kl.True))), `if got != kl.False { t.Fatalf("got %s", kl.ObjString(got)) }`},
		{"unicode", ret(call("cn", irList(sym("$const"), kl.MakeString("λ")), irList(sym("$const"), kl.MakeString("雪")))), `if !kl.IsString(got) || kl.GetString(got) != "λ雪" { t.Fatalf("got %s", kl.ObjString(got)) }`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			compileAndRunGenerated(t, tc.body, `
				for _, mode := range []string{"on", "off"} {
					if err := os.Setenv("SHEN_GO_TYPED_IR", mode); err != nil { t.Fatal(err) }
					kl.ResetTypedIRModeForTest()
					got := kl.Call(&kl.ControlFlow{}, Generated)
					`+tc.check+`
				}
			`)
		})
	}
}

// TestAOTGeneratedDivisionZeroRemainsCatchable verifies that a scalar region
// does not turn Shen's catchable division error into an uncaught Go panic.
func TestAOTGeneratedDivisionZeroRemainsCatchable(t *testing.T) {
	sym := kl.MakeSymbol
	global := irList(sym("$global"), sym("+"))
	div := irList(sym("call"), irList(sym("$global"), sym("/")),
		irList(sym("$const"), kl.MakeNumber(4)), irList(sym("$const"), kl.MakeNumber(0)))
	body := irList(sym("return"), irList(sym("call"), global,
		irList(sym("$const"), kl.MakeNumber(1)), div))
	compileAndRunGenerated(t, body, `
		os.Setenv("SHEN_GO_TYPED_IR", "on")
		kl.ResetTypedIRModeForTest()
		identity := kl.MakeNative(func(e *kl.ControlFlow) { e.Return(e.Get(1)) }, 1)
		got := kl.Try(&kl.ControlFlow{}, Generated).Catch(identity)
		if !kl.IsError(got) { t.Fatalf("division by zero returned %s", kl.ObjString(got)) }
	`)
}

// TestAOTGeneratedRebindingUsesDynamicPrimitive exercises the canonical
// binding guard.  A primitive rebound after generation must be called through
// the normal Obj path, even when typed mode is enabled.
func TestAOTGeneratedRebindingUsesDynamicPrimitive(t *testing.T) {
	sym := kl.MakeSymbol
	body := irList(sym("return"), irList(sym("call"), irList(sym("$global"), sym("+")),
		irList(sym("$const"), kl.MakeNumber(2)), irList(sym("$const"), kl.MakeNumber(3))))
	compileAndRunGenerated(t, body, `
		plus := kl.MakeSymbol("+")
		orig := kl.PrimFunc(plus)
		defer kl.BindSymbolFunc(plus, orig)
		kl.BindSymbolFunc(plus, kl.MakePrimitive("shadow-plus", 2, func(_, _ kl.Obj) kl.Obj { return kl.MakeNumber(99) }))
		os.Setenv("SHEN_GO_TYPED_IR", "on")
		kl.ResetTypedIRModeForTest()
		got := kl.Call(&kl.ControlFlow{}, Generated)
		if n, ok := kl.TypedFloat64(got); !ok || n != 99 { t.Fatalf("rebound primitive got %s", kl.ObjString(got)) }
	`)
}

func compileAndRunGenerated(t *testing.T, body kl.Obj, harness string) {
	t.Helper()
	// Keep the package below the module root so the generated import resolves
	// through this repository's go.mod.
	dir, err := os.MkdirTemp(".", "_aotexec-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	var src bytes.Buffer
	cg := New()
	if err := cg.HandleBodyObj(body, "Generated", &src); err != nil {
		t.Fatal(err)
	}
	cg.HandleSymbol(&src)
	if err := os.WriteFile(filepath.Join(dir, "generated.go"), src.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	runner := "package main\n\nimport (\n\t\"os\"\n\t\"testing\"\n\n\tkl \"github.com/pyrex41/shen-go/kl\"\n)\n\nfunc TestGeneratedExecution(t *testing.T) {\n" + harness + "\n}\n"
	if err := os.WriteFile(filepath.Join(dir, "generated_test.go"), []byte(runner), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated Go execution failed: %v\n%s\nsource:\n%s", err, out, src.String())
	}
}
