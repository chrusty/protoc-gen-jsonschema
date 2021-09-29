package testdata

const Timestamp = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Timestamp",
    "id": "Timestamp",
    "definitions": {
        "Timestamp": {
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

const TimestampFail = `{"timestamp": "twelve oclock"}`

const TimestampPass = `{"timestamp": "1970-01-01T00:00:00Z"}`
