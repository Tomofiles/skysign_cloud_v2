package ports

import (
	"remote-communication/pkg/communication/app"
	c "remote-communication/pkg/communication/domain/communication"
	"remote-communication/pkg/communication/service"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPushCommand(t *testing.T) {
	a := assert.New(t)

	service := userCommunicationServiceMock{}

	service.On("PushCommand", mock.Anything, mock.Anything).Return(DefaultCommunicationCommandID, nil)

	app := app.Application{
		Services: app.Services{
			UserCommunication: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.PushCommandRequest{
		Id:   DefaultCommunicationID,
		Type: skysign_proto.CommandType_ARM,
	}
	response, err := grpc.PushCommand(
		nil,
		request,
	)

	expectResponse := &skysign_proto.PushCommandResponse{
		Id:        DefaultCommunicationID,
		Type:      skysign_proto.CommandType_ARM,
		CommandId: DefaultCommunicationCommandID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestPullCommand(t *testing.T) {
	a := assert.New(t)

	service := edgeCommunicationServiceMock{}

	service.On("PullCommand", mock.Anything, mock.Anything).Return(string(c.CommandTypeARM), nil)

	app := app.Application{
		Services: app.Services{
			EdgeCommunication: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.PullCommandRequest{
		Id:        DefaultCommunicationID,
		CommandId: DefaultCommunicationCommandID,
	}
	response, err := grpc.PullCommand(
		nil,
		request,
	)

	expectResponse := &skysign_proto.PullCommandResponse{
		Id:        DefaultCommunicationID,
		Type:      skysign_proto.CommandType_ARM,
		CommandId: DefaultCommunicationCommandID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestPushUploadMission(t *testing.T) {
	a := assert.New(t)

	service := userCommunicationServiceMock{}

	service.On("PushUploadMission", mock.Anything, mock.Anything).Return(DefaultCommunicationCommandID, nil)

	app := app.Application{
		Services: app.Services{
			UserCommunication: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.PushUploadMissionRequest{
		Id:        DefaultCommunicationID,
		MissionId: DefaultCommunicationMissionID,
	}
	response, err := grpc.PushUploadMission(
		nil,
		request,
	)

	expectResponse := &skysign_proto.PushUploadMissionResponse{
		Id:        DefaultCommunicationID,
		CommandId: DefaultCommunicationCommandID,
		MissionId: DefaultCommunicationMissionID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestPullUploadMission(t *testing.T) {
	a := assert.New(t)

	service := edgeCommunicationServiceMock{}

	service.On("PullUploadMission", mock.Anything, mock.Anything).Return(DefaultCommunicationMissionID, nil)

	app := app.Application{
		Services: app.Services{
			EdgeCommunication: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.PullUploadMissionRequest{
		Id:        DefaultCommunicationID,
		CommandId: DefaultCommunicationCommandID,
	}
	response, err := grpc.PullUploadMission(
		nil,
		request,
	)

	expectResponse := &skysign_proto.PullUploadMissionResponse{
		Id:        DefaultCommunicationID,
		CommandId: DefaultCommunicationCommandID,
		MissionId: DefaultCommunicationMissionID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestPushTelemetry(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultCommunicationCommandID1 = DefaultCommunicationCommandID + "-1"
		DefaultCommunicationCommandID2 = DefaultCommunicationCommandID + "-2"
		DefaultCommunicationCommandID3 = DefaultCommunicationCommandID + "-3"
	)

	var commandIDs = []string{DefaultCommunicationCommandID1, DefaultCommunicationCommandID2, DefaultCommunicationCommandID3}

	service := edgeCommunicationServiceMock{}

	service.On("PushTelemetry", mock.Anything, mock.Anything).Return(commandIDs, nil)

	app := app.Application{
		Services: app.Services{
			EdgeCommunication: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.PushTelemetryRequest{
		Id: DefaultCommunicationID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            c.Armed,
			FlightMode:       "NONE",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}
	response, err := grpc.PushTelemetry(
		nil,
		request,
	)

	expectResponse := &skysign_proto.PushTelemetryResponse{
		Id:         DefaultCommunicationID,
		CommandIds: commandIDs,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestPullTelemetry(t *testing.T) {
	a := assert.New(t)

	var tlm service.UserTelemetry
	tlm = &telemetry{
		latitude:         1.0,
		longitude:        2.0,
		altitude:         3.0,
		relativeAltitude: 4.0,
		speed:            5.0,
		armed:            c.Armed,
		flightMode:       "NONE",
		x:                6.0,
		y:                7.0,
		z:                8.0,
		w:                9.0,
	}

	service := userCommunicationServiceMock{}

	service.On("PullTelemetry", mock.Anything, mock.Anything).Return(tlm, nil)

	app := app.Application{
		Services: app.Services{
			UserCommunication: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.PullTelemetryRequest{
		Id: DefaultCommunicationID,
	}
	response, err := grpc.PullTelemetry(
		nil,
		request,
	)

	expectResponse := &skysign_proto.PullTelemetryResponse{
		Id: DefaultCommunicationID,
		Telemetry: &skysign_proto.Telemetry{
			Latitude:         1.0,
			Longitude:        2.0,
			Altitude:         3.0,
			RelativeAltitude: 4.0,
			Speed:            5.0,
			Armed:            c.Armed,
			FlightMode:       "NONE",
			OrientationX:     6.0,
			OrientationY:     7.0,
			OrientationZ:     8.0,
			OrientationW:     9.0,
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
