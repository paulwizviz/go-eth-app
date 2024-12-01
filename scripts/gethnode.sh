#!/bin/bash

export GETH_NODE_IMAGE_NAME="go-eth-app/gethnode"
export GETH_VER=master
export OS_VER=3.18
export GO_VER=1.22.1-alpine

COMMAND=$1
SUBCOMMAND=$2

function localnode(){
    local cmd=$1
    case $cmd in
        "start")
            docker compose -f ./deployment/local-dev-node.yaml up
            ;;
        "stop")
            docker compose -f ./deployment/local-dev-node.yaml down
            ;;
        *)
            echo "Usage: $0 local [start|stop]"
            ;;
    esac
}

case $COMMAND in
    "build")
        docker compose -f ./build/node/builder.yaml build
        ;;
     "clean")
        docker rmi -f ${GETH_NODE_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "local")
        localnode $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 [build | clean]"
        ;;
esac
