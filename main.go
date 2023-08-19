package main

import (
	"encoding/json"
	"fmt"

	// "net/http"
	// "time"
	// "github.com/gofrs/uuid"
	// "github.com/Shubham-Rasal/blockchain/blockchain"
	// "github.com/Shubham-Rasal/blockchain/blockchain/transactions"
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

func modify(ptr *int) {
	fmt.Printf("%p\n", ptr)
	*ptr = 20
}


func main() {


	var arr []int

	arr = append(arr, 1, 2, 3, 4, 5)

	fmt.Printf("%p\n", &arr[2])
	
	ptr := &arr[2]

	modify(ptr)

	fmt.Println(arr)

 
}


