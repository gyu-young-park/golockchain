package blockchain

import (
	"fmt"
	"github/gyu-young-park/block"
	"github/gyu-young-park/transaction"
)

type Blockchain struct {
	transactionPool []*transaction.Transaction
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
	b := block.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*transaction.Transaction{}
	return b
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := transaction.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}
