package testdata

const EnumWithMessage = `{
    "$ref": "WithFooBarBaz",
    "definitions": {
        "WithFooBarBaz": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "enumField": {
                    "enum": [
                        "Foo",
                        0,
                        "Bar",
                        1,
                        "Baz",
                        2
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
            "id": "WithFooBarBaz"
        }
    }
}`

const EnumWithMessageFail = `{"enumField": 4}`

const EnumWithMessagePass = `{"enumField": 2}`
