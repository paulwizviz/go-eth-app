package main

import (
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	url := "http://localhost:8545"
	addr := "0x24c0d0cB3C5d8EbBcC5C7426872B0693A259e717" // Dev mode address

	bal, err := jrpc.GetBalance(url, 1, addr, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bal)
}
