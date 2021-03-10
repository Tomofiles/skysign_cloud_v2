package rabbitmq

import (
	"context"
	"errors"
	frep "flightreport/pkg/flightreport/domain/flightreport"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

	event := frep.CreatedEvent{}

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event)
	ret := pub.Flush()

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 1)
}

func TestMultipleEventWhenPublish(t *testing.T) {
	a := assert.New(t)

	event1 := frep.CreatedEvent{}
	event2 := frep.FlightoperationCopiedWhenCreatedEvent{}

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event1)
	pub.Publish(event2)
	ret := pub.Flush()

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 2)
}

func TestFanoutExchangeDeclareErrorWhenCreatedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := frep.CreatedEvent{}

	errPub := errors.New("publish error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event)
	ret := pub.Flush()

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenCreatedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := frep.CreatedEvent{}

	errPub := errors.New("publish error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event)
	ret := pub.Flush()

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
}

func TestFanoutExchangeDeclareErrorWhenFlightoperationCopiedWhenCreatedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := frep.FlightoperationCopiedWhenCreatedEvent{}

	errPub := errors.New("publish error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errPub)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event)
	ret := pub.Flush()

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 0)
}

func TestPublishErrorWhenFlightoperationCopiedWhenCreatedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := frep.FlightoperationCopiedWhenCreatedEvent{}

	errPub := errors.New("publish error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(errPub)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event)
	ret := pub.Flush()

	a.Equal(ret, errPub)
	a.Equal(chMock.messageCallCount, 1)
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
