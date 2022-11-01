package formatters

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/CDimonaco/tonio/internal/proto"
	"go.uber.org/zap"
)

type Formatter = func(message core.Message) (*bytes.Buffer, error)

func isJSON(data []byte) bool {
	var i interface{}
	if err := json.Unmarshal(data, &i); err == nil {
		return true
	}
	return false
}

func FormatMessage(
	message core.Message,
	protoRegistry *proto.Registry,
	protoType string,
	logger *zap.SugaredLogger,
) (*bytes.Buffer, error) {
	if protoType != "" && protoRegistry != nil {
		formatter := ProtoMessage(*protoRegistry, protoType)
		output, err := formatter(message)
		if err != nil {
			if !errors.Is(err, ErrNoProtoMessageForDecoding) {
				// fallback to others formatter
				logger.Debugw("could not decode the message as protobuf, fallback to other formatters", "error", err)
			}
		} else {
			return output, nil
		}
	}

	if isJSON(message.Body) {
		return JSONMessage(message)
	}

	return Raw(message)
}
