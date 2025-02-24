#!/bin/sh
go mod tidy
API_PROTO_FILES=$(find protobuf -name "*.proto")
mkdir -p gen-go
protoc --proto_path=./protobuf \
       --go_out=paths=source_relative:./gen-go \
       --go-grpc_out=paths=source_relative:./gen-go \
       ${API_PROTO_FILES}