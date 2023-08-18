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


	account , _ := blockchain.CreateAccount(89)
	a2, _ := blockchain.CreateAccount(90)
	a2.Balance = 1000
	// TidyPrint(a2)

	// transaction := transactions.CreateTransaction(account, a2, 100)
	// fmt.Println(transaction.Signature)

	t2 := transactions.CreateTransaction(a2, account, 20)

	//verify the transaction

	transactions.VerifyTransaction(t2)

	t2.Transfer(&a2, &account)

	fmt.Println(a2.Balance)
	fmt.Println(account.Balance)

 
}
