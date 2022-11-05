package core

import "time"

type Message struct {
	Timestamp   time.Time
	Body        []byte
	ContentType string
	Exchange    string
	Queue       string
	RoutingKeys []string
}
