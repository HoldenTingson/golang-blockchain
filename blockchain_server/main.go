package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockhain: ")
}

func main() {
	port := flag.Uint("port", 5000, "Blockchain Server TCP Port")
	flag.Parse()
	app := NewBlockChainServer(uint16(*port))
	app.Run()
}
