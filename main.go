package main

import (
	"flag"
	"fmt"
	"github/gyu-young-park/blockchain"
	"github/gyu-young-park/server"
	"github/gyu-young-park/wallet"
)

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain server")
	flag.Parse()
	fmt.Println(port)

	app := server.NewBlockChainServer(uint16(*port))
	app.Run()

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
