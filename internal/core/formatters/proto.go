package formatters

import (
	"github.com/jhump/protoreflect/dynamic"
)

func ProtoMessage(messageDescriptor *dynamic.Message, message []byte) ([]byte, error) {
	err := messageDescriptor.Unmarshal(message)
	if err != nil {
		return nil, err
	}

	b, err := messageDescriptor.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return b, nil
}
