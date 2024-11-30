#!/bin/bash

export GETH_TOOLS_IMAGE_NAME="go-eth-app/gethtool"

COMMAND=$1

case $COMMAND in
    "build")
        docker compose -f ./build/tools/builder.yaml build
        ;;
    "clean")
        docker rmi -f ${GETH_TOOLS_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "shell")
        docker run -it --rm $GETH_TOOLS_IMAGE_NAME /bin/bash
        ;;
    *)
        echo "Usage: $0 [build | clean]"
        ;;
esac