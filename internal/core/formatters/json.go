package formatters

import (
	"bytes"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/hokaccha/go-prettyjson"
)

func JSONMessage(messagePayload core.Message) (*bytes.Buffer, error) {
	var output bytes.Buffer

	formatter := prettyjson.NewFormatter()
	prettifiedJSON, err := formatter.Format(messagePayload.Body)

	if err != nil {
		return nil, err
	}

	output.Write(prettifiedJSON)

	output.WriteString("\n\n")

	return &output, nil
}
