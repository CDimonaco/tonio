package composers

import (
	"github.com/CDimonaco/tonio/internal/core"
)

const jsonContentType = "application/json"

func JSON(rawBytes []byte) ([]byte, string, error) {
	if core.IsJSON(rawBytes) {
		return rawBytes, jsonContentType, nil
	}

	return nil, "", nil
}
