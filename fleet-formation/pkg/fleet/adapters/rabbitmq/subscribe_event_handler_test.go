package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

// TestSubscribeEventHandlerFleetIdGaveEvent .
func TestSubscribeEventHandlerFleetIdGaveEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFleetServiceMock{}
	service.On("CreateFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.FleetIdGaveEvent{
		FleetId:          string(DefaultFleetID),
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "flightplan.fleet_id_gave_event"
		QueueName    = "fleet.fleet_id_gave_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.ID, string(DefaultFleetID))
	a.Equal(service.NumberOfVehicles, DefaultFleetNumberOfVehicles)
}

// TestSubscribeEventHandlerFleetIdRemovedEvent .
func TestSubscribeEventHandlerFleetIdRemovedEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFleetServiceMock{}
	service.On("DeleteFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.FleetIdRemovedEvent{
		FleetId: string(DefaultFleetID),
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "flightplan.fleet_id_removed_event"
		QueueName    = "fleet.fleet_id_removed_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.ID, string(DefaultFleetID))
}

// TestSubscribeEventHandlerFleetCopiedEvent .
func TestSubscribeEventHandlerFleetCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetOriginalID = DefaultFleetID + "-new"
		DefaultFleetNewID      = DefaultFleetID + "-new"
	)

	service := manageFleetServiceMock{}
	service.On("CarbonCopyFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.FleetCopiedEvent{
		OriginalFleetId: string(DefaultFleetOriginalID),
		NewFleetId:      string(DefaultFleetNewID),
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "flightoperation.fleet_copied_event"
		QueueName    = "fleet.fleet_copied_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.OriginalID, string(DefaultFleetOriginalID))
	a.Equal(service.NewID, string(DefaultFleetNewID))
}
