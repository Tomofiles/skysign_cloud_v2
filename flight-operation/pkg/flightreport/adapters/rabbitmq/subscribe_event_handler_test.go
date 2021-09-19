package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestSubscribeEventHandlerFlightplanExecutedEvent .
func TestSubscribeEventHandlerFlightplanExecutedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}
	service.On("CreateFlightreport", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

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

	var (
		ExchangeName = "flightoperation.flightoperation_completed_event"
		QueueName    = "flightreport.flightoperation_completed_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.name, DefaultFlightreportName)
	a.Equal(service.description, DefaultFlightreportDescription)
	a.Equal(service.fleetID, string(DefaultFlightreportFleetID))
}
