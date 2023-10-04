package testdata

const MultipleOneOf = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MultipleOneOf",
    "$fullRef": "#/definitions/samples.MultipleOneOf",
    "definitions": {
        "MultipleOneOf": {
            "required": [
                "bar",
                "baz",
                "foo",
                "qux",
                "something"
            ],
            "oneofNames": [
                "choice1",
                "choice2"
            ],
            "properties": {
                "bar": {
                    "$ref": "#/definitions/samples.MultipleOneOf.Bar",
                    "additionalProperties": true
                },
                "baz": {
                    "$ref": "#/definitions/samples.MultipleOneOf.Baz",
                    "additionalProperties": true
                },
                "foo": {
                    "$ref": "#/definitions/samples.MultipleOneOf.Foo",
                    "additionalProperties": true
                },
                "qux": {
                    "$ref": "#/definitions/samples.MultipleOneOf.Qux",
                    "additionalProperties": true
                },
                "something": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "exclusiveGroups": [
                {
                    "required": [
                        "bar",
                        "baz"
                    ]
                },
                {
                    "required": [
                        "foo",
                        "qux"
                    ]
                }
            ],
            "title": "Multiple One Of"
        },
        "samples.MultipleOneOf.Bar": {
            "required": [
                "bar_field"
            ],
            "properties": {
                "bar_field": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Bar"
        },
        "samples.MultipleOneOf.Baz": {
            "required": [
                "baz_field"
            ],
            "properties": {
                "baz_field": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Baz"
        },
        "samples.MultipleOneOf.Foo": {
            "required": [
                "foo_field"
            ],
            "properties": {
                "foo_field": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Foo"
        },
        "samples.MultipleOneOf.Qux": {
            "required": [
                "qux_field"
            ],
            "properties": {
                "qux_field": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Qux"
        }
    }
}`

const MultipleOneOfFail = `{
  "something": true,
  "bar": {"bar_field": 1},
  "baz": {"baz_field": "one"}
}`

const MultipleOneOfPass = `{
  "something": true,
  "bar": {"bar_field": 1}
}`
