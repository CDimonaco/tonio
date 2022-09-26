package rabbit

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	rabbitmq "github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type TonioMessage struct {
	Body        []byte
	ContentType string
}

type Client struct {
	url         string
	logger      *zap.SugaredLogger
	consumer    rabbitmq.Consumer
	pubblisher  rabbitmq.Publisher
	exchange    string
	routingKeys []string
	queue       string
	outc        chan TonioMessage
}

func NewClient(
	host string,
	username string,
	password string,
	exchange string,
	routingKeys []string,
	logger *zap.SugaredLogger,
) (*Client, error) {
	l := logger.With("component", "rabbitClient")
	zl := ZapLogger{l}

	url := fmt.Sprintf("amqp://%s:%s@%s", username, password, host)

	l.Debugw("initializing client", "url", url)

	consumer, err := rabbitmq.NewConsumer(
		url,
		rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogger(zl),
	)

	if err != nil {
		return nil, errors.Wrap(err, "error during rabbitmq consumer init")
	}

	publisher, err := rabbitmq.NewPublisher(
		url,
		rabbitmq.Config{},
		rabbitmq.WithPublisherOptionsLogger(zl),
	)

	if err != nil {
		return nil, errors.Wrap(err, "error during rabbitmq publisher init")
	}

	return &Client{
		url:         url,
		logger:      l,
		consumer:    consumer,
		pubblisher:  *publisher,
		queue:       fmt.Sprintf("tonio.test_queue.%s", strings.Join(routingKeys, ".")),
		exchange:    exchange,
		routingKeys: routingKeys,
		outc:        make(chan TonioMessage),
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

func (c *Client) Consume() (chan TonioMessage, error) {
	err := c.consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			c.logger.Debugf("consumed: %v", string(d.Body))

			c.outc <- TonioMessage{
				Body:        d.Body,
				ContentType: d.ContentType,
			}
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		c.queue,
		c.routingKeys,
		rabbitmq.WithConsumeOptionsBindingExchangeName(c.exchange),
		rabbitmq.WithConsumeOptionsQueueAutoDelete,
	)
	if err != nil {
		return nil, err
	}

	return c.outc, nil
}
