package transactions

import (
	"crypto/ecdsa"
	"golang.org/x/crypto/sha3"
	"encoding/hex"
	"github.com/Shubham-Rasal/blockchain/blockchain"
	"log"
)


func VerifyTransaction(transaction blockchain.Transaction) bool {

	//derive the public key from the signature
	//the public key is a point on the curve
	//the public key is derived from the signature using the recovery id
	//the recovery id is used to determine the sign of the public key
	hasher := sha3.New256()
	hasher.Write([]byte(string(rune(transaction.From.X.Int64()))))
	hasher.Write([]byte(string(rune(transaction.Amount))))
	hasher.Write([]byte(string(rune(transaction.TimeStamp.Unix()))))
	hasher.Write([]byte(string(rune(transaction.Nonce))))
	// hasher.Write([]byte(string(rune(transaction.Nonce))))

	hash := hasher.Sum(nil)

	//compare the hash of the transaction with the hash derived from the signature
	if hex.EncodeToString(hash) != hex.EncodeToString(transaction.TransactionHash) {
		return false
	}

	// Sign creates a recoverable ECDSA signature.
	// The produced signature is in the 65-byte [R || S || V] format where V is 0 or 1.
	//TODO: for now the from field is there in the transaction itself
	//later it will be derived from the signature
	publicKey := transaction.From

	sign := ecdsa.Verify(&publicKey, []byte(transaction.TransactionHash), transaction.Signature.R, transaction.Signature.S)

	log.Printf("Transaction Signature: %t\n", sign)

	return sign
}
	