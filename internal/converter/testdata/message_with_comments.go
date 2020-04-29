package testdata

const MessageWithComments = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "name1": {
            "type": "string",
            "description": "This field is supposed to represent blahblahblah"
        }
    },
    "additionalProperties": true,
    "type": "object",
    "description": "This is a message level comment and talks about what this message is and why you should care about it!"
}`
