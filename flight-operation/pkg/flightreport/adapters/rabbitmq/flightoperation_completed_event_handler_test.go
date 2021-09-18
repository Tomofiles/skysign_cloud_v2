package rabbitmq

import (
	"flight-operation/pkg/flightreport/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

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

	handler := NewFlightoperationCompletedEventHandler(app)

	requestPb := &skysign_proto.FlightoperationCompletedEvent{
		FlightoperationId: string(DefaultFlightreportID),
		Flightoperation: &skysign_proto.Flightoperation{
			Id:          string(DefaultFlightreportID),
			Name:        DefaultFlightreportName,
			Description: DefaultFlightreportDescription,
			FleetId:     string(DefaultFlightreportFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFlightoperationCompletedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.name, DefaultFlightreportName)
	a.Equal(service.description, DefaultFlightreportDescription)
	a.Equal(service.fleetID, string(DefaultFlightreportFleetID))
}
