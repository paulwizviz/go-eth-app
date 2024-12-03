# Tools

Several tools are included in this project to support the task of developing products in this project.

The tools are are built using Docker. Please install Docker.

## Geth Node and Network

A script to build a Geth node in Docker is found here [./build/gethnode/geth.dockerfile](../build/gethnode/geth.dockerfile).

A script to spin up Geth networks are here:

* [Sinmgle dev node](../deployment/gethdev/dev-node.yaml)

### Dev node

A single Geth node running in dev mode is provided. When you start the network, you will see a log. Scroll the log and copy the
developer account as that contains the Eth to enable the issuance of transactions. Here is an example of the log:

```sh
geth-1  | INFO [12-03|14:42:24.525] Maximum peer count                       ETH=50 total=50
geth-1  | INFO [12-03|14:42:24.526] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
geth-1  | INFO [12-03|14:42:24.529] Set global gas cap                       cap=50,000,000
geth-1  | INFO [12-03|14:42:24.654] Using developer account                  address=0xf6aD33E18D1d4cc0ab5926d0FcDA4EEafC430981
geth-1  | INFO [12-03|14:42:24.654] Initializing the KZG library             backend=gokzg
geth-1  | INFO [12-03|14:42:24.678] Allocated trie memory caches             clean=154.00MiB dirty=256.00MiB
geth-1  | INFO [12-03|14:42:24.678] State schema set to default              scheme=path
geth-1  | INFO [12-03|14:42:24.679] Initialising Ethereum protocol           network=1337 dbversion=<nil>
geth-1  | INFO [12-03|14:42:24.679] Initialized path database                cache=154.00MiB buffer=256.00MiB history=90000
```

In the above example, the randomly generated address is `0xf6aD33E18D1d4cc0ab5926d0FcDA4EEafC430981`.

There are two scripts:

* [./scripts/gethnode.sh](../scripts/gethnode.sh) - use this to build a Geth node image.
* [./scripts/network.sh](../scripts/network.sh) - use this to start and stop the dev node.

## Solidity Compiler (solc) and ABI Gen (abigen)

A `solc` and a `abigen` are also provided in this project in the form of a Docker image [./build/tools/tools.dockerfile](../build/tools/tools.dockerfile)

There is a reference script [./scripts/hello.sh](../scripts/hello.sh) demonstrating the steps to compile a solidity contract and generates Go code.