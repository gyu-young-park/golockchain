package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github/gyu-young-park/transaction"
	"time"
)

const BLOCK_HASH_SIZE = sha256.Size

type Block struct {
	Timestamp    int64                      `json:"timestamp"`
	Nonce        int                        `json:"nonce"`
	PreviousHash [BLOCK_HASH_SIZE]byte      `json:"previous_hash"`
	Transactions []*transaction.Transaction `json:"transactions"`
}

func NewBlock(nonce int, previousHash [BLOCK_HASH_SIZE]byte, transactions []*transaction.Transaction) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
	}
}

func (b *Block) PrintInfo() {
	fmt.Printf("nonce		%v\n", b.Nonce)
	fmt.Printf("previousHash	%x\n", b.PreviousHash)
	fmt.Printf("timestamp	%v\n", b.Timestamp)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(Block{
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
		Transactions: b.Transactions,
	})
}

// sha256.Size == 32, block converts to json string that represents byte array. And i use it to make hash value
func (b *Block) Hash() [sha256.Size]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}
