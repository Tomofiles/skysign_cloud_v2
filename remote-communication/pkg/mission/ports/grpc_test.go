package ports

import (
	"remote-communication/pkg/mission/app"
	"remote-communication/pkg/mission/service"
	"testing"

	proto "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUploadMission(t *testing.T) {
	a := assert.New(t)

	s := edgeMissionServiceMock{}

	waypoints := []service.Waypoint{
		&waypoint{
			LatitudeDegree:    1.0,
			LongitudeDegree:   2.0,
			RelativeAltitudeM: 3.0,
			SpeedMS:           4.0,
		},
	}
	s.On("PullMission", mock.Anything, mock.Anything).Return(DefaultMissionID, waypoints, nil)

	app := app.Application{
		Services: app.Services{
			EdgeMission: &s,
		},
	}

	grpc := NewGrpcServer(app)

	request := &proto.GetUploadMissionRequest{
		Id: DefaultMissionUploadID,
	}
	response, err := grpc.GetUploadMission(
		nil,
		request,
	)

	expectResponse := &proto.UploadMission{
		Id: DefaultMissionUploadID,
		Waypoints: []*proto.Waypoint{
			{
				Latitude:       1.0,
				Longitude:      2.0,
				RelativeHeight: 3.0,
				Speed:          4.0,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
