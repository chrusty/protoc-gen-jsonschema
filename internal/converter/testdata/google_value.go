package testdata

const GoogleValue = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "arg": {
            "oneOf": [
                {
                    "type": "array"
                },
                {
                    "type": "boolean"
                },
                {
                    "type": "number"
                },
                {
                    "type": "object"
                },
                {
                    "type": "string"
                }
            ],
            "description": "` + "`Value`" + ` represents a dynamically typed value which can be either\n null, a number, a string, a boolean, a recursive struct value, or a\n list of values. A producer of value is expected to set one of that\n variants, absence of any variant indicates an error.\n\n The JSON representation for ` + "`Value`" + ` is JSON value."
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
