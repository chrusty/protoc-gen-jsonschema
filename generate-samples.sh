#!/usr/bin/env bash
PROTO_DIR="samples/proto"
JSONSCHEMA_DIR="samples/jsonschema"
PATH=$PATH:.

# Generate all of the files:
for PROTO_FILE in `ls ${PROTO_DIR}/*.proto`
do
	protoc --jsonschema_out=${JSONSCHEMA_DIR} --proto_path=${PROTO_DIR} "${PROTO_FILE}"
done
