package blockchain

import (
	"crypto/ecdsa"

	"time"
)

type Transaction struct {
	TransactionHash string
	From            string
	Recipient       string
	Amount          int
	TimeStamp       time.Time
	Nonce           int
	Signature       string
}

type Block struct {
	PreviousBlockHash string
	Transactions      []Transaction
	Nonce             int
	Hash              string
}

type Account struct {
	Balance    int
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	Address    string
}
