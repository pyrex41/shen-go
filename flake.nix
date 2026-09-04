{
  # Pinned Go toolchain for shen-go. Two halves are load-bearing:
  #
  #   * `pkgs.go_1_27` -- at the locked nixpkgs rev the default `pkgs.go` attr is
  #     1.26.5, while go.mod declares `go 1.27`. The explicit versioned attr is
  #     what actually supplies 1.27.
  #   * `env.GOTOOLCHAIN = "local"` -- without it, a `go.mod` directive above the
  #     nix-provided compiler makes GOTOOLCHAIN=auto silently *download* a
  #     toolchain and build with that instead, which makes the pin decorative.
  #
  # `pkgs.git` is required because the nixpkgs darwin stdenv shadows `xcrun`,
  # which breaks the /usr/bin/git shim; without a real git, `go build` of a main
  # package fails with "error obtaining VCS status".
  description = "shen-go development environment";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  outputs = { nixpkgs, ... }: let systems = [ "aarch64-darwin" "aarch64-linux" "x86_64-linux" ]; each = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system}); in {
    # NOTE: buildEnv has no `env` argument, so these package outputs carry the
    # 1.27 compiler but NOT GOTOOLCHAIN=local. Consumers of `packages.toolchain`
    # must set GOTOOLCHAIN=local themselves to get a genuinely pinned build; the
    # devShell below is the fully-pinned entry point.
    packages = each (pkgs: { toolchain = pkgs.buildEnv { name = "shen-go-toolchain"; paths = [ pkgs.go_1_27 pkgs.git ]; }; default = pkgs.buildEnv { name = "shen-go-toolchain"; paths = [ pkgs.go_1_27 pkgs.git ]; }; });
    devShells = each (pkgs: { default = pkgs.mkShell { packages = [ pkgs.go_1_27 pkgs.git ]; env.GOTOOLCHAIN = "local"; }; });
  };
}
