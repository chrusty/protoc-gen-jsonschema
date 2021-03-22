package testdata

const OneOf = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "bar": {
            "properties": {
                "foo": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "baz": {
            "properties": {
                "foo": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    },
    "additionalProperties": true,
    "type": "object",
    "oneOf": [
        {
            "required": [
                "bar"
            ]
        },
        {
            "required": [
                "baz"
            ]
        }
    ]
}`
