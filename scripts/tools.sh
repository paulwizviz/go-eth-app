#!/bin/bash

export GETH_TOOLS_IMAGE_NAME="go-eth-app/gethtool"
export GETH_VER="master"
export GO_VER=1.23.3-bookworm
export OS_VER=22.04
export SOLC_VER=v0.8.28

function build_tools(){
   docker compose -f ./build/tools/builder.yaml build
}

function clean_tools(){
    docker rmi -f ${GETH_TOOLS_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

