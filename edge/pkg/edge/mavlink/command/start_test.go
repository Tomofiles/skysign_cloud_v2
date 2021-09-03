package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mavsdk_rpc_mission "edge/pkg/protos/mission"
)

// TestAdapterStart .
func TestAdapterStart(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_mission.StartMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_SUCCESS,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("StartMission", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterStartInternal(ctx, supportMock, missionMock)

	a.Nil(ret)
	a.Empty(supportMock.message)
}

// TestRequestErrorWhenAdapterStart .
func TestRequestErrorWhenAdapterStart(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	missionMock := &missionServiceClientMock{}
	missionMock.On("StartMission", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	ret := AdapterStartInternal(ctx, supportMock, missionMock)

	a.Equal(ret, ErrRequest)
	a.Equal("start command error: request error", supportMock.message)
}

// TestResponseErrorWhenAdapterStart .
func TestResponseErrorWhenAdapterStart(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	supportMock := &supportMock{}

	response := &mavsdk_rpc_mission.StartMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_BUSY,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("StartMission", mock.Anything, mock.Anything).Return(response, nil)

	ret := AdapterStartInternal(ctx, supportMock, missionMock)

	a.Equal(ret, ErrStartCommand)
	a.Equal("start command error: no start command success", supportMock.message)
}
