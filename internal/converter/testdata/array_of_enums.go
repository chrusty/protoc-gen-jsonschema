package testdata

const ArrayOfEnums = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/ArrayOfEnums",
    "$fullRef": "#/definitions/samples.ArrayOfEnums",
    "definitions": {
        "ArrayOfEnums": {
            "properties": {
                "description": {
                    "type": "string"
                },
                "stuff": {
                    "items": {
                        "$ref": "#/definitions/samples.ArrayOfEnums.inline",
                        "title": "Inline"
                    },
                    "type": "array",
                    "description": "Comment about inline stuff."
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Array Of Enums"
        },
        "samples.ArrayOfEnums.inline": {
            "enum": [
                "FOO",
                0,
                "BAR",
                1,
                "FIZZ",
                2,
                "BUZZ",
                3
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ],
            "title": "Inline"
        }
    }
}`

const ArrayOfEnumsFail = `{
    "description": "something",
    "stuff": [
        "FOOZ"
    ]
}`

const ArrayOfEnumsPass = `{
    "description": "something",
    "stuff": [
       3
    ]
}`
