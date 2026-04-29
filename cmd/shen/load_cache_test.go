package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

var shenTestInit sync.Once

func initShenForLoadCacheTest(tb testing.TB) *kl.ControlFlow {
	tb.Helper()

	ns2_1set = kl.PrimFunc(kl.MakeSymbol("defun"))
	try_1catch = kl.PrimFunc(kl.MakeSymbol("try-catch"))

	var e kl.ControlFlow
	shenTestInit.Do(func() {
		regist(&e)
		kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	})
	resetLoadCache()

	var output bytes.Buffer
	kl.PrimSet(kl.MakeSymbol("*stoutput*"), kl.MakeStream(&output))
	return &e
}

func TestLoadCacheEvaluatesCachedFormsAndInvalidatesOnChange(t *testing.T) {
	e := initShenForLoadCacheTest(t)
	counter := kl.MakeSymbol("load-cache-test-counter")
	kl.PrimSet(counter, kl.MakeInteger(0))

	path := filepath.Join(t.TempDir(), "cached-load.shen")
	writeLoadCacheTestFile(t, path, "(set load-cache-test-counter (+ (value load-cache-test-counter) 1))\n")

	loadShenFile(t, e, path)
	assertIntegerValue(t, counter, 1)

	loadShenFile(t, e, path)
	assertIntegerValue(t, counter, 2)

	hits, misses := loadCacheStats()
	if hits != 1 || misses != 1 {
		t.Fatalf("expected one cache hit and one miss after repeated load, got hits=%d misses=%d", hits, misses)
	}

	writeLoadCacheTestFile(t, path, "(set load-cache-test-counter 10)\n")
	loadShenFile(t, e, path)
	assertIntegerValue(t, counter, 10)

	hits, misses = loadCacheStats()
	if hits != 1 || misses != 2 {
		t.Fatalf("expected source change to invalidate cache, got hits=%d misses=%d", hits, misses)
	}
}

// TestLoadCachePredicateReturnsBooleanForIf reproduces a regression where the
// native reader emitted the symbol `false` (instead of the boolean False) for
// the literal `false` token in a .shen file. Predicates defined as `-> true /
// -> false` then returned a symbol, and `(if Pred ...)` would panic in
// evalIf with "second argument of if should be boolean".
func TestLoadCachePredicateReturnsBooleanForIf(t *testing.T) {
	e := initShenForLoadCacheTest(t)

	path := filepath.Join(t.TempDir(), "predicate.shen")
	writeLoadCacheTestFile(t, path, "(define cache-test-pred X -> false)\n(set cache-test-result (if (cache-test-pred 1) 1 2))\n")

	loadShenFile(t, e, path)
	got := kl.GetInteger(kl.PrimValue(kl.MakeSymbol("cache-test-result")))
	if got != 2 {
		t.Fatalf("cache-test-result = %d, want 2 (predicate returning false should fall to else branch)", got)
	}
}

func TestLoadCacheReinstallsDefunAfterClobbering(t *testing.T) {
	e := initShenForLoadCacheTest(t)

	path := filepath.Join(t.TempDir(), "defun-cache.shen")
	writeLoadCacheTestFile(t, path, "(define cache-test-func -> 42)\n")

	loadShenFile(t, e, path)
	fooSym := kl.MakeSymbol("cache-test-func")
	if got := kl.GetInteger(kl.Call(e, kl.PrimFunc(fooSym))); got != 42 {
		t.Fatalf("first load: cache-test-func returned %d, want 42", got)
	}

	kl.BindSymbolFunc(fooSym, nil)

	loadShenFile(t, e, path)
	if got := kl.GetInteger(kl.Call(e, kl.PrimFunc(fooSym))); got != 42 {
		t.Fatalf("after cache hit reload: cache-test-func returned %d, want 42", got)
	}

	hits, misses := loadCacheStats()
	if hits != 1 || misses != 1 {
		t.Fatalf("expected one cache hit and one miss, got hits=%d misses=%d", hits, misses)
	}
}

// TestLoadCacheFallsBackForDollarSplice covers the path where the native
// reader rejects a file (here, because it contains `($ X)`) and primCachedLoad
// falls back to the kernel `read-file`. The file should still load and its
// effects should be visible.
func TestLoadCacheFallsBackForDollarSplice(t *testing.T) {
	e := initShenForLoadCacheTest(t)

	probe := kl.MakeSymbol("cache-test-fallback-probe")
	kl.PrimSet(probe, kl.MakeInteger(0))

	path := filepath.Join(t.TempDir(), "splice.shen")
	// `($ "ab")` splices to atoms `a` and `b` at top level (which the kernel
	// reads and the kernel evaluator treats as bare symbol references — no
	// effect). The (set ...) form after it is what we observe.
	writeLoadCacheTestFile(t, path, "($ \"ab\")\n(set cache-test-fallback-probe 7)\n")

	loadShenFile(t, e, path)

	got := kl.GetInteger(kl.PrimValue(probe))
	if got != 7 {
		t.Fatalf("fallback path didn't run set form: probe=%d, want 7", got)
	}
}

// TestLoadCacheConcurrentPrimitives stresses the cache lookup/store mutex by
// hammering it from many goroutines. Run with `-race` to validate; it does
// NOT exercise concurrent `(load ...)` calls, since the kl runtime itself is
// not goroutine-safe.
func TestLoadCacheConcurrentPrimitives(t *testing.T) {
	resetLoadCache()
	dir := t.TempDir()
	const fileCount = 8
	const goroutines = 16
	const iterations = 200

	paths := make([]string, fileCount)
	for i := range paths {
		paths[i] = filepath.Join(dir, fmt.Sprintf("concurrent-%d.shen", i))
		if err := os.WriteFile(paths[i], []byte(fmt.Sprintf("(set probe-%d %d)\n", i, i)), 0644); err != nil {
			t.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				p := paths[(g+i)%fileCount]
				entry, ok := lookupLoadCache(p, false)
				if !ok {
					storeLoadCache(entry)
				}
			}
		}(g)
	}
	wg.Wait()

	hits, misses := loadCacheStats()
	totalOps := goroutines * iterations
	if hits+misses != totalOps {
		t.Fatalf("hits+misses=%d, want %d", hits+misses, totalOps)
	}
	// At minimum, every goroutine after the first cache fill should see at
	// least some hits — guard against the mutex being overly aggressive and
	// every call missing.
	if hits == 0 {
		t.Fatalf("expected at least some cache hits, got %d (misses=%d)", hits, misses)
	}
}

func BenchmarkRepeatedCachedLoad(b *testing.B) {
	e := initShenForLoadCacheTest(b)
	counter := kl.MakeSymbol("load-cache-benchmark-counter")
	kl.PrimSet(counter, kl.MakeInteger(0))

	path := filepath.Join(b.TempDir(), "cached-load-bench.shen")
	writeLoadCacheTestFile(b, path, "(set load-cache-benchmark-counter (+ (value load-cache-benchmark-counter) 1))\n")

	loadShenFile(b, e, path)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loadShenFile(b, e, path)
	}
}

// BenchmarkLargeFileCachedLoad exercises a realistic workload: repeatedly
// loading a large generated .shen file from cache. Measures the post-hit cost
// (shen.load-help over already-parsed forms).
func BenchmarkLargeFileCachedLoad(b *testing.B) {
	e := initShenForLoadCacheTest(b)
	path := writeSyntheticShenFile(b, 200)
	loadShenFile(b, e, path)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loadShenFile(b, e, path)
	}
}

// BenchmarkLargeFileColdLoad measures a fresh cold load — cache reset between
// iterations — exercising the native reader plus process-sexprs plus
// shen.load-help.
func BenchmarkLargeFileColdLoad(b *testing.B) {
	e := initShenForLoadCacheTest(b)
	path := writeSyntheticShenFile(b, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		resetLoadCache()
		b.StartTimer()
		loadShenFile(b, e, path)
	}
}

// writeSyntheticShenFile emits a .shen file with `n` `define` forms exercising
// pattern matching, list literals, and let-bindings — similar in shape to a
// typical kernel source file.
func writeSyntheticShenFile(tb testing.TB, n int) string {
	tb.Helper()
	var buf []byte
	for i := 0; i < n; i++ {
		buf = append(buf, []byte(formatSyntheticDefine(i))...)
	}
	path := filepath.Join(tb.TempDir(), "synthetic.shen")
	if err := os.WriteFile(path, buf, 0644); err != nil {
		tb.Fatal(err)
	}
	return path
}

func formatSyntheticDefine(i int) string {
	return fmt.Sprintf(
		"(define synth-fn-%d\n"+
			"  []        -> 0\n"+
			"  [X | Xs]  -> (let Y (+ X %d)\n"+
			"                    Z [Y | (synth-fn-%d Xs)]\n"+
			"                 (cons Y (synth-fn-%d Xs))))\n\n",
		i, i, i, i)
}

func writeLoadCacheTestFile(tb testing.TB, path, content string) {
	tb.Helper()
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		tb.Fatal(err)
	}
}

func loadShenFile(tb testing.TB, e *kl.ControlFlow, path string) {
	tb.Helper()
	res := kl.Call(e, kl.PrimFunc(symload), kl.MakeString(path))
	if res != symloaded {
		tb.Fatalf("load returned %s, want loaded", kl.ObjString(res))
	}
}

func assertIntegerValue(tb testing.TB, symbol kl.Obj, want int) {
	tb.Helper()
	got := kl.GetInteger(kl.PrimValue(symbol))
	if got != want {
		tb.Fatalf("%s = %d, want %d", kl.ObjString(symbol), got, want)
	}
}
