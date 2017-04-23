package testdata

const ArrayOfPrimitives = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "keyWords": {
            "items": {
                "type": "string"
            },
            "type": "array"
        },
        "luckyNumbers": {
            "items": {
                "type": "integer"
            },
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
