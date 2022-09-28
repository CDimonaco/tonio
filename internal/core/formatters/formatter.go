package formatters

import (
	"fmt"
	"os"
	"strings"

	"github.com/hokaccha/go-prettyjson"
	"github.com/jedib0t/go-pretty/text"
)

type Input struct {
	Message     []byte
	ContentType string
	Exchange    string
	Queue       string
	RoutingKeys []string
}

func JSONFormatter(input Input) error {
	os.Stdout.WriteString("\n")
	os.Stdout.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Exchange"),
		text.FgHiWhite.Sprint(input.Exchange),
	))

	os.Stdout.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("ContentType"),
		text.FgHiWhite.Sprint(input.ContentType),
	))

	os.Stdout.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Queue"),
		text.FgHiWhite.Sprint(input.Queue),
	))

	os.Stdout.WriteString(fmt.Sprintf(
		"%s: %s \n",
		text.FgHiGreen.Sprint("Routing keys"),
		text.FgHiWhite.Sprint(strings.Join(input.RoutingKeys, ", ")),
	))

	os.Stdout.WriteString("\n\n")

	formatter := prettyjson.NewFormatter()
	prettifiedJSON, err := formatter.Format(input.Message)

	if err != nil {
		return err
	}

	os.Stdout.Write(prettifiedJSON)

	os.Stdout.WriteString("\n\n")

	return nil
}
