package formatters_test

import (
	"testing"

	"github.com/CDimonaco/tonio/internal/core/formatters"
	prototest "github.com/CDimonaco/tonio/internal/core/formatters/_fixtures"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

type ProtoFormatterTestSuite struct {
	suite.Suite
	dynamic *dynamic.Message
}

func (s *ProtoFormatterTestSuite) SetupSuite() {
	testDynamic, err := dynamic.AsDynamicMessage(&prototest.TestMessage{})
	if err != nil {
		panic(err)
	}

	s.dynamic = testDynamic
}

func (s *ProtoFormatterTestSuite) TestProtoFormattingSuccess() {
	message := prototest.TestMessage{Message: "test message"}
	rawMessage, err := proto.Marshal(&message)
	s.NoError(err)

	result, err := formatters.ProtoMessage(s.dynamic, rawMessage)
	s.NoError(err)
	s.EqualValues(`{"message":"test message"}`, string(result))
}

func (s *ProtoFormatterTestSuite) TestProtoFormattingFailure() {
	result, err := formatters.ProtoMessage(s.dynamic, []byte("testest"))
	s.Error(err)
	s.Nil(result)
}

func TestProtoFormatter(t *testing.T) {
	suite.Run(t, new(ProtoFormatterTestSuite))
}
