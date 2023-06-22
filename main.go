package main

import (
	"encoding/json"
	"fmt"

	// "time"
	// "github.com/gofrs/uuid"
	// blockchain "github.com/Shubham-Rasal/blockchain"
	"github.com/Shubham-Rasal/blockchain/blockchain"
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

	// cmd.Execute()	
	
	//receive
	// rabbitmq.Receive()
	
	// rabbitmq.Send()

	account := blockchain.CreateAccount()
	TidyPrint(account)
// 
}
