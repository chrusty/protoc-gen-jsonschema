package testdata

const JSONFields = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/JSONFields",
    "$fullRef": "#/definitions/samples.JSONFields",
    "definitions": {
        "JSONFields": {
            "required": [
                "otherNumb"
            ],
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
                },
                "otherNumb": {
                    "type": "integer"
                },
                "oldField": {
                    "type": "integer",
                    "deprecated": true
                },
                "withManual": {
                    "type": "integer",
                    "options": {
                        "manualLink": "manual-link"
                    }
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "JSON Fields"
        }
    }
}`

const JSONFieldsFail = `{"someThing": "onetwothree", "other_numb": 123}`

const JSONFieldsPass = `{"someThing": 12345, "otherNumb": 123}`
