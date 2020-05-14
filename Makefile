default: build

build:
	@echo "Generating binary (protoc-gen-jsonschema) ..."
	@mkdir -p bin
	@go build -o bin/protoc-gen-jsonschema cmd/protoc-gen-jsonschema/main.go

install:
	@GO111MODULE=on go get github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema && go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema

build_linux:
	@echo "Generating Linux-amd64 binary (protoc-gen-jsonschema.linux-amd64) ..."
	@GOOS=linux GOARCH=amd64 go build -o protoc-gen-jsonschema.linux-amd64

PROTO_PATH ?= "internal/converter/testdata/proto"
samples:
	@echo "Generating sample JSON-Schemas ..."
	@mkdir -p jsonschemas
	@PATH=./bin:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfMessages.proto 2>/dev/null || echo "No messages found (ArrayOfMessages.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfObjects.proto 2>/dev/null || echo "No messages found (ArrayOfObjects.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfPrimitives.proto 2>/dev/null || echo "No messages found (ArrayOfPrimitives.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Enumception.proto 2>/dev/null || echo "No messages found (Enumception.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ImportedEnum.proto 2>/dev/null || echo "No messages found (ImportedEnum.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/NestedMessage.proto 2>/dev/null || echo "No messages found (NestedMessage.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/NestedObject.proto 2>/dev/null || echo "No messages found (NestedObject.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/PayloadMessage.proto 2>/dev/null || echo "No messages found (PayloadMessage.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/SeveralEnums.proto 2>/dev/null || echo "No messages found (SeveralEnums.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/SeveralMessages.proto 2>/dev/null || echo "No messages found (SeveralMessages.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfEnums.proto 2>/dev/null || echo "No messages found (SeveralMessages.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Maps.proto 2>/dev/null || echo "No messages found (Maps.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/MessageWithComments.proto 2>/dev/null || echo "No messages found (MessageWithComments.proto)"
	@PATH=./bin:$$PATH; protoc -I /usr/include --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/WellKnown.proto

test:
	@go test ./... -cover -v
