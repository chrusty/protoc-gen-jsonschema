module github.com/sixt/protoc-gen-jsonschema

require (
	github.com/alecthomas/jsonschema v0.0.0-20200127222324-dd4542c1f589
	github.com/golang/protobuf v1.3.2
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0
)

replace (
	github.com/alecthomas/jsonschema => github.com/alecthomas/jsonschema v0.0.0-20200127222324-dd4542c1f589
	github.com/iancoleman/orderedmap => github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0
)

go 1.13
