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
                        "enum": [
                            "FOO",
                            0,
                            "BAR",
                            1,
                            "FIZZ",
                            2,
                            "BUZZ",
                            3
                        ]
                    },
                    "type": "array",
                    "title": "Inline"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Array Of Enums"
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
