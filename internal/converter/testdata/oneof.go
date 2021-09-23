package testdata

const OneOf = `{
    "$ref": "OneOf",
    "definitions": {
        "OneOf": {
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
            "id": "OneOf"
        },
        "samples.OneOf.Bar": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "required": [
                "foo"
            ],
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
            "required": [
                "foo"
            ],
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

const OneOfFail = `{
	"something": true,
	"bar": {"foo": 1},
	"baz": {"foo": "one"}
}`

const OneOfPass = `{
	"something": true,
	"bar": {"foo": 1}
}`
