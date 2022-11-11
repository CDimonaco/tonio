package composers

import (
	"errors"

	"github.com/CDimonaco/tonio/internal/proto"
)

const protoContentType = "application/x-protobuf"

var ErrNoProtoMessageForEncoding = errors.New(
	"could not find a message in the registry to marshal the input into a proto",
)

func ProtoMessage(registry proto.Registry, protoType string) Composer {
	return func(rawBytes []byte) ([]byte, string, error) {
		var output []byte
		dynamicMessage := registry.MessageForType(protoType)
		if dynamicMessage == nil {
			return nil, "", ErrNoProtoMessageForEncoding
		}

		err := dynamicMessage.UnmarshalJSON(rawBytes)
		if err != nil {
			return nil, "", err
		}

		output, err = dynamicMessage.Marshal()
		if err != nil {
			return nil, "", err
		}

		return output, protoContentType, nil
	}
}
