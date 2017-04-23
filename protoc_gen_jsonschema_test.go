package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	log "github.com/Sirupsen/logrus"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	assert "github.com/stretchr/testify/assert"
)

var (
	protocBinary         = "/bin/protoc"
	sampleProtoDirectory = "samples/proto"
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
	// testConvertSampleProtos(t, sampleProtos["ArrayOfMessages"])
	testConvertSampleProtos(t, sampleProtos["ArrayOfObjects"])
	testConvertSampleProtos(t, sampleProtos["ArrayOfPrimitives"])
	// testConvertSampleProtos(t, sampleProtos["EnumCeption"])
	testConvertSampleProtos(t, sampleProtos["ImportedEnum"])
	// testConvertSampleProtos(t, sampleProtos["NestedMessage"])
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
	protocCommand := exec.Command(protocBinary, "--descriptor_set_out=/dev/stdout", fmt.Sprintf("--proto_path=%v", sampleProtoDirectory), sampleProtoFileName)
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
	for responseFileIndex, responseFile := range response.File {
		assert.EqualValues(t, sampleProto.expectedJsonSchema[responseFileIndex], *responseFile.Content, "Incorrect JSON-Schema returned")
	}

}

func configureSampleProtos() {
	// ArrayOfMessages:
	sampleProtos["ArrayOfMessages"] = SampleProto{
		protoFileName:   "ArrayOfMessages.proto",
		filesToGenerate: []string{"ArrayOfMessages.proto", "PayloadMessage.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "payload": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "properties": {
                    "complete": {
                        "type": "boolean"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "name": {
                        "type": "string"
                    },
                    "rating": {
                        "type": "number"
                    },
                    "timestamp": {
                        "type": "string"
                    },
                    "topology": {
                        "enum": [
                            "FLAT",
                            0,
                            "NESTED_OBJECT",
                            1,
                            "NESTED_MESSAGE",
                            2,
                            "ARRAY_OF_TYPE",
                            3,
                            "ARRAY_OF_OBJECT",
                            4,
                            "ARRAY_OF_MESSAGE",
                            5
                        ],
                        "oneOf": [
                            {
                                "type": "string"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    }
                },
                "additionalProperties": true,
                "type": "object"
            },
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// ArrayOfObjects:
	sampleProtos["ArrayOfObjects"] = SampleProto{
		protoFileName:   "ArrayOfObjects.proto",
		filesToGenerate: []string{"ArrayOfObjects.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "payload": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "properties": {
                    "complete": {
                        "type": "boolean"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "name": {
                        "type": "string"
                    },
                    "rating": {
                        "type": "number"
                    },
                    "timestamp": {
                        "type": "string"
                    },
                    "topology": {
                        "enum": [
                            "FLAT",
                            0,
                            "NESTED_OBJECT",
                            1,
                            "NESTED_MESSAGE",
                            2,
                            "ARRAY_OF_TYPE",
                            3,
                            "ARRAY_OF_OBJECT",
                            4,
                            "ARRAY_OF_MESSAGE",
                            5
                        ],
                        "oneOf": [
                            {
                                "type": "string"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    }
                },
                "additionalProperties": true,
                "type": "object"
            },
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// ArrayOfPrimitives:
	sampleProtos["ArrayOfPrimitives"] = SampleProto{
		protoFileName:   "ArrayOfPrimitives.proto",
		filesToGenerate: []string{"ArrayOfPrimitives.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "keyWords": {
            "items": {
                "type": "string"
            },
            "type": "array"
        },
        "luckyNumbers": {
            "items": {
                "type": "integer"
            },
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// EnumCeption:
	sampleProtos["EnumCeption"] = SampleProto{
		protoFileName:   "EnumCeption.proto",
		filesToGenerate: []string{"EnumCeption.proto", "PayloadMessage.proto", "ImportedEnum.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "complete": {
            "type": "boolean"
        },
        "failureMode": {
            "enum": [
                "RECURSION_ERROR",
                0,
                "SYNTAX_ERROR",
                1
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        },
        "id": {
            "type": "integer"
        },
        "importedEnum": {
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        },
        "name": {
            "type": "string"
        },
        "payload": {
            "properties": {
                "complete": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "timestamp": {
                    "type": "string"
                },
                "topology": {
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "payloads": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "properties": {
                    "complete": {
                        "type": "boolean"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "name": {
                        "type": "string"
                    },
                    "rating": {
                        "type": "number"
                    },
                    "timestamp": {
                        "type": "string"
                    },
                    "topology": {
                        "enum": [
                            "FLAT",
                            0,
                            "NESTED_OBJECT",
                            1,
                            "NESTED_MESSAGE",
                            2,
                            "ARRAY_OF_TYPE",
                            3,
                            "ARRAY_OF_OBJECT",
                            4,
                            "ARRAY_OF_MESSAGE",
                            5
                        ],
                        "oneOf": [
                            {
                                "type": "string"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    }
                },
                "additionalProperties": true,
                "type": "object"
            },
            "type": "array"
        },
        "rating": {
            "type": "number"
        },
        "timestamp": {
            "type": "string"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// ImportedEnum:
	sampleProtos["ImportedEnum"] = SampleProto{
		protoFileName:   "ImportedEnum.proto",
		filesToGenerate: []string{"ImportedEnum.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "enum": [
        "VALUE_0",
        0,
        "VALUE_1",
        1,
        "VALUE_2",
        2,
        "VALUE_3",
        3
    ],
    "oneOf": [
        {
            "type": "string"
        },
        {
            "type": "integer"
        }
    ]
}`,
		},
	}

	// NestedMessage:
	sampleProtos["NestedMessage"] = SampleProto{
		protoFileName:   "NestedMessage.proto",
		filesToGenerate: []string{"NestedMessage.proto", "PayloadMessage.proto"},
		expectedJsonSchema: []string{

			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "payload": {
            "properties": {
                "complete": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "timestamp": {
                    "type": "string"
                },
                "topology": {
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// NestedObject:
	sampleProtos["NestedObject"] = SampleProto{
		protoFileName:   "NestedObject.proto",
		filesToGenerate: []string{"NestedObject.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "payload": {
            "properties": {
                "complete": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "timestamp": {
                    "type": "string"
                },
                "topology": {
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// PayloadMessage:
	sampleProtos["PayloadMessage"] = SampleProto{
		protoFileName:   "PayloadMessage.proto",
		filesToGenerate: []string{"PayloadMessage.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "complete": {
            "type": "boolean"
        },
        "id": {
            "type": "integer"
        },
        "name": {
            "type": "string"
        },
        "rating": {
            "type": "number"
        },
        "timestamp": {
            "type": "string"
        },
        "topology": {
            "enum": [
                "FLAT",
                0,
                "NESTED_OBJECT",
                1,
                "NESTED_MESSAGE",
                2,
                "ARRAY_OF_TYPE",
                3,
                "ARRAY_OF_OBJECT",
                4,
                "ARRAY_OF_MESSAGE",
                5
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

	// SeveralEnums:
	sampleProtos["SeveralEnums"] = SampleProto{
		protoFileName:   "SeveralEnums.proto",
		filesToGenerate: []string{"SeveralEnums.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "enum": [
        "VALUE_0",
        0,
        "VALUE_1",
        1,
        "VALUE_2",
        2,
        "VALUE_3",
        3
    ],
    "oneOf": [
        {
            "type": "string"
        },
        {
            "type": "integer"
        }
    ]
}`,
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "enum": [
        "VALUE_4",
        0,
        "VALUE_5",
        1,
        "VALUE_6",
        2,
        "VALUE_7",
        3
    ],
    "oneOf": [
        {
            "type": "string"
        },
        {
            "type": "integer"
        }
    ]
}`,
		},
	}

	// SeveralMessages:
	sampleProtos["SeveralMessages"] = SampleProto{
		protoFileName:   "SeveralMessages.proto",
		filesToGenerate: []string{"SeveralMessages.proto"},
		expectedJsonSchema: []string{
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "complete1": {
            "type": "boolean"
        },
        "id1": {
            "type": "integer"
        },
        "name1": {
            "type": "string"
        },
        "rating1": {
            "type": "number"
        },
        "timestamp1": {
            "type": "string"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
			`{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "complete2": {
            "type": "boolean"
        },
        "id2": {
            "type": "integer"
        },
        "name2": {
            "type": "string"
        },
        "rating2": {
            "type": "number"
        },
        "timestamp2": {
            "type": "string"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`,
		},
	}

}
