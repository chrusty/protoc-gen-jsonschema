package testdata

const CyclicalReferenceMessageM = `{
    "$ref": "#/definitions/M",
    "id": "M",
    "definitions": {
        "M": {
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
    "$ref": "#/definitions/Foo",
    "id": "Foo",
    "definitions": {
        "Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
    "$ref": "#/definitions/Bar",
    "id": "Bar",
    "definitions": {
        "Bar": {
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
    "$ref": "#/definitions/Baz",
    "id": "Baz",
    "definitions": {
        "Baz": {
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
            "$schema": "http://json-schema.org/draft-04/schema#",
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
