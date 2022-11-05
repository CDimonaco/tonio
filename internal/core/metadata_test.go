// nolint:lll
package core_test

import (
	"testing"
	"time"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestMetadataExtraction(t *testing.T) {
	tt, err := time.Parse(time.RFC3339, "2022-11-05T17:05:57.993Z")
	assert.NoError(t, err)
	message := core.Message{
		Timestamp:   tt,
		Body:        []byte{},
		ContentType: "application/test",
		Exchange:    "ex",
		Queue:       "queue",
		RoutingKeys: []string{"one", "two"},
	}

	out := core.ExtractMetadata(message)

	assert.EqualValues(t, "\x1b[92mTimestamp\x1b[0m: \x1b[97m2022-11-05T17:05:57Z\x1b[0m \n\x1b[92mExchange\x1b[0m: \x1b[97mex\x1b[0m \n\x1b[92mContentType\x1b[0m: \x1b[97mapplication/test\x1b[0m \n\x1b[92mQueue\x1b[0m: \x1b[97mqueue\x1b[0m \n\x1b[92mRouting keys\x1b[0m: \x1b[97mone, two\x1b[0m \n", out.String())
}
