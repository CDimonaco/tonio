package formatters

import (
	"github.com/hokaccha/go-prettyjson"
)

func JSONMessage(message []byte) ([]byte, error) {
	formatter := prettyjson.NewFormatter()
	prettifiedJSON, err := formatter.Format(message)

	if err != nil {
		return nil, err
	}

	return prettifiedJSON, nil
}
