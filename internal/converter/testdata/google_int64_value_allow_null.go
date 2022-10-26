package testdata

const GoogleInt64ValueAllowNull = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/GoogleInt64ValueAllowNull",
    "definitions": {
        "GoogleInt64ValueAllowNull": {
            "properties": {
                "big_number": {
                    "additionalProperties": true,
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "title": "Google Int 64 Value Allow Null"
        }
    }
}`

const GoogleInt64ValueAllowNullFail = `{"big_number": 12345}`

const GoogleInt64ValueAllowNullPass = `{"big_number": null}`
