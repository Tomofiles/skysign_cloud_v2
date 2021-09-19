package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

func TestHandleFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	service := operateActionServiceMock{}

	service.On("CompleteAction", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateAction: &service,
		},
	}

	handler := NewFlightoperationCompletedEventHandler(app)

	requestPb := &skysign_proto.FlightoperationCompletedEvent{
		Flightoperation: &skysign_proto.Flightoperation{
			FleetId: string(DefaultActionFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightoperationCompletedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(requestPb.Flightoperation.GetFleetId(), service.completeCommand.GetFleetID())
}
