#!/bin/sh
SRC_DIR=.
DST_DIR=.

# Generate Go code for the service and its dependencies
protoc -I=$SRC_DIR \
       --go_out=$DST_DIR \
       --go-grpc_out=$DST_DIR \
       $SRC_DIR/*.proto

# Replace the generated import path with the correct one
sed -i 's|gokits/libs/server/grpc|./grpc|g' $DST_DIR/*.go
