package main

import (
	"github/gyu-young-park/blockchain"
)

func main() {
	blockChain := blockchain.NewBlockchain()

	blockChain.AddTransaction("A", "B", 1.0)
	prevHash := blockChain.GetLastBlock().Hash()
	nonce := blockChain.ProofOfOWork()
	blockChain.CreateBlock(nonce, prevHash)

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	prevHash = blockChain.GetLastBlock().Hash()
	nonce = blockChain.ProofOfOWork()
	blockChain.CreateBlock(nonce, prevHash)
	blockChain.PrintInfo()
}
