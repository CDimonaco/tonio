package formatters

import (
	"bytes"
	"errors"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/CDimonaco/tonio/internal/proto"
)

var ErrNoProtoMessageForDecoding = errors.New("could not find a message in the registry to unmarshal the proto")

func ProtoMessage(registry proto.Registry, protoType string) Formatter {
	return func(message core.Message) (*bytes.Buffer, error) {
		var output bytes.Buffer
		dynamicMessage := registry.MessageForType(protoType)
		if dynamicMessage == nil {
			return nil, ErrNoProtoMessageForDecoding
		}

		err := dynamicMessage.Unmarshal(message.Body)
		if err != nil {
			return nil, err
		}

		b, err := dynamicMessage.MarshalJSON()
		if err != nil {
			return nil, err
		}
		output.Write(b)

		return &output, nil
	}
}
