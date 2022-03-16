package main

import (
	"fmt"
	"github/gyu-young-park/blockchain"
	"github/gyu-young-park/wallet"
)

func main() {
	testBlockchainAddress := "my_block_chain_address"
	blockChain := blockchain.NewBlockchain(testBlockchainAddress)

	w := wallet.NewWallet()
	fmt.Println("1: " + w.PrivateKeyStr())
	fmt.Println(w.PrivateKey())
	fmt.Println("1: " + w.PublicKeyStr())
	fmt.Println(w.PublicKey())

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.PrintInfo()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.PrintInfo()

	fmt.Printf("total bc transaction total value: %v\n", blockChain.CalculateTotalAmount(testBlockchainAddress))
	fmt.Printf("total bc transaction C value: %v\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("total bc transaction D value: %v\n", blockChain.CalculateTotalAmount("D"))
}
