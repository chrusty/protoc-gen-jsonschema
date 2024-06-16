package testdata

const OptionVendorExtension = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionVendorExtension",
    "definitions": {
        "OptionVendorExtension": {
            "properties": {
                "query": {
                    "type": "string",
                    "x-go-custom-tag": "validate:\"required\""
                },
                "result_per_page": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Option Vendor Extension"
        }
    }
}`
