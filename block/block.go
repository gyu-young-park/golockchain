package block

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (b *Block) PrintInfo() {
	fmt.Printf("transactions	%v\n", b.transactions)
	fmt.Printf("nonce		%v\n", b.nonce)
	fmt.Printf("previouseHash	%v\n", b.previousHash)
	fmt.Printf("timestamp	%v\n", b.timestamp)
}
