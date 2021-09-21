package testdata

const HiddenFields = `{
    "$ref": "HiddenFields",
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
