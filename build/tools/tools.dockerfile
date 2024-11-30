ARG OS_VER=22.04
ARG GO_VER=1.23.3-bookworm

FROM golang:${GO_VER} AS builder

ARG GETH_VER
ARG SOLC_VER

WORKDIR /opt

RUN apt update && \
    apt install -y sudo wget gcc g++ cmake make libboost-all-dev git z3 cvc4 nodejs npm protobuf-compiler && \
    cd /opt && \
    git clone  --depth 1 --branch ${GETH_VER} https:///github.com/ethereum/go-ethereum; cd /opt/go-ethereum && \
    go mod tidy; make all

RUN cd /opt && \
    git clone --recursive --branch ${SOLC_VER} https://github.com/ethereum/solidity.git && \
    cd solidity && \
    ./scripts/build.sh

FROM ubuntu:${OS_VER}

COPY --from=builder /opt/go-ethereum/build/bin/abigen /usr/local/bin/abigen
COPY --from=builder /opt/solidity/build/solc/solc /usr/local/bin/solc