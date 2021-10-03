package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
)

// TestSubscribeEventHandlerVehicleCopiedEvent .
func TestSubscribeEventHandlerVehicleCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetID    = NewFleetID()
		DefaultOriginalID = NewVehicleID()
		DefaultNewID      = NewVehicleID()
	)

	service := manageVehicleServiceMock{}
	service.On("CarbonCopyVehicle", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	psm := &pubSubManagerMock{}
	SubscribeEventHandler(nil, psm, app)

	requestPb := &skysign_proto.VehicleCopiedEvent{
		FleetId:           DefaultFleetID,
		OriginalVehicleId: DefaultOriginalID,
		NewVehicleId:      DefaultNewID,
	}
	requestBin, _ := proto.Marshal(requestPb)

	var (
		ExchangeName = "fleet.vehicle_copied_event"
		QueueName    = "vehicle.vehicle_copied_event"
	)

	for _, c := range psm.consumers {
		if c.exchangeName == ExchangeName && c.queueName == QueueName {
			c.handler(requestBin)
		}
	}

	a.Equal(service.FleetID, DefaultFleetID)
	a.Equal(service.OriginalID, DefaultOriginalID)
	a.Equal(service.NewID, DefaultNewID)
}
