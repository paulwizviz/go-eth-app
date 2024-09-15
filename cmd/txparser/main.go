package main

import (
	"fmt"
	"log"
	"net/http"
	"txparser/internal/eth"
)

func main() {
	url := "https://ethereum-rpc.publicnode.com"

	// Create a channel from Ethereum network
	// pass the channel to parser
	ch := eth.ReadNetwork(url)
	parser := eth.NewDefaultParser(ch)

	// Inject parser to REST server
	rest := &eth.RestServer{
		Parser: parser,
	}

	// Setup REST server
	http.HandleFunc("GET /block/current", rest.GetCurrentBlock)
	http.HandleFunc("PUT /subscribe/{address}", rest.Subscribe)
	http.HandleFunc("PUT /transaction/{address}", rest.GetTransactions)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
