{
  description = "A very basic flake";

  inputs = {
    nixpkgs = { url = "github:NixOS/nixpkgs"; };
    flake-utils.url = "github:numtide/flake-utils";
    nix-filter.url = "github:numtide/nix-filter";
    crane = {
      url = "github:ipetkov/crane";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, crane, nix-filter }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs =  import nixpkgs {
          inherit system;
          overlays = [
            nix-filter.overlays.default
          ];
        };
        craneLib = crane.lib.${system};
      in rec {
        packages.nitro-replayer = craneLib.buildPackage {
          src = craneLib.cleanCargoSource ./nitro-replayer;
          nativeBuildInputs = [ pkgs.pkg-config ];
          buildInputs = with pkgs; [ openssl clang zstd ];
          doCheck = false;
          LD_LIBRARY_PATH = (pkgs.lib.strings.makeLibraryPath [
            pkgs.stdenv.cc.cc.lib
            pkgs.llvmPackages.libclang.lib
            pkgs.openssl.dev
          ]);
          PROTOC = "${pkgs.protobuf}/bin/protoc";
          LIBCLANG_PATH = "${pkgs.llvmPackages.libclang.lib}/lib";
          ROCKSDB_LIB_DIR = "${pkgs.rocksdb}/lib";
          RUSTFLAGS = "--cfg=RUSTC_WITHOUT_SPECIALIZATION";
          OPENSSL_NO_VENDOR = "1";
          installPhase = ''
            mkdir -p $out/lib
            mv target/release/libnitro_replayer.so $out/lib/libreplayer.so
          '';
        };
        packages.wasmvm = craneLib.buildPackage {
          src = "${
              pkgs.fetchFromGitHub {
                owner = "CosmWasm";
                repo = "wasmvm";
                rev = "753fb688ce408d17cdf6c6d7c81954d105e07c98";
                hash = "sha256-t8kOaXjAhXZcrn3jlytQlJ8gvbvTkTEw3KcxIS45DLg=";
              }
            }/libwasmvm";
          doCheck = false;
          installPhase = ''
            mkdir -p $out/lib
            mv target/release/libwasmvm.so $out/lib/libwasmvm.${
              builtins.head (pkgs.lib.strings.splitString "-" system)
            }.so
          '';
        };
        packages.darth = pkgs.buildGoModule rec {
          pname = "darth-chain";
          version = "0.1";
          src = pkgs.nix-filter {
            root = ./.;
            include = [
              "go.mod"
              "go.sum"
              "aclmapping"
              "app"
              "assets"
              "cmd"
              "loadtest"
              "oracle"
              "parallelization"
              "proto"
              "store"
              "sync"
              "testutil"
              "types"
              "utils"
              "wasmbinding"
              "x"
            ];
            exclude = [];
          };
          buildInputs = [ pkgs.pkg-config packages.nitro-replayer pkgs.stdenv.cc ];
          vendorHash = "sha256-rHwvxowpu/JCA5SsewD6kKNtrtZX8OBPU9rDLWhJ0Zw=";
          ldflags = [
            "-v -extldflags '-L${packages.nitro-replayer}/lib -lreplayer -L${packages.wasmvm}/lib'"
          ];
          LD_LIBRARY_PATH = (pkgs.lib.strings.makeLibraryPath [
            pkgs.stdenv.cc.cc.lib
            pkgs.llvmPackages.libclang.lib
          ]);
          hardeningEnable = [ "pie" ];
          dontFixup = true;
          doCheck = false;
          CGO_ENABLED = 1;
          meta = {
            mainProgram = "seid";
          };
        };
        apps.local = flake-utils.lib.mkApp {
          drv = pkgs.writeShellApplication {
            name = "darth-local";
            runtimeInputs = [ packages.darth ];
            checkPhase = "";
            text = ''
              # rm -rf ~/.sei
              # seid init demo --chain-id sei
              test_account_name=alice
              # seid keys add $test_account_name --keyring-backend test
              # seid add-genesis-account --keyring-backend test "$test_account_name" 100000000000000000000usei
              # seid gentx --keyring-backend test $test_account_name 70000000000000000000usei --chain-id sei
              # sed -i''' -e 's/mode = "full"/mode = "validator"/g' "$HOME"/.sei/config/config.toml
              # sed -i''' -e 's/indexer = \["null"\]/indexer = \["kv"\]/g' "$HOME"/.sei/config/config.toml
              # KEY=$(jq '.pub_key' ~/.sei/config/priv_validator_key.json -c)
              # jq '.validators = [{}]' ~/.sei/config/genesis.json > ~/.sei/config/tmp_genesis.json
              # jq '.validators[0] += {"power":"70000000000000"}' ~/.sei/config/tmp_genesis.json > ~/.sei/config/tmp_genesis_2.json
              # jq '.validators[0] += {"pub_key":'$KEY'}' ~/.sei/config/tmp_genesis_2.json > ~/.sei/config/tmp_genesis_3.json
              # mv ~/.sei/config/tmp_genesis_3.json ~/.sei/config/genesis.json && rm ~/.sei/config/tmp_genesis.json && rm ~/.sei/config/tmp_genesis_2.json
              # seid collect-gentxs
              # cat ~/.sei/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="300s"' > ~/.sei/config/tmp_genesis.json && mv ~/.sei/config/tmp_genesis.json ~/.sei/config/genesis.json
              # cat ~/.sei/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > ~/.sei/config/tmp_genesis.json && mv ~/.sei/config/tmp_genesis.json ~/.sei/config/genesis.json
              # cat ~/.sei/config/genesis.json | jq '.app_state["distribution"]["params"]["community_tax"]="0.000000000000000000"' > ~/.sei/config/tmp_genesis.json && mv ~/.sei/config/tmp_genesis.json ~/.sei/config/genesis.json
              seid start --chain-id sei
            '';
          };
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [ go gopls cargo clang ];
          LD_LIBRARY_PATH = (pkgs.lib.strings.makeLibraryPath [
            pkgs.stdenv.cc.cc.lib
            pkgs.llvmPackages.libclang.lib
          ]);
          PROTOC = "${pkgs.protobuf}/bin/protoc";
          LIBCLANG_PATH = "${pkgs.llvmPackages.libclang.lib}/lib";
          ROCKSDB_LIB_DIR = "${pkgs.rocksdb}/lib";
        };
      });
}
