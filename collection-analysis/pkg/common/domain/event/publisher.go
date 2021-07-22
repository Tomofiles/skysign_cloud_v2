package event

// Publisher .
type Publisher interface {
	Publish(interface{})
	Flush() error
}
