package event

// PubSubManager .
type PubSubManager interface {
	GetPublisher() (Publisher, ChannelClose, error)
}

// ChannelClose .
type ChannelClose = func() error
