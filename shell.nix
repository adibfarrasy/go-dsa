with import <nixpkgs> {};
let 
in
mkShell {
  buildInputs = [
    # go
  ];

  NIX_LD_LIBRARY_PATH = lib.makeLibraryPath [
    stdenv.cc.cc
    openssl
    # ...
  ];
  NIX_LD = lib.fileContents "${stdenv.cc}/nix-support/dynamic-linker";

  shell = zsh;
  shellHook = ''
    if [ -f config/.env ]; then
      set -a
      source config/.env
      set +a
    fi

    source $HOME/.zshrc
  '';
}
