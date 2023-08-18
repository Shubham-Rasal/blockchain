package main

import (
	"github.com/gofiber/fiber/v2"
    "github.com/Shubham-Rasal/blockchain/blockchain"
	"log"
)


type CreateAccountResponse struct {
    Address string `json:"address"`
    Balance int `json:"balance"`
}

type User struct {
    Address string `json:"address"`
    Balance int `json:"balance"`
}

var users []User

func handleCreateAccount(c *fiber.Ctx) error {


    account , err := blockchain.CreateAccount(80)
    if err != nil {
        return c.JSON(err)
    }    

    users = append(users, User{
        Address: string(account.Address.Hex()),
        Balance: account.Balance,
    })

    return c.JSON(CreateAccountResponse{
        Address: string(account.Address.Hex()),
        Balance: account.Balance,
    })
}

func handleGetAccounts(c *fiber.Ctx) error {
    return c.JSON(users)
}

func main() {
   


    app := fiber.New()
    
	app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})


    app.Get("/create_account" , handleCreateAccount)

    app.Get("/get_accounts" , handleGetAccounts)

	log.Fatal(app.Listen(":3000"))
}
