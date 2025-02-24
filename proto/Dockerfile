FROM golang:1.23

RUN apt update && apt install -y --no-install-recommends curl make git unzip apt-utils
ENV GO111MODULE=on
ENV PROTOC_VERSION=23.3
ENV BUFBUILD_VERSION=1.23.0

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip
RUN unzip protoc-$PROTOC_VERSION-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/

RUN go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
RUN go install github.com/envoyproxy/protoc-gen-validate@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# check proto syntax
RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v$BUFBUILD_VERSION/buf-Linux-x86_64 -o /usr/local/bin/buf
RUN chmod +x /usr/local/bin/buf
RUN apt install -y clang-format
