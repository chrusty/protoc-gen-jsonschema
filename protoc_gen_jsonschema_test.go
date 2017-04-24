package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	log "github.com/Sirupsen/logrus"
	testdata "github.com/chrusty/protoc-gen-jsonschema/testdata"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	assert "github.com/stretchr/testify/assert"
)

var (
	protocBinary         = "/bin/protoc"
	sampleProtoDirectory = "testdata/proto"
	sampleProtos         = make(map[string]SampleProto)
)

type SampleProto struct {
	protoFileName      string
	filesToGenerate    []string
	expectedJsonSchema []string
}

func TestGenerateJsonSchema(t *testing.T) {
	// We only want to see "Info" level logs and above (there's a LOT of debug otherwise):
	log.SetLevel(log.InfoLevel)

	// Make sure we have "protoc" installed and available:
	testForProtocBinary(t)

	// Configure the list of sample protos to test, and their expected JSON-Schemas:
	configureSampleProtos()

	// Convert the protos, compare the results against the expected JSON-Schemas:
	testConvertSampleProtos(t, sampleProtos["ArrayOfMessages"])
	testConvertSampleProtos(t, sampleProtos["ArrayOfObjects"])
	testConvertSampleProtos(t, sampleProtos["ArrayOfPrimitives"])
	testConvertSampleProtos(t, sampleProtos["EnumCeption"])
	testConvertSampleProtos(t, sampleProtos["ImportedEnum"])
	testConvertSampleProtos(t, sampleProtos["NestedMessage"])
	testConvertSampleProtos(t, sampleProtos["NestedObject"])
	testConvertSampleProtos(t, sampleProtos["PayloadMessage"])
	testConvertSampleProtos(t, sampleProtos["SeveralEnums"])
	testConvertSampleProtos(t, sampleProtos["SeveralMessages"])
}

func testForProtocBinary(t *testing.T) {
	path, err := exec.LookPath("protoc")
	if err != nil {
		assert.NoError(t, err, "Can't find 'protoc' binary in $PATH")
	} else {
		protocBinary = path
		log.Infof("Found 'protoc' binary (%v)", protocBinary)
	}
}

func testConvertSampleProto(t *testing.T) {

	// Go through the sample protos:
	for _, sampleProto := range sampleProtos {
		log.Infof("SampleProto: %v", sampleProto.protoFileName)
	}

}

func testConvertSampleProtos(t *testing.T, sampleProto SampleProto) {

	// Open the sample proto file:
	sampleProtoFileName := fmt.Sprintf("%v/%v", sampleProtoDirectory, sampleProto.protoFileName)

	// Prepare to run the "protoc" command (generates a CodeGeneratorRequest):
	protocCommand := exec.Command(protocBinary, "--descriptor_set_out=/dev/stdout", "--include_imports", fmt.Sprintf("--proto_path=%v", sampleProtoDirectory), sampleProtoFileName)
	var protocCommandOutput bytes.Buffer
	protocCommand.Stdout = &protocCommandOutput

	// Run the command:
	err := protocCommand.Run()
	assert.NoError(t, err, "Unable to prepare a codeGeneratorRequest using protoc (%v) for sample proto file (%v)", protocBinary, sampleProtoFileName)

	// Unmarshal the output from the protoc command (should be a "FileDescriptorSet"):
	fileDescriptorSet := new(descriptor.FileDescriptorSet)
	err = proto.Unmarshal(protocCommandOutput.Bytes(), fileDescriptorSet)
	assert.NoError(t, err, "Unable to unmarshal proto FileDescriptorSet for sample proto file (%v)", sampleProtoFileName)

	// Prepare a request:
	codeGeneratorRequest := plugin.CodeGeneratorRequest{
		FileToGenerate: sampleProto.filesToGenerate,
		ProtoFile:      fileDescriptorSet.GetFile(),
	}

	// Perform the conversion:
	response, err := convert(&codeGeneratorRequest)
	assert.NoError(t, err, "Unable to convert sample proto file (%v)", sampleProtoFileName)
	assert.EqualValues(t, len(sampleProto.expectedJsonSchema), len(response.File), "Incorrect number of JSON-Schema files returned")
	if len(sampleProto.expectedJsonSchema) != len(response.File) {
		t.Fail()
	} else {
		for responseFileIndex, responseFile := range response.File {
			assert.EqualValues(t, sampleProto.expectedJsonSchema[responseFileIndex], *responseFile.Content, "Incorrect JSON-Schema returned")
		}
	}

}

func configureSampleProtos() {
	// ArrayOfMessages:
	sampleProtos["ArrayOfMessages"] = SampleProto{
		protoFileName:      "ArrayOfMessages.proto",
		filesToGenerate:    []string{"ArrayOfMessages.proto", "PayloadMessage.proto"},
		expectedJsonSchema: []string{testdata.PayloadMessage, testdata.ArrayOfMessages},
	}

	// ArrayOfObjects:
	sampleProtos["ArrayOfObjects"] = SampleProto{
		protoFileName:      "ArrayOfObjects.proto",
		filesToGenerate:    []string{"ArrayOfObjects.proto"},
		expectedJsonSchema: []string{testdata.ArrayOfObjects},
	}

	// ArrayOfPrimitives:
	sampleProtos["ArrayOfPrimitives"] = SampleProto{
		protoFileName:      "ArrayOfPrimitives.proto",
		filesToGenerate:    []string{"ArrayOfPrimitives.proto"},
		expectedJsonSchema: []string{testdata.ArrayOfPrimitives},
	}

	// EnumCeption:
	sampleProtos["EnumCeption"] = SampleProto{
		protoFileName:      "EnumCeption.proto",
		filesToGenerate:    []string{"EnumCeption.proto", "PayloadMessage.proto", "ImportedEnum.proto"},
		expectedJsonSchema: []string{testdata.PayloadMessage, testdata.ImportedEnum, testdata.EnumCeption},
	}

	// ImportedEnum:
	sampleProtos["ImportedEnum"] = SampleProto{
		protoFileName:      "ImportedEnum.proto",
		filesToGenerate:    []string{"ImportedEnum.proto"},
		expectedJsonSchema: []string{testdata.ImportedEnum},
	}

	// NestedMessage:
	sampleProtos["NestedMessage"] = SampleProto{
		protoFileName:      "NestedMessage.proto",
		filesToGenerate:    []string{"NestedMessage.proto", "PayloadMessage.proto"},
		expectedJsonSchema: []string{testdata.PayloadMessage, testdata.NestedMessage},
	}

	// NestedObject:
	sampleProtos["NestedObject"] = SampleProto{
		protoFileName:      "NestedObject.proto",
		filesToGenerate:    []string{"NestedObject.proto"},
		expectedJsonSchema: []string{testdata.NestedObject},
	}

	// PayloadMessage:
	sampleProtos["PayloadMessage"] = SampleProto{
		protoFileName:      "PayloadMessage.proto",
		filesToGenerate:    []string{"PayloadMessage.proto"},
		expectedJsonSchema: []string{testdata.PayloadMessage},
	}

	// SeveralEnums:
	sampleProtos["SeveralEnums"] = SampleProto{
		protoFileName:      "SeveralEnums.proto",
		filesToGenerate:    []string{"SeveralEnums.proto"},
		expectedJsonSchema: []string{testdata.FirstEnum, testdata.SecondEnum},
	}

	// SeveralMessages:
	sampleProtos["SeveralMessages"] = SampleProto{
		protoFileName:      "SeveralMessages.proto",
		filesToGenerate:    []string{"SeveralMessages.proto"},
		expectedJsonSchema: []string{testdata.FirstMessage, testdata.SecondMessage},
	}

}
