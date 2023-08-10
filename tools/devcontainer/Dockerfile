# FROM golang:1.20-bookworm
FROM ubuntu:22.04

# Basic env
# ENV DEBIAN_FRONTEND noninteractive
# # User env
# ENV PROTOC_VERSION 3.19.6
# ENV ARCH aarch_64
# # Dev env
# ENV PROTOCPATH "/root/protoc"
# ENV PATH "$PATH:$PROTOCPATH/bin"

WORKDIR /home

# Run
RUN \
    ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
    apt-get update -y && \
    apt-get install unzip curl net-tools -y

# RUN \
#     go install golang.org/x/tools/gopls@latest && \
#     go install github.com/ramya-rao-a/go-outline@latest

# RUN \
#     curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${ARCH}.zip" && \
#     unzip "protoc-${PROTOC_VERSION}-linux-${ARCH}.zip" -d "${PROTOCPATH}" && \
#     rm -rf "protoc-${PROTOC_VERSION}-linux-${ARCH}.zip" && \
#     go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 && \
#     go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3 && \
#     go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16


# RUN \
#     apt-get install python3-pip -y

apt-get install python3.11
update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.11 1


apt-get install python3=3.10.6-1~22.04
apt-get install python3-pip