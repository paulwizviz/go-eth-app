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
* [Design](./docs/design.md) - Description of the design of the product.
* [Tools](./docs/tools.md) - Description of the tools to support the development of the project.

## Terms Used in this Project

The terms used in this project are based on the [official documentation](https://ethereum.org/en/developers/docs/).

* Geth client. The Geth client is the default client used in this project. Much of the development in this project is based on the [Geth source code](https://github.com/ethereum/go-ethereum).

* Node. In this project, the Docker container is the default node.

* Unit of Ether. The default unit of Ether is a Wei used in this project is the `Wei`.

| Unit	| Wei Equivalent | Description |
| --- | --- | --- |
| Wei | 10^0 = 1 wei | The smallest unit of Ether. |
| Kwei (Babbage) | 10^3 = 1,000 wei | Thousand wei. |
| Mwei (Lovelace) | 10^6 = 1,000,000 wei | Million wei. |
| Gwei (Shannon) | 10^9 = 1,000,000,000 wei | Billion wei, often used for gas prices. |
| Microether (Szabo) | 10^{12} = 1,000,000,000,000 wei | Trillion wei. |
| Milliether (Finney) | 10^{15} = 1,000,000,000,000,000 wei | Quadrillion wei. |
| Ether | 10^{18} = 1,000,000,000,000,000,000 wei | 1 Ether. |


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