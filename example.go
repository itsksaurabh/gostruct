package main

import "math/big"

type BlockHeader struct {
	Version    int32
	PrevBlock  []byte
	MerkleRoot []byte
	Timestamp  uint32
	Difficulty uint32
	Nonce      uint32
}

type TransactionInput struct {
	Txid      []byte
	Vout      int
	ScriptSig []byte
}

type TransactionOutput struct {
	Value        int
	ScriptPubKey []byte
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

type PublicKey struct {
	X, Y *big.Int
}

type PrivateKey struct {
	D *big.Int
}

type Signature struct {
	R, S *big.Int
}

type UTXO struct {
	TxID   []byte
	Index  int
	Output *TransactionOutput
}

type Address struct {
	Version   byte
	Hash      []byte
	Checksum  []byte
	PublicKey PublicKey
}
