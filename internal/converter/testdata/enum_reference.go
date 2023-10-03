package testdata

const EnumReference1 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageWithEnums",
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
                },
                "var": {
                    "$ref": "#/definitions/samples.MessageWithEnums.DefinedUsedMessage",
                    "additionalProperties": true
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
        "samples.MessageWithEnums.DefinedUnusedEnum": {
            "enum": [
                "Var",
                0
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ],
            "title": "Defined Unused Enum"
        },
        "samples.MessageWithEnums.DefinedUnusedMessage": {
            "properties": {
                "var": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Defined Unused Message"
        },
        "samples.MessageWithEnums.DefinedUnusedMessage.NestedUnusedMessage": {
            "properties": {
                "foo": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Unused Message"
        },
        "samples.MessageWithEnums.DefinedUsedMessage": {
            "properties": {
                "var": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Defined Used Message"
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