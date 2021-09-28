package testdata

const HiddenFields = `{
    "$ref": "#/definitions/HiddenFields",
    "id": "HiddenFields",
    "definitions": {
        "HiddenFields": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "visible1": {
                    "type": "string"
                },
                "visible2": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "HiddenFields"
        }
    }
}`

const HiddenFieldsFail = `{"visible1": 12345}`

const HiddenFieldsPass = `{"visible2": "hello"}`
