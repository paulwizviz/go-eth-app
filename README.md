# Building Decentralised Application (dApp) with Go: An Educational Project

This project provides a practical guide for developers to explore the Ethereum ecosystem through the Go programming language. By leveraging a hypothetical product, this educational initiative demonstrates key concepts and workflows, making it accessible for developers to build and interact with blockchain-based applications.

## Intended Audience

This project is designed for blockchain enthusiasts, students, and experienced developers who wish to explore the Ethereum ecosystem using Go. Whether you're new to decentralized applications (dApps) or looking to deepen your understanding of smart contracts, this educational resource lays the groundwork for building and interacting with Ethereum-based applications.

## Scope of the Project

This project adopts the development of a hypothetical product as a practical and structured approach to demonstrate building Web3 applications with Go. By designing a product with specific features, this educational effort offers a hands-on experience to help developers understand and implement the foundational elements of Ethereum application development. The product includes the following features:

1. **Web3 Application**: An example demonstrating how users can send transactions to the Ethereum network and interact with deployed smart contracts.

2. **Wallet**: Functionality for managing cryptographic keys, enabling users to securely store and access their Ethereum addresses.

3. **Smart Contract Deployment**: Tools for deploying smart contracts to the Ethereum network directly from a Go application.

4. **Blockchain Viewer**: A minimal blockchain explorer for retrieving and displaying Ethereum blockchain data.

By centering the project around this hypothetical product, the approach provides a structured and practical path for learning how to build Web3 applications with Go.

## Project Content

* Application Programming Interfaces 
    * [JSON RPC](./docs/jsonrpc.md)
    * [Go ABI from Solidity](./docs/abi.md)
    * [Contract Operations API](./docs/contractops.md)
* [Design](./docs/design.md) - Description of the design of the product.
* [Tools](./docs/tools.md) - Description of the tools to support the development of the project.

## Terms Used in this Project

The terms used in this project are based on the [official documentation](https://ethereum.org/en/developers/docs/).

* Base fee. Every block has a base fee which acts as a reserve price. The base fee is calculated by a formula that compares the size of the previous block (the amount of gas used for all the transactions) with the target size. The base fee will increase by a maximum of 12.5% per block if the target block size is exceeded.

* Geth client. The Geth client is the default client used in this project. Much of the development in this project is based on the [Geth source code](https://github.com/ethereum/go-ethereum).

* Gas tip cap. This is the maximum price a user is willing to pay above the base price for a transaction to prioritize it. The effective gas tip is the actual tip paid above the base fee of the block.

* Gas fees. These are transaction costs paid on the Ethereum blockchain to perform operations like sending Ether (ETH) or interacting with smart contracts. Gas fees are paid in `gwei`. Gas fees change based on supply and demand. When the network is congested, gas prices are higher, and when there is less traffic, they are lower.

* Max fee. This optional parameter is known as the maxFeePerGas. For a transaction to be executed, the max fee must exceed the sum of the base fee and the tip.

* Node. In this project, the Docker container is the default node.

* Priority fee (tips). The priority fee (tip) incentivizes validators to include a transaction in the block. Small tips give validators a minimal incentive to include a transaction.

* Wei. This is the default unit of Ether used in all API implemented in this project.

## Disclaimer

This project is for educational purposes only. It is not a production-ready solution and requires significant modifications, rigorous security audits, and extensive testing before use in any production environment.

This project is ongoing and may undergo changes without prior notification. By using this project, you acknowledge that you do so at your own risk. The authors of this project accept no liability for any issues or damages resulting from its use. Please use this project as a learning resource and not as a fully functional or secure application.

## Copyright

Unless otherwise specified, this project is copyrighted as follows:

Copyright 2024 The Contributors to go-eth-app

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

For a list of contributors, refer to the CONTRIBUTORS file or the repository's commit history.