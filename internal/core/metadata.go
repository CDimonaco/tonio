package core

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/text"
)

func ExtractMetadata(message Message) bytes.Buffer {
	var output bytes.Buffer

	if !message.Timestamp.IsZero() {
		output.WriteString(fmt.Sprintf(
			"%s: %s \n",
			text.FgHiGreen.Sprint("Timestamp"),
			text.FgHiWhite.Sprint(message.Timestamp.Format(time.RFC3339)),
		))
	}

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Exchange"),
		text.FgHiWhite.Sprint(message.Exchange),
	))

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("ContentType"),
		text.FgHiWhite.Sprint(message.ContentType),
	))

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Queue"),
		text.FgHiWhite.Sprint(message.Queue),
	))

	output.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Routing keys"),
		text.FgHiWhite.Sprint(strings.Join(message.RoutingKeys, ", ")),
	))

	return output
}
