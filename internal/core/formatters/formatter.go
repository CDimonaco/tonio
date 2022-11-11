package formatters

import (
	"bytes"
	"errors"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/CDimonaco/tonio/internal/proto"
	"go.uber.org/zap"
)

type Formatter = func(message []byte) (*bytes.Buffer, error)

func FormatMessage(
	message core.Message,
	protoRegistry *proto.Registry,
	protoType string,
	logger *zap.SugaredLogger,
) (*bytes.Buffer, error) {
	messageBytes := message.Body

	if protoType != "" && protoRegistry != nil {
		formatter := ProtoMessage(*protoRegistry, protoType)
		mb, err := formatter(messageBytes)
		if err != nil {
			if !errors.Is(err, ErrNoProtoMessageForDecoding) {
				// fallback to others formatter
				logger.Debugw("could not decode the message as protobuf, fallback to other formatters", "error", err)
			} else {
				return nil, err
			}
		} else {
			messageBytes = mb.Bytes()
		}
	}

	if core.IsJSON(messageBytes) {
		return JSONMessage(messageBytes)
	}

	return Raw(messageBytes)
}
