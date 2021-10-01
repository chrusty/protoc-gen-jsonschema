package testdata

const HiddenFields = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/HiddenFields",
    "definitions": {
        "HiddenFields": {
            "properties": {
                "visible1": {
                    "type": "string"
                },
                "visible2": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const HiddenFieldsFail = `{"visible1": 12345}`

const HiddenFieldsPass = `{"visible2": "hello"}`
