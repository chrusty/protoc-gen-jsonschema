package testdata

const SelfReference = `{
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
                        "$ref": "Foo"
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
