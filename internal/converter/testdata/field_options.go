package testdata

<<<<<<< HEAD
const FieldOptions = `{
    "$ref": "FieldOptions",
    "definitions": {
        "FieldOptions": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "ignore": {
                    "type": "boolean",
                    "description": "Fields tagged with this will be omitted from generated schemas:"
                },
                "required": {
                    "type": "boolean",
                    "description": "Fields tagged with this will be marked as \"required\" in generated schemas:"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "description": "Custom FieldOptions for protoc-gen-jsonschema:",
            "id": "FieldOptions"
        }
    }
}`

const FieldOptionsFail = `{"ignore": 12345}`

const FieldOptionsPass = `{"required": true}`
=======
const FieldOptions = ``
>>>>>>> 5f61300... resetting test schemas
