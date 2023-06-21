package main

import (
	"encoding/json"
	"fmt"
	"time"
	// "github.com/gofrs/uuid"
	// blockchain "github.com/Shubham-Rasal/blockchain"
	// rabbitmq "github.com/Shubham-Rasal/rabbitmq"
	// "github.com/Shubham-Rasal/cmd"
)

func TidyPrint(data interface{}) {
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		//sleep
		time.Sleep(1 * time.Second)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	// cmd.Execute()

	// var blockChain []blockchain.Block

	// account := blockchain.CreateAccount()
	// recipient := blockchain.CreateAccount()
	// var genesisBlock blockchain.Block
	// genesisBlock.PreviousBlockHash = ""
	// genesisBlock.Nonce = 0
	// genesisBlock.Transactions = append(genesisBlock.Transactions,  blockchain.CreateTransaction(account.PrivateKey, recipient.Address, 100))
	// genesisBlock.Hash = blockchain.HashBlock(genesisBlock)
	// blockChain = append(blockChain, genesisBlock)

	// TidyPrint(blockChain)

	// fmt.Printf("BlockChain : %+v\n" , blockChain)

	// transaction := CreateTransaction("Shubham" , "Arya" , 100)
	// fmt.Println(transaction)

	// rabbitmq.Send()

	//receive
	// rabbitmq.Receive()

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	// fmt.Println(<-c)

}
