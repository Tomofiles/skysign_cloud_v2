package mavlink

import (
	"context"
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	mavsdk_rpc_action "edge-px4/pkg/protos/action"
	mavsdk_rpc_mission "edge-px4/pkg/protos/mission"
)

var (
	ErrRequest = errors.New("request error")
)

type supportMock struct {
	message string
}

func (m *supportMock) NotifyInfo(format string, args ...interface{}) {
	m.message = fmt.Sprintf(format, args...)
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.message = fmt.Sprintf(format, args...)
}

type actionServiceClientMock struct {
	mavsdk_rpc_action.ActionServiceClient
	mock.Mock
}

func (m *actionServiceClientMock) Arm(ctx context.Context, in *mavsdk_rpc_action.ArmRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.ArmResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_action.ArmResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_action.ArmResponse)
	}
	return v, ret.Error(1)
}

func (m *actionServiceClientMock) Disarm(ctx context.Context, in *mavsdk_rpc_action.DisarmRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.DisarmResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_action.DisarmResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_action.DisarmResponse)
	}
	return v, ret.Error(1)
}

func (m *actionServiceClientMock) Takeoff(ctx context.Context, in *mavsdk_rpc_action.TakeoffRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.TakeoffResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_action.TakeoffResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_action.TakeoffResponse)
	}
	return v, ret.Error(1)
}

func (m *actionServiceClientMock) Land(ctx context.Context, in *mavsdk_rpc_action.LandRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.LandResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_action.LandResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_action.LandResponse)
	}
	return v, ret.Error(1)
}

func (m *actionServiceClientMock) ReturnToLaunch(ctx context.Context, in *mavsdk_rpc_action.ReturnToLaunchRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.ReturnToLaunchResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_action.ReturnToLaunchResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_action.ReturnToLaunchResponse)
	}
	return v, ret.Error(1)
}

type missionServiceClientMock struct {
	mavsdk_rpc_mission.MissionServiceClient
	mock.Mock
	MissionItems []*mavsdk_rpc_mission.MissionItem
}

func (m *missionServiceClientMock) UploadMission(ctx context.Context, in *mavsdk_rpc_mission.UploadMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.UploadMissionResponse, error) {
	ret := m.Called(in)
	m.MissionItems = in.MissionItems
	var v *mavsdk_rpc_mission.UploadMissionResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_mission.UploadMissionResponse)
	}
	return v, ret.Error(1)
}

func (m *missionServiceClientMock) StartMission(ctx context.Context, in *mavsdk_rpc_mission.StartMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.StartMissionResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_mission.StartMissionResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_mission.StartMissionResponse)
	}
	return v, ret.Error(1)
}

func (m *missionServiceClientMock) PauseMission(ctx context.Context, in *mavsdk_rpc_mission.PauseMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.PauseMissionResponse, error) {
	ret := m.Called(in)
	var v *mavsdk_rpc_mission.PauseMissionResponse
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*mavsdk_rpc_mission.PauseMissionResponse)
	}
	return v, ret.Error(1)
}
