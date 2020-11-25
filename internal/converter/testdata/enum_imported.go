package testdata

const EnumImport = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "importedEnum": {
            "enum": [
                "VALUE_0",
                0,
                "VALUE_1",
                1,
                "VALUE_2",
                2,
                "VALUE_3",
                3
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
