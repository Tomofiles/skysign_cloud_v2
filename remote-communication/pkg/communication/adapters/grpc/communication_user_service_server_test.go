package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/app"
	c "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/domain/communication"
	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/service"

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

	grpc := NewCommunicationUserServiceServer(app)

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

func TestPushUploadMission(t *testing.T) {
	a := assert.New(t)

	service := userCommunicationServiceMock{}

	service.On("PushUploadMission", mock.Anything, mock.Anything).Return(DefaultCommunicationCommandID, nil)

	app := app.Application{
		Services: app.Services{
			UserCommunication: &service,
		},
	}

	grpc := NewCommunicationUserServiceServer(app)

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

func TestPullTelemetry(t *testing.T) {
	a := assert.New(t)

	var tlm service.UserTelemetry
	tlm = &telemetry{
		latitudeDegree:    1.0,
		longitudeDegree:   2.0,
		altitudeM:         3.0,
		relativeAltitudeM: 4.0,
		speedMS:           5.0,
		armed:             c.Armed,
		flightMode:        "NONE",
		x:                 6.0,
		y:                 7.0,
		z:                 8.0,
		w:                 9.0,
	}

	service := userCommunicationServiceMock{}

	service.On("PullTelemetry", mock.Anything, mock.Anything).Return(tlm, nil)

	app := app.Application{
		Services: app.Services{
			UserCommunication: &service,
		},
	}

	grpc := NewCommunicationUserServiceServer(app)

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
