package testdata

const GoogleValue = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "arg": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
