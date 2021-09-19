package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/domain/fleet"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/service"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/stretchr/testify/mock"
)

const DefaultFleetNumberOfVehicles = 3
const DefaultFleetID = fleet.ID("fleet-id")
const DefaultFleetVehicleID = fleet.VehicleID("vehicle-id")
const DefaultFleetMissionID = fleet.MissionID("mission-id")

type manageFleetServiceMock struct {
	mock.Mock
	ID               string
	NumberOfVehicles int
	OriginalID       string
	NewID            string
}

func (s *manageFleetServiceMock) CreateFleet(
	command service.CreateFleetCommand,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.NumberOfVehicles = command.GetNumberOfVehicles()
	return ret.Error(0)
}

func (s *manageFleetServiceMock) DeleteFleet(
	command service.DeleteFleetCommand,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	return ret.Error(0)
}

func (s *manageFleetServiceMock) CarbonCopyFleet(
	command service.CarbonCopyFleetCommand,
) error {
	ret := s.Called()
	s.OriginalID = command.GetOriginalID()
	s.NewID = command.GetNewID()
	return ret.Error(0)
}

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

type pubSubManagerMock struct {
	consumers       []consumer
	publishHandlers []func(ch crm.Channel, e interface{})
}

func (h *pubSubManagerMock) SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error {
	h.consumers = append(
		h.consumers,
		consumer{
			exchangeName: exchangeName,
			queueName:    queueName,
			handler:      handler,
		})
	return nil
}

func (h *pubSubManagerMock) SetPublishHandler(handler func(ch crm.Channel, e interface{})) error {
	h.publishHandlers = append(h.publishHandlers, handler)
	return nil
}

type consumer struct {
	exchangeName, queueName string
	handler                 func([]byte)
}
