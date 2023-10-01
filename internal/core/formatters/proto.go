package formatters

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/dynamic"
)

func ProtoMessage(messageDescriptor *dynamic.Message, message []byte, ar jsonpb.AnyResolver) ([]byte, error) {
	// m := messageDescriptor.UnmarshalJSON()
	// err := protojson.UN(message, m)
	err := messageDescriptor.UnmarshalMerge(message)
	// err :=

	if err != nil {
		return nil, err
	}

	b, err := messageDescriptor.MarshalJSONPB(&jsonpb.Marshaler{
		AnyResolver: ar,
	})
	if err != nil {
		return nil, err
	}

	return b, nil
}
