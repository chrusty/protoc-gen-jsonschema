package testdata

const SelfReference = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "Foo",
    "definitions": {
        "Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "properties": {},
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$schema": "http://json-schema.org/draft-04/schema#",
                        "$ref": "Foo"
                    },
                    "properties": {},
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Foo"
        }
    }
}`
