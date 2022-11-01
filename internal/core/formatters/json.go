package formatters

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/hokaccha/go-prettyjson"
	"github.com/jedib0t/go-pretty/text"
)

func JSONMessage(messagePayload core.Message) (*bytes.Buffer, error) {
	var output bytes.Buffer

	output.WriteString("\n")
	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Exchange"),
		text.FgHiWhite.Sprint(messagePayload.Exchange),
	))

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("ContentType"),
		text.FgHiWhite.Sprint(messagePayload.ContentType),
	))

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Queue"),
		text.FgHiWhite.Sprint(messagePayload.Queue),
	))

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Routing keys"),
		text.FgHiWhite.Sprint(strings.Join(messagePayload.RoutingKeys, ", ")),
	))

	output.WriteString("\n\n")

	formatter := prettyjson.NewFormatter()
	prettifiedJSON, err := formatter.Format(messagePayload.Body)

	if err != nil {
		return nil, err
	}

	output.Write(prettifiedJSON)

	output.WriteString("\n\n")

	return &output, nil
}
