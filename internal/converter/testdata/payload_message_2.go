package testdata

const PayloadMessage2 = `{
    "$ref": "#/definitions/PayloadMessage2",
    "id": "PayloadMessage2",
    "definitions": {
        "PayloadMessage2": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "required": [
                "name",
                "timestamp",
                "id",
                "rating",
                "complete",
                "topology"
            ],
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
            "id": "PayloadMessage2"
        }
    }
}`

const PayloadMessage2Fail = `{
}`

const PayloadMessage2Pass = `{
    "name": "test",
    "timestamp": "1970-01-01T00:00:00Z",
    "id": 1,
    "rating": 100,
    "complete": true,
    "topology": "FLAT"
}`
