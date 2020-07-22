package testdata

const Proto2Required = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "query"
    ],
    "properties": {
        "query": {
            "type": "string"
        },
        "page_number": {
            "type": "integer"
        },
        "result_per_page": {
            "type": "integer"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
