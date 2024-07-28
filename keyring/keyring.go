package keyring

import (
	"log"

	"aelf.com/aelf-cli/aelfwallet"
	"github.com/zalando/go-keyring"
)

const service = "aelf-cli"

// TODO: change to use environment variable through viper
const user = "anon"

func LoadWallet() (aelfwallet.Wallet, error) {

	mnemonics, err := keyring.Get(service, user)
	if err != nil {
		log.Fatal(err)
	}

	wallet, _ := aelfwallet.CreateWalletFromMnemonic(mnemonics)

	return wallet, nil
}

func SaveWallet(wallet aelfwallet.Wallet) {
	err := keyring.Set(service, user, wallet.Mnemonic)
	if err != nil {
		log.Fatal(err)
	}
}
