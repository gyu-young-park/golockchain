package main

import (
	"flag"
	"fmt"
	"github/gyu-young-park/server"
)

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain server")
	flag.Parse()
	fmt.Println(port)

	app := server.NewBlockChainServer(uint16(*port))
	app.Run()
}
