package main

import (
	"fmt"

	"github.com/gr4yha7/go-blockchain/blockchain"
)

func main() {
	blockchain := blockchain.InitializeBlockchain()

	blockchain.AddBlock("First block")
	blockchain.AddBlock("Second block")
	blockchain.AddBlock("Third block")

	for _, bc := range blockchain.Blocks {
		fmt.Printf("Previous Hash: %x\n", bc.PreviousHash)
		fmt.Printf("Data: %s\n", string(bc.Data))
		fmt.Printf("Hash: %x\n", bc.Hash)
	}

}
