package testdata

const EnumReference1 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageWithEnums",
    "$fullRef": "#/definitions/samples.MessageWithEnums",
    "definitions": {
        "MessageWithEnums": {
            "properties": {
                "enumFieldOne": {
                    "$ref": "#/definitions/samples.EnumOne",
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
                    ],
                    "title": "Enum One"
                },
                "enumFieldTwo": {
                    "$ref": "#/definitions/samples.MessageWithEnums.NestedEnum",
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
                    ],
                    "title": "Nested Enum"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message With Enums"
        },
        "samples.EnumOne": {
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
            ],
            "title": "Enum One"
        },
        "samples.MessageWithEnums.NestedEnum": {
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
            ],
            "title": "Nested Enum"
        }
    }
}`

const EnumReference2 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.EnumOne",
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
    ],
    "title": "Enum One"
}
                `