package converter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/chrusty/protoc-gen-jsonschema/internal/converter/testdata"
)

const (
	sampleProtoDirectory = "testdata/proto"
)

type sampleProto struct {
	Flags              ConverterFlags
	ExpectedJSONSchema []string
	FilesToGenerate    []string
	ProtoFileName      string
	RefPrefix          string
	TargetedMessages   []string
}

func TestGenerateJsonSchema(t *testing.T) {

	// Configure the list of sample protos to test, and their expected JSON-Schemas:
	sampleProtos := configureSampleProtos()

	// Convert the protos, compare the results against the expected JSON-Schemas:
	for name, sampleProto := range sampleProtos {
		t.Run(name, func(t *testing.T) {
			testConvertSampleProto(t, sampleProto)
		})
	}
}

func testConvertSampleProto(t *testing.T, sampleProto sampleProto) {
	t.Helper()

	// Make a Logrus logger:
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetOutput(os.Stderr)

	// Use the logger to make a Converter:
	protoConverter := New(logger)
	protoConverter.Flags = sampleProto.Flags
	protoConverter.refPrefix = sampleProto.RefPrefix

	// Open the sample proto file:
	sampleProtoFileName := fmt.Sprintf("%v/%v", sampleProtoDirectory, sampleProto.ProtoFileName)
	fileDescriptorSet := mustReadProtoFiles(t, sampleProtoDirectory, sampleProto.ProtoFileName)

	// Prepare a request:
	codeGeneratorRequest := plugin.CodeGeneratorRequest{
		FileToGenerate: sampleProto.FilesToGenerate,
		ProtoFile:      fileDescriptorSet.GetFile(),
	}

	if len(sampleProto.TargetedMessages) > 0 {
		arg := fmt.Sprintf("messages=[%s]", strings.Join(sampleProto.TargetedMessages, messageDelimiter))
		codeGeneratorRequest.Parameter = &arg
	}

	// Perform the conversion:
	response, err := protoConverter.convert(&codeGeneratorRequest)
	assert.NoError(t, err, "Unable to convert sample proto file (%v)", sampleProtoFileName)
	assert.Equal(t, len(sampleProto.ExpectedJSONSchema), len(response.File), "Incorrect number of JSON-Schema files returned for sample proto file (%v)", sampleProtoFileName)
	if len(sampleProto.ExpectedJSONSchema) != len(response.File) {
		t.Fail()
	} else {
		for responseFileIndex, responseFile := range response.File {
			assert.Equal(t, strings.TrimSpace(sampleProto.ExpectedJSONSchema[responseFileIndex]), *responseFile.Content, "Incorrect JSON-Schema returned for sample proto file (%v)", sampleProtoFileName)
		}
	}

	// Return now if we have no files:
	if len(response.File) == 0 {
		return
	}

	// Check for the correct prefix:
	if protoConverter.Flags.PrefixSchemaFilesWithPackage {
		assert.Contains(t, response.File[0].GetName(), "samples")
	} else {
		assert.NotContains(t, response.File[0].GetName(), "samples")
	}
}

func configureSampleProtos() map[string]sampleProto {
	return map[string]sampleProto{
		"ArrayOfMessages": {
			ExpectedJSONSchema: []string{testdata.PayloadMessage, testdata.ArrayOfMessages},
			FilesToGenerate:    []string{"ArrayOfMessages.proto", "PayloadMessage.proto"},
			ProtoFileName:      "ArrayOfMessages.proto",
		},
		"ArrayOfObjects": {
			Flags:              ConverterFlags{AllowNullValues: true},
			ExpectedJSONSchema: []string{testdata.ArrayOfObjects},
			FilesToGenerate:    []string{"ArrayOfObjects.proto"},
			ProtoFileName:      "ArrayOfObjects.proto",
		},
		"ArrayOfPrimitives": {
			Flags:              ConverterFlags{AllowNullValues: true},
			ExpectedJSONSchema: []string{testdata.ArrayOfPrimitives},
			FilesToGenerate:    []string{"ArrayOfPrimitives.proto"},
			ProtoFileName:      "ArrayOfPrimitives.proto",
		},
		"BytesPayload": {
			ExpectedJSONSchema: []string{testdata.BytesPayload},
			FilesToGenerate:    []string{"BytesPayload.proto"},
			ProtoFileName:      "BytesPayload.proto",
		},
		"ArrayOfPrimitivesDouble": {
			Flags: ConverterFlags{
				AllowNullValues:           true,
				UseProtoAndJSONFieldNames: true,
			},
			ExpectedJSONSchema: []string{testdata.ArrayOfPrimitivesDouble},
			FilesToGenerate:    []string{"ArrayOfPrimitives.proto"},
			ProtoFileName:      "ArrayOfPrimitives.proto",
		},
		"EnumNestedReference": {
			ExpectedJSONSchema: []string{testdata.EnumNestedReference},
			FilesToGenerate:    []string{"EnumNestedReference.proto"},
			ProtoFileName:      "EnumNestedReference.proto",
		},
		"EnumWithMessage": {
			ExpectedJSONSchema: []string{testdata.EnumWithMessage},
			FilesToGenerate:    []string{"EnumWithMessage.proto"},
			ProtoFileName:      "EnumWithMessage.proto",
		},
		"EnumImport": {
			ExpectedJSONSchema: []string{testdata.EnumImport},
			FilesToGenerate:    []string{"ImportEnum.proto"},
			ProtoFileName:      "ImportEnum.proto",
		},
		"EnumCeption": {
			ExpectedJSONSchema: []string{testdata.PayloadMessage, testdata.ImportedEnum, testdata.EnumCeption},
			FilesToGenerate:    []string{"Enumception.proto", "PayloadMessage.proto", "ImportedEnum.proto"},
			ProtoFileName:      "Enumception.proto",
		},
		"ImportedEnum": {
			ExpectedJSONSchema: []string{testdata.ImportedEnum},
			FilesToGenerate:    []string{"ImportedEnum.proto"},
			ProtoFileName:      "ImportedEnum.proto",
		},
		"NestedMessage": {
			ExpectedJSONSchema: []string{testdata.PayloadMessage, testdata.NestedMessage},
			FilesToGenerate:    []string{"NestedMessage.proto", "PayloadMessage.proto"},
			ProtoFileName:      "NestedMessage.proto",
		},
		"NestedObject": {
			ExpectedJSONSchema: []string{testdata.NestedObject},
			FilesToGenerate:    []string{"NestedObject.proto"},
			ProtoFileName:      "NestedObject.proto",
		},
		"PayloadMessage": {
			ExpectedJSONSchema: []string{testdata.PayloadMessage},
			FilesToGenerate:    []string{"PayloadMessage.proto"},
			ProtoFileName:      "PayloadMessage.proto",
		},
		"SeveralEnums": {
			ExpectedJSONSchema: []string{testdata.FirstEnum, testdata.SecondEnum},
			FilesToGenerate:    []string{"SeveralEnums.proto"},
			ProtoFileName:      "SeveralEnums.proto",
		},
		"SeveralMessages": {
			ExpectedJSONSchema: []string{testdata.FirstMessage, testdata.SecondMessage},
			FilesToGenerate:    []string{"SeveralMessages.proto"},
			ProtoFileName:      "SeveralMessages.proto",
		},
		"ArrayOfEnums": {
			ExpectedJSONSchema: []string{testdata.ArrayOfEnums},
			FilesToGenerate:    []string{"ArrayOfEnums.proto"},
			ProtoFileName:      "ArrayOfEnums.proto",
		},
		"Maps": {
			ExpectedJSONSchema: []string{testdata.Maps},
			FilesToGenerate:    []string{"Maps.proto"},
			ProtoFileName:      "Maps.proto",
		},
		"Comments": {
			ExpectedJSONSchema: []string{testdata.MessageWithComments},
			FilesToGenerate:    []string{"MessageWithComments.proto"},
			ProtoFileName:      "MessageWithComments.proto",
		},
		"SelfReference": {
			ExpectedJSONSchema: []string{testdata.SelfReference},
			FilesToGenerate:    []string{"SelfReference.proto"},
			ProtoFileName:      "SelfReference.proto",
		},
		"CyclicalReference": {
			ExpectedJSONSchema: []string{testdata.CyclicalReferenceMessageM, testdata.CyclicalReferenceMessageFoo, testdata.CyclicalReferenceMessageBar, testdata.CyclicalReferenceMessageBaz},
			FilesToGenerate:    []string{"CyclicalReference.proto"},
			ProtoFileName:      "CyclicalReference.proto",
		},
		"WellKnown": {
			ExpectedJSONSchema: []string{testdata.WellKnown},
			FilesToGenerate:    []string{"WellKnown.proto"},
			ProtoFileName:      "WellKnown.proto",
		},
		"Timestamp": {
			ExpectedJSONSchema: []string{testdata.Timestamp},
			FilesToGenerate:    []string{"Timestamp.proto"},
			ProtoFileName:      "Timestamp.proto",
		},
		"NoPackage": {
			ExpectedJSONSchema: []string{},
			FilesToGenerate:    []string{},
			ProtoFileName:      "NoPackage.proto",
		},
		"PackagePrefix": {
			Flags:              ConverterFlags{PrefixSchemaFilesWithPackage: true},
			ExpectedJSONSchema: []string{testdata.Timestamp},
			FilesToGenerate:    []string{"Timestamp.proto"},
			ProtoFileName:      "Timestamp.proto",
		},
		"Proto2Required": {
			ExpectedJSONSchema: []string{testdata.Proto2Required},
			FilesToGenerate:    []string{"Proto2Required.proto"},
			ProtoFileName:      "Proto2Required.proto",
		},
		"AllRequired": {
			Flags:              ConverterFlags{AllFieldsRequired: true},
			ExpectedJSONSchema: []string{testdata.PayloadMessage2},
			FilesToGenerate:    []string{"PayloadMessage2.proto"},
			ProtoFileName:      "PayloadMessage2.proto",
		},
		"Proto2NestedMessage": {
			ExpectedJSONSchema: []string{testdata.Proto2PayloadMessage, testdata.Proto2NestedMessage},
			FilesToGenerate:    []string{"Proto2PayloadMessage.proto", "Proto2NestedMessage.proto"},
			ProtoFileName:      "Proto2NestedMessage.proto",
		},
		"Proto2NestedObject": {
			Flags:              ConverterFlags{AllFieldsRequired: true},
			ExpectedJSONSchema: []string{testdata.Proto2NestedObject},
			FilesToGenerate:    []string{"Proto2NestedObject.proto"},
			ProtoFileName:      "Proto2NestedObject.proto",
		},
		"TargetedMessages": {
			TargetedMessages:   []string{"MessageKind10", "MessageKind11", "MessageKind12"},
			ExpectedJSONSchema: []string{testdata.MessageKind10, testdata.MessageKind11, testdata.MessageKind12},
			FilesToGenerate:    []string{"TwelveMessages.proto"},
			ProtoFileName:      "TwelveMessages.proto",
		},
		"GoogleValue": {
			ExpectedJSONSchema: []string{testdata.GoogleValue},
			FilesToGenerate:    []string{"GoogleValue.proto"},
			ProtoFileName:      "GoogleValue.proto",
		},
		"JSONFields": {
			Flags:              ConverterFlags{UseJSONFieldnamesOnly: true},
			ExpectedJSONSchema: []string{testdata.JSONFields},
			FilesToGenerate:    []string{"JSONFields.proto"},
			ProtoFileName:      "JSONFields.proto",
		},
		"OneOf": {
			Flags:              ConverterFlags{AllFieldsRequired: true, EnforceOneOf: true},
			ExpectedJSONSchema: []string{testdata.OneOf},
			FilesToGenerate:    []string{"OneOf.proto"},
			ProtoFileName:      "OneOf.proto",
		},
		"HiddenFields": {
			ExpectedJSONSchema: []string{testdata.FieldOptions, testdata.HiddenFields},
			FilesToGenerate:    []string{"options.proto", "HiddenFields.proto"},
			ProtoFileName:      "HiddenFields.proto",
		},
		"Proto3Required": {
			ExpectedJSONSchema: []string{testdata.FieldOptions, testdata.Proto3Required},
			FilesToGenerate:    []string{"options.proto", "Proto3Required.proto"},
			ProtoFileName:      "Proto3Required.proto",
		},
		"RefPrefix": {
			ExpectedJSONSchema: []string{testdata.RefPrefixM, testdata.RefPrefixFoo, testdata.RefPrefixBar, testdata.RefPrefixBaz},
			FilesToGenerate:    []string{"CyclicalReference.proto"},
			ProtoFileName:      "CyclicalReference.proto",
			RefPrefix:          "#/definitions/",
		},
	}
}

// Load the specified .proto files into a FileDescriptorSet. Any errors in loading/parsing will
// immediately fail the test.
func mustReadProtoFiles(t *testing.T, includePath string, filenames ...string) *descriptor.FileDescriptorSet {
	protocBinary, err := exec.LookPath("protoc")
	if err != nil {
		t.Fatalf("Can't find 'protoc' binary in $PATH: %s", err.Error())
	}

	// Use protoc to output descriptor info for the specified .proto files.
	var args []string
	args = append(args, "--descriptor_set_out=/dev/stdout")
	args = append(args, "--include_source_info")
	args = append(args, "--include_imports")
	args = append(args, "-I../../")
	args = append(args, "--proto_path="+includePath)
	args = append(args, filenames...)
	cmd := exec.Command(protocBinary, args...)
	stdoutBuf := bytes.Buffer{}
	stderrBuf := bytes.Buffer{}
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err = cmd.Run()
	if err != nil {
		t.Fatalf("failed to load descriptor set (%s): %s: %s",
			strings.Join(cmd.Args, " "), err.Error(), stderrBuf.String())
	}
	fds := &descriptor.FileDescriptorSet{}
	err = proto.Unmarshal(stdoutBuf.Bytes(), fds)
	if err != nil {
		t.Fatalf("failed to parse protoc output as FileDescriptorSet: %s", err.Error())
	}
	return fds
}
