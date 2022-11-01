package formatters

import (
	"bytes"

	"github.com/CDimonaco/tonio/internal/core"
)

func Raw(message core.Message) (bytes.Buffer, error) {
	var output bytes.Buffer

	output.Write(message.Body)

	output.WriteString("\n\n")

	return output, nil
}
