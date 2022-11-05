package formatters

import (
	"bytes"
)

func Raw(message []byte) (*bytes.Buffer, error) {
	var output bytes.Buffer

	output.Write(message)

	output.WriteString("\n\n")

	return &output, nil
}
