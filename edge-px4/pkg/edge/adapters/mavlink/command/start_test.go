package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_mission "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/protos/mission"
)

// TestAdapterStart .
func TestAdapterStart(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_mission.StartMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_SUCCESS,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("StartMission", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterStartInternal(ctx, missionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterStart .
func TestRequestErrorWhenAdapterStart(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	missionMock := &missionServiceClientMock{}
	missionMock.On("StartMission", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterStartInternal(ctx, missionMock)

	a.Equal("start command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterStart .
func TestResponseErrorWhenAdapterStart(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_mission.StartMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_BUSY,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("StartMission", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterStartInternal(ctx, missionMock)

	a.Equal("start command error: no start command success", ret.Error())
}
