package testdata

const EnumCeption = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "complete": {
            "type": "boolean"
        },
        "failureMode": {
            "enum": [
                "RECURSION_ERROR",
                0,
                "SYNTAX_ERROR",
                1
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        },
        "id": {
            "type": "integer"
        },
        "importedEnum": {
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ]
        },
        "name": {
            "type": "string"
        },
        "payload": {
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
        },
        "payloads": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
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
            },
            "type": "array"
        },
        "rating": {
            "type": "number"
        },
        "timestamp": {
            "type": "string"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
