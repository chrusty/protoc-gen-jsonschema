package testdata

const EnumImport = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/UseImportedEnum",
    "definitions": {
        "UseImportedEnum": {
            "properties": {
                "importedEnum": {
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
                    ],
                    "description": "This is an enum"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const EnumImportFail = `{"importedEnum": "VALUE_4"}`

const EnumImportPass = `{"importedEnum": "VALUE_3"}`
