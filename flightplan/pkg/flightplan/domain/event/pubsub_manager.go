package event

import "context"

// PubSubManager .
type PubSubManager interface {
	GetPublisher() (Publisher, ConnectionClose, error)
	SetConsumer(ctx context.Context, exchangeName string, handler Handler) error
}

// ConnectionClose .
type ConnectionClose = func()

// Handler .
type Handler = func([]byte)
