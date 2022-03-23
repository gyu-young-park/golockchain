package server

import (
	"github/gyu-young-park/blockchain"
	"github/gyu-young-park/wallet"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockChainSever struct {
	port uint16
}

func NewBlockChainServer(port uint16) *BlockChainSever {
	return &BlockChainSever{port: port}
}

func (server *BlockChainSever) Port() uint16 {
	return server.port
}

func (server *BlockChainSever) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minerWallet.BlockchainAddress(), server.port)
		cache["blockchain"] = bc
		log.Printf("private_key: %v", minerWallet.PrivateKeyStr())
		log.Printf("public_key: %v", minerWallet.PublicKeyStr())
		log.Printf("blockchain_address: %v", minerWallet.BlockchainAddress())
	}
	return bc
}

func (server *BlockChainSever) GetChain(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bc := server.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (server *BlockChainSever) Run() {
	http.HandleFunc("/", server.GetChain)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(server.port)), nil))
}
