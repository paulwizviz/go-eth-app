#!/bin/sh

geth init /genesis.json
geth --http --http.addr 0.0.0.0 --http.api personal,eth,net,web3 --dev