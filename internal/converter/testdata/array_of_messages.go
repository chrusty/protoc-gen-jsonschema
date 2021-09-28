package testdata

const ArrayOfMessages = `{
    "$ref": "#/definitions/ArrayOfMessages",
    "id": "ArrayOfMessages",
    "definitions": {
        "ArrayOfMessages": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "description": {
                    "type": "string"
                },
                "payload": {
                    "items": {
                        "$ref": "#/definitions/samples.PayloadMessage"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "ArrayOfMessages"
        },
        "samples.PayloadMessage": {
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
            "id": "samples.PayloadMessage"
        }
    }
}`

const ArrayOfMessagesFail = `{
    "description": "something",
    "payload": [
        {"topology": "cruft"}
    ]
}`

const ArrayOfMessagesPass = `{
    "description": "something",
    "payload": [
        {"topology": "ARRAY_OF_MESSAGE"}
    ]
}`
