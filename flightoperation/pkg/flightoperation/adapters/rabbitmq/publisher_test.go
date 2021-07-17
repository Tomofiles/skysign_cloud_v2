package rabbitmq

import (
	"context"
	"errors"
	fope "flightoperation/pkg/flightoperation/domain/flightoperation"
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

	event := fope.FlightoperationCompletedEvent{}

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

	event1 := fope.FlightoperationCompletedEvent{}
	event2 := fope.FleetCopiedEvent{}

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

func TestFanoutExchangeDeclareErrorWhenFlightoperationCompletedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fope.FlightoperationCompletedEvent{}

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

func TestPublishErrorWhenFlightoperationCompletedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fope.FlightoperationCompletedEvent{}

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

func TestFanoutExchangeDeclareErrorWhenFleetCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fope.FleetCopiedEvent{}

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

func TestPublishErrorWhenFleetCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fope.FleetCopiedEvent{}

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
