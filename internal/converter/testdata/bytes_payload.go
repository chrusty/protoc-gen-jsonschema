package testdata

const BytesPayload = `{
    "$ref": "BytesPayload",
    "definitions": {
        "BytesPayload": {
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
            "type": "object",
            "id": "BytesPayload"
        }
    }
}`
