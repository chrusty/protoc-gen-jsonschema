package testdata

const GoogleValue = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/GoogleValue",
    "$fullRef": "#/definitions/samples.GoogleValue",
    "definitions": {
        "GoogleValue": {
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
                    "title": "Value",
                    "description": "` + "`Value`" + ` represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values. A producer of value is expected to set one of these variants. Absence of any variant indicates an error. The JSON representation for ` + "`Value`" + ` is JSON value."
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Google Value"
        }
    }
}`

const GoogleValueFail = `{"arg": null}`

const GoogleValuePass = `{"arg": 12345}`
