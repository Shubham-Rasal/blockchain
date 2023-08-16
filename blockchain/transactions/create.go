package transactions

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/sha3"
	"time"
	_rand "math/rand"
	"math"
	"github.com/Shubham-Rasal/blockchain/blockchain"	

	
)


func CreateTransaction(account blockchain.Account, to blockchain.Account, amount int) blockchain.Transaction {


	var transaction blockchain.Transaction
	transaction.From = account.PublicKey
	transaction.Recipient = to.PublicKey
	transaction.Amount = amount
	//store unix time in the transaction
	transaction.TimeStamp = time.Now()
	//generate a random number for the transaction
	transaction.Nonce = _rand.Intn(math.MaxInt64)
	//generate a transaction hash
	hasher := sha3.New256()
	hasher.Write([]byte(string(rune(transaction.From.X.Int64()))))
	hasher.Write([]byte(string(rune(transaction.Amount))))
	hasher.Write([]byte(string(rune(transaction.TimeStamp.Unix()))))
	hasher.Write([]byte(string(rune(transaction.Nonce))))
	transaction.TransactionHash = hasher.Sum(nil)
	//sign the transaction hash using the private key
	r, s, err := ecdsa.Sign(rand.Reader, account.PrivateKey, []byte(transaction.TransactionHash))
	if err != nil {
		fmt.Println(err)
	}

	//the signature is a pair of numbers (r,s) which are the coordinates of the point on the curve
	//the signature is encoded as a string
	//v is the recovery id which is used to recover the public key from the signature
	transaction.Signature = blockchain.Signature{
		R: r,
		S: s,
		V: 0,
	}

	return transaction
}
