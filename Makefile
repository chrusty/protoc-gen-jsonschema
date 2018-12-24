default: darwin linux windows

build:
	@echo "Generating binary (protoc-gen-jsonschema) ..."
	@go build -o protoc-gen-jsonschema

samples:
	@echo "Generating sample JSON-Schemas ..."
	@mkdir -p jsonschemas
	@PATH=.:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=testdata/proto testdata/proto/ArrayOfMessages.proto 2>/dev/null || echo "No messages found (ArrayOfMessages.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=testdata/proto testdata/proto/ArrayOfObjects.proto 2>/dev/null || echo "No messages found (ArrayOfObjects.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=allow_null_values:jsonschemas --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto 2>/dev/null || echo "No messages found (ArrayOfPrimitives.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=testdata/proto testdata/proto/Enumception.proto 2>/dev/null || echo "No messages found (Enumception.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=testdata/proto testdata/proto/ImportedEnum.proto 2>/dev/null || echo "No messages found (ImportedEnum.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=testdata/proto testdata/proto/NestedMessage.proto 2>/dev/null || echo "No messages found (NestedMessage.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=testdata/proto testdata/proto/NestedObject.proto 2>/dev/null || echo "No messages found (NestedObject.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=testdata/proto testdata/proto/PayloadMessage.proto 2>/dev/null || echo "No messages found (PayloadMessage.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=testdata/proto testdata/proto/SeveralEnums.proto 2>/dev/null || echo "No messages found (SeveralEnums.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=disallow_bigints_as_strings:jsonschemas --proto_path=testdata/proto testdata/proto/SeveralMessages.proto 2>/dev/null || echo "No messages found (SeveralMessages.proto)"
	@PATH=.:$$PATH; protoc --jsonschema_out=string_to_string_map:jsonschemas --proto_path=testdata/proto testdata/proto/StringToStringMap.proto 2>/dev/null || echo "No messages found (StringToStringMap.proto)"

test:
	@go test
