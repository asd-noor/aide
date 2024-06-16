{
  description = "Go Functional Programming Package DevShell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
	  hardeningDisable = [ "fortify" ];
          buildInputs = with pkgs; [
	    go_1_21
	  ];
        };
      }
    );
}
