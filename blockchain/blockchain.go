package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"github/gyu-young-park/block"
	"github/gyu-young-park/transaction"
	"github/gyu-young-park/utils"
	"log"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Blockchain struct {
	transactionPool   []*transaction.Transaction
	chain             []*block.Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &block.Block{
		Nonce: 0,
	}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
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

func (bc *Blockchain) AddTransaction(transaction *transaction.Transaction, senderPublicKey *ecdsa.PublicKey, signature *utils.Signature) bool {
	// when recipient was miner, not should not verify signature
	if transaction.SenderBlockchainAddress == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, transaction)
		return true
	}
	if bc.VerifyTransactionSignature(senderPublicKey, signature, transaction) {
		// sender transaction all value is same with sender wallet balance
		/* now i commented, because of testing easily
		if bc.CalculateTotalAmount(transaction.SenderBlockchainAddress) < transaction.Value {
			log.Println("ERROR: Not enough balance in wallet")
			return false
		}
		*/
		bc.transactionPool = append(bc.transactionPool, transaction)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
	}
	return false
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, signature *utils.Signature, transaction *transaction.Transaction) bool {
	m, _ := transaction.MarshalJSON()
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], signature.R, signature.S)
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

func (bc *Blockchain) validProof(nonce int, previousHash [32]byte, transaction []*transaction.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := block.Block{Timestamp: 0, Nonce: nonce, PreviousHash: previousHash, Transactions: transaction}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) proofOfOWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.GetLastBlock().Hash()
	nonce := 0
	for !bc.validProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	transaction := transaction.NewTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	bc.AddTransaction(transaction, nil, nil)
	nonce := bc.proofOfOWork()
	previousHash := bc.GetLastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.Transactions {
			value := t.Value
			if blockchainAddress == t.RecipientBlockchainAddress {
				totalAmount += value
			}

			if blockchainAddress == t.SenderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}
