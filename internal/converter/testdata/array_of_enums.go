package testdata

const ArrayOfEnums = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/ArrayOfEnums",
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
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object"
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
