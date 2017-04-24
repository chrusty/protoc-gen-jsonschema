#!/usr/bin/env bash
PROTO_DIR="testdata/proto"
JSONSCHEMA_DIR="jsonschemas"
PATH=$PATH:.

# Ensure that the JSONSchema directory exists
mkdir -p $JSONSCHEMA_DIR

# Generate all of the files:
for PROTO_FILE in `ls ${PROTO_DIR}/*.proto`
do
	protoc --jsonschema_out=${JSONSCHEMA_DIR} --proto_path=${PROTO_DIR} "${PROTO_FILE}"
done
