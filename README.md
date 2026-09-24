# Shen/Go

A [Go](https://go.dev) port of [Shen](https://shen-language.github.io/) — pattern matching, optional types, Prolog, and a compiler-compiler on a single binary.

## Build and run

Go 1.27 or newer.

```
make shen
./shen                 # REPL
./shen eval -e '(+ 1 2)'
./shen script path/to/file.shen
```

Windows: `make shen-exe`. The binary has no runtime dependencies.

A standard library (filter, mapc, fold, sort, …) loads at startup from Tarver's StLib. Set `SHEN_NO_STDLIB=1` to skip it. See `cmd/shen/stlib/PROVENANCE.md`.

## Tests

Three layers, three questions:

| Command | Question |
|---------|----------|
| `make certify` | Does this port implement Shen? Official kernel suite, **134/134**. |
| `make test` | Do our VM/CLI internals work? Fast Go tests; skip cert with `-short`. |
| `make test-all` | Both of the above. Prefer this over `go test ./...` (that would cache cert). |

**Do the ports agree?** That suite lives in [Bifrost](https://github.com/pyrex41/bifrost), not here. From a sibling checkout:

```
cd ../bifrost
go run . --impls shen-go
```

Kernel 134 stays in this repo. Bifrost is the cross-port corpus (predicates, equality, tuples, property-vector, …).

**Do the natives match the kernel?** `kl.InstallKernelFast` rebinds 58 kernel functions to Go. [`kl/equiv.json`](kl/equiv.json) is the audited table. It holds one row per rebound symbol, with the native, its arity, the globals it touches, a `verified` verdict, and the differential cases behind that verdict in `inputs`. Every case is run twice, once against the kernel's own KL definition and once against the native, and the two results must agree. `go test ./kl -run TestEquiv` rebuilds the table and fails if the committed file differs; `EQUIV_WRITE=1` rewrites it. `kl equiv-check kl/equiv.json` replays the stored cases on any host, prints exactly one `equiv NAME ok|FAIL` line per row plus a summary, and exits non-zero on a FAIL. Where a native and the kernel disagree the row is marked unverified rather than the native changed; the `reason`, `disagreement` and `note` fields say what differed. The top-level `harness` object documents what the verdicts mean and the vocabulary the cases use.

## Optional: compile hot files to Go

The REPL and `define` run on the bytecode VM. For a hot file you can AOT-compile to a plugin and load it at startup (Linux/macOS; same Go module and toolchain as `./shen`):

```
make precompile FILE=bench/hot.shen OUT=hot.so
./shen -precompiled hot.so
```

After `(load F)`, a plugin compiled from the same bytes is re-engaged automatically. Details and flags: `SPIKE-compile-to-go.md`.

Profile a run with `./shen -cpuprofile FILE script …`.

A caught or uncaught Shen error leaves nothing on stdout. To see where the interpreter recovered a raised error (expression, message and Go stack, on stderr), run with `SHEN_DEBUG_RECOVER=1`. A Go panic inside a native is a bug rather than an error path, and is always traced to stderr.

## Bootstrap

To regenerate `cmd/shen/*.go` from the KLambda kernel:

```
cd compiled
kl
(load-file "script.kl")
cd ..
make shen
```

`kl` is a small KLambda interpreter. `script.kl` loads the kernel in upstream `install.lsp` order (S42 has no `shen.initialise`; do not reorder), compiles Shen → KL → IR → Go. Provenance: `kernel/klambda/PROVENANCE.md`.

## Nix

Optional. `nix develop` or `direnv allow` for a pinned toolchain. `nix shell .#toolchain` is what [Bifrost](https://github.com/pyrex41/bifrost) composes. Nix is never required at runtime.

## Learn Shen

- [shenlanguage.org](https://www.shenlanguage.org/)
- [Community wiki](https://github.com/Shen-Language/wiki/wiki)

## License

- Shen © 2010–2015 Mark Tarver — [license](http://www.shenlanguage.org/license.pdf)
- shen-go © 2017–2022 Arthur Mao — [BSD 3-Clause](http://opensource.org/licenses/BSD-3-Clause)
