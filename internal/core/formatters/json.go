package formatters

import (
	"bytes"

	"github.com/hokaccha/go-prettyjson"
)

func JSONMessage(message []byte) (*bytes.Buffer, error) {
	var output bytes.Buffer

	formatter := prettyjson.NewFormatter()
	prettifiedJSON, err := formatter.Format(message)

	if err != nil {
		return nil, err
	}

	output.Write(prettifiedJSON)

	output.WriteString("\n\n")

	return &output, nil
}
