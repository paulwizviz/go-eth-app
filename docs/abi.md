# Go ABI from Solidity

This approach involve generating Go ABI using [Geth Abigen](https://geth.ethereum.org/docs/tools/abigen).

## Tools

This project includes scripts to build tools -- `solc` and `abigen`-- packaged in [Docker image](../build/tools/tools.dockerfile).

* [solc](https://github.com/ethereum/solidity) - A Solidity compiler to generate EVM byte code and `abi` file.
* [abigen](https://geth.ethereum.org/docs/tools/abigen) - A generator to convert ABI to Go source code.

## Working Examples

This project contains a reference Solidity contact [/solidity/hello/hello.sol](../solidity/hello/hello.sol) use to generate Go source code. The process to generate Go source are found in this [./scripts/hello.sh](../scripts/hello.sh). 

The steps to generate Go source code are as follows:

* STEP 1 - Build the docker tool image, run the command `./scripts/hello.sh build`.
* STEP 2 - Generate the ABI, run the command `./scripts/hello.sh compile`.
* STEP 3 - Generate the Go file, run the command `./scripts/hello.sh abi`.

This is what happens when you generate Go source code from the reference contract:

* `./scripts/hello.sh compile` generates ABI and BIN files found in `./solidity/abi/hello`.
* `./scripts/hello.sh abi` generates Go source code in `internal/hello/hello.go`.

When you have generated the Go package `./internal/hello`, you can start writing application to interact with solidity deployed in an Ethereum node. The following illustrates how you can do it.

```Go
import "github.com/ethereum/go-ethereum/ethclient"

nodeurl := "<url to node>"
conn, err := ethclient.Dial(nodeurl)

// Generated factory
contract, err := hello.NewHelloWorld("<contract address>", conn)
```
