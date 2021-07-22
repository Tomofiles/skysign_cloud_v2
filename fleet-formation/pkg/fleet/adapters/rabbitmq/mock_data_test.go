package rabbitmq

import (
	"context"
	crm "fleet-formation/pkg/common/adapters/rabbitmq"
	"fleet-formation/pkg/fleet/domain/fleet"

	"github.com/stretchr/testify/mock"
)

const DefaultFleetID = fleet.ID("fleet-id")
const DefaultFleetVehicleID = fleet.VehicleID("vehicle-id")
const DefaultFleetMissionID = fleet.MissionID("mission-id")

type channelMockPublish struct {
	mock.Mock
	message          crm.Message
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

func (ch *channelMockPublish) Publish(queue string, message crm.Message) error {
	ret := ch.Called()
	ch.message = message
	ch.messageCallCount++
	return ret.Error(0)
}

func (ch *channelMockPublish) Consume(ctx context.Context, queue string) (<-chan crm.Message, error) {
	panic("implement me")
}

func (ch *channelMockPublish) Close() error {
	panic("implement me")
}

type publishHandlerMock struct {
	publishHandlers []func(ch crm.Channel, e interface{})
}

func (h *publishHandlerMock) SetPublishHandler(handler func(ch crm.Channel, e interface{})) error {
	h.publishHandlers = append(h.publishHandlers, handler)
	return nil
}
