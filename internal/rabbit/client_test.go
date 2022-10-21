package rabbit_test

import (
	"sync"
	"testing"

	"github.com/CDimonaco/tonio/internal/rabbit"
	"github.com/stretchr/testify/suite"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger //nolint

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
func (s *RabbitClientSuite) TestClientConsuming() {
	wg := new(sync.WaitGroup)

	exchange := "test_ex"
	keys := []string{"messages"}

	client, err := rabbit.NewClient(
		"localhost",
		"tonio",
		"tonio",
		exchange,
		"direct",
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
