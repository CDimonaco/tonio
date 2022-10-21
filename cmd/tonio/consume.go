package tonio

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/CDimonaco/tonio/internal/core/formatters"
	"github.com/CDimonaco/tonio/internal/rabbit"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var consumeCmd = &cobra.Command{ //nolint
	Use:   "consume [routing keys]",
	Short: "Consume messages from an exchange",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := newLogger(debug)
		client, err := rabbit.NewClient(
			host,
			username,
			password,
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

		group.Go(func() error {
			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)

			select {
			case s := <-sigc:
				{
					logger.Debugw("received signal, shutdown", "sig", s)
					done()
				}
			case <-groupCtx.Done():
				{
					os.Stdout.WriteString("bye!")
					return groupCtx.Err()
				}
			}

			return nil
		})

		group.Go(func() error {
			for m := range msgc {
				formatterInput := formatters.Input{
					Message:     m.Body,
					ContentType: m.ContentType,
					Exchange:    exchange,
					Queue:       m.Queue,
					RoutingKeys: args,
				}

				out, err := formatters.JSONMessage(formatterInput)
				if err != nil {
					return err
				}

				os.Stdout.Write(out.Bytes())
			}

			return nil
		})

		return group.Wait()
	},
}
