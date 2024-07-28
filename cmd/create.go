/*
Copyright © 2024 Lee Zhi Feng <zifcrypto.eth@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"aelf.com/aelf-cli/aelfwallet"
	"aelf.com/aelf-cli/keyring"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new wallet",
	Long:  `Creates an aelf HD Wallet which will be saved onto your keyring. For example: aelf create`,
	Run: func(cmd *cobra.Command, args []string) {
		wallet, error := aelfwallet.CreateWallet()
		if error != nil {
			fmt.Println(error)
			return
		}

		fmt.Println("Mnemonic: ", wallet.Mnemonic)
		fmt.Println("Address: ", wallet.Address)
		fmt.Println("Public Key: ", wallet.PublicKey)
		fmt.Println("Private Key: ", wallet.PrivateKey)

		prompt := promptui.Select{
			Label: "Save in keyring?",
			Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == "Yes" {
			keyring.SaveWallet(wallet)
			fmt.Printf("Saved in keyring.\n")
		} else {
			fmt.Printf("You chose not to save.\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
