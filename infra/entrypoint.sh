#!/bin/sh

BINARY="./main"
GO_ENTRYPOINT="./cmd/service/main.go"
BUILD_CMD="go build -o ${BINARY} ${GO_ENTRYPOINT}" 

CompileDaemon \
    --build="${BUILD_CMD}" \
    --command="${BINARY}"