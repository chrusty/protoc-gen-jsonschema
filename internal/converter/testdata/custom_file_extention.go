package testdata

const CustomFileExtention = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/CustomFileExtention",
    "definitions": {
        "CustomFileExtention": {
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