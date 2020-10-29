package testdata

const JSONFields = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "name": {
            "type": "string"
        },
        "time_stamp": {
            "type": "string"
        },
        "identifier": {
            "type": "integer"
        },
        "rating": {
            "type": "number"
        },
        "complete": {
            "type": "boolean"
        }
    },
    "additionalProperties": true,
    "type": "object"
}
`
