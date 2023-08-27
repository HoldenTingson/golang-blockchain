# Golang Blockchain

A simple blockchain implementation in golang

# How to Use

1. Start the blockchain server

```Go
$ cd blockchain_server
$ go run main.go blockchain_server.go -port PORT
```

2. Start the wallet server

```Go
$ cd wallet_server
$ go run main.go wallet_server.go -port PORT
```

3. Use API endpoints on the Blockchain Server to interact with the blockchain, such as mining blocks and getting the blockchain detail.

4. Use API endpoints on the Wallet Server to retrieve user wallet information, get the wallet balance, and create bitcoin transactions.

# Configuration

By default, the blockchain server runs on port 5000, and the wallet server runs on port 8080. You can change these ports by specifying the port number using -port.

# Requirements

* github.com/btcsuite/btcutil
* golang.org/x/crypto
