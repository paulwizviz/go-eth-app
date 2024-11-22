package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"paulwizviz/go-eth-app/internal/eth"
	rest "paulwizviz/go-eth-app/internal/http"
	"syscall"
	"time"
)

const EthUrl = "https://ethereum-rpc.publicnode.com"

func main() {
	ctx := context.Background()
	notify, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create a channel from Ethereum network
	// pass the channel to parser
	ch := eth.ReadNetwork(ctx, EthUrl)
	parser := eth.NewDefaultParser(ch)

	// Inject parser to REST server
	rest := &rest.RestServer{
		Parser: parser,
	}

	// Setup REST server
	http.HandleFunc("GET /", rest.GetCurrentBlock)
	http.HandleFunc("GET /addresses", rest.GetAddresses)
	http.HandleFunc("GET /addresses/{address}", rest.GetTransactions)
	http.HandleFunc("GET /addresses/{address}/subscribe", rest.Subscribe)

	server := &http.Server{
		Addr:        "0.0.0.0:8080",
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	log.Println("Server is listening on port 8080...")
	go server.ListenAndServe()

	<-notify.Done()

	fmt.Println("")
	log.Printf("Shutting down server...")
	shutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	server.Shutdown(shutCtx)
	log.Println("Bye!")
}
