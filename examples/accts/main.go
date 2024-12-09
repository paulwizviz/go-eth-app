package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"paulwizviz/go-eth-app/internal/jrpc"
	"syscall"
)

func main() {
	url := "http://localhost:8545"

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		i := 0
		for {
			accts, err := jrpc.Accounts(ctx, url, uint(i))
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(i, accts)
			i++
		}
	}()

	go func() {
		for range 1_000_000_000 {
		}
		cancel()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c
}
