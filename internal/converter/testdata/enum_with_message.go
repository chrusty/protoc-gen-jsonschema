package testdata

const EnumWithMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/WithFooBarBaz",
    "$fullRef": "#/definitions/samples.WithFooBarBaz",
    "definitions": {
        "WithFooBarBaz": {
            "properties": {
                "enumField": {
                    "$ref": "#/definitions/samples.FooBarBaz",
                    "title": "Foo Bar Baz"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "With Foo Bar Baz"
        },
        "samples.FooBarBaz": {
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
            "title": "Foo Bar Baz"
        }
    }
}`

const EnumWithMessageFail = `{"enumField": 4}`

const EnumWithMessagePass = `{"enumField": 2}`

const EnumWithMessageEnum = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.FooBarBaz",
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
    "title": "Foo Bar Baz"
}`


const EnumWithMessageEnumFail = `{}`

const EnumWithMessageEnumPass = `2`
