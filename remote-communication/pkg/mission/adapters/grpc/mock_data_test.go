package grpc

import (
	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/service"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
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

type serviceRegistrarMock struct {
	descs []*grpc.ServiceDesc
	impls []interface{}
}

func (s *serviceRegistrarMock) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.descs = append(s.descs, desc)
	s.impls = append(s.impls, impl)
}
