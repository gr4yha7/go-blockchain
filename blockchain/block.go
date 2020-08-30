package blockchain

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash[:]
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
