package testdata

const (
	CyclicalReferenceMessageM = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "foo": {
            "$ref": "samples.Foo",
            "additionalProperties": true,
            "type": "object"
        }
    },
    "additionalProperties": true,
    "type": "object",
    "definitions": {
        "samples.Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "properties": {},
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$schema": "http://json-schema.org/draft-04/schema#",
                        "properties": {
                            "id": {
                                "properties": {},
                                "type": "integer"
                            },
                            "baz": {
                                "properties": {
                                    "enabled": {
                                        "properties": {},
                                        "type": "boolean"
                                    },
                                    "foo": {
                                        "$ref": "samples.Foo",
                                        "additionalProperties": true,
                                        "type": "object"
                                    }
                                },
                                "additionalProperties": true,
                                "type": "object"
                            }
                        },
                        "additionalProperties": true,
                        "type": "object"
                    },
                    "properties": {},
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "samples.Foo"
        }
    }
}`

	CyclicalReferenceMessageFoo = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "Foo",
    "definitions": {
        "Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "properties": {},
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$schema": "http://json-schema.org/draft-04/schema#",
                        "properties": {
                            "id": {
                                "properties": {},
                                "type": "integer"
                            },
                            "baz": {
                                "properties": {
                                    "enabled": {
                                        "properties": {},
                                        "type": "boolean"
                                    },
                                    "foo": {
                                        "$ref": "Foo",
                                        "additionalProperties": true,
                                        "type": "object"
                                    }
                                },
                                "additionalProperties": true,
                                "type": "object"
                            }
                        },
                        "additionalProperties": true,
                        "type": "object"
                    },
                    "properties": {},
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Foo"
        }
    }
}`

	CyclicalReferenceMessageBar = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "Bar",
    "definitions": {
        "Bar": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "id": {
                    "properties": {},
                    "type": "integer"
                },
                "baz": {
                    "properties": {
                        "enabled": {
                            "properties": {},
                            "type": "boolean"
                        },
                        "foo": {
                            "properties": {
                                "name": {
                                    "properties": {},
                                    "type": "string"
                                },
                                "bar": {
                                    "items": {
                                        "$schema": "http://json-schema.org/draft-04/schema#",
                                        "$ref": "Bar"
                                    },
                                    "properties": {},
                                    "type": "array"
                                }
                            },
                            "additionalProperties": true,
                            "type": "object"
                        }
                    },
                    "additionalProperties": true,
                    "type": "object"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Bar"
        }
    }
}`

	CyclicalReferenceMessageBaz = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "Baz",
    "definitions": {
        "Baz": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "enabled": {
                    "properties": {},
                    "type": "boolean"
                },
                "foo": {
                    "properties": {
                        "name": {
                            "properties": {},
                            "type": "string"
                        },
                        "bar": {
                            "items": {
                                "$schema": "http://json-schema.org/draft-04/schema#",
                                "properties": {
                                    "id": {
                                        "properties": {},
                                        "type": "integer"
                                    },
                                    "baz": {
                                        "$ref": "Baz",
                                        "additionalProperties": true,
                                        "type": "object"
                                    }
                                },
                                "additionalProperties": true,
                                "type": "object"
                            },
                            "properties": {},
                            "type": "array"
                        }
                    },
                    "additionalProperties": true,
                    "type": "object"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Baz"
        }
    }
}`
)
