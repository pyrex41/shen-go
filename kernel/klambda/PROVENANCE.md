# ShenOSKernel provenance

This directory vendors the 23-file community ShenOSKernel 43.2 KLambda kernel, byte-for-byte from the
`ShenOSKernel-43.2.tar.gz` release asset of Shen-Language/shen-sources. `(version)` is `"43.2"`. The Go port keeps its compiler in `src/compiler.shen`; `compiler.kl` is not part of the community release archive and is not vendored here.

- Upstream release: `shen-43.2` (tag object `7b7563e5`, commit `e6947f507ebc`)
- Archive: https://github.com/Shen-Language/shen-sources/releases/download/shen-43.2/ShenOSKernel-43.2.tar.gz
- Archive SHA-256: `23c71508f0de24785e1f127ea4fab440081904d81c463090bc642ee840417f32`
- Previous import: `ShenOSKernel-43.1-ec16a9f.tar.gz` (upstream `ec16a9f`, 43.1 plus the type checker deref-once change), SHA-256 `060066a8f84d9a2f7d41cb3b37b986a8d4653279752b739b5063cd648d9baacd`

The generated Go modules are produced from these KLambda files by
`compiled/script.kl`. Shen/Go supplies byte-stream host operations and binds
the 43.1 public character-stream names (`char-stream?`, `read-unit-string`,
and `write-string`) while reporting that its streams are not character
streams.

The 43.1 upgrade replaces the complete prior release kernel tree. Its port upgrade
guide requires the public character-stream interfaces and normal
`shen.initialise` boot, both of which are covered by the generated kernel and
the native primitive registry.
