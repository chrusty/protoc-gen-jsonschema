package testdata

const FieldOptions = `{
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
    "description": "Custom FieldOptions for protoc-gen-jsonschema:"
}`
