{
  description = "ManeKani front-end";

  # Use the unstable nixpkgs to use the latest set of node packages
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/master";

    flake-utils.url = "github:numtide/flake-utils";
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Set the major version of Node.js
            nodejs-19_x

            nodePackages.pnpm

            nodePackages.typescript
            nodePackages.typescript-language-server

            nodePackages.tailwindcss
          ];
        };
      });
}
