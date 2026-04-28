package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestNativeReaderMatchesShen(t *testing.T) {
	e := initShenForLoadCacheTest(t)

	cases := []struct {
		name string
		src  string
	}{
		{"empty", ""},
		{"just-comments", "\\\\ a line comment\n\\* block *\\\n"},
		{"simple-list", "(foo bar baz)\n"},
		{"strings", "(foo \"hello\" \"with c#10;newline\")\n"},
		{"numbers", "(1 2 3 -4 +5 6.5 -7.25 1e3 -2.5e-2)\n"},
		{"brackets", "[1 2 3] [a | b] [] [a b | (c d)]\n"},
		{"nested", "(define f X -> (let Y [1 X] [Y | rest]))\n"},
		{"type-sig", "(define f { number --> number } X -> X)\n"},
		{"colon-eq", "(let X := 1 X)\n"},
		{"vector-empty", "(<>)\n"},
		{"misc-syms", "(@p X Y) (@s S Ss) (foo? bar! baz/qux)\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), tc.name+".shen")
			if err := os.WriteFile(path, []byte(tc.src), 0644); err != nil {
				t.Fatal(err)
			}

			native, err := parseShenFile(path)
			if err != nil {
				t.Fatalf("native reader error: %v", err)
			}

			shenForms := readShenViaKernel(t, e, path)
			if !shenObjEqual(native, shenForms) {
				t.Fatalf("native reader output differs from kernel reader\n  src:    %q\n  native: %s\n  kernel: %s", tc.src, kl.ObjString(native), kl.ObjString(shenForms))
			}
		})
	}
}

// readShenViaKernel parses a file using the kernel reader, but stops short of
// `process-sexprs` to get the same level of structure as parseShenFile returns.
func TestNativeReaderMatchesKernelSources(t *testing.T) {
	e := initShenForLoadCacheTest(t)

	// Walk the bundled kernel sources and verify the native reader matches the
	// kernel reader byte-for-byte (post-<s-exprs>, pre-process-sexprs).
	sourceDir := "../../kernel/sources"
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		t.Skipf("kernel sources unavailable: %v", err)
	}
	checked := 0
	for _, ent := range entries {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		if filepath.Ext(name) != ".shen" {
			continue
		}
		path := filepath.Join(sourceDir, name)
		t.Run(name, func(t *testing.T) {
			native, err := parseShenFile(path)
			if err != nil {
				t.Fatalf("native reader error on %s: %v", name, err)
			}
			shenForms := readShenViaKernel(t, e, path)
			if !shenObjEqual(native, shenForms) {
				t.Fatalf("native reader differs from kernel reader on %s", name)
			}
		})
		checked++
	}
	if checked == 0 {
		t.Fatal("no .shen files found in kernel/sources")
	}
}

func readShenViaKernel(tb testing.TB, e *kl.ControlFlow, path string) kl.Obj {
	tb.Helper()
	bytelist := kl.Call(e, kl.PrimFunc(kl.MakeSymbol("read-file-as-bytelist")), kl.MakeString(path))
	// The kernel exposes <s-exprs> via shen.<s-exprs>, but it must be invoked
	// via `compile`. Use `read-from-string-unprocessed`: it does
	// (compile <s-exprs> bytes), no process-sexprs.
	str := bytelistToString(e, bytelist)
	return kl.Call(e, kl.PrimFunc(kl.MakeSymbol("read-from-string-unprocessed")), str)
}

func bytelistToString(e *kl.ControlFlow, bytelist kl.Obj) kl.Obj {
	return kl.Call(e, kl.PrimFunc(kl.MakeSymbol("shen.bytes->string")), bytelist)
}

// shenObjEqual structurally compares two kl.Obj values, treating numbers,
// strings, symbols, booleans, and nested cons cells.
func shenObjEqual(a, b kl.Obj) bool {
	if a == b {
		return true
	}
	if a == kl.Nil || b == kl.Nil {
		return a == b
	}
	if kl.IsString(a) && kl.IsString(b) {
		return kl.GetString(a) == kl.GetString(b)
	}
	if kl.IsNumber(a) && kl.IsNumber(b) {
		return kl.ObjString(a) == kl.ObjString(b)
	}
	if kl.IsSymbol(a) && kl.IsSymbol(b) {
		return kl.GetSymbol(a) == kl.GetSymbol(b)
	}
	// Try as pairs.
	aSlice, aOK := tryListToSlice(a)
	bSlice, bOK := tryListToSlice(b)
	if !aOK || !bOK {
		return false
	}
	if len(aSlice) != len(bSlice) {
		return false
	}
	for i := range aSlice {
		if !shenObjEqual(aSlice[i], bSlice[i]) {
			return false
		}
	}
	return true
}

func tryListToSlice(o kl.Obj) ([]kl.Obj, bool) {
	defer func() {
		recover()
	}()
	return kl.ListToSlice(o), true
}
