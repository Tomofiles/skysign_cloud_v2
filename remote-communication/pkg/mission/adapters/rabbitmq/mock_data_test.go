package rabbitmq

import (
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
