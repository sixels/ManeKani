{
  description = "TODO";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

    naersk.url = "github:nmattia/naersk/master";
    nixpkgs-mozilla = {
      url = "github:mozilla/nixpkgs-mozilla";
      flake = false;
    };

    flake-utils.url = "github:numtide/flake-utils";
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, naersk, flake-utils, nixpkgs-mozilla, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = (import nixpkgs) {
          inherit system;
          overlays = [ (import nixpkgs-mozilla) ];
        };

        toolchain = (pkgs.rustChannelOf {
          rustToolchain = ./rust-toolchain;
          sha256 = "sha256-DzNEaW724O8/B8844tt5AVHmSjSQ3cmzlU4BP90oRlY=";
        }).rust;

        # environment = { ; };

        naersk' = pkgs.callPackage naersk {
          cargo = toolchain;
          rustc = toolchain;
        };

      in rec {
        # For `nix build` & `nix run`:
        defaultPackage = naersk'.buildPackage {
          inherit (import ./nix/default.nix {
            inherit pkgs;
            inherit system;
          })
          ;
          src = ./.;

          nativeBuildInputs = with pkgs; [ pkg-config clippy ];
          buildInputs = with pkgs; [ openssl ];
        };

        # For `nix develop`
        devShell = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            toolchain
            rustfmt
            nixpkgs-fmt
            pkg-config
            clippy
            diesel-cli
          ];
          buildInputs = with pkgs; [ openssl postgresql ];

          shellHook = ''
            echo $0
            export PGDATA="$PWD/nix/pgdata" \
                   PGUSER="$USER"

            # Having issues with direnv triggering `EXIT` all the time
            # trap "'$PWD/nix/pg.sh' down" EXIT
            $PWD/nix/pg.sh up
          '';
        };
      });

}
