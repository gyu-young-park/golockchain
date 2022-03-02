package main

import (
	"github/gyu-young-park/blockchain"
)

func main() {
	blockChain := blockchain.NewBlockchain()
	prevHash := blockChain.GetLastBlock().Hash()
	blockChain.CreateBlock(5, prevHash)

	prevHash = blockChain.GetLastBlock().Hash()
	blockChain.CreateBlock(10, prevHash)
	blockChain.PrintInfo()
}
