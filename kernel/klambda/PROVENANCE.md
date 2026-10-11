# Provenance

Everything under `kernel/` except this file is the community
**ShenOSKernel 42.2** release, vendored byte-for-byte.

- Repo: [`Shen-Language/shen-sources`](https://github.com/Shen-Language/shen-sources)
- Tag: `shen-42.2`
- Commit: `73507e069fe147aefb728d632c4ce40bc6a7b314` (2026-10-01, "Prepare Shen kernel 42.2 release")
- Release asset: `ShenOSKernel-42.2.tar.gz`
- Archive SHA-256: `1d02273003654d34ec020ae67cb214ecec410dee3350419ff807d82b652678d8`
  (the checksum pinned in shen-scheme 0.50's Makefile)
- `(version)` is `"42.2"` (set by `shen.initialise-environment` in `init.kl`).

This is the kernel upstream [shen-scheme 0.50](https://github.com/tizoc/shen-scheme/tree/v0.50)
ships, so shen-go and shen-scheme run the same KLambda.

To check the tree: unpack the tarball and run
`diff -r ShenOSKernel-42.2 kernel`; the only difference is this file.
`kernel/CHANGELOG.md` is upstream's changelog and is kept pristine; shen-go's
vendoring history is recorded below instead.

## Lineage choice: community kernel, adopted wholesale

Tracking issue pyrex41/shen-go#61 (umbrella pyrex41/bifrost#26,
pyrex41/yggdrasil#32) offered two options for moving to 42.2:

- **(a) adopt ShenOSKernel-42.2 wholesale** -- chosen;
- (b) port the 42.1/42.2 changes onto Tarver's S42.0 base as a derived kernel.

Option (a) replaces the Tarver S42 lineage that shen-go vendored from
2026-07-14 to 2026-10 (see "History" below). The structural consequences for
shen-go:

- **`shen.initialise` is back.** Every module in `klambda/` is a plain list of
  `(defun ...)` forms; all top-level state (globals, the arity table, lambda
  forms, datatype declarations) is set up by `shen.initialise` in `init.kl`.
  Load order no longer matters. shen-go runs the generated `Main` thunks, swaps
  in its native FNV-1a `hash`, installs the native kernel hot paths
  (`kl.InstallKernelFast`) and then calls `shen.initialise`
  (`cmd/shen/main.go:regist`, `kl/boot.go:BootKernel`, `compiled/script.kl`).
- **`get`/`put`/`unput` sit on the community dictionary layer** (`dict.kl`):
  `*property-vector*` is a `shen.dict` (an absvector tagged
  `shen.dictionary`, with capacity, count and alist buckets of
  `(Key . ((Attr . Value) ...))`). shen-go's native `get`/`put`/`unput`,
  `arity` and `fn` (`kl/kernelfast.go`) read and write that layout; they are
  audited against the KL definitions by `kl/equiv.json`.
- **Lambda forms live in the property dict** (`shen.lambda-form`), not in a
  `shen.*lambdatable*` alist. shen-go's native `fn` reads them on every call
  with no cache, so the 42.0 requirement that a redefinition through
  `update-lambda-table` (including removal of a stale lambda form) is seen by
  `(fn F)`, partial and dynamic application holds trivially.
- **REPL entry points are the community ones** (`shen.repl` / `shen.loop`,
  with `shen.toplevel-display-exception`). shen-go's Go REPL loop mirrors
  `shen.loop` and reports errors through `shen.toplevel-display-exception`.

## 42.1 reader / macroexpand audit

`doc/port-upgrades.md` (42.1) requires ports that override the reader,
`macroexpand` or streams to keep the source-form expansion order. shen-go's
overrides:

- `read-file` (`cmd/shen/nativereader.go`) replaces only the tokenise/parse
  stage (`<s-exprs>`, unchanged in 42.2) and hands the raw forms to the kernel's
  own `shen.process-sexprs`, so source-form handlers and `shen.macroexpand-h`
  (4 arguments) run exactly as in the kernel.
- `read`, `lineread`, `shen.read-loop`, `shen.try-parse` and `macroexpand` are
  not overridden. The Go REPL calls the kernel's `shen.read-evaluate-print`.
  Input at EOF without a trailing newline is handled by the kernel's read loop.
- The load cache (`cmd/shen/loadcache.go`) memoises `read-file` results per
  unchanged file; it wraps the same function and does not touch expansion.

The upstream source-forms extension tests (`tests/extensions/source-forms`)
pass, 27/27.

## Modules booted

17 KLambda modules are compiled to Go and booted: `sys`, `writer`, `core`,
`reader`, `declarations`, `toplevel`, `macros`, `load`, `prolog`, `sequent`,
`track`, `t-star`, `yacc`, `types`, `dict`, `init` and
`extension-launcher` (the shared `shen script/eval/repl` launcher CLI that
Yggdrasil stage 1 relies on).

Vendored but not booted, as before: `stlib.kl` (shen-go loads Tarver's StLib
from `cmd/shen/stlib`, see its PROVENANCE), `extension-features.kl`,
`extension-expand-dynamic.kl`, `extension-programmable-pattern-matching.kl`,
`extension-namespaces.kl` and `extension-type-annotations.kl`.

## History

- 2026-10: ShenOSKernel 42.2 (community), this vendoring. Replaced Tarver S42.
- 2026-08-27: Tarver S42 (2026-08-25 refresh), mirror `pyrex41/shen-upstream`
  tag `s42-pristine-20260825`, archive SHA-256
  `30abdc7e5a1e27b7a20109c1ed141e4712885e31f24d9710d16415fbbd4dfb23`.
- 2026-07-14: Tarver S41.2 (2026-07-11 refresh), replacing the community
  ShenOSKernel-41.2.
- Before that: community ShenOSKernel-41.2 and earlier, vendored byte-identically.

### Historical caveat (Tarver lineage, 2026-07 to 2026-10)

Kept for reference; it described the kernel this vendoring replaced.

> Upstream **reused the "41.2" version number for a restructured kernel**. This
> is NOT the community ShenOSKernel-41.2 (github.com/Shen-Language/shen-sources,
> tag `shen-41.2`) that this repository previously vendored byte-identically. It
> is Tarver's own reference-implementation kernel -- the one his CommonLisp port
> (`backend.lsp` / `install.lsp`) builds from. The two share most `defun`s but
> differ structurally in the porter-facing initialisation layer:
>
> - **No `shen.initialise` wrapper.** The community kernel (since 22.0)
>   collected every top-level `(set ...)`/`(put ... arity ...)` form into a
>   single `shen.initialise` function in `init.kl`, making load order
>   irrelevant. Tarver's kernel keeps those forms **inline** in
>   `declarations.kl` (globals + arity table + external-symbols) and `types.kl`
>   (datatype `(declare ...)` forms), so the kernel is **load-order dependent**
>   again. shen-go loaded the modules in upstream `install.lsp` order and ran
>   each module's top-level forms as its generated `Main` thunk was called.
> - **`get`/`put`/`unput` moved to `sys.kl`** and were re-implemented on top
>   of the standard `hash`/`limit`/`<-vector`/`vector->` primitives, using the
>   pure-KL helpers `shen.change-pointer-value` / `shen.remove-pointer`. No
>   separate dictionary layer.
> - **REPL entry is `shen.shen` / `shen.loop`** (not the community
>   `shen.repl`), and there is no `shen.toplevel-display-exception` --
>   `shen.loop` prints the raw error string.
>
> That vendoring also carried `backend.kl` (Tarver's KLambda->CommonLisp code
> generator, not booted) and community extensions on top of Tarver's kernel;
> `dict.kl`, `init.kl`, `stlib.kl` and the shen-cl `compiler.kl` artifact had
> been removed. All of that is superseded by the wholesale 42.2 tree.
