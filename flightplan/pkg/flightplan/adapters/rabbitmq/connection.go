package rabbitmq

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const urlTemplate = "amqp://%s:%s@%s:%s/"
const urlTLSTemplate = "amqps://%s:%s@%s:%s/"

var (
	host     = "localhost"
	user     = "guest"
	password = "guest"
	port     = "5672"
	sslmode  = "disable"
)

// NewRabbitMQConnection .
func NewRabbitMQConnection() (Connection, error) {
	if envHost := os.Getenv("MQ_HOST"); envHost != "" {
		host = envHost
	}
	if envUser := os.Getenv("MQ_USERNAME"); envUser != "" {
		user = envUser
	}
	if envPassword := os.Getenv("MQ_PASSWORD"); envPassword != "" {
		password = envPassword
	}
	if envPort := os.Getenv("MQ_PORT"); envPort != "" {
		port = envPort
	}
	if envSslmode := os.Getenv("MQ_SSL_ENABLED"); envSslmode != "" {
		sslmode = envSslmode
	}

	var url string
	if sslmode == "enable" {
		url = fmt.Sprintf(urlTLSTemplate, user, password, host, port)
	} else {
		url = fmt.Sprintf(urlTemplate, user, password, host, port)
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to RabbitMQ")
	}

	connection := NewConnection(conn)

	return connection, nil
}

// Connection .
type Connection interface {
	GetChannel() (Channel, error)
	Close() error
}

// Channel .
type Channel interface {
	FanoutExchangeDeclare(exchange string) error
	QueueDeclareAndBind(exchange, queue string) error
	Publish(queue string, message Message) error
	Consume(ctx context.Context, queue string) (<-chan Message, error)
	Close() error
}

// Message .
type Message = []byte
