# ShenOSKernel provenance

This directory vendors the 23-file community ShenOSKernel 43.4 KLambda kernel, byte-for-byte from the
`ShenOSKernel-43.4.tar.gz` release asset of Shen-Language/shen-sources. `(version)` is `"43.4"`. The Go port keeps its compiler in `src/compiler.shen`; `compiler.kl` is not part of the community release archive and is not vendored here.

- Upstream release: `shen-43.4` (tag object `40fd335893cf7caf4c3a285e122ea0de06696430`, commit `f34b307b490880eedcfe1f5e5d3cc24ca3d2602e`)
- Archive: https://github.com/Shen-Language/shen-sources/releases/download/shen-43.4/ShenOSKernel-43.4.tar.gz
- Archive SHA-256: `15e7972e6bf0cd65d37c2eba85750b8839d2dbadda41027eae0c88e5017a2756`
- Previous import: `shen-43.2` (archive SHA-256 `23c71508f0de24785e1f127ea4fab440081904d81c463090bc642ee840417f32`)

The generated Go modules are produced from these KLambda files by
`compiled/script.kl`. Shen/Go supplies byte-stream host operations and binds
the 43.1 public character-stream names (`char-stream?`, `read-unit-string`,
and `write-string`) while reporting that its streams are not character
streams.

The 43.1 upgrade replaces the complete prior release kernel tree. Its port upgrade
guide requires the public character-stream interfaces and normal
`shen.initialise` boot, both of which are covered by the generated kernel and
the native primitive registry.
