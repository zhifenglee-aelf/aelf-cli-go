module aelf.com/aelf-cli

go 1.22.3

require (
	github.com/AElfProject/aelf-sdk.go v1.3.0
	github.com/btcsuite/btcutil v1.0.1
	github.com/foxnut/go-hdwallet v0.0.0-20200602072018-8db9c730e77c
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.8.1
	github.com/zalando/go-keyring v0.2.5
)

require (
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cpacia/bchutil v0.0.0-20181003130114-b126f6a35b6c // indirect
	github.com/danieljoos/wincred v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ethereum/go-ethereum v1.9.13 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/haltingstate/secp256k1-go v0.0.0-20151224084235-572209b26df6 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	golang.org/x/crypto v0.0.0-20200427165652-729f1e841bcc // indirect
	golang.org/x/sys v0.8.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace aelf.com/aelf-cli/aelfwallet => ./aelfwallet

replace aelf.com/aelf-cli/keyring => ./keyring
