package ports

import (
	"fleet-formation/pkg/vehicle/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleVehicleCopiedEvent(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultOriginalID = DefaultVehicleID + "-original"
		DefaultNewID      = DefaultVehicleID + "-new"
	)

	service := manageVehicleServiceMock{}

	service.On("CarbonCopyVehicle", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageVehicle: &service,
		},
	}

	handler := NewEventHandler(app)

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
