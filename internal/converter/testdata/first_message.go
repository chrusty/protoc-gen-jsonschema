package testdata

const FirstMessage = `{
    "$ref": "#/definitions/FirstMessage",
    "id": "FirstMessage",
    "definitions": {
        "FirstMessage": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name1": {
                    "type": "string"
                },
                "timestamp1": {
                    "type": "string"
                },
                "id1": {
                    "type": "integer"
                },
                "rating1": {
                    "type": "number"
                },
                "complete1": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "FirstMessage"
        }
    }
}`

const FirstMessageFail = `{"complete1": "hello"}`

const FirstMessagePass = `{"complete1": true}`
