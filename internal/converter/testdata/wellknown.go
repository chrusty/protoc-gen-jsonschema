package testdata

const WellKnown = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "string_value": {
            "additionalProperties": true,
            "type": "string"
        },
        "map_of_integers": {
            "additionalProperties": {
                "additionalProperties": true,
                "type": "integer"
            },
            "type": "object"
        },
        "map_of_scalar_integers": {
            "additionalProperties": {
                "type": "integer"
            },
            "type": "object"
        },
        "list_of_integers": {
            "items": {
                "type": "integer",
                "description": "Wrapper message for ` + "`int32`" + `.\n\n The JSON representation for ` + "`Int32Value`" + ` is JSON number."
            },
            "type": "array"
        },
        "duration": {
            "pattern": "^([0-9]+\\.?[0-9]*|\\.[0-9]+)(ns|ms|s|m|h|d)$",
            "type": "string",
            "description": "This is a duration:",
            "format": "regex"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
