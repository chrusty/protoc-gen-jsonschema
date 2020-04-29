package testdata

const ArrayOfEnums = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
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
}`
