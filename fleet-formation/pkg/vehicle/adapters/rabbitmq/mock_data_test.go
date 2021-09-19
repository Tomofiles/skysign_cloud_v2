package rabbitmq

import (
	"context"

	v "github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/domain/vehicle"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/service"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/stretchr/testify/mock"
)

const DefaultVehicleID = v.ID("vehicle-id")
const DefaultVehicleCommunicationID = v.CommunicationID("communication-id")
const DefaultFleetID = v.FleetID("fleet-id")

type manageVehicleServiceMock struct {
	mock.Mock
	OriginalID string
	NewID      string
	FleetID    string
}

func (s *manageVehicleServiceMock) GetVehicle(
	command service.GetVehicleCommand,
	model service.RetrievedModel,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) ListVehicles(
	model service.RetrievedModel,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) CreateVehicle(
	command service.CreateVehicleCommand,
	createdID service.CreatedID,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) UpdateVehicle(
	command service.UpdateVehicleCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) DeleteVehicle(
	command service.DeleteVehicleCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageVehicleServiceMock) CarbonCopyVehicle(
	command service.CarbonCopyVehicleCommand,
) error {
	ret := s.Called()
	s.OriginalID = command.GetOriginalID()
	s.NewID = command.GetNewID()
	s.FleetID = command.GetFleetID()
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
