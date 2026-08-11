# Shen/Go, a Go port of the Shen language

Shen is a portable functional programming language by [Mark Tarver](http://marktarver.com) that offers

- pattern matching,
- λ calculus consistency,
- macros,
- optional lazy evaluation,
- static type checking,
- an integrated fully functional Prolog,
- and an inbuilt compiler-compiler.

shen-go is a port of the Shen language that runs on top of Go implementations.

## Building

Make sure you have [Go installed](https://golang.org/doc/install).

```
make shen
```

or for windows

```
make shen-exe
```

## Running

```
./shen
```

This binary has no dependency, you can move it to any where you want.

## Standard library

The Shen kernel ships no standard library, so shen-go bundles one from Tarver's
`Lib/StLib` (vendored under `cmd/shen/stlib/`, `go:embed`ed into the binary and
loaded at startup). Functions like `filter`, `mapc`, `take`/`drop`,
`foldl`/`foldr`, `sort`, and the Maths/Strings/Vectors/Tuples helpers work out
of the box — no separate install step. Loading adds ~0.12s to startup; set
`SHEN_NO_STDLIB=1` to skip it (kernel functions only). See
`cmd/shen/stlib/PROVENANCE.md`.

## Performance: compile hot files to Go (optional)

By default every function runs on the bytecode VM — instant to define, ideal for
the REPL. For hot, compute-heavy code you can instead **compile a whole file to
native Go** ahead of time and load it at startup. The two coexist in one process:
a precompiled function overrides only its own binding; everything else (including
anything you type at the REPL) stays on the VM, and the two call each other freely.

**1. Precompile a file (offline, ~1s per function):**

```
make precompile FILE=bench/hot.shen OUT=hot.so
```

This runs Shen → KL → IR → Go → `go build -buildmode=plugin`, emitting `hot.so`
plus `hot.so.fns` (an arity manifest, headed with the source path and the sha256
of the source bytes). It accepts a `.shen` or a `.kl` file.

**2. Run with the compiled functions loaded at startup (~0.3s, no compile cost):**

```
./shen -precompiled hot.so        # repeatable, or comma-separated
```

The functions in the `.so` now run as compiled Go; a bad/missing `.so` just warns
and the REPL continues on the VM. `(load-native "hot.so")` does the same from
inside a session or from a script.

Measured on this kernel (darwin/arm64): a recursion-heavy benchmark runs ~1.6–1.9×
faster compiled than on the VM. The compiled code still uses the runtime's boxed
values and trampoline, so it does not reach SBCL-class native speed — it is a
precompile-the-hot-path option, not a replacement for the VM.

> **Constraint:** a `.so` and the `shen` that loads it must be built from the same
> module and Go toolchain (both from this tree). Go plugins are Linux/macOS only
> and cannot be unloaded (redefining a function means rebuilding the `.so`).

### Using precompiled code with script runners

Naming the `.so` is all it takes — including for scripts that `(load ...)` the
same file, which is what test-suite runners do:

```
make precompile FILE=path/to/hot.shen OUT=/tmp/hot.so
./shen -precompiled /tmp/hot.so script path/to/run-tests.shen
```

Shen's `load` is a definer: `(load "hot.shen")` re-runs every `define` in the
file, which rebinds exactly the symbols a plugin just bound. Left alone, that
throws the compiled versions away — and since suite runners open with
`(load "shen/world/prng.shen")`, `-precompiled` used to buy them **nothing**.

So `shen` closes the loop: **after a successful `(load F)`, any loaded plugin
that was compiled from `F` is re-engaged automatically.** The default path is the
native one; there is nothing to remember and no ordering to get right.

Freshness is decided by **content**, never by path or mtime: `make precompile`
stamps the source's sha256 into `hot.so.fns`, and a plugin is re-engaged only if
the file just loaded hashes to exactly that. Edit `hot.shen` and the hashes
diverge, so the freshly loaded (correct) VM definitions stay in force and you get
one warning on stderr telling you to rebuild. Stale compiled code never runs
silently.

| flag / env | effect |
|---|---|
| *(default)* | re-engage a matching plugin after `(load ...)` |
| `-auto-native=false` | never re-engage; `load` wins, as before |
| `SHEN_NO_AUTO_NATIVE=1` | same, for harnesses that cannot change argv |
| `SHEN_AUTO_NATIVE_VERBOSE=1` | narrate each engagement on stderr |

Nothing is printed on the happy path (a suite's stdout is often compared against
a golden file), and nothing is ever discovered from disk: only the `.so`s you
named with `-precompiled` or `(load-native ...)` are candidates.

> Measured on `bench/hot.shen` (darwin/arm64, min CPU seconds of 5 runs):
> VM only 1.33s; `-precompiled` + `(load ...)` 1.33s (no gain, the old behavior);
> auto-engaged (this default) **0.89s, 1.49×** — the same as the manual
> `(load ...)` + `(load-native ...)` ordering, without the manual step.

The manual form still works if you want explicit control (`-auto-native=false`
plus `(load-native "/tmp/hot.so")` after your loads). Big-int-heavy functions are
supported, but continue to use Shen's boxed numeric representation, so measure
the workload before and after enabling the plugin. To capture a profile for
further optimization, add `-cpuprofile FILE` (and optionally `-memprofile FILE`)
before `script`:

```
./shen -cpuprofile /tmp/prng.pprof script path/to/run-tests.shen
go tool pprof -top ./shen /tmp/prng.pprof
```

## Testing

Testing has two clearly separate tiers:

### 1. Canonical Shen kernel certification (the official suite)

This is the **official ShenOSKernel acceptance test suite** (shipped under
`kernel/tests`) — the external bar for "Certified" on
[shen-language.github.io](https://shen-language.github.io). It validates that
this port correctly implements the Shen language, independent of any tests we
wrote. It currently passes **134/134 (100%)**.

```
make certify
```

This builds the `shen` binary and runs `kernel/tests/runme.shen` end-to-end,
asserting a clean 100% pass rate. You can also run it by hand:

```
make shen
cd kernel/tests
../../shen
(load "runme.shen")
```

### 2. Our own Go unit tests (additional coverage)

These are **our** tests — they exercise the `kl` evaluator/bytecode VM and the
`shen` CLI directly, beyond what the canonical suite covers. They are fast and
do not run the certification suite.

```
make test
```

These are result-cached by Go: rerun with no source change and you'll see
`ok ... (cached)` and a sub-second finish. The certification suite is
intentionally **not** cached (`-count=1`) — it execs the freshly-built binary
against `kernel/tests`, inputs Go's cache can't track, so a cached PASS could be
stale. `make certify` always runs for real.

### Everything at once

```
make test-all   # our unit tests (cached) + the canonical certification (always run)
```

(Prefer `make test-all` over a bare `go test ./...`: the latter would cache the
certification result, which can go stale against kernel or VM changes.)

## How to bootstrap

You can just do
```
cd compiled
kl
(load-file "script.kl")
cd ..
make shen
``` 
Explanation : 

`kl` implements a simple klambda interpreter in Go, which is used to bootstrap
`shen`. `compiled/script.kl` drives the whole regeneration: it loads the KLambda
kernel, loads the KL->IR compiler (`src/compiler.shen`), compiles each module to
its intermediate representation (`*.tmp`), and emits the Go sources under
`cmd/shen/`. The full transformation path is Shen -> KL -> IR -> Go.

The vendored kernel (Tarver's S41.2 refresh — see `kernel/klambda/PROVENANCE.md`)
is **load-order dependent**: there is no `shen.initialise`, so each module runs
its own top-level `(set ...)`/`(put ... arity ...)`/`(declare ...)` forms as it
loads. `script.kl` loads the modules in upstream `install.lsp` order; do not
reorder them. After regeneration:

```
make shen
```


## Learn Shen
* [Official website of Shen](http://shenlanguage.org/)
* [Shen Community Wiki](https://github.com/Shen-Language/wiki/wiki)

## License

- Shen, Copyright © 2010-2015 Mark Tarver - [License](http://www.shenlanguage.org/license.pdf).
- shen-go, Copyright © 2017-2022 Arthur Mao under [BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause).
