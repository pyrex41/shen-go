package main

// Per-process load cache for (load "FILE.shen").
//
// In shen-go the load pipeline is, per the kernel:
//
//	(load F) -> (read-file F) -> shen.load-help -> for each form: shen.shen->kl
//	                                                              then eval-kl
//
// read-file is the expensive, interpreted half: it reads the raw bytes (native,
// cheap) and then *tokenizes, parses and lowers* them (shen.<s-exprs>,
// shen.process-sexprs) entirely in interpreted Shen/KL. For files that do not
// change between loads in the same process (e.g. a host re-loading the same
// kernel/library file, or a REPL session that loads a file repeatedly) that
// tokenize/parse/lower work is pure waste.
//
// This cache memoizes the *result* of read-file (the fully tokenized, parsed and
// lowered form list) keyed by absolute path + a content hash of the file bytes.
// On a cache hit we return the cached form list and skip the interpreted
// tokenize/parse/lower. We do NOT cache evaluation: load-help / shen->kl /
// eval-kl still run on every load, so every observable side effect of a load —
// defun arity registration, datatype/type declarations, value sets, prints — is
// replayed identically. A cached reload is therefore observably indistinguishable
// from a cold load; only the redundant parsing is elided.
//
// Invalidation is by content hash: any change to the file bytes (size or
// content) produces a different key and forces a cold re-parse. We deliberately
// do not key on mtime alone — content hashing is immune to mtime resolution and
// to edit-in-place-with-preserved-mtime, and the byte read it requires is cheap
// next to the parse it saves.
//
// Error results are never cached: a read-file that raises (reader error, missing
// file) falls straight through to the kernel every time.

import (
	"crypto/sha256"
	"os"
	"sync"

	"github.com/pyrex41/shen-go/kl"
)

// loadCacheEntry is one memoized parse: the content hash it was parsed from and
// the form list read-file produced.
type loadCacheEntry struct {
	hash  [32]byte
	forms kl.Obj
}

// loadCache memoizes read-file results keyed by absolute file path. It is
// per-process and guarded by a mutex so a future concurrent host cannot corrupt
// it; the common single-threaded REPL/load path takes the lock uncontended.
type loadCacheT struct {
	mu      sync.Mutex
	entries map[string]loadCacheEntry
	hits    int
	misses  int
}

var loadCache = &loadCacheT{entries: make(map[string]loadCacheEntry)}

// stats returns the (hit, miss) counters. Used by tests to prove a second load
// of an unchanged file took the cached path.
func (c *loadCacheT) stats() (hits, misses int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.hits, c.misses
}

// reset clears the cache and counters. Used by tests.
func (c *loadCacheT) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[string]loadCacheEntry)
	c.hits = 0
	c.misses = 0
}

// lookup returns the cached forms for path iff an entry exists whose hash matches
// the current file content. Reading the file to hash it is the only I/O the fast
// path does; on a hit we skip the whole interpreted parse.
//
// The path is resolved through *home-directory* (kl.ResolveHomePath) exactly as
// the kernel `open` does, and the RESOLVED path is the cache key: the same
// relative argument under two different *home-directory* values denotes two
// different files and must not share a cache entry.
func (c *loadCacheT) lookup(path string) (kl.Obj, bool) {
	resolved := kl.ResolveHomePath(path)
	data, err := os.ReadFile(resolved)
	if err != nil {
		// Can't hash it (missing/unreadable): let the kernel read-file produce
		// the canonical error. Never cache.
		return nil, false
	}
	h := sha256.Sum256(data)

	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.entries[resolved]; ok && e.hash == h {
		c.hits++
		return e.forms, true
	}
	c.misses++
	return nil, false
}

// store records forms as the parse of path's current content. It re-hashes the
// file rather than trusting a hash threaded from lookup so the stored key always
// reflects what was actually parsed. Like lookup, it keys on the
// *home-directory*-resolved path.
func (c *loadCacheT) store(path string, forms kl.Obj) {
	resolved := kl.ResolveHomePath(path)
	data, err := os.ReadFile(resolved)
	if err != nil {
		return
	}
	h := sha256.Sum256(data)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[resolved] = loadCacheEntry{hash: h, forms: forms}
}

// installLoadCache rebinds the kernel's read-file with a caching wrapper. It must
// run AFTER regist (which binds the interpreted read-file) so we can capture the
// original and delegate to it on a miss. The wrapper preserves read-file's exact
// contract: same argument (the file path), same return value (the form list, or
// an error object the caller can catch).
func installLoadCache() {
	readFileSym := kl.MakeSymbol("read-file")
	orig := kl.PrimFunc(readFileSym) // the interpreted kernel read-file

	wrapper := kl.MakeNative(func(e *kl.ControlFlow) {
		arg := e.Get(1)
		path := kl.GetString(arg)

		if forms, ok := loadCache.lookup(path); ok {
			e.Return(forms)
			return
		}

		// Miss: run the real read-file (tokenize/parse/lower). Only cache a
		// clean parse — never an error result.
		forms := kl.Call(e, orig, arg)
		if !kl.IsError(forms) {
			loadCache.store(path, forms)
		}
		e.Return(forms)
	}, 1)

	kl.BindSymbolFunc(readFileSym, wrapper)
}
