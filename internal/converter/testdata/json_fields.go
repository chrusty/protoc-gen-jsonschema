package testdata

const JSONFields = `{
    "$ref": "#/definitions/JSONFields",
    "id": "JSONFields",
    "definitions": {
        "JSONFields": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "identifier": {
                    "type": "integer"
                },
                "someThing": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "snakeNumb": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "JSONFields"
        }
    }
}`

const JSONFieldsFail = `{"someThing": "onetwothree"}`

const JSONFieldsPass = `{"someThing": 12345}`
