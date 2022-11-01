package proto

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"go.uber.org/zap"
)

type Registry struct {
	descriptors []*desc.FileDescriptor
	logger      *zap.SugaredLogger
}

func NewRegistry(protoPath string, logger *zap.SugaredLogger) (*Registry, error) {
	p := &protoparse.Parser{
		ImportPaths: []string{protoPath},
	}

	logger.Debugw("proto import path", "path", protoPath)

	var protoFiles []string

	err := filepath.Walk(protoPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(path, ".proto") {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Debugw("protobuf files found in proto folder", "files", strings.Join(protoFiles, ","))

	resolved, err := protoparse.ResolveFilenames([]string{protoPath}, protoFiles...)
	if err != nil {
		return nil, err
	}

	logger.Debugw("protobuf file resolved", "files", strings.Join(resolved, ","))

	descs, err := p.ParseFiles(resolved...)
	if err != nil {
		return nil, err
	}

	return &Registry{descriptors: descs, logger: logger}, nil
}

func (d *Registry) MessageForType(protoType string) *dynamic.Message {
	for _, descriptor := range d.descriptors {
		if messageDescriptor := descriptor.FindMessage(protoType); messageDescriptor != nil {
			d.logger.Debugw("found proto message for type", "type", protoType)
			return dynamic.NewMessage(messageDescriptor)
		}
	}
	return nil
}
