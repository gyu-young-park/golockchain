package blockchain

import (
	"fmt"
	"github/gyu-young-park/block"
	"github/gyu-young-park/transaction"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
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

func (bc *Blockchain) CopyTransactionPool() []*transaction.Transaction {
	transactions := make([]*transaction.Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, transaction.NewTransaction(
			t.SenderBlockchainAddress,
			t.RecipientBlockchainAddress,
			t.Value,
		))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transaction []*transaction.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := block.Block{0, nonce, previousHash, transaction}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfOWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.GetLastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}
