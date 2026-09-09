# Nix interpreter package

The flake retains the existing default/toolchain package and development shell,
and exposes a separate runnable interpreter package:

```sh
nix build .#shen-go
./result/bin/shen eval -e '(+ 1 2)'
nix run .#shen-go -- eval -e '(reverse [1 2 3])'
nix flake check
```

Remote consumers select `packages.${system}.shen-go`, whose executable is `shen`.
It embeds its runtime resources and can run outside this source checkout.
Package builds use Go 1.27 with `GOTOOLCHAIN=local`, `go.mod`/`go.sum` and a fixed
vendor hash. `flake.lock` pins the compiler's Nixpkgs revision.

Supported flake systems are aarch64 Darwin and aarch64/x86_64 Linux. The locked
Nixpkgs no longer supports Intel Darwin. Build checks on macOS do not establish
Linux execution evidence.

The interpreter smoke check evaluates arithmetic and an embedded standard-library
function from a temporary working directory. It is packaging evidence, separate
from interpreter conformance and downstream application validation.

The package also runs the existing CLI tests, including the native-plugin case.
That case compiles its launcher and plugin under separate build deadlines before
measuring execution, so cold compiler work does not consume the runtime budget.

Use `nix develop` for the existing Go development shell. Its default package remains
the toolchain for compatibility; downstream executable consumers must use `shen-go`.
When module dependencies change, update the package's vendor hash and rebuild.
