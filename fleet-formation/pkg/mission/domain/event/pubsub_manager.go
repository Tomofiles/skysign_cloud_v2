package event

import "context"

// PubSubManager .
type PubSubManager interface {
	GetPublisher() (Publisher, ChannelClose, error)
	SetConsumer(ctx context.Context, exchangeName, queueName string, handler Handler) error
}

// ChannelClose .
type ChannelClose = func() error

// Handler .
type Handler = func([]byte)
