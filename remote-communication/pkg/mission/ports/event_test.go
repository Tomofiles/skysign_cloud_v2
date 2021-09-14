package ports

import (
	"remote-communication/pkg/mission/app"
	"remote-communication/pkg/mission/service"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleCopiedMissionCreatedEvent(t *testing.T) {
	a := assert.New(t)

	s := manageMissionServiceMock{}

	s.On("CreateMission", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &s,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.CopiedMissionCreatedEvent{
		MissionId: DefaultMissionID,
		Mission: &skysign_proto.Mission{
			Id: DefaultMissionID,
			Navigation: &skysign_proto.Navigation{
				UploadId: DefaultMissionUploadID,
				Waypoints: []*skysign_proto.Waypoint{
					{
						Latitude:       11.0,
						Longitude:      21.0,
						RelativeHeight: 31.0,
						Speed:          41.0,
					},
				},
			},
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleCopiedMissionCreatedEvent(
		nil,
		requestBin,
	)

	expectWaypoints := []service.Waypoint{
		&waypoint{
			LatitudeDegree:  11.0,
			LongitudeDegree: 21.0,
			RelativeHeightM: 31.0,
			SpeedMS:         41.0,
		},
	}

	a.Nil(err)
	a.Equal(s.ID, DefaultMissionUploadID)
	a.Equal(s.Waypoints, expectWaypoints)
}
