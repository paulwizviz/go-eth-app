package main

import (
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	url := "http://localhost:8545"
	addr := "0x2adc25665018aa1fe0e6bc666dac8fc2697ff9ba" // Dev mode address

	bal, err := jrpc.GetBalance(url, 1, addr, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bal)
}
