package tonio

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/CDimonaco/tonio/internal/core/formatters"
	"github.com/CDimonaco/tonio/internal/proto"
	"github.com/CDimonaco/tonio/internal/rabbit"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func init() {
	consumeCmd.PersistentFlags().StringVarP(&exchangeType, "type", "t", "direct", "RabbitMQ exchange type")
	consumeCmd.PersistentFlags().BoolVar(&durable, "durable", true, "Durable exchange")
	consumeCmd.PersistentFlags().StringVar(&protoMessageType, "proto-type", "", "Full qualified name of protobuf message")
	consumeCmd.PersistentFlags().StringVar(&protoFilesPath, "proto-files-path", "", "Path to proto files")
}

func protoModeEnabled(protoRegistry *proto.Registry, protoType string) bool {
	return protoType != "" && protoRegistry != nil
}

func protoForMessageType(registry *proto.Registry, protoMessageType string) *dynamic.Message {
	return registry.MessageForType(protoMessageType)
}

var consumeCmd = &cobra.Command{
	Use:     "consume [routing keys]",
	Short:   "Consume messages from an exchange",
	Args:    cobra.MinimumNArgs(1),
	PreRunE: initializeProtoRegistry,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := newLogger(debug)
		client, err := rabbit.NewClient(
			connection,
			exchange,
			exchangeType,
			durable,
			args,
			logger,
		)
		if err != nil {
			return err
		}

		msgc, err := client.Consume()
		if err != nil {
			return err
		}

		ctx, done := context.WithCancel(context.Background())
		group, groupCtx := errgroup.WithContext(ctx)

		defer close(msgc)
		defer func() {
			err := client.Close()
			if err != nil {
				logger.Errorw("error during rabbitmq client cosing", "error", err)
			}
		}()

		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)

		group.Go(func() error {
			for {
				select {
				case s := <-sigc:
					{
						logger.Debugw("received signal, shutdown", "sig", s)
						done()
					}
				case <-groupCtx.Done():
					{
						os.Stdout.WriteString("bye! \n")
						return groupCtx.Err()
					}
				}
			}
		})

		group.Go(func() error {
			for {
				select {
				case <-groupCtx.Done():
					{
						return nil
					}
				case m := <-msgc:
					{
						var output bytes.Buffer

						output.WriteString("\033[2J")
						output.WriteString("\u001b[0;0H")

						meta := core.ExtractMetadata(m)
						_, _ = meta.WriteTo(&output)

						message := m.Body

						if protoModeEnabled(protoRegistry, protoMessageType) {
							protoMessageDescriptor := protoForMessageType(protoRegistry, protoMessageType)

							if protoMessageDescriptor != nil {
								mb, err := formatters.ProtoMessage(protoMessageDescriptor, message)
								if err != nil {
									logger.Errorw(
										"could not decode proto message",
										"proto_type",
										protoMessageType,
										"error",
										err,
									)
									continue
								}
								message = mb
							} else {
								logger.Debugw(
									"skipping proto decoding, message type not found in registry",
									"proto_type",
									protoMessageType,
								)
							}
						}

						if core.IsJSON(message) {
							m, err := formatters.JSONMessage(message)
							if err != nil {
								logger.Errorw(
									"could not decode json message",
									"error",
									err,
								)

								continue
							}

							message = m
						}

						_, _ = output.Write(message)
						_, _ = output.WriteTo(os.Stdout)
					}
				}
			}
		})

		if err := group.Wait(); !errors.Is(err, context.Canceled) {
			return err
		}

		return nil
	},
}
