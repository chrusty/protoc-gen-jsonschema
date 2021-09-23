package testdata

const ImportedEnum = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
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
}`

const ImportedEnumFail = `"VALUE_5"`

const ImportedEnumPass = `"VALUE_3"`
