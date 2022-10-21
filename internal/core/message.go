package core

type TonioMessage struct {
	Body        []byte
	ContentType string
	Queue       string
}
