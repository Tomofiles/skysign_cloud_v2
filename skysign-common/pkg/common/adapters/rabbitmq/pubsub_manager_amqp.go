package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
)

// PublishHandler .
type PublishHandler = func(Channel, interface{})

// PubSubManagerSetter .
type PubSubManagerSetter interface {
	SetPublishHandler(handler PublishHandler) error
}

// PubSubManager .
type PubSubManager struct {
	conn        Connection
	pubHandlers []PublishHandler
}

// NewPubSubManager .
func NewPubSubManager(conn Connection) *PubSubManager {
	return &PubSubManager{
		conn:        conn,
		pubHandlers: []PublishHandler{},
	}
}

// GetPublisher .
func (psm *PubSubManager) GetPublisher() (event.Publisher, event.ChannelClose, error) {
	ch, close, err := getChannel(psm.conn)
	if err != nil {
		return nil, nil, err
	}

	return NewPublisher(ch, psm.pubHandlers), close, nil
}

// SetPublishHandler .
func (psm *PubSubManager) SetPublishHandler(handler PublishHandler) error {
	psm.pubHandlers = append(psm.pubHandlers, handler)
	return nil
}

// SetConsumer .
func (psm *PubSubManager) SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error {
	ch, close, err := getChannel(psm.conn)
	if err != nil {
		return err
	}

	if err := ch.FanoutExchangeDeclare(
		exchangeName,
	); err != nil {
		close()
		return err
	}

	if err = ch.QueueDeclareAndBind(
		exchangeName,
		queueName,
	); err != nil {
		close()
		return err
	}

	messageCh, err := ch.Consume(
		ctx,
		queueName,
	)
	if err != nil {
		close()
		return err
	}

	go func() {
		defer close()
		for {
			select {
			case <-ctx.Done():
				return
			case message, ok := <-messageCh:
				if !ok {
					return
				}

				handler(message)
			}
		}
	}()

	return nil
}

func getChannel(conn Connection) (Channel, event.ChannelClose, error) {
	ch, err := conn.GetChannel()
	if err != nil {
		return nil, nil, err
	}
	var chClose = func() error {
		return ch.Close()
	}
	return ch, chClose, nil
}
