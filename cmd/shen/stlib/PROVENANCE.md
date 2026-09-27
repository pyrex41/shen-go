# Provenance — Shen standard library (StLib)

These sources are vendored from Mark Tarver's **S42 (2026-08-25 refresh)**
`Lib/StLib`, the same lineage as the kernel (see
`kernel/klambda/PROVENANCE.md`).

Canonical source (designated mirror of Tarver's uploads):

- Repo: `pyrex41/shen-upstream` (formerly `pyrex41/shen-s41.1`; old URLs redirect)
- Tag: `s42-pristine-20260825`
- Archive SHA-256: `30abdc7e5a1e27b7a20109c1ed141e4712885e31f24d9710d16415fbbd4dfb23`
- Files vendored from `Lib/StLib/` in that tag, byte-identical.

## Why this exists

The Shen **kernel** ships no standard library — `map`, `append`, `reverse`,
`element?` etc. are kernel functions, but `filter`, `mapc`, `take`/`drop`,
`foldl`/`foldr`, `sort`, and the Maths/Strings/Vectors/Tuples/Symbols helpers
live only in StLib. Before this, shen-go shipped **no** stdlib at all (the
community `stlib.kl` was vendored historically but never actually booted). This
directory is compiled into the image at startup (see `cmd/shen/stlib.go`) so those
functions work out of the box.

## What is loaded, and in what order

`cmd/shen/stlib.go` uses the subset and order of upstream `install.shen`
(`stlibInstallOrder`): Symbols → Maths (macros, then the `.dtype`s and their
sources) → Lists → Strings → Vectors → IO → Tuples → `package-stlib.shen`, then
declares the `stlib` externals as system functions (the tail of `install.shen`).
`python3 scripts/generate-stlib.py` compiles this ordered set to
`cmd/shen/stlibcompiled/` and snapshots the source digest, arities, package
metadata, and six macro forms. Startup replays the reader effects before
exposing the compiled functions. A source-digest mismatch falls back to the
interpreted load; `SHEN_STDLIB_INTERPRETED=1` selects that path explicitly.
Both paths turn type-checking off. Set `SHEN_NO_STDLIB=1` to skip StLib entirely.

## Files present but NOT loaded by default

The full `Lib/StLib` tree is vendored for provenance/source-of-truth. These are
present but not in the default boot set (they are optional/extra upstream
modules, matching upstream `install.shen`, which also omits them):
`Calendar/date.shen`, `Data/data.shen`, `Maths/r.shen`, `Strings/regex.shen`,
`Strings/smartmem.shen`, and the pre-built `Lists/lists.shen.kl`.

## Patches

None. The sources are byte-identical to the mirror tag; no upstream StLib source
required modification to load on shen-go.
