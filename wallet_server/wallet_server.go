package main

import (
	"github/gyu-young-park/utils"
	"github/gyu-young-park/wallet"
	"io"
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write([]byte("hello world"))
	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method)
	switch req.Method {
	case http.MethodPost:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Add("Conent-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.WalletIndexHandler)
	http.HandleFunc("/wallet", utils.WrapperCORS(ws.Wallet))
	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port)), nil)
	if err != nil {
		log.Fatal("ERROR: Start Wallet Server Error!")
	}
}
