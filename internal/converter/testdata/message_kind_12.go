package testdata

const MessageKind12 = `{
    "$ref": "#/definitions/MessageKind12",
    "id": "MessageKind12",
    "definitions": {
        "MessageKind12": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "f": {
                    "$ref": "#/definitions/samples.MessageKind11",
                    "additionalProperties": true
                },
                "kind5": {
                    "$ref": "#/definitions/samples.MessageKind5",
                    "additionalProperties": true
                },
                "kind6": {
                    "$ref": "#/definitions/samples.MessageKind6",
                    "additionalProperties": true
                },
                "kind7": {
                    "$ref": "#/definitions/samples.MessageKind7",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "MessageKind12"
        },
        "samples.MessageKind1": {
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
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind1"
        },
        "samples.MessageKind11": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "ones": {
                    "items": {
                        "$ref": "#/definitions/samples.MessageKind1"
                    },
                    "type": "array"
                },
                "kind2": {
                    "$ref": "#/definitions/samples.MessageKind2",
                    "additionalProperties": true
                },
                "kind3": {
                    "$ref": "#/definitions/samples.MessageKind3",
                    "additionalProperties": true
                },
                "kind4": {
                    "$ref": "#/definitions/samples.MessageKind4",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind11"
        },
        "samples.MessageKind2": {
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
                "isa": {
                    "type": "boolean"
                },
                "hasa": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind2"
        },
        "samples.MessageKind3": {
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
                "someProp": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind3"
        },
        "samples.MessageKind4": {
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
                "special": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind4"
        },
        "samples.MessageKind5": {
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
                "foo": {
                    "type": "number"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind5"
        },
        "samples.MessageKind6": {
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
                "bar": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind6"
        },
        "samples.MessageKind7": {
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
                "baz": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.MessageKind7"
        }
    }
}`
