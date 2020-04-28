package testdata

const ArrayOfObjects = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "payload": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "properties": {
                    "name": {
                        "properties": {},
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "string"
                            }
                        ]
                    },
                    "timestamp": {
                        "properties": {},
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "string"
                            }
                        ]
                    },
                    "id": {
                        "properties": {},
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    },
                    "rating": {
                        "properties": {},
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "number"
                            }
                        ]
                    },
                    "complete": {
                        "properties": {},
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "boolean"
                            }
                        ]
                    },
                    "topology": {
                        "properties": {},
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
                            },
                            {
                                "type": "null"
                            }
                        ]
                    }
                },
                "additionalProperties": true,
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "object"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        }
    },
    "additionalProperties": true,
    "oneOf": [
        {
            "type": "null"
        },
        {
            "type": "object"
        }
    ]
}`
