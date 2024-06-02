package testdata

const NumericFormatBigIntAsString = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/NumericFormat",
    "definitions": {
        "NumericFormat": {
            "properties": {
                "int_val": {
                    "type": "integer",
                    "format": "int32"
                },
                "long_val": {
                    "type": "integer",
                    "format": "int64"
                },
                "float_val": {
                    "type": "number",
                    "format": "float"
                },
                "double_val": {
                    "type": "number",
                    "format": "double"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Numeric Format"
        }
    }
}`
