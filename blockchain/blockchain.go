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
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) PrintInfo() {
	for i := 0; i < len(bc.chain); i++ {
		fmt.Printf("===========================\n")
		bc.chain[i].PrintInfo()
		fmt.Printf("===========================\n")
	}
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *block.Block {
	b := block.NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
