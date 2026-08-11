package main

// Tests for the default path to native (autonative.go): after (load FILE), a
// plugin compiled from exactly FILE must be re-engaged, and a plugin compiled
// from a *different* version of FILE must not be.
//
// Two tiers, for the same reason the repo splits `make test` from `make certify`:
//   - TestEngageDecision is pure logic and runs everywhere in milliseconds. It is
//     the freshness guarantee itself (install only on a content match).
//   - TestAutoNativeEngagesAfterLoad drives the real binary against a real Go
//     plugin. It needs -buildmode=plugin (Linux/macOS) and about a minute of
//     compilation, so it is skipped under -short (which is what CI runs).

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestEngageDecision pins the decision table. The stale case is the one that
// matters: a plugin whose recorded source name matches but whose content hash
// does not must NEVER be installed — the file on disk has been edited since the
// .so was built, so the compiled code would silently answer with the old
// definitions.
func TestEngageDecision(t *testing.T) {
	const freshHash = "aaaa"
	const otherHash = "bbbb"

	cases := []struct {
		name string
		prov provenance
		sum  string
		base string
		want engageDecision
	}{
		{
			name: "exact content match engages",
			prov: provenance{source: "/build/host/prng.shen", sha256: freshHash},
			sum:  freshHash, base: "prng.shen",
			want: engageInstall,
		},
		{
			name: "content match engages even from another path",
			// The .so was built in one checkout and the identical file is
			// loaded from another: same bytes, so the compiled code is a
			// faithful compilation of what was just loaded.
			prov: provenance{source: "/build/host/prng.shen", sha256: freshHash},
			sum:  freshHash, base: "prng-copy.shen",
			want: engageInstall,
		},
		{
			name: "same name, edited content is stale",
			prov: provenance{source: "/build/host/prng.shen", sha256: freshHash},
			sum:  otherHash, base: "prng.shen",
			want: engageStale,
		},
		{
			name: "unrelated file is skipped",
			prov: provenance{source: "/build/host/prng.shen", sha256: freshHash},
			sum:  otherHash, base: "eventlog.shen",
			want: engageSkip,
		},
		{
			name: "manifest without provenance never engages",
			// A .so built before the manifest carried a source hash: we
			// cannot prove it is fresh, so we leave it alone.
			prov: provenance{},
			sum:  freshHash, base: "prng.shen",
			want: engageSkip,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := &nativeModule{so: "/tmp/x.so", prov: c.prov}
			if got := m.decide(c.sum, c.base); got != c.want {
				t.Fatalf("decide(%q, %q) = %d, want %d", c.sum, c.base, got, c.want)
			}
		})
	}
}

// TestProvenanceParsing checks the manifest header round-trips, including the
// requirement that unknown keys are ignored (so the format can grow) and that
// only a hash makes a plugin eligible.
func TestProvenanceParsing(t *testing.T) {
	var p provenance
	if p.ok() {
		t.Fatal("empty provenance must not be usable")
	}
	p.absorb("#source /a/b/c.shen")
	p.absorb("#future-key whatever")
	p.absorb("#sha256 deadbeef")
	if p.source != "/a/b/c.shen" || p.sha256 != "deadbeef" {
		t.Fatalf("parsed %+v", p)
	}
	if !p.ok() {
		t.Fatal("provenance with a hash must be usable")
	}
}

// TestAutoNativeEngagesAfterLoad is the end-to-end proof of the issue-#20 ask.
//
// Setup: a source file defining (probe) -> "vm", and a plugin (built from the
// hand-written fixture in cmd/plugintest/probeplugin) whose Install binds probe
// to a compiled version returning "native". The .fns manifest is stamped with the
// source's path and content hash, exactly as `make precompile` does.
//
// Then, running `shen -precompiled probe.so script driver.shen` where driver.shen
// does `(load "mod.shen")` before calling (probe):
//
//	default          -> "native"  (the plugin was re-engaged after the load)
//	-auto-native=false -> "vm"    (opt-out: load wins, as it did before)
//	source edited    -> "vm2" + a stale warning (never the stale compiled code)
func TestAutoNativeEngagesAfterLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a Go plugin (tens of seconds); run without -short")
	}
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Skipf("go plugins are not supported on %s", runtime.GOOS)
	}

	dir := t.TempDir()
	srcPath := filepath.Join(dir, "mod.shen")
	soPath := filepath.Join(dir, "probe.so")

	writeSource := func(body string) {
		if err := os.WriteFile(srcPath, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	// The VM definition of probe answers "vm"; the plugin answers "native".
	writeSource("(define probe -> \"vm\")\n")

	// Build the plugin from the fixture package.
	build := exec.Command("go", "build", "-buildmode=plugin", "-o", soPath, "./../plugintest/probeplugin")
	if out, err := build.CombinedOutput(); err != nil {
		t.Skipf("go build -buildmode=plugin unavailable here: %v\n%s", err, out)
	}

	// Stamp the manifest the way `make precompile` does.
	writeManifest := func() {
		data, err := os.ReadFile(srcPath)
		if err != nil {
			t.Fatal(err)
		}
		manifest := fmt.Sprintf("#source %s\n#sha256 %x\nprobe 0\n", srcPath, sha256.Sum256(data))
		if err := os.WriteFile(soPath+".fns", []byte(manifest), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	writeManifest()

	// `script` evaluates forms without echoing their values, so print the
	// marker explicitly.
	driver := filepath.Join(dir, "driver.shen")
	script := fmt.Sprintf("(load \"%s\")\n(output \"answered-by ~A~%%\" (probe))\n", srcPath)
	if err := os.WriteFile(driver, []byte(script), 0o644); err != nil {
		t.Fatal(err)
	}

	// Default: the plugin must be back in force after the load.
	out, _ := runCLI(t, "-precompiled", soPath, "script", driver)
	if !strings.Contains(out, "native") {
		t.Fatalf("auto-engage did not re-install the plugin after (load ...):\n%s", out)
	}

	// Opt-out: the load wins, exactly as it did before this feature.
	out, _ = runCLI(t, "-auto-native=false", "-precompiled", soPath, "script", driver)
	if !strings.Contains(out, "vm") || strings.Contains(out, "native") {
		t.Fatalf("-auto-native=false should have left the VM definition in force:\n%s", out)
	}

	// Diverged source: the .so is now stale. The freshly loaded definition must
	// win and the user must be told, rather than silently running old code.
	writeSource("(define probe -> \"vm2\")\n")
	out, _ = runCLI(t, "-precompiled", soPath, "script", driver)
	if strings.Contains(out, "native") {
		t.Fatalf("stale plugin was engaged for an edited source:\n%s", out)
	}
	if !strings.Contains(out, "vm2") {
		t.Fatalf("expected the edited VM definition to answer, got:\n%s", out)
	}
	if !strings.Contains(out, "compiled from a different version") {
		t.Fatalf("expected a staleness warning on stderr, got:\n%s", out)
	}
}
