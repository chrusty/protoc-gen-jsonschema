package testdata

const EnumCeption = `{
    "$ref": "Enumception",
    "definitions": {
        "Enumception": {
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
                "payload": {
                    "$ref": "samples.PayloadMessage",
                    "additionalProperties": true
                },
                "payloads": {
                    "items": {
                        "$ref": "samples.PayloadMessage"
                    },
                    "type": "array"
                },
                "importedEnum": {
                    "enum": [
                        "VALUE_0",
                        0,
                        "VALUE_1",
                        1,
                        "VALUE_2",
                        2,
                        "VALUE_3",
                        3
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
            "id": "Enumception"
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

const EnumCeptionFail = `{"payloads": [ {"topology": "MAP"} ]}`

const EnumCeptionPass = `{"payloads": [ {"topology": "ARRAY_OF_MESSAGE"} ]}`
