package rabbit

import "go.uber.org/zap"

type ZapLogger struct {
	*zap.SugaredLogger
}

func (l ZapLogger) Tracef(msg string, args ...interface{}) {
	l.Errorf(msg, args...)
}
