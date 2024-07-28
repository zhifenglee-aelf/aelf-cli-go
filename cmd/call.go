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
	"encoding/hex"
	"fmt"

	"aelf.com/aelf-cli/keyring"
	"github.com/AElfProject/aelf-sdk.go/client"
	"github.com/spf13/cobra"
)

var Endpoint string = "https://aelf-public-node.aelf.io"

//aelf-command call 2JYYALQT1dwX6izcnedRJL3ZNvaHNjEYFfPg4tk59wtau8mi77 -e https://tdvw-test-node.aelf.io GetOwner

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call [contract-address|contract-name] [method] [params]",
	Short: "Call a read-only method on a contract.",
	Long:  `Call a read-only method on a contract. For example: aelf-command call ASh2Wt7nSEmYqnGxPPzp4pnVDU4uhj1XW9Se5VeZcX2UDdyjx -a 2CpoZzMKR912WvbF3kwG2vcyZQzDC4i52cMN1kad128gRY6Kwj -e https://tdvw-test-node.aelf.io GetBalance`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}

		return nil
		//return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call called")
		fmt.Println("Endpoint: ", Endpoint)

		wallet, _ := keyring.LoadWallet()

		var aelf = client.AElfClient{
			Host:       cmd.Flag("endpoint").Value.String(),
			Version:    "1.0",
			PrivateKey: wallet.PrivateKey,
		}

		empty := []byte{}
		txn, error := aelf.CreateTransaction(wallet.Address, args[0], args[1], empty)
		if error != nil {
			fmt.Println(error)
			return
		}

		signedTxn, error := aelf.SignTransaction(wallet.PrivateKey, txn)
		if error != nil {
			fmt.Println(error)
			return
		}

		result, error := aelf.ExecuteTransaction(hex.EncodeToString(signedTxn))
		if error != nil {
			fmt.Println(error)
			return
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(callCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	callCmd.PersistentFlags().StringVarP(&Endpoint, "endpoint", "e", "https://aelf-public-node.aelf.io", "endpoint to connect to")
}
