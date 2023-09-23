package proto_test

import (
	"testing"

	internalproto "github.com/CDimonaco/tonio/internal/proto"
	"github.com/stretchr/testify/assert"
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
