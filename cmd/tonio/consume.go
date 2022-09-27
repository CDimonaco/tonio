package tonio

import (
	"os"

	"github.com/spf13/cobra"
)

var consumeCmd = &cobra.Command{ //nolint
	Use:   "consume",
	Short: "Consume messages from an exchange",
	Run: func(cmd *cobra.Command, args []string) {
		os.Stdout.WriteString("consume command")
	},
}
