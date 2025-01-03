
[[decentralisation]]


## TODO

- [ ] Store and load the tree from leveldb
- [ ] Account Creation with node adding to the tree
- [ ] rlp encoding of the account
- [ ] hashing the accounts 
- [ ] merkel proofs
- [ ] Translate Transactions into tree updates
- [ ] Construct blocks from the transactions
- [ ] Persist the blockchain
- [ ] p2p communication between nodes
- [ ] 


## Features

- [ ] automated simulation of transactions  
- [ ] accounts and send money
- [ ] nodes talk to each other
- [ ] visualization using  frontend
- [ ] measure backend performance
- [ ] simulate separate nodes


### node architecture

**- The execution client** (also known as the Execution Engine, EL client or formerly the Eth1 client) listens to new transactions broadcasted in the network, executes them in EVM, and holds the latest state and database of all current Ethereum data.
- **The consensus client** (also known as the Beacon Node, CL client or formerly the Eth2 client) implements the proof-of-stake consensus algorithm, which enables the network to achieve agreement based on validated data from the execution client

- transactions - An Ethereum transaction refers to an action initiated by an externally-owned account, in other words an account managed by a human, not a contract. For example, if Bob sends Alice 1 ETH, Bob's account must be debited and Alice's must be credited. This state-changing action takes place within a transaction.
- merkel tree implementation for the blockchain

## world state

- merkel patricia tree
- persist state in database
https://easythereentropy.wordpress.com/2014/06/04/understanding-the-ethereum-trie/
https://ethereum.org/en/developers/docs/data-structures-and-encoding/patricia-merkle-trie/

The blockchain is a history of set of actions taken on the initial state having all the records which has brought it to the present state

Each transaction on the network of ethereum is translated to a CRUD operation on the tree data structure which represents the state of the system.
- If you create an account, a new leaf node is added to the tree
- if you transfer some eth, your balance atrribute is reduced and all the node related to the you and reciever are updated
- each leaf node is either a contract or an externally owned account
- the root nodes hash is derived from all the child nodes which is called a merkel proof method


### Merkel Patrichia Tree Implementation
[[Ethereum Merkel Patricia Tree]]

####  Database for storing the world state

Just like traditional databases use b trees to represent the index of the records, I think ethereum uses a special data structure called a merkel patricia tree to represent the world state along with all its key-value pairs.

The reason for using this data structure is that the ds has an inherit security feature called the merkel proofs and speed provided by the patricia shortcut.

While going through the geth codebase, I found that the project uses leveldb which is a blazingly fast key value store written in cpp. But geth was using a go implementation for the same.

Since I wanted to learn about new things anyway, I think I should give it a shot. Also wanted to implement interfaces and abstraction so that I can easily swap out the db in future if the **leveldb** become too cumbersome.

How the world state is stored in the database :

 Go-ethereum uses levelDB, and parity uses rocksDB to store states. They are key-value storages. Keys and values saved in the storage are not the key-values of the Ethereum state. The value that is stored in the storage is the content of MPT node while the key is the hash of this node.


https://github.com/syndtr/goleveldb

https://medium.com/cybermiles/diving-into-ethereums-world-state-c893102030ed


## Accounts

- Have a public and private address to identify 
- balance and nonce
- crud should be available


The blockchain supports only ownable account with no support for contract account. Each account is just a pair of ecdsa  (public and private) keys. 


## Transactions and blocks

- actions that change the world state
- type - transfer money from one account to other
- transactions will be bundled intro fixed size block
- each transaction has - from, to amount

Transactions are integral part of the blockchain are the only way for the account holders to manipulate the state of the blockchain. 

### Proofs

Transactions need to be cryptographically signed and anyone should be able to verify it easily based on the signature. In ethereum this is achieved by using digital signature algorithms and hash functions.

I am exploring the ecliptic curve digital signature algoritm which is used by geth for verification.

- [ ] implement signature
- [ ] verify signature
- [ ] get public address from the signature using some algo
- [ ] Produce many dummy transactions
- [ ] Implement queue for storing the incoming transaction

## Consensus 

- network of nodes
- achieved through proof of stake

These are mechanism to protect against sybil attacks. Atleast 66% of nodes involved must agree on the world state

#### Proof of work

In this consensus mechanism nodes try to solve a cryptographic puzzle. Whoever solves this puzzle is rewarded and is responsible for adding a new block to the chain. Other nodes can easily verify the answer of the puzzle solved by the winner is valid or not.

This puzzle which is talked about everywhere is simple to understand. Hashing functions are one way functions which give you a fixed size value given any input of any size. Also, whenever you give it the same input it will always give you the same output.

One simple example of a hash function would be a mod operator. For eg, let's always mod an input with 7. 

1 - > 1 % 7 = 1
7- > 7 % 7 = 0
701 -> 701 % 7 = 1

As you can see whatever input i give to this hash function it will always give me a number between zero and 6. 
Something interesting is happening here as well. If you noticed 1 and 701 are both giving 1 as their hashes. This means that if you have the output you can't determine from what kind of input it came from. This is the one way part of the hash function.


## Execution 

- Listens for and executes the transactions and maintains the local state 
- give this state to the consensus client to update in the network


#### Gossip Methods

- find block number
- send transaction 

#### State Methods

- get balance
- call contracts (optinal)

#### History Methods

- get uncle block
- get old transactions and blocks


## Interface for client application

The present ethereum network uses a JSON-RPC specification to interface with the network. 
rpc stands for remote procedure call and its a way for a program to call a procedure as if it was part of the same program on the same computer.

Rpc protocols use message passing to coordinate the call and params which are passeed back and forth between the calle and server.

Popular implementation are 
gRPC and JSON-RPC

### gRPC

Developed by google , uses http2 for transport and protocol buffers as an interface description language.

Due to complexity, grpc is not supported in the browsers, so it requieres a proxy which limits its usecase to communication between microservices


## Simulation of Transactions

Using a go server producers send thousands of transactions to the message queue. Not required but I wanted to learn about it so...

Transactions are produced as random from one user to other and pushed  into the message queue










