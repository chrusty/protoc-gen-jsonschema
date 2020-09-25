default: build

.PHONY: build
build:
	@echo "Generating binary (protoc-gen-jsonschema) ..."
	@mkdir -p bin
	@go build -o bin/protoc-gen-jsonschema cmd/protoc-gen-jsonschema/main.go

.PHONY: fmt
fmt:
	@gofmt -s -w .
	@goimports -w -local github.com/chrusty/protoc-gen-jsonschema .

.PHONY: install
install:
	@GO111MODULE=on go get github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema && go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema

.PHONY: build_linux
build_linux:
	@echo "Generating Linux-amd64 binary (protoc-gen-jsonschema.linux-amd64) ..."
	@GOOS=linux GOARCH=amd64 go build -o protoc-gen-jsonschema.linux-amd64

PROTO_PATH ?= "internal/converter/testdata/proto"
.PHONY: samples
samples:
	@echo "Generating sample JSON-Schemas ..."
	@mkdir -p jsonschemas
	@PATH=./bin:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfMessages.proto || echo "No messages found (ArrayOfMessages.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfObjects.proto || echo "No messages found (ArrayOfObjects.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfPrimitives.proto || echo "No messages found (ArrayOfPrimitives.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Enumception.proto || echo "No messages found (Enumception.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ImportedEnum.proto || echo "No messages found (ImportedEnum.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/NestedMessage.proto || echo "No messages found (NestedMessage.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/NestedObject.proto || echo "No messages found (NestedObject.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/PayloadMessage.proto || echo "No messages found (PayloadMessage.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/SeveralEnums.proto || echo "No messages found (SeveralEnums.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/SeveralMessages.proto || echo "No messages found (SeveralMessages.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Timestamp.proto || echo "No messages found (Timestamp.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=all_fields_required:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/PayloadMessage2.proto || echo "No messages found (PayloadMessage2.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/ArrayOfEnums.proto || echo "No messages found (SeveralMessages.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Maps.proto || echo "No messages found (Maps.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/MessageWithComments.proto || echo "No messages found (MessageWithComments.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Proto2Required.proto || echo "No messages found (Proto2Required.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Proto2NestedMessage.proto || echo "No messages found (Proto2NestedMessage.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/GoogleValue.proto || echo "No messages found (GoogleValue.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=all_fields_required:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/Proto2NestedObject.proto || echo "No messages found (Proto2NestedObject.proto)"
	@PATH=./bin:$$PATH; protoc -I /usr/include --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/WellKnown.proto || echo "No messages found (WellKnown.proto)"
	@PATH=./bin:$$PATH; protoc --jsonschema_out=jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/NoPackage.proto
	@PATH=./bin:$$PATH; protoc --jsonschema_out=messages=[MessageKind10+MessageKind11+MessageKind12]:jsonschemas --proto_path=${PROTO_PATH} ${PROTO_PATH}/TwelveMessages.proto || echo "No messages found (TwelveMessages.proto)"

.PHONY: test
test:
	@go test ./... -cover -v
