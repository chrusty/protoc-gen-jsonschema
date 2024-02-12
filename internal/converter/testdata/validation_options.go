package testdata

const ValidationOptions = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/ValidationOptions",
    "definitions": {
        "ValidationOptions": {
            "required": [
                "requiredString"
            ],
            "properties": {
                "stringWithLengthConstraints": {
                    "maxLength": 10,
                    "minLength": 5,
                    "type": "string"
                },
                "luckyNumbersWithArrayConstraints": {
                    "items": {
                        "type": "integer"
                    },
                    "maxItems": 6,
                    "minItems": 2,
                    "type": "array"
                },
                "requiredString": {
                    "type": "string"
                },
                "stringRegex": {
                    "pattern": "^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$",
                    "type": "string",
                    "format": "regex"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Validation Options"
        }
    }
}`

const ValidationOptionsFail = `{
	"stringWithLengthConstraints": "this string is way too long",
	"luckyNumbersWithArrayConstraints": [1]
}`

const ValidationOptionsPass = `{
	"stringWithLengthConstraints": "thisisok",
	"luckyNumbersWithArrayConstraints": [1,2,3,4],
    "requiredString": "I am set!"
}`
