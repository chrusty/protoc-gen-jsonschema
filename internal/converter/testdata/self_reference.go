package testdata

const SelfReference = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "Foo",
    "definitions": {
        "Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$schema": "http://json-schema.org/draft-04/schema#",
                        "$ref": "Foo"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Foo"
        }
    }
}`
