package testdata

const Proto3Required = `{
    "$ref": "#/definitions/Proto3Required",
    "id": "Proto3Required",
    "definitions": {
        "Proto3Required": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "required": [
                "query",
                "page_number"
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
            "type": "object",
            "id": "Proto3Required"
        }
    }
}`

const Proto3RequiredFail = `{
	"page_number": 4
}`

const Proto3RequiredPass = `{
	"query": "what?",
	"page_number": 4
}`
