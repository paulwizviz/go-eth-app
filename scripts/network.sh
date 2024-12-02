#!/bin/bash

export NETWORK_NAME=go-eth-app_network

COMMAND=$1
SUBCOMMAND=$2

function devnode(){
    local cmd=$1
    case $cmd in
        "start")
            docker compose -f ./deployment/dev-node.yaml up
            ;;
        "stop")
            docker compose -f ./deployment/dev-node.yaml down
            ;;
        *)
            echo "Usage: $0 dev [start|stop]"
            ;;
    esac
}

case $COMMAND in
    "dev")
        devnode $SUBCOMMAND
        ;;
    *)
        echo "Usage $0 [dev]"
        ;;
esac