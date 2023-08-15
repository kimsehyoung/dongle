FROM ubuntu:22.04

# ENV
ENV PROTOC_VERSION 3.19.6
ENV GO_VERSION 1.20
ENV PATH "$PATH:/usr/local/protoc/bin"
ENV PATH "$PATH:/usr/local/go/bin"


WORKDIR /home


# Common Packages
RUN \
    ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
    apt-get update -y && \
    apt-get install unzip curl git net-tools -y

# Protoc
RUN \
    [ "$(uname -m)" = "x86_64" ] && ARCH="x86_64" || ARCH="aarch_64" && \
    curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${ARCH}.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-${ARCH}.zip -d /usr/local/protoc && \
    rm -rf protoc-${PROTOC_VERSION}-linux-${ARCH}.zip

# Golang
RUN \
    [ "$(uname -m)" = "x86_64" ] && ARCH="amd64" || ARCH="arm64" && \
    curl -LO https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    rm -rf go${GO_VERSION}.linux-${ARCH}.tar.gz && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16

# Python
RUN \
    apt-get install python3.11 python3-pip -y && \
    apt-get install ffmpeg -y
