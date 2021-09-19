package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/service"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

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
