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
	// TidyPrint(account)

	transaction := transactions.CreateTransaction(account.PrivateKey, "0x123", 100)
	fmt.Println(transaction.Signature)

	t2 := transactions.CreateTransaction(a2.PrivateKey, "0x12343", 100)

	//verify the transaction

	res := transactions.VerifyTransaction(t2)

	fmt.Print(res)

	
	

// 
}
