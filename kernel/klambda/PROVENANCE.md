# Provenance

The KLambda kernel under this directory is vendored from Mark Tarver's
**S41.2 (2026-07-11 refresh)**.

Canonical source (the designated mirror of Tarver's uploads):

- Repo: `pyrex41/shen-s41.1` (private)
- Tag: `s41.2-pristine-20260711`
- Commit: `11fc51b`
- Files vendored from `KLambda/*.kl` in that tag; verified byte-identical.

Upstream origin (secondary detail — what the mirror imported):

- URL: https://www.shenlanguage.org/Download/S41.2.zip
- Last-Modified: **2026-07-11** (re-upload; see caveat below)
- Archive SHA-256: `51becbfd60fa8c93c3f8ae5b20b948eaa84c4b1d14ad2f5d2a056002a53ee836`

## Caveat: reused version number, different lineage

Upstream **reused the "41.2" version number for a restructured kernel**. This is
NOT the community ShenOSKernel-41.2 (github.com/Shen-Language/shen-sources, tag
`shen-41.2`) that this repository previously vendored byte-identically. It is
Tarver's own reference-implementation kernel — the one his CommonLisp port
(`backend.lsp` / `install.lsp`) builds from. The two share most `defun`s but
differ structurally in the porter-facing initialisation layer:

- **No `shen.initialise` wrapper.** The community kernel (since 22.0) collected
  every top-level `(set ...)`/`(put ... arity ...)` form into a single
  `shen.initialise` function in `init.kl`, making load order irrelevant. Tarver's
  kernel keeps those forms **inline** in `declarations.kl` (globals + arity
  table + external-symbols) and `types.kl` (datatype `(declare ...)` forms), so
  the kernel is **load-order dependent** again. shen-go now loads the modules in
  upstream `install.lsp` order and runs each module's top-level forms as its
  generated `Main` thunk is called (see `cmd/shen/main.go:regist` and
  `compiled/script.kl`).
- **`get`/`put`/`unput` moved to `sys.kl`** and re-implemented on top of the
  standard `hash`/`limit`/`<-vector`/`vector->` primitives, using the pure-KL
  helpers `shen.change-pointer-value` / `shen.remove-pointer` (also in `sys.kl`).
  No separate dictionary layer, and no native pointer primitives are required.
- **REPL entry is `shen.shen` / `shen.loop`** (not the community `shen.repl`),
  and there is no `shen.toplevel-display-exception` — `shen.loop` prints the raw
  error string. shen-go's Go REPL loop mirrors that (`cmd/shen/main.go:repl`).

## Files vendored

15 KLambda modules from `S41/KLambda/`:

- 14 that shen-go boots and compiles to Go: `core`, `declarations`, `load`,
  `macros`, `prolog`, `reader`, `sequent`, `sys`, `t-star`, `toplevel`, `track`,
  `types`, `writer`, `yacc`.
- `backend.kl` — Tarver's KLambda->CommonLisp code generator (the `cl.*`
  functions). Vendored for completeness/provenance but **NOT** part of shen-go's
  boot or codegen: it is irrelevant to the Go runtime, which has its own
  KL->IR->Go path (`src/compiler.shen` + `codegen/`).

## Removed vs the previous (community 41.2) vendoring

- `dict.kl` — the dictionary layer was dropped upstream; `get`/`put`/`unput`
  now live in `sys.kl` (see above).
- `init.kl` — no longer exists upstream; initialisation is inline (see above).
- `stlib.kl` — the standard library is not part of this KLambda set and was
  never loaded by shen-go's boot.
- `compiler.kl` — this was a shen-cl build artifact, not part of Tarver's
  kernel; shen-go never loaded it.

## Extensions (community additions, retained)

`extension-launcher.kl`, `extension-features.kl`, `extension-expand-dynamic.kl`
and `extension-programmable-pattern-matching.kl` are Bruno Deferrari's community
extensions, not part of Tarver's kernel. They are retained on top of the vendored
kernel. Only `extension-launcher.kl` is booted (it provides the shared
`shen script/eval/repl` launcher CLI that yggdrasil stage-1 relies on); the other
three are vendored but not booted, as before. They were written against the
community kernel API and have not been re-validated against this kernel beyond
the launcher.
