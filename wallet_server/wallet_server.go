package main

import (
	"log"
	"net/http"
	"strconv"
)

type WalletServer struct {
	Port    uint16
	Gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{
		Port:    port,
		Gateway: gateway,
	}
}

func (ws *WalletServer) WalletIndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello world"))
	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.WalletIndexHandler)
	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port)), nil)
	if err != nil {
		log.Fatal("ERROR: Start Wallet Server Error!")
	}
}
