package main

import (
	"fmt"
	"structs/pkg/blockchain/blockchain"
)

func main() {
	chain := blockchain.CreateBlockchain(1) //why it's always false when I set it to 0 ?
	chain.AddBlock("Rodion", "Dima", 300)
	chain.AddBlock("Dima", "Anastasia", 120)
	fmt.Println(chain.IsValid())
}
