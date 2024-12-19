package main

import (
	"fmt"
	"log"

	"github.com/paulwizviz/go-eth-app/internal/wallet"
)

func main() {

	acct, err := wallet.KeystoreCreate("./tmp", "abcdef")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(acct.URL.Path)

	privkey, err := wallet.KeystoreRecoverPrivKey(acct.URL.Path, "abcdef")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T", privkey)
}
