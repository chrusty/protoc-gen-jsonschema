// protoc plugin which converts .proto to JSON schema
// It is spawned by protoc and generates JSON-schema files.
// "Heavily influenced" by Google's "protog-gen-bq-schema"
//
// usage:
//  $ bin/protoc --jsonschema_out=path/to/outdir foo.proto
//
package main

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/sirupsen/logrus"
	"github.com/sixt/protoc-gen-jsonschema/internal/converter"
)

func main() {

	// Make a Logrus logger (default to INFO):
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stderr)

	// Use the logger to make a Converter:
	protoConverter := converter.New(logger)

	// Convert the generator request:
	var ok = true
	logger.Debug("Processing code generator request")
	res, err := protoConverter.ConvertFrom(os.Stdin)
	if err != nil {
		ok = false
		if res == nil {
			message := fmt.Sprintf("Failed to read input: %v", err)
			res = &plugin.CodeGeneratorResponse{
				Error: &message,
			}
		}
	}

	logger.Debug("Serializing code generator response")
	data, err := proto.Marshal(res)
	if err != nil {
		logger.WithError(err).Fatal("Cannot marshal response")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		logger.WithError(err).Fatal("Failed to write response")
	}

	if ok {
		logger.Debug("Succeeded to process code generator request")
	} else {
		logger.Warn("Failed to process code generator but successfully sent the error to protoc")
		os.Exit(1)
	}
}
