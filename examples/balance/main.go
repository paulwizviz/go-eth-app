package main

import (
	"context"
	"fmt"
	"log"
	"paulwizviz/go-eth-app/internal/jrpc"
)

func main() {
	url := "http://localhost:8545"
	addr := "0x1a12FAD7cA2fb46417Ffc3dE3Cda72219459292C"

	bal, err := jrpc.GetBalance(context.TODO(), url, 1, addr, jrpc.BlockTagLATEST)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bal)
}
