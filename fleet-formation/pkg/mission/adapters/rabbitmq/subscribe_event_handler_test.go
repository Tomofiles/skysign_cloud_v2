package rabbitmq

import (
	"fleet-formation/pkg/mission/app"
	m "fleet-formation/pkg/mission/domain/mission"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

// TestSubscribeEventHandlerMissionCopiedEvent .
func TestSubscribeEventHandlerMissionCopiedEvent(t *testing.T) {
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
				ID: string(DefaultMissionID),
			},
		),
	}
	service.On("CarbonCopyMission", mock.Anything).Return(missionModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.MissionCopiedEvent{
		FleetId:           string(DefaultMissionID),
		OriginalMissionId: string(DefaultOriginalID),
		NewMissionId:      string(DefaultNewID),
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "fleet.mission_copied_event"
		QueueName    = "mission.mission_copied_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.OriginalID, string(DefaultOriginalID))
	a.Equal(service.NewID, string(DefaultNewID))
}
