package testdata

const CyclicalReferenceMessageM = `{
    "$ref": "M",
    "definitions": {
        "M": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "foo": {
                    "$ref": "samples.Foo",
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
                    "$ref": "samples.Baz",
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
                    "$ref": "samples.Foo",
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
                        "$ref": "samples.Bar"
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
    "$ref": "Foo",
    "definitions": {
        "Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$ref": "samples.Bar"
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
                    "$ref": "samples.Baz",
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
                    "$ref": "Foo",
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
    "$ref": "Bar",
    "definitions": {
        "Bar": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "samples.Baz",
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
                    "$ref": "samples.Foo",
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
                        "$ref": "Bar"
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
    "$ref": "Baz",
    "definitions": {
        "Baz": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "samples.Foo",
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
                    "$ref": "Baz",
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
                        "$ref": "samples.Bar"
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
