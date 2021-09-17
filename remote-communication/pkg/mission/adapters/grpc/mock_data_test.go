package grpc

import (
	"remote-communication/pkg/mission/service"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = "mission-id"
const DefaultMissionUploadID = "upload-id"

type edgeMissionServiceMock struct {
	mock.Mock
}

func (s *edgeMissionServiceMock) PullMission(
	command service.PullMissionCommand,
	pulledMission service.PulledMission,
) error {
	ret := s.Called()
	var id string
	if ret.Get(0) != nil {
		id = ret.Get(0).(string)
	}
	var f []service.Waypoint
	if ret.Get(1) != nil {
		f = ret.Get(1).([]service.Waypoint)
	}
	pulledMission(id, f)
	return ret.Error(2)
}
