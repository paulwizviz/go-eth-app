# Geth Nodes and Networks

This project includes Docker based Geth nodes configured to run in Dev node and a local network comprising of two nodes.

The Geth Docker nodes used in this project are derived from this image [ethereum/client-go](https://hub.docker.com/r/ethereum/client-go).

## Dev node

Use this [./scripts/gethnode.sh](../scripts/gethnode.sh) to activate the `dev node`. The command to activate the node is `./scripts/gethnode.sh dev start`.

When you start the network, you will see a log. Scroll the log and copy the developer account as that contains the Eth to enable the issuance of transactions. Here is an example of the log:

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

In the above example, the randomly generated account address is `0xf6aD33E18D1d4cc0ab5926d0FcDA4EEafC430981`. The address is regenerated when the node is re-activated. This account is unlocked that can be transacted without signing.

## Two Nodes Local Network

**TO DO**



