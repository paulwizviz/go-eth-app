# Ethereum JSON RPC

Every Ethereum client provides a uniform way to interact with an Ethereum client via [JSON RPC protocol](https://www.jsonrpc.org/specification).

## Packages

The Go Ethereum project provides out-of-the-box JSON-RPC packages. These are:
 
* `github.com/ethereum/go-ethereum/rpc` - low level JSON-RPC.
* `github.com/ethereum/go-ethereum/ethclient` - a package that is built on top of the `rpc` package.

Here is an example using `rpc` package:

```go
    client, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	// JSON-RPC: eth_blockNumber
	var blockNumber string
	err = client.CallContext(context.Background(), &blockNumber, "eth_blockNumber")
	if err != nil {
		log.Fatalf("Failed to fetch block number: %v", err)
	}

	fmt.Printf("Current block number: %s\n", blockNumber)
```

Here is an example using `ethclient` package:

```go
    // Connect to the Ethereum node
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	// Get the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch block number: %v", err)
	}

	fmt.Printf("Current block number: %d\n", blockNumber)
```

There is an implementation of a JSON-RPC in this project [./internal/jrpc](../internal/jrpc/). Typically, there is no need to create a JSON-ROC client from scratch. This implementation in this project is intended primarily for eductional purpose.

## Working Examples

You will find a working example [here](../examples/jrpc/main.go). This example demonstrate the steps to connect to a node JSON-RPC server `https://ethereum-rpc.publicnode.com`.

## References

* [The Ethereum JSON RPC specification](https://ethereum.github.io/execution-apis/api-documentation/)
* [List of Ethereum RPC Nodes](https://ethereumnodes.com/)