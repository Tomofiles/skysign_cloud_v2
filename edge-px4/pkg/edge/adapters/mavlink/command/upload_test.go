package mavlink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"edge-px4/pkg/edge/domain/model"
	mavsdk_rpc_mission "edge-px4/pkg/protos/mission"
)

// TestAdapterUploadNoItems .
func TestAdapterUploadNoItems(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_mission.UploadMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_SUCCESS,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("UploadMission", mock.Anything, mock.Anything).Return(response, nil)

	model := &model.Mission{}
	ret := AdapterUploadInternal(ctx, missionMock, model)

	expectMissionItems := []*mavsdk_rpc_mission.MissionItem{}
	a.Nil(ret)
	a.Equal(expectMissionItems, missionMock.MissionItems)
}

// TestAdapterUploadSingleItems .
func TestAdapterUploadSingleItems(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_mission.UploadMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_SUCCESS,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("UploadMission", mock.Anything, mock.Anything).Return(response, nil)

	model := &model.Mission{
		Waypoints: []*model.Waypoints{
			{
				Latitude:       1.0,
				Longitude:      2.0,
				RelativeHeight: 3.0,
				Speed:          4.0,
			},
		},
	}
	ret := AdapterUploadInternal(ctx, missionMock, model)

	expectMissionItems := []*mavsdk_rpc_mission.MissionItem{
		{
			LatitudeDeg:       1.0,
			LongitudeDeg:      2.0,
			RelativeAltitudeM: 3.0,
			SpeedMS:           4.0,
		},
	}
	a.Nil(ret)
	a.Equal(expectMissionItems, missionMock.MissionItems)
}

// TestAdapterUploadMultipleItems .
func TestAdapterUploadMultipleItems(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_mission.UploadMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_SUCCESS,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("UploadMission", mock.Anything, mock.Anything).Return(response, nil)

	model := &model.Mission{
		Waypoints: []*model.Waypoints{
			{
				Latitude:       11.0,
				Longitude:      21.0,
				RelativeHeight: 31.0,
				Speed:          41.0,
			},
			{
				Latitude:       12.0,
				Longitude:      22.0,
				RelativeHeight: 32.0,
				Speed:          42.0,
			},
			{
				Latitude:       13.0,
				Longitude:      23.0,
				RelativeHeight: 33.0,
				Speed:          43.0,
			},
		},
	}
	ret := AdapterUploadInternal(ctx, missionMock, model)

	expectMissionItems := []*mavsdk_rpc_mission.MissionItem{
		{
			LatitudeDeg:       11.0,
			LongitudeDeg:      21.0,
			RelativeAltitudeM: 31.0,
			SpeedMS:           41.0,
		},
		{
			LatitudeDeg:       12.0,
			LongitudeDeg:      22.0,
			RelativeAltitudeM: 32.0,
			SpeedMS:           42.0,
		},
		{
			LatitudeDeg:       13.0,
			LongitudeDeg:      23.0,
			RelativeAltitudeM: 33.0,
			SpeedMS:           43.0,
		},
	}
	a.Nil(ret)
	a.Equal(expectMissionItems, missionMock.MissionItems)
}

// TestRequestErrorWhenAdapterUpload .
func TestRequestErrorWhenAdapterUpload(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	missionMock := &missionServiceClientMock{}
	missionMock.On("UploadMission", mock.Anything, mock.Anything).Return(nil, ErrRequest)

	model := &model.Mission{}
	ret := AdapterUploadInternal(ctx, missionMock, model)

	a.Equal("upload command error: request error", ret.Error())
}

// TestResponseErrorWhenAdapterUpload .
func TestResponseErrorWhenAdapterUpload(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	response := &mavsdk_rpc_mission.UploadMissionResponse{
		MissionResult: &mavsdk_rpc_mission.MissionResult{
			Result: mavsdk_rpc_mission.MissionResult_BUSY,
		},
	}
	missionMock := &missionServiceClientMock{}
	missionMock.On("UploadMission", mock.Anything, mock.Anything).Return(response, nil)

	model := &model.Mission{}
	ret := AdapterUploadInternal(ctx, missionMock, model)

	a.Equal("upload command error: no upload command success", ret.Error())
}
