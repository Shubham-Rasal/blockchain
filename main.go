package main

import (
	"encoding/json"
	"fmt"

	// "net/http"
	// "time"
	// "github.com/gofrs/uuid"
	// blockchain "github.com/Shubham-Rasal/blockchain"
	"github.com/Shubham-Rasal/blockchain/blockchain"
	"github.com/Shubham-Rasal/blockchain/blockchain/transactions"
	// "github.com/Shubham-Rasal/blockchain/merkel"
	// rabbitmq "github.com/Shubham-Rasal/blockchain/rabbitmq"
	// "github.com/Shubham-Rasal/cmd"
)

func TidyPrint(data interface{}) {
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}



func main() {


	account := blockchain.CreateAccount()
	a2 := blockchain.CreateAccount()
	a2.Balance = 1000
	TidyPrint(a2)

	transaction := transactions.CreateTransaction(account.PrivateKey, a2.PublicKey, 100)
	fmt.Println(transaction.Signature)

	t2 := transactions.CreateTransaction(a2.PrivateKey, account.PublicKey, 20)

	//verify the transaction

	res := transactions.VerifyTransaction(t2)

	fmt.Print(res)

	
	

// 
}
