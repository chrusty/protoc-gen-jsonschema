Protobuf to JSON-Schema compiler
================================
This takes protobuf definitions and converts them into JSONSchemas, which can be used to dynamically validate JSON messages.

This will hopefully be useful for people who define their data using ProtoBuf, but use JSON for the "wire" format.

"Heavily influenced" by [Google's protobuf-to-BigQuery-schema compiler](https://github.com/GoogleCloudPlatform/protoc-gen-bq-schema).


Installation
------------
`GO111MODULE=on go get github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema && go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema`


Links
-----
* [About JSON Schema](http://json-schema.org/)
* [Popular GoLang JSON-Schema validation library](https://github.com/xeipuuv/gojsonschema)
* [Another GoLang JSON-Schema validation library](https://github.com/lestrrat/go-jsschema)


Usage
-----
protoc-gen-jsonschema is designed to run like any other proto generator. The following examples show how to use options flags to enable different generator behaviours (more examples in the Makefile too).

* Allow NULL values (by default, JSONSchemas will reject NULL values unless we explicitly allow them):
    `protoc --jsonschema_out=allow_null_values:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Disallow additional properties (JSONSchemas won't validate JSON containing extra parameters):
    `protoc --jsonschema_out=disallow_additional_properties:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Disallow permissive validation of big-integers as strings (eg scientific notation):
    `protoc --jsonschema_out=disallow_bigints_as_strings:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Prefix generated schema files with their package name (as a directory):
    `protoc --jsonschema_out=prefix_schema_files_with_package:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Require all fields (because proto3 doesn't accommodate this):
     `protoc --jsonschema_out=all_fields_required:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`
* Enable debug logging:
    `protoc --jsonschema_out=debug:. --proto_path=testdata/proto testdata/proto/ArrayOfPrimitives.proto`


Sample protos (for testing)
---------------------------
* Proto with a simple (flat) structure: [samples.PayloadMessage](internal/converter/testdata/proto/PayloadMessage.proto)
* Proto containing a nested object (defined internally): [samples.NestedObject](internal/converter/testdata/proto/NestedObject.proto)
* Proto containing a nested message (defined in a different proto file): [samples.NestedMessage](internal/converter/testdata/proto/NestedMessage.proto)
* Proto containing an array of a primitive types (string, int): [samples.ArrayOfPrimitives](internal/converter/testdata/proto/ArrayOfPrimitives.proto)
* Proto containing an array of objects (internally defined): [samples.ArrayOfObjects](internal/converter/testdata/proto/ArrayOfObjects.proto)
* Proto containing an array of messages (defined in a different proto file): [samples.ArrayOfMessage](internal/converter/testdata/proto/ArrayOfMessage.proto)
* Proto containing multi-level enums (flat and nested and arrays): [samples.Enumception](internal/converter/testdata/proto/Enumception.proto)
* Proto containing a stand-alone enum: [samples.ImportedEnum](internal/converter/testdata/proto/ImportedEnum.proto)
* Proto containing 2 stand-alone enums: [samples.FirstEnum, samples.SecondEnum](internal/converter/testdata/proto/SeveralEnums.proto)
* Proto containing 2 messages: [samples.FirstMessage, samples.SecondMessage](internal/converter/testdata/proto/SeveralMessages.proto)
