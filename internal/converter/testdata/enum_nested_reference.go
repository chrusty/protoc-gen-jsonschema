package testdata

const EnumNestedReference = `{
    "$ref": "Msg",
    "definitions": {
        "Msg": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "nestedEnumField": {
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
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
            "type": "object",
            "id": "Msg"
        }
    }
}`

const EnumNestedReferenceFail = `{"nestedEnumField": 8}`

const EnumNestedReferencePass = `{"nestedEnumField": "FLAT"}`
