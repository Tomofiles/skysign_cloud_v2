package rabbitmq

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// SampleEvent1 .
type SampleEvent1 struct {
	message string
}

// SampleEvent2 .
type SampleEvent2 struct {
	message string
}

func TestNoneEventWhenPublish(t *testing.T) {
	a := assert.New(t)

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	ret := pub.Flush()

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 0)
}

func TestSingleEventWhenPublish(t *testing.T) {
	a := assert.New(t)

	event := SampleEvent1{
		message: "aaa",
	}

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	var handledEvent interface{}
	psm.SetPublishHandler(
		func(c Channel, i interface{}) {
			c.Publish("", []byte{})
			handledEvent = i
		},
	)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event)
	ret := pub.Flush()

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
	a.IsType(SampleEvent1{}, handledEvent)
}

func TestMultipleEventWhenPublish(t *testing.T) {
	a := assert.New(t)

	event1 := SampleEvent1{
		message: "aaa",
	}
	event2 := SampleEvent2{
		message: "bbb",
	}

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	var handledEvents []interface{}
	psm.SetPublishHandler(
		func(c Channel, i interface{}) {
			if event, ok := i.(SampleEvent1); ok {
				c.Publish("", []byte{})
				handledEvents = append(handledEvents, event)
			}
		},
	)
	psm.SetPublishHandler(
		func(c Channel, i interface{}) {
			if event, ok := i.(SampleEvent2); ok {
				c.Publish("", []byte{})
				handledEvents = append(handledEvents, event)
			}
		},
	)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event1)
	pub.Publish(event2)
	ret := pub.Flush()

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 2)
	a.IsType(SampleEvent1{}, handledEvents[0])
	a.IsType(SampleEvent2{}, handledEvents[1])
}

type channelMockPublish struct {
	mock.Mock
	message          Message
	messageCallCount int
	isClose          bool
}

func (ch *channelMockPublish) FanoutExchangeDeclare(exchange string) error {
	ret := ch.Called()
	return ret.Error(0)
}

func (ch *channelMockPublish) QueueDeclareAndBind(exchange, queue string) error {
	panic("implement me")
}

func (ch *channelMockPublish) Publish(queue string, message Message) error {
	ret := ch.Called()
	ch.message = message
	ch.messageCallCount++
	return ret.Error(0)
}

func (ch *channelMockPublish) Consume(ctx context.Context, queue string) (<-chan Message, error) {
	panic("implement me")
}

func (ch *channelMockPublish) Close() error {
	panic("implement me")
}
