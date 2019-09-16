package converter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/chrusty/protoc-gen-jsonschema/internal/converter/testdata"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	protocBinary         = "/bin/protoc"
	sampleProtoDirectory = "testdata/proto"
	sampleProtos         = make(map[string]sampleProto)
)

type sampleProto struct {
	AllowNullValues           bool
	ExpectedJSONSchema        []string
	FilesToGenerate           []string
	ProtoFileName             string
	UseProtoAndJSONFieldNames bool
}

func TestGenerateJsonSchema(t *testing.T) {

	// Make sure we have "protoc" installed and available:
	testForProtocBinary(t)

	// Configure the list of sample protos to test, and their expected JSON-Schemas:
	configureSampleProtos()

	// Convert the protos, compare the results against the expected JSON-Schemas:
	testConvertSampleProto(t, sampleProtos["Comments"])
	testConvertSampleProto(t, sampleProtos["ArrayOfMessages"])
	testConvertSampleProto(t, sampleProtos["ArrayOfObjects"])
	testConvertSampleProto(t, sampleProtos["ArrayOfPrimitives"])
	testConvertSampleProto(t, sampleProtos["ArrayOfPrimitivesDouble"])
	testConvertSampleProto(t, sampleProtos["EnumCeption"])
	testConvertSampleProto(t, sampleProtos["ImportedEnum"])
	testConvertSampleProto(t, sampleProtos["NestedMessage"])
	testConvertSampleProto(t, sampleProtos["NestedObject"])
	testConvertSampleProto(t, sampleProtos["PayloadMessage"])
	testConvertSampleProto(t, sampleProtos["SeveralEnums"])
	testConvertSampleProto(t, sampleProtos["SeveralMessages"])
	testConvertSampleProto(t, sampleProtos["ArrayOfEnums"])
	testConvertSampleProto(t, sampleProtos["Maps"])
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

func testConvertSampleProto(t *testing.T, sampleProto sampleProto) {

	// Make a Logrus logger:
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetOutput(os.Stderr)

	// Use the logger to make a Converter:
	protoConverter := New(logger)
	protoConverter.AllowNullValues = sampleProto.AllowNullValues
	protoConverter.UseProtoAndJSONFieldnames = sampleProto.UseProtoAndJSONFieldNames

	// Open the sample proto file:
	sampleProtoFileName := fmt.Sprintf("%v/%v", sampleProtoDirectory, sampleProto.ProtoFileName)

	// Prepare to run the "protoc" command (generates a CodeGeneratorRequest):
	protocCommand := exec.Command(protocBinary, "--descriptor_set_out=/dev/stdout", "--include_imports", fmt.Sprintf("--proto_path=%v", sampleProtoDirectory), sampleProtoFileName)
	var protocCommandOutput bytes.Buffer
	errChan := &bytes.Buffer{}
	protocCommand.Stdout = &protocCommandOutput
	protocCommand.Stderr = errChan
	// Run the command:
	err := protocCommand.Run()
	assert.NoError(t, err, "Unable to prepare a codeGeneratorRequest using protoc (%v) for sample proto file (%v) (%s)", protocBinary, sampleProtoFileName, strings.TrimSpace(errChan.String()))

	// Unmarshal the output from the protoc command (should be a "FileDescriptorSet"):
	fileDescriptorSet := new(descriptor.FileDescriptorSet)
	err = proto.Unmarshal(protocCommandOutput.Bytes(), fileDescriptorSet)
	assert.NoError(t, err, "Unable to unmarshal proto FileDescriptorSet for sample proto file (%v)", sampleProtoFileName)

	// Prepare a request:
	codeGeneratorRequest := plugin.CodeGeneratorRequest{
		FileToGenerate: sampleProto.FilesToGenerate,
		ProtoFile:      fileDescriptorSet.GetFile(),
	}

	// Perform the conversion:
	response, err := protoConverter.convert(&codeGeneratorRequest)
	assert.NoError(t, err, "Unable to convert sample proto file (%v)", sampleProtoFileName)
	assert.Equal(t, len(sampleProto.ExpectedJSONSchema), len(response.File), "Incorrect number of JSON-Schema files returned for sample proto file (%v)", sampleProtoFileName)
	if len(sampleProto.ExpectedJSONSchema) != len(response.File) {
		t.Fail()
	} else {
		for responseFileIndex, responseFile := range response.File {
			assert.Equal(t, sampleProto.ExpectedJSONSchema[responseFileIndex], *responseFile.Content, "Incorrect JSON-Schema returned for sample proto file (%v)", sampleProtoFileName)
		}
	}

}

func configureSampleProtos() {
	// ArrayOfMessages:
	sampleProtos["ArrayOfMessages"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.PayloadMessage, testdata.ArrayOfMessages},
		FilesToGenerate:    []string{"ArrayOfMessages.proto", "PayloadMessage.proto"},
		ProtoFileName:      "ArrayOfMessages.proto",
	}

	// ArrayOfObjects:
	sampleProtos["ArrayOfObjects"] = sampleProto{
		AllowNullValues:    true,
		ExpectedJSONSchema: []string{testdata.ArrayOfObjects},
		FilesToGenerate:    []string{"ArrayOfObjects.proto"},
		ProtoFileName:      "ArrayOfObjects.proto",
	}

	// ArrayOfPrimitives:
	sampleProtos["ArrayOfPrimitives"] = sampleProto{
		AllowNullValues:    true,
		ExpectedJSONSchema: []string{testdata.ArrayOfPrimitives},
		FilesToGenerate:    []string{"ArrayOfPrimitives.proto"},
		ProtoFileName:      "ArrayOfPrimitives.proto",
	}

	// ArrayOfPrimitives:
	sampleProtos["ArrayOfPrimitivesDouble"] = sampleProto{
		AllowNullValues:           true,
		ExpectedJSONSchema:        []string{testdata.ArrayOfPrimitivesDouble},
		FilesToGenerate:           []string{"ArrayOfPrimitives.proto"},
		ProtoFileName:             "ArrayOfPrimitives.proto",
		UseProtoAndJSONFieldNames: true,
	}

	// EnumCeption:
	sampleProtos["EnumCeption"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.PayloadMessage, testdata.ImportedEnum, testdata.EnumCeption},
		FilesToGenerate:    []string{"Enumception.proto", "PayloadMessage.proto", "ImportedEnum.proto"},
		ProtoFileName:      "Enumception.proto",
	}

	// ImportedEnum:
	sampleProtos["ImportedEnum"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.ImportedEnum},
		FilesToGenerate:    []string{"ImportedEnum.proto"},
		ProtoFileName:      "ImportedEnum.proto",
	}

	// NestedMessage:
	sampleProtos["NestedMessage"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.PayloadMessage, testdata.NestedMessage},
		FilesToGenerate:    []string{"NestedMessage.proto", "PayloadMessage.proto"},
		ProtoFileName:      "NestedMessage.proto",
	}

	// NestedObject:
	sampleProtos["NestedObject"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.NestedObject},
		FilesToGenerate:    []string{"NestedObject.proto"},
		ProtoFileName:      "NestedObject.proto",
	}

	// PayloadMessage:
	sampleProtos["PayloadMessage"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.PayloadMessage},
		FilesToGenerate:    []string{"PayloadMessage.proto"},
		ProtoFileName:      "PayloadMessage.proto",
	}

	// SeveralEnums:
	sampleProtos["SeveralEnums"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.FirstEnum, testdata.SecondEnum},
		FilesToGenerate:    []string{"SeveralEnums.proto"},
		ProtoFileName:      "SeveralEnums.proto",
	}

	// SeveralMessages:
	sampleProtos["SeveralMessages"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.FirstMessage, testdata.SecondMessage},
		FilesToGenerate:    []string{"SeveralMessages.proto"},
		ProtoFileName:      "SeveralMessages.proto",
	}

	// ArrayOfEnums:
	sampleProtos["ArrayOfEnums"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.ArrayOfEnums},
		FilesToGenerate:    []string{"ArrayOfEnums.proto"},
		ProtoFileName:      "ArrayOfEnums.proto",
	}

	// Maps:
	sampleProtos["Maps"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.Maps},
		FilesToGenerate:    []string{"Maps.proto"},
		ProtoFileName:      "Maps.proto",
	}

	// Comments:
	sampleProtos["Comments"] = sampleProto{
		AllowNullValues:    false,
		ExpectedJSONSchema: []string{testdata.MessageWithComments},
		FilesToGenerate:    []string{"MessageWithComments.proto"},
		ProtoFileName:      "MessageWithComments.proto",
	}
}
