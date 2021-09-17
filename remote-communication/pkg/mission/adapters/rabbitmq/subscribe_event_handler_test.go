package rabbitmq

import (
	"remote-communication/pkg/mission/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestSubscribeEventHandleCopiedMissionCreatedEvent .
func TestSubscribeEventHandleCopiedMissionCreatedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageMissionServiceMock{}
	service.On("CreateMission", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageMission: &service,
		},
	}

	psm := &publishHandlerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.CopiedMissionCreatedEvent{
		MissionId: DefaultMissionID,
		Mission: &skysign_proto.Mission{
			Id: DefaultMissionID,
			Navigation: &skysign_proto.Navigation{
				UploadId:  DefaultMissionUploadID,
				Waypoints: []*skysign_proto.Waypoint{},
			},
		},
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "mission.copied_mission_created_event"
		QueueName    = "uploadmission.copied_mission_created_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.ID, DefaultMissionUploadID)
}
