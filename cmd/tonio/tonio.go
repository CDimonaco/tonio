// nolint:gochecknoglobals
package tonio

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	host         string
	username     string
	password     string
	exchange     string
	exchangeType string
	debug        bool
	durable      bool
)

func newLogger(debug bool) *zap.SugaredLogger {
	conf := zap.NewProductionConfig()
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	conf.Level = zap.NewAtomicLevelAt(zap.WarnLevel)

	if debug {
		conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	l, err := conf.Build()
	if err != nil {
		panic(err)
	}

	return l.Sugar()
}

var TonioCmd = &cobra.Command{
	Use:     "tonio",
	Short:   "RabbitMq command line utility for message handling",
	Version: "0.1.0",
}

func init() {
	TonioCmd.PersistentFlags().StringVar(&host, "host", "", "RabbitMq host")
	TonioCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "RabbitMq username")
	TonioCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "RabbitMq password")
	TonioCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug output")
	TonioCmd.PersistentFlags().StringVarP(&exchange, "exchange", "e", "", "RabbitMq exchange")
	TonioCmd.PersistentFlags().StringVarP(&exchangeType, "type", "t", "direct", "RabbitMq exchange type")
	TonioCmd.PersistentFlags().BoolVar(&durable, "durable", true, "Durable exchange")

	_ = TonioCmd.MarkPersistentFlagRequired("host")
	_ = TonioCmd.MarkPersistentFlagRequired("username")
	_ = TonioCmd.MarkPersistentFlagRequired("password")
	_ = TonioCmd.MarkPersistentFlagRequired("exchange")

	TonioCmd.AddCommand(consumeCmd)
}
