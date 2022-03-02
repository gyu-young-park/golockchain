package blockchain

import (
	"fmt"
	"github/gyu-young-park/block"
)

type Blockchain struct {
	transactionPool []string
	chain           []*block.Block
}

func NewBlockchain() *Blockchain {
	b := &block.Block{
		Nonce: 0,
	}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) PrintInfo() {
	for i := 0; i < len(bc.chain); i++ {
		fmt.Printf("===========================\n")
		bc.chain[i].PrintInfo()
		fmt.Printf("===========================\n")
	}
}

func (bc *Blockchain) GetLastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [block.BLOCK_HASH_SIZE]byte) *block.Block {
	b := block.NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
