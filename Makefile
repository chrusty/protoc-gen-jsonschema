default: darwin linux windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o protoc-gen-jsonschema.darwin-amd64

linux:
	GOOS=linux GOARCH=amd64 go build -o protoc-gen-jsonschema.linux-amd64

windows:
	GOOS=windows GOARCH=amd64 go build -o protoc-gen-jsonschema.windows-amd64.exe
