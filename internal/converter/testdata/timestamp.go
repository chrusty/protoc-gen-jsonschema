package testdata

const Timestamp = `{
    "$ref": "Timestamp",
    "definitions": {
        "Timestamp": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "timestamp": {
                    "type": "string",
                    "format": "date-time"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Timestamp"
        }
    }
}`
