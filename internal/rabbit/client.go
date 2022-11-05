package rabbit

import (
	"fmt"
	"strings"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/pkg/errors"
	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type Client struct {
	connection      string
	logger          *zap.SugaredLogger
	consumer        rabbitmq.Consumer
	pubblisher      rabbitmq.Publisher
	exchange        string
	routingKeys     []string
	queue           string
	exchangeKind    string
	exchangeDurable bool
	outc            chan core.Message
}

func NewClient(
	connection string,
	exchange string,
	exchangeKind string,
	exchangeDurable bool,
	routingKeys []string,
	logger *zap.SugaredLogger,
) (*Client, error) {
	l := logger.With("component", "rabbitClient")
	zl := ZapLogger{l}

	l.Debugw("initializing client", "url", connection)

	c, err := rabbitmq.NewConsumer(
		connection,
		rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogger(zl),
	)

	if err != nil {
		return nil, errors.Wrap(err, "error during rabbitmq consumer init")
	}

	p, err := rabbitmq.NewPublisher(
		connection,
		rabbitmq.Config{},
		rabbitmq.WithPublisherOptionsLogger(zl),
	)

	if err != nil {
		return nil, errors.Wrap(err, "error during rabbitmq publisher init")
	}

	return &Client{
		connection:      connection,
		logger:          l,
		consumer:        c,
		exchangeKind:    exchangeKind,
		exchangeDurable: exchangeDurable,
		pubblisher:      *p,
		queue:           fmt.Sprintf("tonio.test_queue.%s", strings.Join(routingKeys, ".")),
		exchange:        exchange,
		routingKeys:     routingKeys,
		outc:            make(chan core.Message),
	}, nil
}

// Close, closed the client connection and execute cleanup operations
func (c *Client) Close() error {
	err := c.consumer.Close()
	if err != nil {
		return errors.Wrap(err, "error during consumer closing")
	}

	err = c.pubblisher.Close()
	if err != nil {
		return errors.Wrap(err, "error during publisher closing")
	}

	return nil
}

func (c *Client) Consume() (chan core.Message, error) {
	consumeOpts := []func(*rabbitmq.ConsumeOptions){
		rabbitmq.WithConsumeOptionsBindingExchangeName(c.exchange),
		rabbitmq.WithConsumeOptionsBindingExchangeKind(c.exchangeKind),
		rabbitmq.WithConsumeOptionsQueueAutoDelete,
	}

	if c.exchangeDurable {
		consumeOpts = append(consumeOpts, rabbitmq.WithConsumeOptionsBindingExchangeDurable)
	}

	err := c.consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			c.logger.Debugf("consumed: %v", string(d.Body))

			c.outc <- core.Message{
				Body:        d.Body,
				ContentType: d.ContentType,
				Queue:       c.queue,
				Exchange:    c.exchange,
				RoutingKeys: c.routingKeys,
				Timestamp:   d.Timestamp,
				Headers:     d.Headers,
			}
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		c.queue,
		c.routingKeys,
		consumeOpts...,
	)
	if err != nil {
		return nil, err
	}

	return c.outc, nil
}
