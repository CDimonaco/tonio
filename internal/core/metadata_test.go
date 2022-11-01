// nolint:lll
package core_test

import (
	"testing"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestMetadataExtraction(t *testing.T) {
	message := core.Message{
		Body:        []byte{},
		ContentType: "application/test",
		Exchange:    "ex",
		Queue:       "queue",
		RoutingKeys: []string{"one", "two"},
	}

	out := core.ExtractMetadata(message)

	assert.EqualValues(t, "\x1b[92mExchange\x1b[0m: \x1b[97mex\x1b[0m \n\x1b[92mContentType\x1b[0m: \x1b[97mapplication/test\x1b[0m \n\x1b[92mQueue\x1b[0m: \x1b[97mqueue\x1b[0m \n\x1b[92mRouting keys\x1b[0m: \x1b[97mone, two\x1b[0m \n", out.String())
}
