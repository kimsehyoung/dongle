FROM golang:1.20

# Basic env
ENV DEBIAN_FRONTEND noninteractive
# User env
ENV OS linux
ENV PROTOC_VERSION 3.19.6
ENV ARCH aarch_64
# Dev env
ENV PROTOCPATH "/root/protoc"
ENV PATH "$PATH:$PROTOCPATH/bin"

WORKDIR /home

# Run
RUN ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime

RUN apt-get update -y && \
    apt-get install unzip curl -y

RUN \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/ramya-rao-a/go-outline@latest


RUN \
    curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip" && \
    unzip "protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip" -d "${PROTOCPATH}" && \
    rm -rf "protoc-${PROTOC_VERSION}-linux-${ARCH}.zip" && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16

