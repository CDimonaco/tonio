// nolint:gochecknoglobals
package tonio

import (
	"github.com/CDimonaco/tonio/internal/proto"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	connection       string
	exchange         string
	exchangeType     string
	debug            bool
	durable          bool
	protoFilesPath   string
	protoMessageType string
	protoRegistry    *proto.Registry
)

func initializeProtoRegistry(cmd *cobra.Command, args []string) error {
	if protoFilesPath != "" {
		r, err := proto.NewRegistry(protoFilesPath, newLogger(debug))
		if err != nil {
			return err
		}

		protoRegistry = r
	}

	return nil
}

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
	Short:   "RabbitMQ command line utility for message handling",
	Version: "0.1.0",
}

func init() {
	TonioCmd.PersistentFlags().StringVarP(&connection, "connection", "c", "", "RabbitMQ connection")
	TonioCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug output")
	TonioCmd.PersistentFlags().StringVarP(&exchange, "exchange", "e", "", "RabbitMQ exchange")

	_ = TonioCmd.MarkPersistentFlagRequired("connection")

	TonioCmd.AddCommand(consumeCmd)
	TonioCmd.AddCommand(produceCmd)
}
