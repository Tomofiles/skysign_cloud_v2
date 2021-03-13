package rabbitmq

import (
	"context"
	"errors"
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
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

	event := flightplan.CreatedEvent{}

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

	event1 := flightplan.CreatedEvent{}
	event2 := flightplan.DeletedEvent{}
	event3 := flightplan.CopiedEvent{}
	event4 := fleet.VehicleCopiedWhenFlightplanCopiedEvent{}
	event5 := fleet.MissionCopiedWhenFlightplanCopiedEvent{}

	connMock := &connectionMockCommon{}
	chMock := &channelMockPublish{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("Publish", mock.Anything).Return(nil)

	psm := NewPubSubManager(connMock)

	pub, _, _ := psm.GetPublisher()

	pub.Publish(event1)
	pub.Publish(event2)
	pub.Publish(event3)
	pub.Publish(event4)
	pub.Publish(event5)
	ret := pub.Flush()

	a.Nil(ret)
	a.Equal(chMock.messageCallCount, 5)
}

func TestFanoutExchangeDeclareErrorWhenCreatedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := flightplan.CreatedEvent{}

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

	event := flightplan.CreatedEvent{}

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

func TestFanoutExchangeDeclareErrorWhenDeletedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := flightplan.DeletedEvent{}

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

func TestPublishErrorWhenDeletedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := flightplan.DeletedEvent{}

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

func TestFanoutExchangeDeclareErrorWhenCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := flightplan.CopiedEvent{}

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

func TestPublishErrorWhenCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := flightplan.CopiedEvent{}

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

func TestFanoutExchangeDeclareErrorWhenVehicleCopiedWhenFlightplanCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fleet.VehicleCopiedWhenFlightplanCopiedEvent{}

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

func TestPublishErrorWhenVehicleCopiedWhenFlightplanCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fleet.VehicleCopiedWhenFlightplanCopiedEvent{}

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

func TestFanoutExchangeDeclareErrorWhenMissionCopiedWhenFlightplanCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fleet.MissionCopiedWhenFlightplanCopiedEvent{}

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

func TestPublishErrorWhenMissionCopiedWhenFlightplanCopiedEventPublish(t *testing.T) {
	a := assert.New(t)

	event := fleet.MissionCopiedWhenFlightplanCopiedEvent{}

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
