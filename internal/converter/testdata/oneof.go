package testdata

const OneOf = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "bar": {
            "$ref": "samples.OneOf.Bar",
            "additionalProperties": true
        },
        "baz": {
            "$ref": "samples.OneOf.Baz",
            "additionalProperties": true
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
    ],
    "definitions": {
        "samples.OneOf.Bar": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "foo": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.OneOf.Bar"
        },
        "samples.OneOf.Baz": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "foo": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.OneOf.Baz"
        }
    }
}`
