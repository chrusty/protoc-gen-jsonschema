package testdata

const RefPrefixM = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "foo": {
            "$ref": "#/definitions/samples.Foo",
            "additionalProperties": true
        }
    },
    "additionalProperties": true,
    "type": "object",
    "definitions": {
        "samples.Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$schema": "http://json-schema.org/draft-04/schema#",
                        "properties": {
                            "id": {
                                "type": "integer"
                            },
                            "baz": {
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
                            }
                        },
                        "additionalProperties": true,
                        "type": "object"
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

const RefPrefixFoo = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Foo",
    "definitions": {
        "Foo": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name": {
                    "type": "string"
                },
                "bar": {
                    "items": {
                        "$schema": "http://json-schema.org/draft-04/schema#",
                        "properties": {
                            "id": {
                                "type": "integer"
                            },
                            "baz": {
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
                        },
                        "additionalProperties": true,
                        "type": "object"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "id": "Foo"
        }
    }
}`

const RefPrefixBar = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Bar",
    "definitions": {
        "Bar": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "properties": {
                        "enabled": {
                            "type": "boolean"
                        },
                        "foo": {
                            "properties": {
                                "name": {
                                    "type": "string"
                                },
                                "bar": {
                                    "items": {
                                        "$schema": "http://json-schema.org/draft-04/schema#",
                                        "$ref": "#/definitions/Bar"
                                    },
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

const RefPrefixBaz = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Baz",
    "definitions": {
        "Baz": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "properties": {
                        "name": {
                            "type": "string"
                        },
                        "bar": {
                            "items": {
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
                                "type": "object"
                            },
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
