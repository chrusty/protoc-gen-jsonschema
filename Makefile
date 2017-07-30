default: darwin linux windows

darwin:
	@echo "Generating MacOS binary (protoc-gen-jsonschema.darwin-amd64) ..."
	@GOOS=darwin GOARCH=amd64 go build -o protoc-gen-jsonschema.darwin-amd64

linux:
	@echo "Generating Linux binary (protoc-gen-jsonschema.linux-amd64) ..."
	@GOOS=linux GOARCH=amd64 go build -o protoc-gen-jsonschema.linux-amd64

windows:
	@echo "Generating Windows binary (protoc-gen-jsonschema.windows-amd64.exe) ..."
	@GOOS=windows GOARCH=amd64 go build -o protoc-gen-jsonschema.windows-amd64.exe

samples:
	@echo "Generating sample JSON-Schemas ..."
	@mkdir -p jsonschemas
	@PATH=.:$$PATH; for PROTO_FILE in `ls testdata/proto/*.proto`; do protoc --jsonschema_out=disallow_additional_properties:jsonschemas --proto_path=testdata/proto $$PROTO_FILE; done
	# disallow_additional_properties
	# disallow_bigints_as_strings
