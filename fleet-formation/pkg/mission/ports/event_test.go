package ports

import (
	"fleet-formation/pkg/mission/app"
	m "fleet-formation/pkg/mission/domain/mission"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleMissionCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultMissionID + "-original"
		DefaultNewID      = DefaultMissionID + "-new"
	)

	service := manageMissionServiceMock{}

	missionModel := &missionModelMock{
		mission: m.AssembleFrom(
			nil,
			&missionComponentMock{
				ID:   string(DefaultMissionID),
				Name: DefaultMissionName,
				Navigation: navigationComponentMock{
					TakeoffPointGroundAltitudeM: DefaultMissionTakeoffPointGroundAltitudeM,
					Waypoints:                   []waypointComponentMock{},
					UploadID:                    string(DefaultMissionUploadID),
				},
				Version: string(DefaultMissionVersion),
			},
		),
	}
	service.On("CarbonCopyMission", mock.Anything).Return(missionModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.MissionCopiedEvent{
		OriginalMissionId: DefaultOriginalID,
		NewMissionId:      DefaultNewID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleMissionCopiedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.OriginalID, DefaultOriginalID)
	a.Equal(service.NewID, DefaultNewID)
}
