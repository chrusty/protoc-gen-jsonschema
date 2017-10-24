Protobuf to JSON-Schema compiler
================================
This takes protobuf definitions and converts them into JSONSchemas, which can be used to dynamically validate JSON messages.

This will hopefully be useful for people who define their data using ProtoBuf, but use JSON for the "wire" format.

"Heavily influenced" by [Google's protobuf-to-BigQuery-schema compiler](https://github.com/GoogleCloudPlatform/protoc-gen-bq-schema).


Links
-----
* [About JSON Schema](http://json-schema.org/)
* [Popular GoLang JSON-Schema validation library](https://github.com/xeipuuv/gojsonschema)
* [Another GoLang JSON-Schema validation library](https://github.com/lestrrat/go-jsschema)


Usage
-----
* Allow NULL values (by default, JSONSchemas will reject NULL values unless we explicitly allow them):
    `protoc --jsonschema_out=allow_null_values:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Disallow additional properties (JSONSchemas won't validate JSON containing extra parameters):
    `protoc --jsonschema_out=disallow_additional_properties:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Disallow permissive validation of big-integers as strings (eg scientific notation):
    `protoc --jsonschema_out=disallow_bigints_as_strings:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Enable debug logging:
    `protoc --jsonschema_out=debug:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`


Sample protos (for testing)
---------------------------
* Proto with a simple (flat) structure: [samples.PayloadMessage](testdata/proto/PayloadMessage.proto)
* Proto containing a nested object (defined internally): [samples.NestedObject](testdata/proto/NestedObject.proto)
* Proto containing a nested message (defined in a different proto file): [samples.NestedMessage](testdata/proto/NestedMessage.proto)
* Proto containing an array of a primitive types (string, int): [samples.ArrayOfPrimitives](testdata/proto/ArrayOfPrimitives.proto)
* Proto containing an array of objects (internally defined): [samples.ArrayOfObjects](testdata/proto/ArrayOfObjects.proto)
* Proto containing an array of messages (defined in a different proto file): [samples.ArrayOfMessage](testdata/proto/ArrayOfMessage.proto)
* Proto containing multi-level enums (flat and nested and arrays): [samples.Enumception](testdata/proto/Enumception.proto)
* Proto containing a stand-alone enum: [samples.ImportedEnum](testdata/proto/ImportedEnum.proto)
* Proto containing 2 stand-alone enums: [samples.FirstEnum, samples.SecondEnum](testdata/proto/SeveralEnums.proto)
* Proto containing 2 messages: [samples.FirstMessage, samples.SecondMessage](testdata/proto/SeveralMessages.proto)
