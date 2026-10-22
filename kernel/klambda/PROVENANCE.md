# ShenOSKernel provenance

This directory vendors the 23-file community ShenOSKernel 43.1 KLambda kernel as built at upstream
Shen-Language/shen-sources master `ec16a9f`: the 43.1 release plus the type checker's
deref-once change (Shen-Language/shen-sources#117). `(version)` is still `"43.1"`. The Go port keeps its compiler in `src/compiler.shen`; `compiler.kl` is not part of the community release archive and is not vendored here.

- Upstream commit: `ec16a9f` (Shen-Language/shen-sources master, merge of #117), on top of tag `shen-43.1` (`715824a52b55e9c81a9501df8548ff3edd8411f5`)
- Build: `ShenOSKernel-43.1-ec16a9f.tar.gz` from https://github.com/pyrex41/shen-sources/releases/tag/upstream-ec16a9f (unmodified upstream sources, built with the repo's `make klambda`/`make release`; `kernel/VERSION` is `43.1-ec16a9f`)
- Archive SHA-256: `060066a8f84d9a2f7d41cb3b37b986a8d4653279752b739b5063cd648d9baacd`
- Previous import: the `shen-43.1` release archive, SHA-256 `900eedef44f440af78f9395ed06e689ce1604546fc309f5bc7c8bc84cc651920`

The generated Go modules are produced from these KLambda files by
`compiled/script.kl`. Shen/Go supplies byte-stream host operations and binds
the 43.1 public character-stream names (`char-stream?`, `read-unit-string`,
and `write-string`) while reporting that its streams are not character
streams.

The 43.1 upgrade replaces the complete prior release kernel tree. Its port upgrade
guide requires the public character-stream interfaces and normal
`shen.initialise` boot, both of which are covered by the generated kernel and
the native primitive registry.
