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
plus `hot.so.fns` (an arity manifest). It accepts a `.shen` or a `.kl` file.

**2. Run with the compiled functions loaded at startup (~0.3s, no compile cost):**

```
./shen -precompiled hot.so        # repeatable, or comma-separated
```

The functions in the `.so` now run as compiled Go; a bad/missing `.so` just warns
and the REPL continues on the VM. `(load-native "hot.so")` does the same from
inside a session.

Measured on this kernel (darwin/arm64): a recursion-heavy benchmark runs ~1.6–1.9×
faster compiled than on the VM. The compiled code still uses the runtime's boxed
values and trampoline, so it does not reach SBCL-class native speed — it is a
precompile-the-hot-path option, not a replacement for the VM.

> **Constraint:** a `.so` and the `shen` that loads it must be built from the same
> module and Go toolchain (both from this tree). Go plugins are Linux/macOS only
> and cannot be unloaded (redefining a function means rebuilding the `.so`).

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

`kl` implement a simple klambda interpreter in Go, which can be used to bootstrap `shen`

```
(load-file "../kernel/klambda/toplevel.kl")
(load-file "../kernel/klambda/core.kl")
(load-file "../kernel/klambda/sys.kl")
(load-file "../kernel/klambda/sequent.kl")
(load-file "../kernel/klambda/yacc.kl")
(load-file "../kernel/klambda/reader.kl")
(load-file "../kernel/klambda/prolog.kl")
(load-file "../kernel/klambda/track.kl")
(load-file "../kernel/klambda/load.kl")
(load-file "../kernel/klambda/writer.kl")
(load-file "../kernel/klambda/macros.kl")
(load-file "../kernel/klambda/declarations.kl")
(load-file "../kernel/klambda/t-star.kl")
(load-file "../kernel/klambda/types.kl")
(load-file "../kernel/klambda/dict.kl")
(load-file "../kernel/klambda/init.kl")
(shen.initialise)
```

`shen` source files is generated from the `.kl` files. The full transformation path is Shen -> KL -> IR -> Go.

The file `src/compiler.shen` is a transpiler from KL to an intermediate representation(IR), load it:

```
(load "../src/compiler.shen")
```

Compile the klambda to the intermediate representation:

```
(set *maximum-print-sequence-size* 100000)
(compile-file "../kernel/klambda/sys.kl" "sys.tmp")
(compile-file "../kernel/klambda/writer.kl" "writer.tmp")
(compile-file "../kernel/klambda/core.kl" "core.tmp")
(compile-file "../kernel/klambda/reader.kl" "reader.tmp")
(compile-file "../kernel/klambda/declarations.kl" "declarations.tmp")
(compile-file "../kernel/klambda/toplevel.kl" "toplevel.tmp")
(compile-file "../kernel/klambda/macros.kl" "macros.tmp")
(compile-file "../kernel/klambda/load.kl" "load.tmp")
(compile-file "../kernel/klambda/prolog.kl" "prolog.tmp")
(compile-file "../kernel/klambda/sequent.kl" "sequent.tmp")
(compile-file "../kernel/klambda/track.kl" "track.tmp")
(compile-file "../kernel/klambda/t-star.kl" "t-star.tmp")
(compile-file "../kernel/klambda/yacc.kl" "yacc.tmp")
(compile-file "../kernel/klambda/types.kl" "types.tmp")
(compile-file "../kernel/klambda/dict.kl" "dict.tmp")
(compile-file "../kernel/klambda/init.kl" "init.tmp")
```

And generate the Go files from the intermediate representation:

Use `compiled/bctogo.shen` to generate the Go files from the intermediate representation.

Now the shen source files are available, built it:

```
make shen
```


## Learn Shen
* [Official website of Shen](http://shenlanguage.org/)
* [Shen Community Wiki](https://github.com/Shen-Language/wiki/wiki)

## License

- Shen, Copyright © 2010-2015 Mark Tarver - [License](http://www.shenlanguage.org/license.pdf).
- shen-go, Copyright © 2017-2022 Arthur Mao under [BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause).
