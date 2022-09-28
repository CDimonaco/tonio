package formatters

type Input struct {
	Message     []byte
	ContentType string
	Exchange    string
	Queue       string
	RoutingKeys []string
}
