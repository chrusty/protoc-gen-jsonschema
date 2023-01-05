package jsonschema

import (
	"io"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/AppliedIntuition/protoc-gen-jsonschema/internal/converter"
)

type ProtoConverter interface {
	ConvertFrom(rd io.Reader) (*pluginpb.CodeGeneratorResponse, error)
}

func New(logger *logrus.Logger) ProtoConverter {
	return converter.New(logger)
}
