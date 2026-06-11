# Provenance

The `kernel/` tree (including this directory) is vendored verbatim from the
official ShenOSKernel **41.2** release:

- Release tag: `shen-41.2`
- URL: https://github.com/Shen-Language/shen-sources/releases/tag/shen-41.2
- Archive: `ShenOSKernel-41.2.zip`
- SHA-256: `49f1b85d02348d9b3ebc461570c5c56cc066270ab81e35d5257625fb9d17fe82`

Every file under `kernel/` is **byte-identical** to the corresponding file in
that archive, with exactly one addition (see below). This was verified by
re-extracting the archive (after checking its SHA-256) and running a recursive
`diff` against `kernel/`.

## Exception: `compiler.kl`

`klambda/compiler.kl` is **not** part of the ShenOSKernel release. It is a
generated artifact of [shen-cl](https://github.com/Shen-Language/shen-cl)
(the KLambda output of compiling shen-cl's `compiler.shen`). The copy here is
the fresh shen-cl generation against ShenOSKernel-41.2, taken byte-identical
from `ShenOSKernel-41.2/klambda/compiler.kl` as produced by a current shen-cl
build. Because it is regenerated rather than released, its gensym variable
numbering (`V412` vs `V414`, etc.) can differ between otherwise-equivalent
generations.

## Note: `extension-programmable-pattern-matching.kl`

New in 41.2. It is an **opt-in** extension: it is vendored here for
completeness but is deliberately NOT part of the boot/load lists
(`compiled/script.kl`, `compiled/precompile.kl`, `cmd/kl/runtests.kl`); the
canonical boot remains the same module set as 41.1.
