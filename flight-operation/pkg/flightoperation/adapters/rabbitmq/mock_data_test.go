package rabbitmq

import (
	"context"

	fo "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/domain/flightoperation"
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/service"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightoperationID = fo.ID("flightoperation-id")
const DefaultFlightoperationName = "flightoperation-name"
const DefaultFlightoperationDescription = "flightoperation-description"
const DefaultFlightoperationFleetID = fo.FleetID("fleet-id")

type manageFlightoperationServiceMock struct {
	mock.Mock
	name, description, fleetID string
}

func (s *manageFlightoperationServiceMock) GetFlightoperation(
	command service.GetFlightoperationCommand,
	model service.RetrievedModel,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageFlightoperationServiceMock) ListFlightoperations(
	model service.RetrievedModel,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageFlightoperationServiceMock) CreateFlightoperation(
	command service.CreateFlightoperationCommand,
) error {
	s.name = command.GetFlightoperation().GetName()
	s.description = command.GetFlightoperation().GetDescription()
	s.fleetID = command.GetFlightoperation().GetFleetID()
	ret := s.Called()
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
