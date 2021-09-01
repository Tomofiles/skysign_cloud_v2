package ports

import (
	"context"
)

// PubSubManagerSetter .
type PubSubManagerSetter interface {
	SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error
}
