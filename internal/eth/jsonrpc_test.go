package eth

import (
	"fmt"
)

func Example_ethBlockNum() {

	url := "https://ethereum-rpc.publicnode.com"
	blocknumber, err := currentBlock(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(blocknumber.Int64() != int64(0))

	txns, err := getBlockTransactions(url, blocknumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(txns) != 0)

	// Output:
	// true
	// true
}
