package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "time"
	// "github.com/gofrs/uuid"
	// blockchain "github.com/Shubham-Rasal/blockchain"
	"github.com/Shubham-Rasal/blockchain/blockchain"
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


func handler(w http.ResponseWriter, r *http.Request) {
	
}


func main() {

	// cmd.Execute()	
	
	//receive
	// rabbitmq.Receive()
	
	// rabbitmq.Send()

	account := blockchain.CreateAccount()
	TidyPrint(account)

	// merkelTree := merkel.MerkelTree{}
	// fmt.Println(merkelTree)
	// merkelTree.Init([]string{"a"})
	// merkelTree.Build()
	// fmt.Println(merkelTree)
	// merkelTree.Delete("a")

	//conver 546 to hex
	fmt.Println(len("0x6088B06c5a187058434655B71057a9ee93E13d0d"))
	fmt.Println(len(account.Address))

	http.HandleFunc("/ping" , handler)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
// 
}
