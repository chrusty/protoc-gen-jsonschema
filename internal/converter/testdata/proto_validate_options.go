package testdata

const ProtoValidateOptions = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/ProtoValidateOptions",
    "definitions": {
        "ProtoValidateOptions": {
            "required": [
                "requiredString"
            ],
            "properties": {
                "stringWithLengthConstraints": {
                    "maxLength": 10,
                    "minLength": 5,
                    "type": "string"
                },
                "requiredString": {
                    "type": "string"
                },
                "stringRegex": {
                    "pattern": "gr(a|e)y",
                    "type": "string",
                    "format": "regex"
                },
                "int32Range": {
                    "maximum": 20,
                    "exclusiveMaximum": true,
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "integer"
                },
                "int32RangeInc": {
                    "maximum": 5,
                    "minimum": 1,
                    "type": "integer"
                },
                "int32Gt": {
                    "minimum": 5,
                    "exclusiveMinimum": true,
                    "type": "integer"
                },
                "int32Lt": {
                    "maximum": 6,
                    "exclusiveMaximum": true,
                    "type": "integer"
                },
                "int32Gte": {
                    "minimum": 7,
                    "type": "integer"
                },
                "int32Lte": {
                    "maximum": 8,
                    "type": "integer"
                },
                "int64Range": {
                    "maximum": 20,
                    "exclusiveMaximum": true,
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "string"
                },
                "int64RangeInc": {
                    "maximum": 8,
                    "minimum": 3,
                    "type": "string"
                },
                "int64Gt": {
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "string"
                },
                "int64Lt": {
                    "maximum": 11,
                    "exclusiveMaximum": true,
                    "type": "string"
                },
                "int64Gte": {
                    "minimum": 12,
                    "type": "string"
                },
                "int64Lte": {
                    "maximum": 13,
                    "type": "string"
                },
                "floatRange": {
                    "maximum": 20,
                    "exclusiveMaximum": true,
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "number"
                },
                "floatRangeInc": {
                    "maximum": 8,
                    "minimum": 3,
                    "type": "number"
                },
                "floatGt": {
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "number"
                },
                "floatLt": {
                    "maximum": 11,
                    "exclusiveMaximum": true,
                    "type": "number"
                },
                "floatGte": {
                    "minimum": 12,
                    "type": "number"
                },
                "floatLte": {
                    "maximum": 13,
                    "type": "number"
                },
                "someStringsButNotTooMany": {
                    "items": {
                        "type": "string"
                    },
                    "maxItems": 6,
                    "minItems": 2,
                    "type": "array"
                },
                "someIntsButNotTooMany": {
                    "items": {
                        "type": "integer"
                    },
                    "maxItems": 21,
                    "minItems": 4,
                    "type": "array"
                },
                "email": {
                    "type": "string",
                    "format": "email"
                },
                "uuid": {
                    "type": "string",
                    "format": "uuid"
                },
                "uri": {
                    "type": "string",
                    "format": "uri"
                },
                "hostname": {
                    "type": "string",
                    "format": "hostname"
                },
                "ipv4": {
                    "type": "string",
                    "format": "ipv4"
                },
                "ipv6": {
                    "type": "string",
                    "format": "ipv6"
                },
                "uint32Range": {
                    "maximum": 20,
                    "exclusiveMaximum": true,
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "integer"
                },
                "uint32RangeInc": {
                    "maximum": 5,
                    "minimum": 1,
                    "type": "integer"
                },
                "uint32Gt": {
                    "minimum": 5,
                    "exclusiveMinimum": true,
                    "type": "integer"
                },
                "uint32Lt": {
                    "maximum": 6,
                    "exclusiveMaximum": true,
                    "type": "integer"
                },
                "uint32Gte": {
                    "minimum": 7,
                    "type": "integer"
                },
                "uint32Lte": {
                    "maximum": 8,
                    "type": "integer"
                },
                "uint64Range": {
                    "maximum": 20,
                    "exclusiveMaximum": true,
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "string"
                },
                "uint64RangeInc": {
                    "maximum": 8,
                    "minimum": 3,
                    "type": "string"
                },
                "uint64Gt": {
                    "minimum": 10,
                    "exclusiveMinimum": true,
                    "type": "string"
                },
                "uint64Lt": {
                    "maximum": 11,
                    "exclusiveMaximum": true,
                    "type": "string"
                },
                "uint64Gte": {
                    "minimum": 12,
                    "type": "string"
                },
                "uint64Lte": {
                    "maximum": 13,
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto Validate Options"
        }
    }
}`

const ProtoValidateOptionsFail = `{
	"stringWithLengthConstraints": "this string is way too long",
	"someIntsButNotTooMany": [1]
}`

const ProtoValidateOptionsPass = `{
	"stringWithLengthConstraints": "thisisok",
    "requiredString": "I am set!",
    "stringRegex": "grey",

    "int32Range": 13,
    "int32RangeInc": 1,
    "int32Gt": 6,
    "int32Lt": 5,
    "int32Gte": 7,
    "int32Gte": 8,

    "int64Range": "19",
    "int64RangeInc": "8",
    "int64Gt": "11",
    "int64Lt": "10",
    "int64Gte": "12",
    "int64Lte": "13",
    "floatRangeInc": 7.2,

    "someStringsButNotTooMany": ["one", "two"],
    "someIntsButNotTooMany": [1,2,3,4],

    "email": "test@example.com",
    "hostname": "server001",
    "ipv4": "127.0.0.1"
}`
