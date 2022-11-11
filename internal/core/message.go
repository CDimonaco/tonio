package core

import (
	"encoding/json"
	"time"
)

type Message struct {
	Timestamp   time.Time
	Headers     map[string]interface{}
	Body        []byte
	ContentType string
	Exchange    string
	Queue       string
	RoutingKeys []string
}

func IsJSON(data []byte) bool {
	var i interface{}
	if err := json.Unmarshal(data, &i); err == nil {
		return true
	}
	return false
}
