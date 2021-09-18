package testdata

const OneOf = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "something"
    ],
    "properties": {
        "bar": {
            "required": [
                "foo"
            ],
            "properties": {
                "foo": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "baz": {
            "required": [
                "foo"
            ],
            "properties": {
                "foo": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "something": {
            "type": "boolean"
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
