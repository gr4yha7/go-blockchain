package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return &block
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis block", []byte{})
}

func (c *Blockchain) AddBlock(data string) {
	prevHash := c.Blocks[len(c.Blocks)-1].Hash
	block := CreateBlock(data, prevHash)
	c.Blocks = append(c.Blocks, block)
}

func InitializeBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {
	blockchain := InitializeBlockchain()

	blockchain.AddBlock("First block")
	blockchain.AddBlock("Second block")
	blockchain.AddBlock("Third block")

	for _, bc := range blockchain.Blocks {
		fmt.Printf("Previous Hash: %x\n", bc.PreviousHash)
		fmt.Printf("Data: %s\n", string(bc.Data))
		fmt.Printf("Hash: %x\n", bc.Hash)
	}
}
