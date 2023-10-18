package testdata

const EnumCeption = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/Enumception",
    "$fullRef": "#/definitions/samples.Enumception",
    "definitions": {
        "Enumception": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "failureMode": {
                    "$ref": "#/definitions/samples.Enumception.FailureModes",
                    "title": "Failure Modes"
                },
                "payload": {
                    "$ref": "#/definitions/samples.PayloadMessage",
                    "additionalProperties": true
                },
                "payloads": {
                    "items": {
                        "$ref": "#/definitions/samples.PayloadMessage"
                    },
                    "type": "array"
                },
                "importedEnum": {
                    "$ref": "#/definitions/samples.ImportedEnum",
                    "title": "Imported Enum"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Enumception"
        },
        "samples.Enumception.FailureModes": {
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
            ],
            "title": "Failure Modes",
            "description": "FailureModes enum"
        },
        "samples.ImportedEnum": {
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
                    "description": "Zero",
                    "const": "VALUE_0"
                },
                {
                    "description": "Zero",
                    "const": 0
                },
                {
                    "description": "One",
                    "const": "VALUE_1"
                },
                {
                    "description": "One",
                    "const": 1
                },
                {
                    "description": "Two",
                    "const": "VALUE_2"
                },
                {
                    "description": "Two",
                    "const": 2
                },
                {
                    "description": "Three",
                    "const": "VALUE_3"
                },
                {
                    "description": "Three",
                    "const": 3
                }
            ],
            "title": "Imported Enum",
            "description": "This is an enum"
        },
        "samples.PayloadMessage": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "topology": {
                    "$ref": "#/definitions/samples.PayloadMessage.Topology",
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Payload Message"
        },
        "samples.PayloadMessage.Topology": {
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
            ],
            "title": "Topology"
        }
    }
}`

const EnumCeptionFail = `{"payloads": [ {"topology": "MAP"} ]}`

const EnumCeptionPass = `{"payloads": [ {"topology": "ARRAY_OF_MESSAGE"} ]}`
