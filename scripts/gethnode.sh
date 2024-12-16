#!/bin/bash

if [ "$(basename $(realpath .))" != "go-eth-app" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
fi

export NETWORK_NAME=go-eth-app_network

COMMAND=$1
SUBCOMMAND=$2

function dev_node(){
    local cmd=$1
    case $cmd in
        "start")
            docker compose -f ./deployment/gethdev/dev-node.yaml up
            ;;
        "stop")
            docker compose -f ./deployment/gethdev/dev-node.yaml down
            ;;
        *)
            echo "Usage: $0 dev [start | stop]"
            ;;
    esac
}

case $COMMAND in
    "dev")
        dev_node $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 [dev]"
        ;;
esac
