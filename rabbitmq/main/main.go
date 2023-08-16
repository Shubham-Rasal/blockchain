package main

import "github.com/Shubham-Rasal/blockchain/rabbitmq"

func main() {

	rabbitmq.Send()
	rabbitmq.Receive()


}
