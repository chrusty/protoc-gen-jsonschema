package testdata

const MessageWithComments = `{
    "$ref": "#/definitions/MessageWithComments",
    "id": "MessageWithComments",
    "definitions": {
        "MessageWithComments": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "name1": {
                    "type": "string",
                    "description": "This field is supposed to represent blahblahblah"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "description": "This is a message level comment and talks about what this message is and why you should care about it!",
            "id": "MessageWithComments"
        }
    }
}`

const MessageWithCommentsFail = `{"name1": 12345}`
