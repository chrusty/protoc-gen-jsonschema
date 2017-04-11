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
* Message with a simple (flat) structure: [samples.PayloadMessage](PayloadMessage.proto)
* Message containing a nested object (defined internally): [ssamples.NestedObject](NestedObject.proto)
* Message containing a nested message (defined in a different proto file): [samples.NestedMessage](NestedMessage.proto)
* Message containing an array of a primitive types (string, int): [samples.ArrayOfPrimitives](ArrayOfPrimitives.proto)
* Message containing an array of objects (internally defined): [samples.ArrayOfObjects](ArrayOfObjects.proto)
* Message containing an array of messages (defined in a different proto file): [samples.ArrayOfMessage](ArrayOfMessage.proto)
* Message containing multi-level enums (flat and nested and arrays): [samples.Enumception](Enumception.proto)
