package blockchain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"
)

type Transaction struct {
	TransactionHash []byte
	From            ecdsa.PublicKey
	Recipient       ecdsa.PublicKey
	Amount          int
	TimeStamp       time.Time
	Nonce           int
	Signature       Signature
}

type Signature struct {
	R *big.Int
	S *big.Int
	V int
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
	Address    Address
}

type Hash [32]byte

type Address [20]byte

func (a Address) Hex() []byte {
	var buf [len(a)*2 + 2]byte
	copy(buf[:2], "0x")
	hex.Encode(buf[2:], a[:])
	return buf[:]
}

// SetBytes sets the address to the value of b.
// If b is larger than len(a), b will be cropped from the left.
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

//trims the hash to 32 bytes
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}


func (t *Transaction) Transfer(a *Account , b *Account){
	
	amountToTransfer := t.Amount

	if a.Balance < amountToTransfer {
		fmt.Println("Insufficient Balance")
		return
	}

	a.Balance -= amountToTransfer
	b.Balance += amountToTransfer

	log.Printf("Transferred %d from %x to %x\n", amountToTransfer, a.Address.Hex(), b.Address.Hex())
	log.Printf("New Balance of %x: %d\n", a.Address.Hex(), a.Balance)
	log.Printf("New Balance of %x: %d\n", b.Address.Hex(), b.Balance)
	
}
