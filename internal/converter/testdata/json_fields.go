package testdata

const JSONFields = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/JSONFields",
    "definitions": {
        "JSONFields": {
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
            "type": "object"
        }
    }
}`

const JSONFieldsFail = `{"someThing": "onetwothree"}`

const JSONFieldsPass = `{"someThing": 12345}`
