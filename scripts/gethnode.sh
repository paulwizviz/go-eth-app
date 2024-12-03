#!/bin/bash

if [ "$(basename $(realpath .))" != "go-eth-app" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
fi

COMMAND=$1

case $COMMAND in
    "build")
        docker compose -f ./build/gethnode/builder.yaml build
        ;;
     "clean")
        docker rmi -f ${GETH_NODE_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo "Usage: $0 [build | clean]"
        ;;
esac
