package testdata

const OptionEnumsTrimPrefix = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.Scheme",
    "enum": [
        "UNSPECIFIED",
        "HTTP",
        "HTTPS"
    ],
    "type": "string",
    "title": "Scheme"
}`

const OptionEnumsTrimPrefixPass = `"HTTP"`

const OptionEnumsTrimPrefixFail = `4`
