package testdata

const CyclicalReferenceMessageM = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/M",
    "id": "M",
    "definitions": {
        "M": {
            "properties": {
                "foo": {
                    "$ref": "#/definitions/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "M"
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
            "type": "object",
            "id": "samples.Bar"
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
            "type": "object",
            "id": "samples.Baz"
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
            "type": "object",
            "id": "samples.Foo"
        }
    }
}`

const CyclicalReferenceMessageFoo = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Foo",
    "id": "Foo",
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
            "type": "object",
            "id": "Foo"
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
            "type": "object",
            "id": "samples.Bar"
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
            "type": "object",
            "id": "samples.Baz"
        }
    }
}`

const CyclicalReferenceMessageBar = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Bar",
    "id": "Bar",
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
            "type": "object",
            "id": "Bar"
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
            "type": "object",
            "id": "samples.Baz"
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
            "type": "object",
            "id": "samples.Foo"
        }
    }
}`

const CyclicalReferenceMessageBaz = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Baz",
    "id": "Baz",
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
            "type": "object",
            "id": "Baz"
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
            "type": "object",
            "id": "samples.Bar"
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
            "type": "object",
            "id": "samples.Foo"
        }
    }
}`
