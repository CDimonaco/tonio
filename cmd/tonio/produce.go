// nolint:gochecknoglobals
package tonio

import "github.com/spf13/cobra"

var (
	headers []string
)

func init() {
	produceCmd.PersistentFlags().StringVar(&protoMessageType, "proto-type", "", "Full qualified name of protobuf message")
	produceCmd.PersistentFlags().StringVar(&protoFilesPath, "proto-files-path", "", "Path to proto files")
	produceCmd.Flags().StringArrayVar(&headers, "header", []string{}, "Header in format [KEY]:[VALUE]. Use multiple times to add more headers.")

}

var produceCmd = &cobra.Command{ //nolint
	Use:     "INPUT | produce [routing keys]",
	Short:   "Produce message, use with command piping",
	Args:    cobra.MinimumNArgs(1),
	PreRunE: initializeProtoRegistry,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
