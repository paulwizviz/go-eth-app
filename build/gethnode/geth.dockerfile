ARG GO_VER=1.22.1-alpine
ARG OS_VER=3.18

FROM golang:${GO_VER}${OS_VER} AS builder

ARG GETH_VER

RUN apk update && \
    apk add git make && \
    cd /opt && \
    git clone  --depth 1 --branch ${GETH_VER} https:///github.com/ethereum/go-ethereum; cd /opt/go-ethereum && \
    go mod tidy; make all 

FROM alpine:${OS_VER}

COPY --from=builder /opt/go-ethereum/build/bin/* /usr/local/bin