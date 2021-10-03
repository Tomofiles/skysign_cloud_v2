package rabbitmq

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleVehicleCopiedEvent(t *testing.T) {
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

	handler := NewVehicleCopiedEventHandler(app)

	requestPb := &skysign_proto.VehicleCopiedEvent{
		OriginalVehicleId: DefaultOriginalID,
		NewVehicleId:      DefaultNewID,
		FleetId:           DefaultFleetID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleVehicleCopiedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.OriginalID, DefaultOriginalID)
	a.Equal(service.NewID, DefaultNewID)
	a.Equal(service.FleetID, DefaultFleetID)
}
