# About

This project can be used as an example to understand Go Compiler. 
This package uses Go's [ast package](https://pkg.go.dev/go/ast) to retrieve Go Struct properties. It prints all struct names defined in a `.go` file.

The Go's ast package is used to explore the syntax tree representation of a Go package, and can be used to perform static analysis, code linting, metaprogramming, and anything that requires a structured interpretation of Go source code.


# Usage

To print all struct names defined in a `.go` file, run the following command:
```
$ go run main.go --fpath ./example.go 
```

> --fpath is the file path to read from. Defaults to "./example.go"

### sample output

```
 ~/go/src/itsksaurabh/getstruct/ [main] go run main.go --fpath ./example.go 
========================= List of Struct names in ./example.go file =========================
* BlockHeader
* TransactionInput
* TransactionOutput
* MerkleNode
* PublicKey
* PrivateKey
* Signature
* UTXO
* Address
```
