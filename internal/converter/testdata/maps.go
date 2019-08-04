package testdata

const Maps = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "map_of_strings": {
		"type": "object",
		"additionalProperties": {
			"type": "string"
		}
	},
    "map_of_ints": {
		"type": "object",
		"additionalProperties": {
			"type": "integer"
		}
	},
	"map_of_messages": {
		"type": "object",
		"additionalProperties": {
			"properties": {
				"complete": {
					"type": "boolean"
				},
				"id": {
					"type": "integer"
				},
				"name": {
					"type": "string"
				},
				"rating": {
					"type": "number"
				},
				"timestamp": {
					"type": "string"
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
			"type": "object"
		}
	}
}`
