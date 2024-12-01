#!/bin/bash

if [ "$(basename $(realpath .))" != "go-eth-app" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/tools.sh
fi

function compile_sol(){
    local sol="Hello.sol" # Change this to meet your script
    docker run -v $(PWD)/solidity/hello/$sol:/opt/solidity/$sol \
            -v $(PWD)/solidity/abi/hello/:/opt/abi \
            ${GETH_TOOLS_IMAGE_NAME} solc --abi --bin /opt/solidity/$sol -o /opt/abi
}

function abi(){
    local abi="HelloWorld.abi" # You will find this in /solidity/abi
    local pkg="hello"
    local type="HelloWorld"
    docker run -v ${PWD}/solidity/abi/$pkg/$abi:/opt/abi/$abi \
               -v ${PWD}/internal/contract/$pkg:/opt/contract/$pkg \
               ${GETH_TOOLS_IMAGE_NAME} abigen --abi /opt/abi/$abi --pkg $pkg --type $type --out /opt/contract/$pkg/"${pkg}.go"
}

COMMAND=$1

case $COMMAND in
    "abi")
        abi
        ;;
    "build")
        build_tools
        ;;
    "clean")
        clean_tools
        rm -rf $PWD/internal/hello
        ;;
    "compile")
        compile_sol
        ;;
    "shell")
        docker run -it --rm $GETH_TOOLS_IMAGE_NAME /bin/bash
        ;;
    *)
        echo "Usage: $0 [build | clean]"
        ;;
esac