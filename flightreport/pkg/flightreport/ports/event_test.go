package ports

import (
	"flightreport/pkg/flightreport/app"
	"flightreport/pkg/skysign_proto"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleFlightoperationCompletedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	service.On("CreateFlightreport", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FlightoperationCompletedEvent{
		FlightoperationId: string(DefaultFlightoperationID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightoperationCompletedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, string(DefaultFlightoperationID))
}
