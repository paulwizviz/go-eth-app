# Solidity Compiler and Abigen

This projet use Docker based solidity compiler (`solc`) and `abigen` tools to support the compilation of Solidity contracts and generating Go binding for Solidity contracts. The compiler is derived from the Docker image [ethereum/solc](https://hub.docker.com/r/ethereum/solc/) and the `abigen` is derived from the image [ethereum/client-go](https://hub.docker.com/r/ethereum/client-go). This project also include a [script](../../build/tools/tools.dockerfile) to build an image from source 

## Working Examples

This project contains an example Solidity contact [/solidity/hello/hello.sol](../solidity/hello/hello.sol) use to generate Go binding or demonstrate contract deployment. The process to generate Go source are found in this [./scripts/hello.sh](../scripts/hello.sh). 

The steps to compile and generate Go binding are as follows:

* STEP 1 - Run the command `./scripts/hello.sh compile` to generate ABI and BIN files found in the folder `./solidity/abi/hello`.
* STEP 2 - Run the command `./scripts/hello.sh abi` to generate Go binding in `internal/hello/hello.go`.

Having generated the Go binding `./internal/hello`, start writing dApp to interact with the Solidity deployed in an Ethereum node. The following illustrates the steps to write the dApp.

```Go
import "github.com/ethereum/go-ethereum/ethclient"

nodeurl := "<url to node>"
conn, err := ethclient.Dial(nodeurl)

// Generated factory
contract, err := hello.NewHelloWorld("<contract address>", conn)
```
