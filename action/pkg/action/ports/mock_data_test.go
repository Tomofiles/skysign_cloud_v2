package ports

import (
	act "action/pkg/action/domain/action"
	"action/pkg/action/service"

	"github.com/stretchr/testify/mock"
)

const DefaultActionID = act.ID("action-id")
const DefaultActionCommunicationID = act.CommunicationID("communication-id")
const DefaultActionFlightplanID = act.FlightplanID("flightplan-id")

type manageActionServiceMock struct {
	mock.Mock
	requestDpo service.CreateActionRequestDpo
}

func (s *manageActionServiceMock) CreateAction(
	requestDpo service.CreateActionRequestDpo,
) error {
	ret := s.Called()
	s.requestDpo = requestDpo
	return ret.Error(0)
}

func (s *manageActionServiceMock) GetTrajectory(
	requestDpo service.GetTrajectoryRequestDpo,
	responseEachDpo service.GetTrajectoryResponseDpo,
) error {
	ret := s.Called()
	if snapshots := ret.Get(0); snapshots != nil {
		for _, s := range snapshots.([]act.TelemetrySnapshot) {
			responseEachDpo(s)
		}
	}
	return ret.Error(1)
}

type operateActionServiceMock struct {
	mock.Mock
	completeRequestDpo  service.CompleteActionRequestDpo
	telemetryRequestDpo service.PushTelemetryRequestDpo
}

func (s *operateActionServiceMock) CompleteAction(
	requestDpo service.CompleteActionRequestDpo,
) error {
	ret := s.Called()
	s.completeRequestDpo = requestDpo
	return ret.Error(0)
}

func (s *operateActionServiceMock) PushTelemetry(
	requestDpo service.PushTelemetryRequestDpo,
) error {
	ret := s.Called()
	s.telemetryRequestDpo = requestDpo
	return ret.Error(0)
}
