#!/bin/bash

GIT_ROOT_PATH=$(git rev-parse --show-toplevel)
GEN_PATH="${GIT_ROOT_PATH}/api/proto/gen/go"
PROTO_PATH="${GIT_ROOT_PATH}/api/proto/protos"

for PROTO_FILE in $(find $PROTO_PATH -maxdepth 1 -name "*.proto")
do
    SERVICE_NAME=$(basename $PROTO_FILE .proto)
    mkdir -p "${GEN_PATH}/${SERVICE_NAME}/yes"

    protoc --proto_path "${PROTO_PATH}" \
            --go_out "${GEN_PATH}/${SERVICE_NAME}/yes" --go_opt paths=source_relative \
            --go-grpc_out "${GEN_PATH}/${SERVICE_NAME}/yes" --go-grpc_opt paths=source_relative \
            --grpc-gateway_out "${GEN_PATH}/${SERVICE_NAME}/yes" \
            --grpc-gateway_opt logtostderr=true \
            --grpc-gateway_opt paths=source_relative \
            ${PROTO_FILE}
done

#  --grpc-gateway_opt generate_unbound_methods=true
# https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/#generate_unbound_methods
