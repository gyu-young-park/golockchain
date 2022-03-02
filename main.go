package main

import (
	"github/gyu-young-park/blockchain"
	"log"
)

func main() {
	log.Println("start golockchain")

	blockchainInstance := blockchain.NewBlockchain()
	blockchainInstance.CreateBlock(2, "hash 1")
	blockchainInstance.CreateBlock(5, "hash 2")
	blockchainInstance.PrintInfo()
}
