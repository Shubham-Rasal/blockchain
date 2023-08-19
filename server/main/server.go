package main

import (
	"github.com/gofiber/fiber/v2"
    "github.com/Shubham-Rasal/blockchain/blockchain"
    "github.com/Shubham-Rasal/blockchain/blockchain/transactions"
	"log"
)


type CreateAccountResponse struct {
    Address string `json:"address"`
    Balance int `json:"balance"`
}

type User struct {
    Address string `json:"address"`
    Account blockchain.Account `json:"account"`
}

//create a slice of pointers to User this prevents creating a copy of the user
var users []*User

func handleCreateAccount(c *fiber.Ctx) error {
  
    account , err := blockchain.CreateAccount(100)
    if err != nil {
        return c.JSON(err)
    }    

    users = append(users, &User{
        Address: string(account.Address.Hex()),
        Account: account,
    })

    return c.JSON(CreateAccountResponse{
        Address: string(account.Address.Hex()),
        Balance: account.Balance,
    })
}



func handleGetAccounts(c *fiber.Ctx) error {
    var accounts []CreateAccountResponse
    for _ , user := range users {
        accounts = append(accounts, CreateAccountResponse{
            Address: user.Address,
            Balance: user.Account.Balance,
        })
    }

    return c.JSON(accounts)
}


type TransactionRequest struct {
    From string `json:"from"`
    To string `json:"to"`
    Amount int `json:"amount"`
}

func handleTransaction(c *fiber.Ctx) error {
    var transactionRequest TransactionRequest
    if err := c.BodyParser(&transactionRequest); err != nil {
        return c.JSON(err)
    }

    fromIndex := -1

    
    for idx , user := range users {
        if user.Address == transactionRequest.From {
            fromIndex = idx
            log.Printf("initial address: %p\n" , &user.Account)
            break
        }
    }

    if fromIndex == -1 {
        return c.JSON(fiber.Map{
            "message": "from account not found",
        })
    }


    from := &users[fromIndex].Account    

    toIndex := -1

    //range creates a copy of the user and return 
    // in order to get the original user we need to use the index
    // then dereference the pointer to get the original user
    for idx , user := range users {
        if user.Address == transactionRequest.To {
            toIndex = idx
            break
        }
    }

    if toIndex == -1 {
        return c.JSON(fiber.Map{
            "message": "to account not found",
        })
    }

    to := &users[toIndex].Account  

    if *from == *to {
        return c.JSON(fiber.Map{
            "message": "from and to accounts cannot be the same",
        })
    }

    transaction := transactions.CreateTransaction(from, to, transactionRequest.Amount)

    log.Printf("final address: %p\n" , from)

    // log.Printf("New Balance of %s: %d\n", string(from.Address.Hex()), from.Balance)
	// log.Printf("New Balance of %s: %d\n", string(to.Address.Hex()), to.Balance)

    return c.JSON(transaction)

}

    


func main() {
   


    app := fiber.New()
    
	app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

    //TODO: add options to send initial balance
    app.Get("/create_account" , handleCreateAccount)

    app.Get("/get_accounts" , handleGetAccounts)

    //create a transaction
    app.Post("/create_transaction" , handleTransaction)

	log.Fatal(app.Listen(":3000"))
}
