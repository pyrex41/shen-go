# Provenance — Shen standard library (StLib)

These sources are vendored from Mark Tarver's **S41.2 (2026-07-11 refresh)**
`Lib/StLib`, the same lineage as the kernel (see
`kernel/klambda/PROVENANCE.md`).

Canonical source (designated mirror of Tarver's uploads):

- Repo: `pyrex41/shen-upstream` (formerly `pyrex41/shen-s41.1`; old URLs redirect)
- Tag: `s41.2-pristine-20260711`
- Commit: `11fc51b`
- Files vendored from `Lib/StLib/` in that tag, byte-identical.

## Why this exists

The Shen **kernel** ships no standard library — `map`, `append`, `reverse`,
`element?` etc. are kernel functions, but `filter`, `mapc`, `take`/`drop`,
`foldl`/`foldr`, `sort`, and the Maths/Strings/Vectors/Tuples/Symbols helpers
live only in StLib. Before this, shen-go shipped **no** stdlib at all (the
community `stlib.kl` was vendored historically but never actually booted). This
directory is loaded into the image at startup (see `cmd/shen/stlib.go`) so those
functions work out of the box.

## What is loaded, and in what order

`cmd/shen/stlib.go` loads the subset and order of upstream `install.shen`
(`stlibInstallOrder`): Symbols → Maths (macros, then the `.dtype`s and their
sources) → Lists → Strings → Vectors → IO → Tuples → `package-stlib.shen`, then
declares the `stlib` externals as system functions (the tail of `install.shen`).
It is loaded with **type-checking off** — the sources are known-good and the
kernel would otherwise type-check the whole library on every startup (~0.37s vs
~0.12s). Set `SHEN_NO_STDLIB=1` to skip loading entirely.

## Files present but NOT loaded by default

The full `Lib/StLib` tree is vendored for provenance/source-of-truth. These are
present but not in the default boot set (they are optional/extra upstream
modules, matching upstream `install.shen`, which also omits them):
`Calendar/date.shen`, `Data/data.shen`, `Maths/r.shen`, `Strings/regex.shen`,
`Strings/smartmem.shen`, and the pre-built `Lists/lists.shen.kl`.

## Patches

None. The sources are byte-identical to the mirror tag; no upstream StLib source
required modification to load on shen-go.
