package testdata

const CyclicalReferenceMessageM = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/M",
    "definitions": {
        "M": {
            "properties": {
                "foo": {
                    "$ref": "#/definitions/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/definitions/samples.Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/definitions/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Foo": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$ref": "#/definitions/samples.Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const CyclicalReferenceMessageFoo = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/Foo",
    "definitions": {
        "Foo": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$ref": "#/definitions/samples.Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/definitions/samples.Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/definitions/Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const CyclicalReferenceMessageBar = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/Bar",
    "definitions": {
        "Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/definitions/samples.Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/definitions/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Foo": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$ref": "#/definitions/Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const CyclicalReferenceMessageBaz = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/Baz",
    "definitions": {
        "Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/definitions/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/definitions/Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.Foo": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$ref": "#/definitions/samples.Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`
