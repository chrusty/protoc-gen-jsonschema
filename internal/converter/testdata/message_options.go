package testdata

const MessageOptions = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageOptions",
    "definitions": {
        "MessageOptions": {
            "properties": {
                "ignore": {
                    "type": "boolean",
                    "description": "Messages tagged with this will not be processed:"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "description": "Custom MessageOptions:"
        }
    }
}`

const MessageOptionsFail = `{"ignore": 12345}`

const MessageOptionsPass = `{"ignore": true}`
