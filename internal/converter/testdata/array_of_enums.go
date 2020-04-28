package testdata

const ArrayOfEnums = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "properties": {},
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
            "properties": {},
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
