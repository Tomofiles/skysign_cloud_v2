package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

// TestAdapterPause .
func TestAdapterPause(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_mission.PauseMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_SUCCESS,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("PauseMission", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterPauseInternal(ctx, supportMock, missionMock)

	a.Nil(ret)
}

// TestRequestErrorWhenAdapterPause .
func TestRequestErrorWhenAdapterPause(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	missionMock := &missionServiceClientMock{}
	missionMock.On("PauseMission", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterPauseInternal(ctx, supportMock, missionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("pause command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterPause .
func TestResponseErrorWhenAdapterPause(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_mission.PauseMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_BUSY,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("PauseMission", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterPauseInternal(ctx, supportMock, missionMock)

	a.Equal(ret, ErrPauseCommand)
	a.Equal("pause command error: no pause command success", supportMock.message)
}
