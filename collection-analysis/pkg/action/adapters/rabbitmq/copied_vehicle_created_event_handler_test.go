package rabbitmq

import (
	"collection-analysis/pkg/action/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestHandleCopiedVehicleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageActionServiceMock{}

	service.On("CreateAction", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageAction: &service,
		},
	}

	handler := NewCopiedVehicleCreatedEventHandler(app)

	requestPb := &skysign_proto.CopiedVehicleCreatedEvent{
		VehicleId:       string(DefaultActionID),
		CommunicationId: string(DefaultActionCommunicationID),
		FleetId:         string(DefaultActionFleetID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleCopiedVehicleCreatedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.GetVehicleId(), service.command.GetID())
	a.Equal(requestPb.GetCommunicationId(), service.command.GetCommunicationID())
	a.Equal(requestPb.GetFleetId(), service.command.GetFleetID())
}
