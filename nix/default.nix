{ lib, naersk, stdenv, clangStdenv, hostPlatform, targetPlatform, pkg-config
, libiconv, rustfmt, cargo, rustc, clippy }:

let cargoToml = (builtins.fromTOML (builtins.readFile ../Cargo.toml));

in naersk.lib."${targetPlatform.system}".buildPackage rec {
  src = ./.;
}
