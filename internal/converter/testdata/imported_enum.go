package testdata

const ImportedEnum = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
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
            "title": "Zero",
            "const": "VALUE_0"
        },
        {
            "title": "Zero",
            "const": 0
        },
        {
            "title": "One",
            "const": "VALUE_1"
        },
        {
            "title": "One",
            "const": 1
        },
        {
            "title": "Two",
            "const": "VALUE_2"
        },
        {
            "title": "Two",
            "const": 2
        },
        {
            "title": "Three",
            "const": "VALUE_3"
        },
        {
            "title": "Three",
            "const": 3
        }
    ],
    "description": "This is an enum"
}`

const ImportedEnumFail = `"VALUE_5"`

const ImportedEnumPass = `"VALUE_3"`
