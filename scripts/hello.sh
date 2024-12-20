#!/bin/bash

if [ "$(basename $(realpath .))" != "go-eth-app" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
fi

function compile_sol(){
    local sol="Hello.sol" # Change this to meet your script
    docker run -v $(PWD)/solidity/hello/$sol:/opt/solidity/$sol \
            -v $(PWD)/solidity/abi/hello/:/opt/abi \
            ${SOLC_TOOL} --abi --bin /opt/solidity/$sol -o /opt/abi --evm-version paris
}

function abi(){
    local abi="HelloWorld.abi" # You will find this in /solidity/abi
    local bin="HelloWorld.bin"
    local pkg="hello"
    local type="HelloWorld"
    docker run -v ${PWD}/solidity/abi/$pkg/$abi:/opt/abi/$abi \
               -v ${PWD}/solidity/abi/$pkg/$bin:/opt/api/$bin \
               -v ${PWD}/internal/contract/$pkg:/opt/contract/$pkg \
               ${GO_CLIENT_TOOL} abigen --abi /opt/abi/$abi --bin /opt/api/$bin --pkg $pkg --type $type --out /opt/contract/$pkg/"${pkg}.go"
}

COMMAND=$1

case $COMMAND in
    "abi")
        abi
        ;;
    "clean")
        rm -rf $PWD/solidity/abi/hello
        rm -rf $PWD/internal/hello
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "compile")
        compile_sol
        ;;
    *)
        echo "Usage: $0 [abi | compile]"
        ;;
esac