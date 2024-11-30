#!/bin/bash

export GETH_NODE_IMAGE_NAME="go-eth-app/gethnode"
export GETH_VER=master
export OS_VER=3.18
export GO_VER=1.22.1-alpine

COMMAND=$1

case $COMMAND in
    "build")
        docker compose -f ./build/node/builder.yaml build
        ;;
     "clean")
        docker rmi -f ${GETH_NODE_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo "Usage: $0 [build | clean]"
        ;;
esac
