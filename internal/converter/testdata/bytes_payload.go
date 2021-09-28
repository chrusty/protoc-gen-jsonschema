package testdata

const BytesPayload = `{
    "$ref": "#/definitions/BytesPayload",
    "id": "BytesPayload",
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

const BytesPayloadFail = `{"payload": 12345}`
