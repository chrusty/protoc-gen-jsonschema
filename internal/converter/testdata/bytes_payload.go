package testdata

const BytesPayload = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "payload": {
            "type": "string",
            "format": "binary",
            "binaryEncoding": "base64"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
