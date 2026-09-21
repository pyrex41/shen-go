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

## Optional: compile hot files to Go

The REPL and `define` run on the bytecode VM. For a hot file you can AOT-compile to a plugin and load it at startup (Linux/macOS; same Go module and toolchain as `./shen`):

```
make precompile FILE=bench/hot.shen OUT=hot.so
./shen -precompiled hot.so
```

After `(load F)`, a plugin compiled from the same bytes is re-engaged automatically. Details and flags: `SPIKE-compile-to-go.md`.

Profile a run with `./shen -cpuprofile FILE script …`.

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
