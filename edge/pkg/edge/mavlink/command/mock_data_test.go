package mavlink

import (
	"context"
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	mavsdk_rpc_action "edge/pkg/protos/action"
	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

var (
	ErrRequest = errors.New("request error")
)

type supportMock struct {
	message string
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.message = fmt.Sprintf(format, args...)
}

type actionServiceClientMock struct {
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

func (m *actionServiceClientMock) Reboot(ctx context.Context, in *mavsdk_rpc_action.RebootRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.RebootResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) Kill(ctx context.Context, in *mavsdk_rpc_action.KillRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.KillResponse, error) {
	panic("implement me")
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

func (m *actionServiceClientMock) TransitionToFixedWing(ctx context.Context, in *mavsdk_rpc_action.TransitionToFixedWingRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.TransitionToFixedWingResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) TransitionToMulticopter(ctx context.Context, in *mavsdk_rpc_action.TransitionToMulticopterRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.TransitionToMulticopterResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) GetTakeoffAltitude(ctx context.Context, in *mavsdk_rpc_action.GetTakeoffAltitudeRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.GetTakeoffAltitudeResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) SetTakeoffAltitude(ctx context.Context, in *mavsdk_rpc_action.SetTakeoffAltitudeRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.SetTakeoffAltitudeResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) GetMaximumSpeed(ctx context.Context, in *mavsdk_rpc_action.GetMaximumSpeedRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.GetMaximumSpeedResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) SetMaximumSpeed(ctx context.Context, in *mavsdk_rpc_action.SetMaximumSpeedRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.SetMaximumSpeedResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) GetReturnToLaunchAltitude(ctx context.Context, in *mavsdk_rpc_action.GetReturnToLaunchAltitudeRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.GetReturnToLaunchAltitudeResponse, error) {
	panic("implement me")
}

func (m *actionServiceClientMock) SetReturnToLaunchAltitude(ctx context.Context, in *mavsdk_rpc_action.SetReturnToLaunchAltitudeRequest, opts ...grpc.CallOption) (*mavsdk_rpc_action.SetReturnToLaunchAltitudeResponse, error) {
	panic("implement me")
}

type missionServiceClientMock struct {
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

func (m *missionServiceClientMock) CancelMissionUpload(ctx context.Context, in *mavsdk_rpc_mission.CancelMissionUploadRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.CancelMissionUploadResponse, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) DownloadMission(ctx context.Context, in *mavsdk_rpc_mission.DownloadMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.DownloadMissionResponse, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) CancelMissionDownload(ctx context.Context, in *mavsdk_rpc_mission.CancelMissionDownloadRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.CancelMissionDownloadResponse, error) {
	panic("implement me")
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

func (m *missionServiceClientMock) ClearMission(ctx context.Context, in *mavsdk_rpc_mission.ClearMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.ClearMissionResponse, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) SetCurrentMissionItemIndex(ctx context.Context, in *mavsdk_rpc_mission.SetCurrentMissionItemIndexRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.SetCurrentMissionItemIndexResponse, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) IsMissionFinished(ctx context.Context, in *mavsdk_rpc_mission.IsMissionFinishedRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.IsMissionFinishedResponse, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) SubscribeMissionProgress(ctx context.Context, in *mavsdk_rpc_mission.SubscribeMissionProgressRequest, opts ...grpc.CallOption) (mavsdk_rpc_mission.MissionService_SubscribeMissionProgressClient, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) GetReturnToLaunchAfterMission(ctx context.Context, in *mavsdk_rpc_mission.GetReturnToLaunchAfterMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.GetReturnToLaunchAfterMissionResponse, error) {
	panic("implement me")
}

func (m *missionServiceClientMock) SetReturnToLaunchAfterMission(ctx context.Context, in *mavsdk_rpc_mission.SetReturnToLaunchAfterMissionRequest, opts ...grpc.CallOption) (*mavsdk_rpc_mission.SetReturnToLaunchAfterMissionResponse, error) {
	panic("implement me")
}
