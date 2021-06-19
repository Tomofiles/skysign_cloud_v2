package ports

// import (
// 	"testing"
// 	"vehicle/pkg/skysign_proto"
// 	"vehicle/pkg/vehicle/app"

// 	"github.com/golang/protobuf/proto"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestHandleVehicleCopiedWhenFlightplanCopiedEvent(t *testing.T) {
// 	a := assert.New(t)

// 	var (
// 		DefaultOriginalID = DefaultVehicleID + "-original"
// 		DefaultNewID      = DefaultVehicleID + "-new"
// 	)

// 	service := manageVehicleServiceMock{}

// 	service.On("CarbonCopyVehicle", mock.Anything).Return(nil)

// 	app := app.Application{
// 		Services: app.Services{
// 			ManageVehicle: &service,
// 		},
// 	}

// 	handler := NewEventHandler(app)

// 	requestPb := &skysign_proto.VehicleCopiedWhenFlightplanCopiedEvent{
// 		OriginalVehicleId: DefaultOriginalID,
// 		NewVehicleId:      DefaultNewID,
// 		FlightplanId:      DefaultFlightplanID,
// 	}
// 	requestBin, _ := proto.Marshal(requestPb)
// 	err := handler.HandleVehicleCopiedWhenFlightplanCopiedEvent(
// 		nil,
// 		requestBin,
// 	)

// 	a.Nil(err)
// 	a.Equal(service.OriginalID, DefaultOriginalID)
// 	a.Equal(service.NewID, DefaultNewID)
// 	a.Equal(service.FlightplanID, DefaultFlightplanID)
// }
