package rabbit_test

import (
	"sync"
	"testing"
	"time"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/CDimonaco/tonio/internal/rabbit"
	"github.com/stretchr/testify/suite"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	testLogger = l.Sugar()
}

const amqpConnection = "amqp://tonio:tonio@localhost:5672"

type RabbitClientSuite struct {
	suite.Suite
	publisher rabbitmq.Publisher
}

func TestRabbitClientSuite(t *testing.T) {
	suite.Run(t, new(RabbitClientSuite))
}

func (s *RabbitClientSuite) SetupSuite() {
	pub, err := rabbitmq.NewPublisher(
		amqpConnection,
		rabbitmq.Config{},
	)

	s.NoError(err)

	s.publisher = *pub
}

func (s *RabbitClientSuite) TestClientProducing() {
	wg := new(sync.WaitGroup)

	exchange := "test_ex"
	keys := []string{"messages_sent"}

	client, err := rabbit.NewClient(
		amqpConnection,
		exchange,
		"direct",
		false,
		keys,
		testLogger,
	)
	s.NoError(err)

	out, err := client.Consume()
	s.NoError(err)

	tt, err := time.Parse(time.RFC3339, "2022-11-05T17:05:57.993Z")
	s.NoError(err)

	message := core.Message{
		Timestamp: tt,
		Headers: map[string]interface{}{
			"test": "test",
		},
		Body:        []byte("hello proto?"),
		ContentType: "text",
		Exchange:    exchange,
		RoutingKeys: keys,
	}

	err = client.Produce(message)
	s.NoError(err)

	wg.Add(1)

	go func() {
		defer wg.Done()
		result := <-out
		s.EqualValues("hello proto?", string(result.Body))
		s.EqualValues("text", result.ContentType)
		s.EqualValues(map[string]interface{}{
			"test": "test",
		}, result.Headers)
		s.EqualValues(tt.Format(time.RFC3339), result.Timestamp.UTC().Format(time.RFC3339))
	}()

	wg.Wait()

	err = client.Close()
	s.NoError(err)

}

func (s *RabbitClientSuite) TestClientConsuming() {
	wg := new(sync.WaitGroup)

	exchange := "test_ex"
	keys := []string{"messages"}

	client, err := rabbit.NewClient(
		amqpConnection,
		exchange,
		"direct",
		false,
		keys,
		testLogger,
	)
	s.NoError(err)

	out, err := client.Consume()

	s.NoError(err)

	err = s.publisher.Publish(
		[]byte("hello moto"),
		keys,
		rabbitmq.WithPublishOptionsExchange(exchange),
	)
	s.NoError(err)

	wg.Add(1)

	go func() {
		defer wg.Done()
		result := <-out
		s.EqualValues("hello moto", string(result.Body))
	}()

	wg.Wait()

	err = client.Close()
	s.NoError(err)
}
