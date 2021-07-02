package ports

import (
	"mission/pkg/mission/app"
	"mission/pkg/skysign_proto"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleMissionCopiedWhenFlightplanCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultMissionID + "-original"
		DefaultNewID      = DefaultMissionID + "-new"
	)

	service := manageMissionServiceMock{}

	service.On("CarbonCopyMission", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.MissionCopiedWhenFlightplanCopiedEvent{
		OriginalMissionId: DefaultOriginalID,
		NewMissionId:      DefaultNewID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleMissionCopiedWhenFlightplanCopiedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.OriginalID, DefaultOriginalID)
	a.Equal(service.NewID, DefaultNewID)
}
