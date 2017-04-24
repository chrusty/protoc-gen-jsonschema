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
