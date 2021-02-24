package ports

import (
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/skysign_proto"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleCreatedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFleetServiceMock{}

	service.On("CreateFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FlightplanCreatedEvent{
		FlightplanId: DefaultFlightplanID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleCreatedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, DefaultFlightplanID)
}

func TestHandleDeletedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFleetServiceMock{}

	service.On("DeleteFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FlightplanDeletedEvent{
		FlightplanId: DefaultFlightplanID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleDeletedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, DefaultFlightplanID)
}
