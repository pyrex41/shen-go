package main

import (
	"crypto/sha256"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/tiancaiamao/shen-go/kl"
)

type loadCacheKey struct {
	path      string
	typecheck bool
}

type loadCacheFreshness struct {
	size    int64
	modTime time.Time
	hash    [sha256.Size]byte
}

type loadCacheEntry struct {
	freshness  loadCacheFreshness
	parsed     kl.Obj
	typecheck  bool
	sourcePath string
}

var (
	loadCacheMu     sync.Mutex
	loadCache       = make(map[loadCacheKey]loadCacheEntry)
	loadCacheHits   int
	loadCacheMisses int
)

func installLoadCache() {
	kl.BindSymbolFunc(symload, kl.MakeNative(primCachedLoad, 1))
}

func primCachedLoad(e *kl.ControlFlow) {
	file := kl.GetString(e.Get(1))
	typecheckFlag := kl.PrimValue(symshen_4_dtc_d)
	typecheck := typecheckFlag != kl.False
	start := kl.PrimGetTime(symrun)

	entry, ok := lookupLoadCache(file, typecheck)
	if !ok {
		entry.parsed = readShenForms(e, entry.sourcePath)
		storeLoadCache(entry)
	}

	kl.Call(e, kl.PrimFunc(symshen_4load_1help), typecheckFlag, entry.parsed)

	printLoadRunTime(e, start)
	if typecheck {
		printLoadTypecheckSummary(e)
	}
	e.Return(symloaded)
}

func lookupLoadCache(file string, typecheck bool) (loadCacheEntry, bool) {
	sourcePath, freshness, err := statLoadCacheFile(file)
	if err != nil {
		panic(kl.MakeError(err.Error()))
	}

	key := loadCacheKey{path: sourcePath, typecheck: typecheck}
	loadCacheMu.Lock()
	defer loadCacheMu.Unlock()

	entry, ok := loadCache[key]
	if ok && entry.freshness == freshness {
		loadCacheHits++
		return entry, true
	}

	loadCacheMisses++
	return loadCacheEntry{
		freshness:  freshness,
		typecheck:  typecheck,
		sourcePath: sourcePath,
	}, false
}

func statLoadCacheFile(file string) (string, loadCacheFreshness, error) {
	absPath, err := filepath.Abs(file)
	if err != nil {
		return "", loadCacheFreshness{}, err
	}
	sourcePath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return "", loadCacheFreshness{}, err
	}
	info, err := os.Stat(sourcePath)
	if err != nil {
		return "", loadCacheFreshness{}, err
	}
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return "", loadCacheFreshness{}, err
	}
	return sourcePath, loadCacheFreshness{
		size:    info.Size(),
		modTime: info.ModTime(),
		hash:    sha256.Sum256(content),
	}, nil
}

// readShenForms reads a .shen source file and returns the same value the
// kernel `read-file` would produce — a list of forms after `process-sexprs`.
// On supported syntax it uses the Go-native parser to skip the kernel's slow
// byte-list and parser-combinator passes; otherwise it falls back to
// `read-file` so behavior matches exactly.
func readShenForms(e *kl.ControlFlow, path string) kl.Obj {
	sexprs, err := parseShenFile(path)
	if err != nil {
		return kl.Call(e, kl.PrimFunc(symread_1file), kl.MakeString(path))
	}
	return kl.Call(e, kl.PrimFunc(symshen_4process_1sexprs), sexprs)
}

func storeLoadCache(entry loadCacheEntry) {
	key := loadCacheKey{path: entry.sourcePath, typecheck: entry.typecheck}
	loadCacheMu.Lock()
	loadCache[key] = entry
	loadCacheMu.Unlock()
}

func printLoadRunTime(e *kl.ControlFlow, start kl.Obj) {
	elapsed := kl.PrimNumberSubtract(kl.PrimGetTime(symrun), start)
	secs := kl.PrimStringConcat(kl.PrimStr(elapsed), kl.MakeString(" secs\n"))
	msg := kl.PrimStringConcat(kl.MakeString("\nrun time: "), secs)
	kl.Call(e, kl.PrimFunc(sympr), msg, kl.Call(e, kl.PrimFunc(symstoutput)))
}

func printLoadTypecheckSummary(e *kl.ControlFlow) {
	inferences := kl.Call(e, kl.PrimFunc(syminferences))
	detail := kl.Call(e, kl.PrimFunc(symshen_4app), inferences, kl.MakeString(" inferences\n"), symshen_4a)
	msg := kl.PrimStringConcat(kl.MakeString("\ntypechecked in "), detail)
	kl.Call(e, kl.PrimFunc(sympr), msg, kl.Call(e, kl.PrimFunc(symstoutput)))
}

func resetLoadCache() {
	loadCacheMu.Lock()
	loadCache = make(map[loadCacheKey]loadCacheEntry)
	loadCacheHits = 0
	loadCacheMisses = 0
	loadCacheMu.Unlock()
}

func loadCacheStats() (int, int) {
	loadCacheMu.Lock()
	defer loadCacheMu.Unlock()
	return loadCacheHits, loadCacheMisses
}
