# AElf CLI

A command-line interface (CLI) tool for interacting with the AElf blockchain.

## Description

This project provides a CLI tool to interact with the AElf blockchain. It allows users to create, sign, and execute transactions using their wallet credentials.

## Installation

To install the CLI tool, follow these steps:

1. Clone the repository:
    ```sh
    git clone git@github.com:zhifenglee-aelf/aelf-cli-go.git
    cd aelf-cli-go
    ```

2. Build the project:
    ```sh
    go build -o aelf
    ```

3. Run the CLI tool:
    ```sh
    ./aelf
    ```

## Usage

The CLI tool provides various commands to interact with the AElf blockchain. 

Below is an example of how to use the `create` command:

```sh
./aelf create
```

This creates a wallet and stores the wallet on the user's local system keyring.

```sh
./aelf wallet
```

This shows the current wallet that will be used for transactions.
