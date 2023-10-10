package testdata

const Maps = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Maps",
    "$fullRef": "#/definitions/samples.Maps",
    "definitions": {
        "Maps": {
            "properties": {
                "map_of_strings_to_strings": {
                    "mapKey": "string",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "type": "object"
                },
                "map_of_strings_to_ints": {
                    "mapKey": "string",
                    "additionalProperties": {
                        "type": "integer"
                    },
                    "type": "object"
                },
                "map_of_strings_to_messages": {
                    "mapKey": "string",
                    "additionalProperties": {
                        "$ref": "#/definitions/samples.PayloadMessage",
                        "additionalProperties": true
                    },
                    "type": "object"
                },
                "map_of_ints_to_strings": {
                    "mapKey": "integer",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "type": "object"
                },
                "map_of_ints_to_ints": {
                    "mapKey": "integer",
                    "additionalProperties": {
                        "type": "integer"
                    },
                    "type": "object"
                },
                "map_of_ints_to_messages": {
                    "mapKey": "integer",
                    "additionalProperties": {
                        "$ref": "#/definitions/samples.PayloadMessage",
                        "additionalProperties": true
                    },
                    "type": "object"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Maps"
        },
        "samples.PayloadMessage": {
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
                    "$ref": "#/definitions/samples.PayloadMessage.Topology",
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
                    ],
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Payload Message"
        },
        "samples.PayloadMessage.Topology": {
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
            ],
            "title": "Topology"
        }
    }
}`

const MapsFail = `{
	"map_of_strings_to_strings": {
		"one": 1,
		"two": 2,
		"three": 3
	}
}`

const MapsPass = `{
	"map_of_strings_to_strings": {
		"one": "1",
		"two": "2",
		"three": "3"
	}
}`
