package core

type Message struct {
	Body        []byte
	ContentType string
	Exchange    string
	Queue       string
	RoutingKeys []string
}
