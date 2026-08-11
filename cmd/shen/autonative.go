package main

// The default path to native for hot *loaded* modules.
//
// The problem this solves (issue #20): a plugin binds its compiled functions at
// startup, but Shen's `load` is a definer — `(load "prng.shen")` re-runs every
// `define` in the file and rebinds those same symbols back onto the bytecode VM,
// silently discarding the compiled versions. Suite runners open with exactly
// that (`(load "shen/world/prng.shen")` at the top of run-tests.shen), so
// `-precompiled` alone bought them nothing and the only way to get the native
// code was to remember to call `(load-native "prng.so")` again *after* every
// load. That is not a default path; it is a footgun with a workaround.
//
// So: after any successful `(load F)`, if F is the source a loaded plugin was
// compiled from, re-install that plugin. The user's intent is already explicit —
// they named the .so on the command line (`-precompiled`) or called load-native —
// and this just keeps that intent true across a load that would otherwise undo
// it. Nothing is discovered from disk, so no code runs that the user did not ask
// for.
//
// Staleness: "is F the source this plugin was compiled from" is answered by
// CONTENT, not by path or mtime. `make precompile` records the sha256 of the
// source bytes in the .so's .fns manifest; we hash the file we just loaded and
// require an exact match. That makes a false positive impossible in the way that
// matters: if the .shen file has been edited since the .so was built, the hashes
// differ, we do NOT install, and the freshly-loaded VM definitions — which are
// what the source actually says — stay in force. A path/mtime check would have
// to guess; a content hash does not. When the file's *name* matches a plugin's
// recorded source but the content does not, that is a stale plugin the user
// probably meant to rebuild, so we say so once on stderr.
//
// Opt-out: `-auto-native=false`, or SHEN_NO_AUTO_NATIVE=1 in the environment for
// harnesses that cannot change the command line. With auto-engagement off the
// old manual ordering (`(load ...)` then `(load-native ...)`) still works.
//
// Nothing is printed on the happy path — a suite runner's stdout is often
// compared against a golden file, so this feature must be invisible unless
// something is wrong. SHEN_AUTO_NATIVE_VERBOSE=1 narrates each engagement on
// stderr when you want to confirm it fired.

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pyrex41/shen-go/kl"
)

// provenance is the header a `make precompile` manifest carries about the source
// file the plugin was built from:
//
//	#source /abs/path/to/prng.shen
//	#sha256 <hex of that file's bytes at precompile time>
//
// A manifest without these (built before provenance existed) yields a zero
// provenance, whose ok() is false: such a plugin can still be loaded and used,
// it just cannot be auto-engaged, because we have no way to prove it is not
// stale.
type provenance struct {
	source string
	sha256 string
}

func (p provenance) ok() bool { return p.sha256 != "" }

// absorb parses one "#key value" manifest header line into p. Unknown keys are
// ignored so the manifest format can grow without breaking older binaries.
func (p *provenance) absorb(line string) {
	key, value, found := strings.Cut(strings.TrimPrefix(line, "#"), " ")
	if !found {
		return
	}
	value = strings.TrimSpace(value)
	switch key {
	case "source":
		p.source = value
	case "sha256":
		p.sha256 = value
	}
}

// nativeModule is one loaded plugin plus the provenance of its source.
type nativeModule struct {
	so   string
	prov provenance
	// warnedStale remembers which loaded paths we have already complained
	// about, so a suite that loads the same diverged file in a loop prints
	// one warning rather than one per load.
	warnedStale map[string]bool
}

// nativeModules is every plugin loaded in this process, in load order. Written
// only by registerNativeModule (startup -precompiled and load-native), read by
// engageNativeFor. Single-threaded, like the rest of the loader.
var nativeModules []*nativeModule

// autoNative is the master switch for auto-engagement, resolved once in main
// from -auto-native and SHEN_NO_AUTO_NATIVE.
var autoNative = true

// autoNativeVerbose narrates engagements on stderr (SHEN_AUTO_NATIVE_VERBOSE=1).
// Off by default: the happy path must not perturb a suite's output.
var autoNativeVerbose = os.Getenv("SHEN_AUTO_NATIVE_VERBOSE") != ""

// engageDecision is what to do with one registered plugin when some file has
// just been loaded.
type engageDecision int

const (
	// engageSkip: this plugin has nothing to do with the loaded file.
	engageSkip engageDecision = iota
	// engageInstall: the loaded file is byte-for-byte the source this plugin
	// was compiled from — re-bind its compiled functions.
	engageInstall
	// engageStale: same file name, different content. The plugin is out of
	// date with respect to the source; say so and stay on the VM.
	engageStale
)

// decide is the whole safety argument in one function: install only on an exact
// content match, never on a name match alone.
//
// sum is the hex sha256 of the file just loaded; base is its base name.
func (m *nativeModule) decide(sum, base string) engageDecision {
	if !m.prov.ok() {
		// A manifest with no recorded source hash (a .so built before
		// provenance existed). We cannot prove freshness, so we do not
		// auto-engage it — the user can still (load-native ...) explicitly.
		return engageSkip
	}
	if m.prov.sha256 == sum {
		return engageInstall
	}
	if filepath.Base(m.prov.source) == base {
		return engageStale
	}
	return engageSkip
}

// registerNativeModule records (or refreshes) a loaded plugin. Keyed by .so
// path: re-loading the same .so updates the entry in place rather than growing
// the registry, so a script that calls load-native in a loop stays bounded.
func registerNativeModule(soPath string, prov provenance) {
	for _, m := range nativeModules {
		if m.so == soPath {
			m.prov = prov
			return
		}
	}
	nativeModules = append(nativeModules, &nativeModule{so: soPath, prov: prov})
}

// installAutoNative wraps the kernel's `load` so a successful load is followed
// by re-engaging any plugin compiled from exactly that file.
//
// Must run AFTER regist (which binds the kernel `load`) so we can capture and
// delegate to it. The wrapper preserves load's contract exactly: same argument,
// same return value, same raised errors — the only addition happens after load
// has already produced its result.
func installAutoNative() {
	loadSym := kl.MakeSymbol("load")
	orig := kl.PrimFunc(loadSym)

	wrapper := kl.MakeNative(func(e *kl.ControlFlow) {
		arg := e.Get(1)
		res := kl.Call(e, orig, arg)
		// Only a clean load can have rebound anything worth re-engaging;
		// a failed load leaves the image alone.
		if !kl.IsError(res) {
			engageNativeFor(e, kl.GetString(arg))
		}
		e.Return(res)
	}, 1)

	kl.BindSymbolFunc(loadSym, wrapper)
}

// engageNativeFor re-installs every loaded plugin whose recorded source content
// hash equals the current content of the file just loaded.
//
// Cost when nothing is precompiled (the overwhelmingly common case): one branch.
// With plugins loaded it is one file read + one sha256 per load, next to a load
// that just tokenized, macroexpanded and evaluated that same file.
func engageNativeFor(e *kl.ControlFlow, path string) {
	if !autoNative || len(nativeModules) == 0 {
		return
	}
	// Resolve exactly as the kernel `open` does (honoring *home-directory*), so
	// we hash the same bytes load just read.
	resolved := kl.ResolveHomePath(path)
	data, err := os.ReadFile(resolved)
	if err != nil {
		return
	}
	sum := fmt.Sprintf("%x", sha256.Sum256(data))
	base := filepath.Base(resolved)

	// Snapshot: installPlugin does not touch the registry, but iterate over a
	// copy of the header anyway so this stays correct if that ever changes.
	mods := nativeModules
	for _, m := range mods {
		switch m.decide(sum, base) {
		case engageInstall:
			if _, err := installPlugin(e, m.so); err != nil {
				fmt.Fprintf(os.Stderr, "warning: could not re-engage %s after loading %s: %v (continuing on the VM)\n", m.so, path, err)
				continue
			}
			if autoNativeVerbose {
				fmt.Fprintf(os.Stderr, "auto-native: re-engaged %s after loading %s\n", m.so, path)
			}
		case engageStale:
			m.warnStale(path)
		}
	}
}

// warnStale reports, once per loaded path, that a plugin was compiled from a
// different version of the file just loaded. We deliberately do not install it:
// the source on disk is the truth, and running the old compiled code instead
// would be a silent, extremely confusing wrong answer.
func (m *nativeModule) warnStale(path string) {
	if m.warnedStale == nil {
		m.warnedStale = make(map[string]bool)
	}
	if m.warnedStale[path] {
		return
	}
	m.warnedStale[path] = true
	fmt.Fprintf(os.Stderr,
		"warning: %s was compiled from a different version of %s; staying on the VM for it (rebuild with `make precompile FILE=%s OUT=%s`)\n",
		m.so, path, path, m.so)
}
