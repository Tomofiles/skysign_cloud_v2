package rabbitmq

import (
	"flight-operation/pkg/flightoperation/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

// TestSubscribeEventHandlerFlightplanExecutedEvent .
func TestSubscribeEventHandlerFlightplanExecutedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}
	service.On("CreateFlightoperation", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.FlightplanExecutedEvent{
		Flightplan: &skysign_proto.Flightplan{
			Name:        DefaultFlightoperationName,
			Description: DefaultFlightoperationDescription,
			FleetId:     string(DefaultFlightoperationFleetID),
		},
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "flightplan.flightplan_executed_event"
		QueueName    = "flightoperation.flightplan_executed_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.name, DefaultFlightoperationName)
	a.Equal(service.description, DefaultFlightoperationDescription)
	a.Equal(service.fleetID, string(DefaultFlightoperationFleetID))
}
