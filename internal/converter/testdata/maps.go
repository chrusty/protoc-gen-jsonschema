package testdata

const Maps = `{
    "$ref": "Maps",
    "definitions": {
        "Maps": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "map_of_strings": {
                    "additionalProperties": {
                        "type": "string"
                    },
                    "type": "object"
                },
                "map_of_ints": {
                    "additionalProperties": {
                        "type": "integer"
                    },
                    "type": "object"
                },
                "map_of_messages": {
                    "additionalProperties": {
                        "$ref": "samples.PayloadMessage",
                        "additionalProperties": true
                    },
                    "type": "object"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Maps"
        },
        "samples.PayloadMessage": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "topology": {
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.PayloadMessage"
        }
    }
}`

const MapsFail = `{
	"map_of_strings": {
		"one": 1,
		"two": 2,
		"three": 3
	}
}`

const MapsPass = `{
	"map_of_strings": {
		"one": "1",
		"two": "2",
		"three": "3"
	}
}`
