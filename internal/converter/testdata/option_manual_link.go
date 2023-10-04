package testdata

const OptionManualLink = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionManualLink",
    "$fullRef": "#/definitions/samples.OptionManualLink",
    "definitions": {
        "OptionManualLink": {
            "properties": {
                "name2": {
                    "type": "string"
                },
                "id2": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "options": {
                "manualLink": "https://www.google.com"
            },
            "title": "Option Manual Link"
        }
    }
}`
