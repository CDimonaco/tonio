package proto_test

import (
	"io/ioutil"
	"testing"

	internalproto "github.com/CDimonaco/tonio/internal/proto"
	prototest "github.com/CDimonaco/tonio/internal/proto/_fixtures"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var testLogger *zap.SugaredLogger //nolint

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	testLogger = l.Sugar()
}

func TestMessageForTypeSuccess(t *testing.T) {
	registry, err := internalproto.NewRegistry("./_fixtures", testLogger)
	assert.NoError(t, err)

	message := registry.MessageForType("Tonio.Test.TestMessage")
	assert.NotNil(t, message)
}

func TestMessageForTypeNotFound(t *testing.T) {
	registry, err := internalproto.NewRegistry("./_fixtures", testLogger)
	assert.NoError(t, err)

	message := registry.MessageForType("Tonio.Test.TestMessage2")
	assert.Nil(t, message)
}

func TestAsddf(t *testing.T) {
	test := prototest.TestMessage{Message: "ciao"}

	b, err := proto.Marshal(&test)
	assert.NoError(t, err)

	ioutil.WriteFile("./test.b", b, 777)

}
