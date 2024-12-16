#!/bin/sh

geth version
geth --http --http.addr 0.0.0.0 --http.api personal,eth,net,web3 --dev