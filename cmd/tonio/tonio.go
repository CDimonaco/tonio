package tonio

import (
	"github.com/spf13/cobra"
)

var (
	host        string   //nolint
	username    string   //nolint
	password    string   //nolint
	exchange    string   //nolint
	routingKeys []string //nolint
)

var TonioCmd = &cobra.Command{ //nolint
	Use:     "tonio",
	Short:   "RabbitMq command line utility for message handling",
	Version: "0.1.0",
}

func init() {
	TonioCmd.PersistentFlags().StringVarP(&host, "host", "c", "", "RabbitMq host")
	TonioCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "RabbitMq username")
	TonioCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "RabbitMq password")
	TonioCmd.PersistentFlags().StringVarP(&exchange, "exchange", "e", "", "RabbitMq exchange")
	TonioCmd.PersistentFlags().StringSliceVarP(&routingKeys, "routing", "r", []string{}, "RabbitMq routing keys")

	_ = TonioCmd.MarkPersistentFlagRequired("host")
	_ = TonioCmd.MarkPersistentFlagRequired("username")
	_ = TonioCmd.MarkPersistentFlagRequired("password")
	_ = TonioCmd.MarkPersistentFlagRequired("exchange")
	_ = TonioCmd.MarkPersistentFlagRequired("routing")

	TonioCmd.AddCommand(consumeCmd)
}
