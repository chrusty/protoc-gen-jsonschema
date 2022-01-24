package testdata

const MessageWithComments = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageWithComments",
    "definitions": {
        "MessageWithComments": {
            "properties": {
                "name1": {
                    "type": "string",
                    "description": "This field is supposed to represent blahblahblah"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "This is a detached leading comment (which becomes the title)",
            "description": "This is a detached leading comment (which becomes the title)\n\nThis is a message level comment and talks about what this message is and why you should care about it!"
        }
    }
}`

const MessageWithCommentsFail = `{"name1": 12345}`
