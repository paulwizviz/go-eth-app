#!/bin/bash

if [ "$(basename $(realpath .))" != "go-eth-app" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
fi

function build_tools(){
   docker compose -f ./build/tools/builder.yaml build
}

function clean_tools(){
    docker rmi -f ${GETH_TOOLS_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

