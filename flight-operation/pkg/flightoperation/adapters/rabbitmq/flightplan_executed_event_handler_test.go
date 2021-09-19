package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightoperation/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleFlightplanExecutedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	service.On("CreateFlightoperation", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	handler := NewFlightplanExecutedEventHandler(app)

	requestPb := &skysign_proto.FlightplanExecutedEvent{
		Flightplan: &skysign_proto.Flightplan{
			Name:        DefaultFlightoperationName,
			Description: DefaultFlightoperationDescription,
			FleetId:     string(DefaultFlightoperationFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightplanExecutedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.name, DefaultFlightoperationName)
	a.Equal(service.description, DefaultFlightoperationDescription)
	a.Equal(service.fleetID, string(DefaultFlightoperationFleetID))
}
