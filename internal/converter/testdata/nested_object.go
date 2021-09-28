package testdata

const NestedObject = `{
    "$ref": "#/definitions/NestedObject",
    "id": "NestedObject",
    "definitions": {
        "NestedObject": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "payload": {
                    "$ref": "#/definitions/samples.NestedObject.NestedPayload",
                    "additionalProperties": true
                },
                "description": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "NestedObject"
        },
        "samples.NestedObject.NestedPayload": {
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "type": "object",
            "id": "samples.NestedObject.NestedPayload"
        }
    }
}`

const NestedObjectFail = `{"payload": false}`

const NestedObjectPass = `{
	"payload": {
	  "topology": "NESTED_OBJECT"
	}
}`
