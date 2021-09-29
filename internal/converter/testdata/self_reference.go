package testdata

const SelfReference = `{
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
                        "$ref": "#/definitions/Foo"
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

const SelfReferenceFail = `{
	"bar": [
		{
			"name": false
		}
	]
}`

const SelfReferencePass = `{
	"bar": [
		{
			"name": "referenced-bar",
			"bar": [
				{
					"name": "barception"
				}
			]
		}
	]
}`
