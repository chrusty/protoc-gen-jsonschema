package testdata

const WellKnown = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "string_value": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "map_of_integers": {
            "additionalProperties": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
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
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
            },
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
