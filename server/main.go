package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain server")
	flag.Parse()
	fmt.Println(port)

	app := NewBlockChainServer(uint16(*port))
	app.Run()
}
