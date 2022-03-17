package main

import (
	"fmt"
	"github/gyu-young-park/blockchain"
	"github/gyu-young-park/wallet"
)

func main() {
	blockchainWallet := wallet.NewWallet()
	blockChain := blockchain.NewBlockchain(blockchainWallet.BlockchainAddress())

	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := walletA.NewTransactionForSignature(walletB.BlockchainAddress(), 1.0)
	isAdded := blockChain.AddTransaction(t.Transaction, t.SenderPublicKey, t.GenerateSignature())
	fmt.Println("Added?", isAdded)

	blockChain.Mining()
	blockChain.PrintInfo()
}
