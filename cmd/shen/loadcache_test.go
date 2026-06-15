package main

// Regression + benchmark for the per-process load cache (loadcache.go).
//
// These tests boot the full Shen kernel in-process (the same sequence main()
// runs) so they can drive the real (load "FILE") path and inspect the cache
// counters directly — something the subprocess CLI tests can't do.

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

// bootShen initializes a Shen runtime exactly like main(): register every
// kernel overlay, install the load cache, swap in the native hash, then run
// shen.initialise. The returned ControlFlow can evaluate (load ...) and
// arbitrary forms. Each call is independent (fresh symbol table is process
// global, but a fresh ControlFlow is fine for sequential eval here).
func bootShen(t testing.TB) *kl.ControlFlow {
	t.Helper()
	ns2_1set = kl.PrimFunc(kl.MakeSymbol("defun"))
	try_1catch = kl.PrimFunc(kl.MakeSymbol("try-catch"))
	kl.BindSymbolFunc(kl.MakeSymbol("load-native"), loadNative)

	var e kl.ControlFlow
	regist(&e)
	installNativeReader()
	installLoadCache()
	kl.BindSymbolFunc(kl.MakeSymbol("hash"), kl.MakePrimitive("hash", 2, kl.PrimHash))
	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	return &e
}

// evalForm reads one Shen form from src and evaluates it against e, returning the
// result Obj. It fails the test on a reader or eval error.
func evalForm(t testing.TB, e *kl.ControlFlow, src string) kl.Obj {
	t.Helper()
	res := kl.Eval(e, mustRead(t, src))
	if kl.IsError(res) {
		t.Fatalf("eval %q errored: %s", src, kl.GetString(kl.PrimErrorToString(res)))
	}
	return res
}

func mustRead(t testing.TB, src string) kl.Obj {
	t.Helper()
	exp, err := kl.NewSexpReader(strings.NewReader(src), false).Read()
	if err != nil {
		t.Fatalf("read %q: %v", src, err)
	}
	return exp
}

func writeFile(t testing.TB, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// TestLoadCacheEvaluatesCachedFormsAndInvalidatesOnChange is the core regression.
// It proves three things:
//  1. A repeated (load F) of an UNCHANGED file takes the cached parse path
//     (cache hit) yet still re-evaluates — the redefined value from the cold
//     load is still in effect, i.e. the cached reload is observably identical.
//  2. CHANGING the file invalidates the cache (cache miss) and the new
//     definition takes effect — eval side effects (defun, value) replay on the
//     cold re-parse.
//  3. Function arity metadata is (re)registered on every load, so a function
//     defined by a loaded file is applicable through the Shen reader
//     (partial-application path) after both a cached and a cold load.
func TestLoadCacheEvaluatesCachedFormsAndInvalidatesOnChange(t *testing.T) {
	e := bootShen(t)
	loadCache.reset()

	dir := t.TempDir()
	path := filepath.Join(dir, "probe.shen")
	loadExpr := `(load "` + path + `")`

	// --- cold load #1: define a 0-arity probe returning 41 and a 2-arity adder.
	writeFile(t, path, `(defun cache-probe () 41)
(defun cache-add (X Y) (+ X Y))`)
	evalForm(t, e, loadExpr)

	if got := kl.ObjString(evalForm(t, e, "(cache-probe)")); got != "41" {
		t.Fatalf("after cold load, (cache-probe) = %s, want 41", got)
	}
	// Arity metadata must be registered: applying cache-add through the reader
	// (which resolves applications via stored arity) must work.
	if got := kl.ObjString(evalForm(t, e, "(cache-add 2 3)")); got != "5" {
		t.Fatalf("after cold load, (cache-add 2 3) = %s, want 5", got)
	}

	h1, m1 := loadCache.stats()
	if m1 != 1 || h1 != 0 {
		t.Fatalf("after cold load: hits=%d misses=%d, want hits=0 misses=1", h1, m1)
	}

	// --- cached load #2: SAME bytes. Must be a cache hit, must still evaluate.
	evalForm(t, e, loadExpr)
	h2, m2 := loadCache.stats()
	if h2 != 1 || m2 != 1 {
		t.Fatalf("after unchanged reload: hits=%d misses=%d, want hits=1 misses=1 (cache hit)", h2, m2)
	}
	// Observably identical: the function is still defined and behaves the same.
	if got := kl.ObjString(evalForm(t, e, "(cache-probe)")); got != "41" {
		t.Fatalf("after cached reload, (cache-probe) = %s, want 41", got)
	}
	if got := kl.ObjString(evalForm(t, e, "(cache-add 4 5)")); got != "9" {
		t.Fatalf("after cached reload, (cache-add 4 5) = %s, want 9 (arity replayed)", got)
	}

	// --- change the file: cache must invalidate and the NEW body take effect.
	writeFile(t, path, `(defun cache-probe () 42)
(defun cache-add (X Y) (+ X Y))`)
	evalForm(t, e, loadExpr)
	h3, m3 := loadCache.stats()
	if m3 != 2 {
		t.Fatalf("after changed reload: misses=%d, want 2 (content-hash invalidation)", m3)
	}
	if h3 != 1 {
		t.Fatalf("after changed reload: hits=%d, want still 1 (changed file is not a hit)", h3)
	}
	if got := kl.ObjString(evalForm(t, e, "(cache-probe)")); got != "42" {
		t.Fatalf("after changed reload, (cache-probe) = %s, want 42 (eval replayed new body)", got)
	}

	// --- load the now-current bytes again: hit again.
	evalForm(t, e, loadExpr)
	h4, m4 := loadCache.stats()
	if h4 != 2 || m4 != 2 {
		t.Fatalf("after re-loading current bytes: hits=%d misses=%d, want hits=2 misses=2", h4, m4)
	}
}

// TestLoadHonorsHomeDirectory is the regression for the PR #16 review finding:
// the native read-file (and the load cache) must resolve a relative (load "F")
// through *home-directory* exactly as the kernel `open` does, NOT read F from
// the process CWD. The repro puts the SAME relative filename in both the CWD and
// *home-directory*, with different definitions, and asserts the
// *home-directory* copy is the one that loads — matching the interpreted reader.
func TestLoadHonorsHomeDirectory(t *testing.T) {
	e := bootShen(t)
	loadCache.reset()

	homeSym := kl.MakeSymbol("*home-directory*")
	orig := kl.PrimValue(homeSym)
	defer kl.PrimSet(homeSym, orig) // don't leak the home dir into other tests

	cwdDir := t.TempDir()
	homeDir := t.TempDir()
	// Same relative name in both places, different definitions.
	writeFile(t, filepath.Join(cwdDir, "target.shen"), `(defun home-probe () 111)`)
	writeFile(t, filepath.Join(homeDir, "target.shen"), `(defun home-probe () 222)`)

	t.Chdir(cwdDir) // process CWD now contains the 111 copy
	// *home-directory* must end with a separator: `open` concatenates literally.
	kl.PrimSet(homeSym, kl.MakeString(homeDir+string(filepath.Separator)))

	evalForm(t, e, `(load "target.shen")`)
	if got := kl.ObjString(evalForm(t, e, "(home-probe)")); got != "222" {
		t.Fatalf("(load \"target.shen\") read the CWD copy, not *home-directory*: (home-probe) = %s, want 222", got)
	}

	// A cached reload of the same (home, path) must hit and still be 222.
	evalForm(t, e, `(load "target.shen")`)
	if got := kl.ObjString(evalForm(t, e, "(home-probe)")); got != "222" {
		t.Fatalf("after cached reload, (home-probe) = %s, want 222", got)
	}

	// Point *home-directory* at a third dir whose copy differs: the cache key is
	// the resolved path, so this must NOT collide with the previous entry.
	homeDir2 := t.TempDir()
	writeFile(t, filepath.Join(homeDir2, "target.shen"), `(defun home-probe () 333)`)
	kl.PrimSet(homeSym, kl.MakeString(homeDir2+string(filepath.Separator)))
	evalForm(t, e, `(load "target.shen")`)
	if got := kl.ObjString(evalForm(t, e, "(home-probe)")); got != "333" {
		t.Fatalf("after switching *home-directory*, (home-probe) = %s, want 333 (cache must key on resolved path)", got)
	}
}

// BenchmarkRepeatedCachedLoad measures the steady-state cost of re-(load)ing an
// unchanged file — the path the cache accelerates. Compare against
// BenchmarkColdLoad below (which forces a re-parse every iteration).
func BenchmarkRepeatedCachedLoad(b *testing.B) {
	e := bootShen(b)
	loadCache.reset()

	dir := b.TempDir()
	path := filepath.Join(dir, "bench.shen")
	writeFile(b, path, benchSource)
	loadExpr := `(load "` + path + `")`
	form := mustRead(b, loadExpr)

	// Prime the cache once so every measured iteration is a hit.
	kl.Eval(e, form)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kl.Eval(e, form)
	}
}

// BenchmarkColdLoad measures the same load with the cache disabled (reset before
// each iteration), i.e. a full tokenize/parse/lower every time. This is the
// baseline the cache improves on.
func BenchmarkColdLoad(b *testing.B) {
	e := bootShen(b)

	dir := b.TempDir()
	path := filepath.Join(dir, "bench.shen")
	writeFile(b, path, benchSource)
	loadExpr := `(load "` + path + `")`
	form := mustRead(b, loadExpr)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loadCache.reset() // force a cold parse each iteration
		kl.Eval(e, form)
	}
}

// benchSource is a few non-trivial defuns: enough tokenize/parse/lower work that
// the cache win is visible, without depending on any external file.
const benchSource = `(defun bench-fib (N)
  (if (< N 2)
      N
      (+ (bench-fib (- N 1)) (bench-fib (- N 2)))))

(defun bench-map (F Xs)
  (if (empty? Xs)
      ()
      (cons (F (head Xs)) (bench-map F (tail Xs)))))

(defun bench-sum (Xs)
  (if (empty? Xs)
      0
      (+ (head Xs) (bench-sum (tail Xs)))))

(defun bench-range (A B)
  (if (>= A B)
      ()
      (cons A (bench-range (+ A 1) B))))
`
