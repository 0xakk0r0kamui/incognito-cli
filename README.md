[![Go Report Card](https://goreportcard.com/badge/github.com/incognitochain/incognito-cli)](https://goreportcard.com/report/github.com/incognitochain/incognito-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/incognitochain/incognito-cli/blob/main/LICENSE)

incognito-cli
=============
A command line tool for the Incognito network

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->

# Usage
<!-- usage -->
## Installation
Install to the `$GOPATH` folder.
```shell
$ go install
```
This command will install the CLI application into your `GOPATH` folder. Alternatively, you can build and install the binary file
into a desired folder by the following command.
```shell
$ go build -o PATH/TO/YOUR/FOLDER/appName
```
If you have issues with these commands, try to clean the golang module cache first.
```shell
go clean --modcache
```

## Usage
There are two options for you to run the Incognito CLI by:
1. Downloading the pre-compiled executable binary file, you can find it in the [releases](https://github.com/incognitochain/incognito-cli/releases).
2. Compiling your own executable binary file from source as in the Installation instruction above.

Then execute the binary file with the following commands.

```shell
$ incognito-cli help
NAME:
   incognito-cli - A simple CLI application for the Incognito network

USAGE:
   incognito-cli [global options] command [command options] [arguments...]

VERSION:
   v0.0.2

DESCRIPTION:
   A simple CLI application for the Incognito network. With this tool, you can run some basic functions on your computer to interact with the Incognito network such as checking balances, transferring PRV or tokens, consolidating and converting your UTXOs, transferring tokens, manipulating with the pDEX, etc.

AUTHOR:
   Incognito Devs Team

COMMANDS:
   help, h  Shows a list of commands or help for one command
   ACCOUNTS:
     balance                  Check the balance of an account.
     consolidate, csl         Consolidate UTXOs of an account.
     generateaccount, genacc  Generate a new Incognito account.
     history, hst             Retrieve the history of an account.
     keyinfo                  Print all related-keys of a private key.
     submitkey, sub           Submit an ota key to the full-node.
     utxo                     Print the UTXOs of an account.
   COMMITTEES:
     checkrewards    Get all rewards of a payment address.
     withdrawreward  Withdraw the reward of a privateKey w.r.t to a tokenID.
   PDEX:
     pdecheckprice   Check the price between two tokenIDs
     pdecontribute   Create a pDEX contributing transaction
     pdeshare        Retrieve the share amount of a pDEX pair
     pdetrade        Create a trade transaction
     pdetradestatus  Get the status of a trade
     pdewithdraw     Create a pDEX withdrawal transaction
   TRANSACTIONS:
     checkreceiver  Check if an OTA key is a receiver of a transaction.
     convert        Convert UTXOs of an account w.r.t a tokenID.
     convertall     Convert UTXOs of an account for all assets.
     send           Send an amount of PRV or token from one wallet to another wallet.

GLOBAL OPTIONS:
   --clientVersion value         version of the incclient (default: 2)
   --debug value                 whether to enable the debug mode (0 - disabled, != 0 - enabled) (default: 1)
   --host value                  custom full-node host
   --network value, --net value  network environment (mainnet, testnet, testnet1, devnet, local, custom) (default: "mainnet")
   --help, -h                    show help (default: false)
   --version, -v                 print the version (default: false)

COPYRIGHT:
   This tool is developed and maintained by the Incognito Devs Team. It is free for anyone. However, any commercial usages should be acknowledged by the Incognito Devs Team.
```
# Commands
<!-- commands -->
* [`ACCOUNTS`](#accounts)
  * [`balance`](#balance)
  * [`consolidate`](#consolidate)
  * [`generateaccount`](#generateaccount)
  * [`history`](#history)
  * [`keyinfo`](#keyinfo)
  * [`submitkey`](#submitkey)
  * [`utxo`](#utxo)
* [`COMMITTEES`](#committees)
  * [`checkrewards`](#checkrewards)
  * [`withdrawreward`](#withdrawreward)
* [`PDEX`](#pdex)
  * [`pdecheckprice`](#pdecheckprice)
  * [`pdecontribute`](#pdecontribute)
  * [`pdeshare`](#pdeshare)
  * [`pdetrade`](#pdetrade)
  * [`pdetradestatus`](#pdetradestatus)
  * [`pdewithdraw`](#pdewithdraw)
* [`TRANSACTIONS`](#transactions)
  * [`checkreceiver`](#checkreceiver)
  * [`convert`](#convert)
  * [`convertall`](#convertall)
  * [`send`](#send)
## ACCOUNTS
### balance
Check the balance of an account.
```shell
$ incognito-cli help balance
NAME:
   incognito-cli balance - Check the balance of an account.

USAGE:
   balance --privateKey PRIVATE_KEY [--tokenID TOKEN_ID]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   
```

### consolidate
This function helps consolidate UTXOs of an account. It consolidates a version of UTXOs at a time, users need to specify which version they need to consolidate. Please note that this process is time-consuming and requires a considerable amount of CPU.
```shell
$ incognito-cli help consolidate
NAME:
   incognito-cli consolidate - Consolidate UTXOs of an account.

USAGE:
   consolidate --privateKey PRIVATE_KEY [--tokenID TOKEN_ID] [--version VERSION] [--numThreads NUM_THREADS] [--enableLog ENABLE_LOG] [--logFile LOG_FILE]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

DESCRIPTION:
   This function helps consolidate UTXOs of an account. It consolidates a version of UTXOs at a time, users need to specify which version they need to consolidate. Please note that this process is time-consuming and requires a considerable amount of CPU.

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --version value, -v value           version of the transaction (1 or 2) (default: 2)
   --numThreads value                  number of threads used in this action (default: 4)
   --enableLog                         enable log for this action (default: false)
   --logFile value                     location of the log file (default: "os.Stdout")
   
```

### generateaccount
This function helps generate a new mnemonic phrase and its Incognito account.
```shell
$ incognito-cli help generateaccount
NAME:
   incognito-cli generateaccount - Generate a new Incognito account.

USAGE:
   generateaccount [--numShards NUM_SHARDS]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

DESCRIPTION:
   This function helps generate a new mnemonic phrase and its Incognito account.

OPTIONS:
   --numShards value  the number of shard (default: 8)
   
```

### history
This function helps retrieve the history of an account w.r.t a tokenID. Please note that this process is time-consuming and requires a considerable amount of CPU.
```shell
$ incognito-cli help history
NAME:
   incognito-cli history - Retrieve the history of an account.

USAGE:
   history --privateKey PRIVATE_KEY [--tokenID TOKEN_ID] [--numThreads NUM_THREADS] [--enableLog ENABLE_LOG] [--logFile LOG_FILE] [--csvFile CSV_FILE]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

DESCRIPTION:
   This function helps retrieve the history of an account w.r.t a tokenID. Please note that this process is time-consuming and requires a considerable amount of CPU.

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --numThreads value                  number of threads used in this action (default: 4)
   --enableLog                         enable log for this action (default: false)
   --logFile value                     location of the log file (default: "os.Stdout")
   --csvFile value, --csv value        the csv file location to store the history
   
```

### keyinfo
Print all related-keys of a private key.
```shell
$ incognito-cli help keyinfo
NAME:
   incognito-cli keyinfo - Print all related-keys of a private key.

USAGE:
   keyinfo --privateKey PRIVATE_KEY

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   
```

### submitkey
This function submits an otaKey to the full-node to use the full-node's cache. If an access token is provided, it will submit the ota key in an authorized manner. See https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/accounts/submit_key.md for more details.
```shell
$ incognito-cli help submitkey
NAME:
   incognito-cli submitkey - Submit an ota key to the full-node.

USAGE:
   submitkey --otaKey OTA_KEY [--accessToken ACCESS_TOKEN] [--fromHeight FROM_HEIGHT] [--isReset IS_RESET]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

DESCRIPTION:
   This function submits an otaKey to the full-node to use the full-node's cache. If an access token is provided, it will submit the ota key in an authorized manner. See https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/accounts/submit_key.md for more details.

OPTIONS:
   --otaKey value, --ota value  a base58-encoded ota key
   --accessToken value          a 64-character long hex-encoded authorized access token
   --fromHeight value           the beacon height at which the full-node will sync from (default: 0)
   --isReset                    whether the full-node should reset the cache for this ota key (default: false)
   
```

### utxo
Print the UTXOs of an account.
```shell
$ incognito-cli help utxo
NAME:
   incognito-cli utxo - Print the UTXOs of an account.

USAGE:
   utxo --privateKey PRIVATE_KEY [--tokenID TOKEN_ID]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   ACCOUNTS

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   
```

## COMMITTEES
### checkrewards
Get all rewards of a payment address.
```shell
$ incognito-cli help checkrewards
NAME:
   incognito-cli checkrewards - Get all rewards of a payment address.

USAGE:
   checkrewards --address ADDRESS

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   COMMITTEES

OPTIONS:
   --address value, --addr value  a base58-encoded payment address
   
```

### withdrawreward
Withdraw the reward of a privateKey w.r.t to a tokenID.
```shell
$ incognito-cli help withdrawreward
NAME:
   incognito-cli withdrawreward - Withdraw the reward of a privateKey w.r.t to a tokenID.

USAGE:
   withdrawreward --privateKey PRIVATE_KEY [--address ADDRESS] [--tokenID TOKEN_ID] [--version VERSION]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   COMMITTEES

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --address value, --addr value       the payment address of a candidate (default: the payment address of the privateKey)
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --version value, -v value           version of the transaction (1 or 2) (default: 2)
   
```

## PDEX
### pdecheckprice
This function checks the price of a pair of tokenIds. It must be supplied with the selling amount since the pDEX uses the AMM algorithm.
```shell
$ incognito-cli help pdecheckprice
NAME:
   incognito-cli pdecheckprice - Check the price between two tokenIDs

USAGE:
   pdecheckprice --sellTokenID SELL_TOKEN_ID --buyTokenID BUY_TOKEN_ID --sellingAmount SELLING_AMOUNT

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   PDEX

DESCRIPTION:
   This function checks the price of a pair of tokenIds. It must be supplied with the selling amount since the pDEX uses the AMM algorithm.

OPTIONS:
   --sellTokenID value    ID of the token to sell
   --buyTokenID value     ID of the token to buy
   --sellingAmount value  the amount of sellTokenID wished to sell (default: 0)
   
```

### pdecontribute
This function creates a pDEX contributing transaction. See more about this transaction: https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/pdex/contribute.md
```shell
$ incognito-cli help pdecontribute
NAME:
   incognito-cli pdecontribute - Create a pDEX contributing transaction

USAGE:
   pdecontribute --privateKey PRIVATE_KEY --pairId PAIR_ID [--tokenID TOKEN_ID] --amount AMOUNT [--version VERSION]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   PDEX

DESCRIPTION:
   This function creates a pDEX contributing transaction. See more about this transaction: https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/pdex/contribute.md

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --pairId value                      the ID of the contributing pair (see https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/pdex/contribute.md)
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --amount value, --amt value         the amount of the action (default: 0)
   --version value, -v value           version of the transaction (1 or 2) (default: 2)
   
```

### pdeshare
This function returns the share amount of a user within a pDEX pair.
```shell
$ incognito-cli help pdeshare
NAME:
   incognito-cli pdeshare - Retrieve the share amount of a pDEX pair

USAGE:
   pdeshare --address ADDRESS --tokenID1 TOKEN_ID_1 [--tokenID2 TOKEN_ID_2]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   PDEX

DESCRIPTION:
   This function returns the share amount of a user within a pDEX pair.

OPTIONS:
   --address value, --addr value  a base58-encoded payment address
   --tokenID1 value               ID of the first token
   --tokenID2 value               ID of the second token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   
```

### pdetrade
This function creates a trade transaction on the pDEX.
```shell
$ incognito-cli help pdetrade
NAME:
   incognito-cli pdetrade - Create a trade transaction

USAGE:
   pdetrade --privateKey PRIVATE_KEY --sellTokenID SELL_TOKEN_ID --buyTokenID BUY_TOKEN_ID --sellingAmount SELLING_AMOUNT [--minAcceptAmount MIN_ACCEPT_AMOUNT] [--tradingFee TRADING_FEE]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   PDEX

DESCRIPTION:
   This function creates a trade transaction on the pDEX.

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --sellTokenID value                 ID of the token to sell
   --buyTokenID value                  ID of the token to buy
   --sellingAmount value               the amount of sellTokenID wished to sell (default: 0)
   --minAcceptAmount value             the minimum acceptable amount of buyTokenID wished to receive (default: 0)
   --tradingFee value                  the trading fee (measured in nano PRV) (default: 0)
   
```

### pdetradestatus
This function returns the status of a trade (1: successful, 2: failed). If a `not found` error occurs, it means that the trade has not been acknowledged by the beacon chain. Just wait and check again later.
```shell
$ incognito-cli help pdetradestatus
NAME:
   incognito-cli pdetradestatus - Get the status of a trade

USAGE:
   pdetradestatus --txHash TX_HASH

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   PDEX

DESCRIPTION:
   This function returns the status of a trade (1: successful, 2: failed). If a `not found` error occurs, it means that the trade has not been acknowledged by the beacon chain. Just wait and check again later.

OPTIONS:
   --txHash value  the transaction hash
   
```

### pdewithdraw
This function creates a transaction withdrawing an amount of `shared` from the pDEX. See more about this transaction: https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/pdex/withdrawal.md
```shell
$ incognito-cli help pdewithdraw
NAME:
   incognito-cli pdewithdraw - Create a pDEX withdrawal transaction

USAGE:
   pdewithdraw --privateKey PRIVATE_KEY --amount AMOUNT --tokenID1 TOKEN_ID_1 [--tokenID2 TOKEN_ID_2] [--version VERSION]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   PDEX

DESCRIPTION:
   This function creates a transaction withdrawing an amount of `shared` from the pDEX. See more about this transaction: https://github.com/incognitochain/go-incognito-sdk-v2/blob/master/tutorials/docs/pdex/withdrawal.md

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --amount value, --amt value         the amount of the action (default: 0)
   --tokenID1 value                    ID of the first token
   --tokenID2 value                    ID of the second token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --version value, -v value           version of the transaction (1 or 2) (default: 2)
   
```

## TRANSACTIONS
### checkreceiver
This function checks if an OTA key is a receiver of a transaction. If so, it will try to decrypt the received outputs and return the receiving info.
```shell
$ incognito-cli help checkreceiver
NAME:
   incognito-cli checkreceiver - Check if an OTA key is a receiver of a transaction.

USAGE:
   checkreceiver --txHash TX_HASH --otaKey OTA_KEY [--readonlyKey READONLY_KEY]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   TRANSACTIONS

DESCRIPTION:
   This function checks if an OTA key is a receiver of a transaction. If so, it will try to decrypt the received outputs and return the receiving info.

OPTIONS:
   --txHash value                   the transaction hash
   --otaKey value, --ota value      a base58-encoded ota key
   --readonlyKey value, --ro value  a base58-encoded read-only key
   
```

### convert
This function helps convert UTXOs v1 of a user to UTXO v2 w.r.t a tokenID. Please note that this process is time-consuming and requires a considerable amount of CPU.
```shell
$ incognito-cli help convert
NAME:
   incognito-cli convert - Convert UTXOs of an account w.r.t a tokenID.

USAGE:
   convert --privateKey PRIVATE_KEY [--tokenID TOKEN_ID] [--numThreads NUM_THREADS] [--enableLog ENABLE_LOG] [--logFile LOG_FILE]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   TRANSACTIONS

DESCRIPTION:
   This function helps convert UTXOs v1 of a user to UTXO v2 w.r.t a tokenID. Please note that this process is time-consuming and requires a considerable amount of CPU.

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --numThreads value                  number of threads used in this action (default: 4)
   --enableLog                         enable log for this action (default: false)
   --logFile value                     location of the log file (default: "os.Stdout")
   
```

### convertall
This function helps convert UTXOs v1 of a user to UTXO v2 for all assets. It will automatically check for all UTXOs v1 of all tokens and convert them. Please note that this process is time-consuming and requires a considerable amount of CPU.
```shell
$ incognito-cli help convertall
NAME:
   incognito-cli convertall - Convert UTXOs of an account for all assets.

USAGE:
   convertall --privateKey PRIVATE_KEY [--numThreads NUM_THREADS] [--logFile LOG_FILE]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   TRANSACTIONS

DESCRIPTION:
   This function helps convert UTXOs v1 of a user to UTXO v2 for all assets. It will automatically check for all UTXOs v1 of all tokens and convert them. Please note that this process is time-consuming and requires a considerable amount of CPU.

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --numThreads value                  number of threads used in this action (default: 4)
   --logFile value                     location of the log file (default: "os.Stdout")
   
```

### send
This function sends an amount of PRV or token from one wallet to another wallet. By default, it used 100 nano PRVs to pay the transaction fee.
```shell
$ incognito-cli help send
NAME:
   incognito-cli send - Send an amount of PRV or token from one wallet to another wallet.

USAGE:
   send --privateKey PRIVATE_KEY --address ADDRESS --amount AMOUNT [--tokenID TOKEN_ID] [--fee FEE] [--version VERSION]

   OPTIONAL flags are denoted by a [] bracket.

CATEGORY:
   TRANSACTIONS

DESCRIPTION:
   This function sends an amount of PRV or token from one wallet to another wallet. By default, it used 100 nano PRVs to pay the transaction fee.

OPTIONS:
   --privateKey value, --prvKey value  a base58-encoded private key
   --address value, --addr value       a base58-encoded payment address
   --amount value, --amt value         the amount of the action (default: 0)
   --tokenID value                     ID of the token (default: "0000000000000000000000000000000000000000000000000000000000000004")
   --fee value                         the PRV amount for paying the transaction fee (default: 100)
   --version value, -v value           version of the transaction (1 or 2) (default: 2)
   
```

<!-- commandsstop -->
