package rabbitmq

import (
	"context"
	"remote-communication/pkg/mission/service"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = "mission-id"
const DefaultMissionUploadID = "upload-id"

type manageMissionServiceMock struct {
	mock.Mock
	ID        string
	Waypoints []service.Waypoint
}

func (s *manageMissionServiceMock) CreateMission(
	command service.CreateMissionCommand,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.Waypoints = command.GetWaypoints()
	return ret.Error(0)
}

type publishHandlerMock struct {
	consumers []consumer
}

func (h *publishHandlerMock) SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error {
	h.consumers = append(
		h.consumers,
		consumer{
			exchangeName: exchangeName,
			queueName:    queueName,
			handler:      handler,
		})
	return nil
}

type consumer struct {
	exchangeName, queueName string
	handler                 func([]byte)
}

type eventHandlerMock struct {
	events1 []byte
}

func (h *eventHandlerMock) HandleCopiedMissionCreatedEvent(
	ctx context.Context,
	event []byte,
) error {
	h.events1 = append(h.events1, event...)
	return nil
}
