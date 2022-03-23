package server

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockChainSever struct {
	port uint16
}

func NewBlockChainServer(port uint16) *BlockChainSever {
	return &BlockChainSever{port: port}
}

func (server *BlockChainSever) Port() uint16 {
	return server.port
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world") // response
}

func (server *BlockChainSever) Run() {
	http.HandleFunc("/", HelloWorld)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(server.port)), nil))
}
