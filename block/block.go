package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

const (
	BLOCK_HASH_SIZE = sha256.Size
)

type Block struct {
	Nonce        int                   `json:"nonce"`
	PreviousHash [BLOCK_HASH_SIZE]byte `json:"previous_hash"`
	Timestamp    int64                 `json:"timestamp"`
	Transactions []string              `json:"transactions"`
}

func NewBlock(nonce int, previousHash [BLOCK_HASH_SIZE]byte) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
}

func (b *Block) PrintInfo() {
	fmt.Printf("transactions	%v\n", b.Transactions)
	fmt.Printf("nonce		%v\n", b.Nonce)
	fmt.Printf("previouseHash	%x\n", b.PreviousHash)
	fmt.Printf("timestamp	%v\n", b.Timestamp)
}

// sha256.Size == 32, block converts to json string that represents byte array. And i use it to make hash value
func (b *Block) Hash() [sha256.Size]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}
