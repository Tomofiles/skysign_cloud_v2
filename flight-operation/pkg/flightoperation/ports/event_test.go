package ports

import (
	"flight-operation/pkg/flightoperation/app"
	"flight-operation/pkg/skysign_proto"
	"testing"

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

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FlightplanExecutedEvent{
		Flightplan: &skysign_proto.Flightplan{
			Name:        DefaultName,
			Description: DefaultDescription,
			FleetId:     string(DefaultFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightplanExecutedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.name, DefaultName)
	a.Equal(service.description, DefaultDescription)
	a.Equal(service.fleetID, string(DefaultFleetID))
}
