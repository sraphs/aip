#!/bin/bash

set -eux
cd $(dirname $0)

SCHEMA=testdata
OUT_PATH=testdata

[ -d ${OUT_PATH} ] || mkdir ${OUT_PATH}

INCLUDE_PATH=third_party

API_PROTO_FILES=$(find ${SCHEMA} -name "*.proto")
protoc \
    --proto_path ./${SCHEMA} \
    --proto_path ./${INCLUDE_PATH} \
    --go_out=paths=source_relative:${OUT_PATH} \
    --go-grpc_out=paths=source_relative:${OUT_PATH} \
    ${API_PROTO_FILES}

cd -
