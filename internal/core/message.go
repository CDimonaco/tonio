package core

import "time"

type Message struct {
	Timestamp   time.Time
	Headers     map[string]interface{}
	Body        []byte
	ContentType string
	Exchange    string
	Queue       string
	RoutingKeys []string
}
