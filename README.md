# Overview

This is a simple project to connect to an Ethereum network to get transaction data, parse it and persists the parse data

## Requirements

* Using only JSON-RPC
* Parse Transaction
* Store Transaction in memory

## Solution

1. The layout of this project follows the principles outline [here](https://paulwizviz.github.io/go/2022/12/23/go-proverb-architecture.html)

2. The project is organised around two packages `eth` -- containing operations to interact with Ethereum network -- and `store` -- containing a mock of a key value store.

3. The assumption is the `Parser` is the business logic responsible for storing data read from the Ethereum network, so that consumer of the object is able to read data from store without making RPC call to the Ethereum network.