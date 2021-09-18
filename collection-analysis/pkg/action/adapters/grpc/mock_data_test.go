package grpc

import (
	act "collection-analysis/pkg/action/domain/action"
	"collection-analysis/pkg/action/service"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
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

type serviceRegistrarMock struct {
	descs []*grpc.ServiceDesc
	impls []interface{}
}

func (s *serviceRegistrarMock) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.descs = append(s.descs, desc)
	s.impls = append(s.impls, impl)
}
