# ShenOSKernel provenance

This directory vendors the community ShenOSKernel 43.1 release wholesale.

- Tag: `shen-43.1`
- Commit: `715824a52b55e9c81a9501df8548ff3edd8411f5`
- Release archive SHA-256: `900eedef44f440af78f9395ed06e689ce1604546fc309f5bc7c8bc84cc651920`
- Kernel source checkout used for this import: `Shen-Language/shen-sources`

The generated Go modules are produced from these KLambda files by
`compiled/script.kl`. Shen/Go supplies byte-stream host operations and binds
the 43.1 public character-stream names (`char-stream?`, `read-unit-string`,
and `write-string`) while reporting that its streams are not character
streams.

The 43.1 upgrade replaces the complete prior kernel tree. Its port upgrade
guide requires the public character-stream interfaces and normal
`shen.initialise` boot, both of which are covered by the generated kernel and
the native primitive registry.
