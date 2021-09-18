package rabbitmq

import (
	act "collection-analysis/pkg/action/domain/action"
	"collection-analysis/pkg/action/service"
	"context"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/stretchr/testify/mock"
)

const DefaultActionID = act.ID("action-id")
const DefaultActionCommunicationID = act.CommunicationID("communication-id")
const DefaultActionFleetID = act.FleetID("fleet-id")

type manageActionServiceMock struct {
	mock.Mock
	command service.CreateActionCommand
}

func (s *manageActionServiceMock) CreateAction(
	command service.CreateActionCommand,
) error {
	ret := s.Called()
	s.command = command
	return ret.Error(0)
}

func (s *manageActionServiceMock) GetTrajectory(
	command service.GetTrajectoryCommand,
	telemetry service.TelemetrySnapshot,
) error {
	ret := s.Called()
	if snapshots := ret.Get(0); snapshots != nil {
		for _, s := range snapshots.([]act.TelemetrySnapshot) {
			telemetry(s)
		}
	}
	return ret.Error(1)
}

type operateActionServiceMock struct {
	mock.Mock
	completeCommand  service.CompleteActionCommand
	telemetryCommand service.PushTelemetryCommand
}

func (s *operateActionServiceMock) CompleteAction(
	command service.CompleteActionCommand,
) error {
	ret := s.Called()
	s.completeCommand = command
	return ret.Error(0)
}

func (s *operateActionServiceMock) PushTelemetry(
	command service.PushTelemetryCommand,
) error {
	ret := s.Called()
	s.telemetryCommand = command
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
