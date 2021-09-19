package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleFleetIDRemovedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFleetServiceMock{}

	service.On("DeleteFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	handler := NewFleetIDRemovedEventHandler(app)

	requestPb := &skysign_proto.FleetIdRemovedEvent{
		FleetId: string(DefaultFleetID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFleetIDRemovedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, string(DefaultFleetID))
}
