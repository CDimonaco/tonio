package formatters

import (
	"bytes"

	"github.com/CDimonaco/tonio/internal/core"
)

type Formatter = func(message core.Message) (bytes.Buffer, error)
