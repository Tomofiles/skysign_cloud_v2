package ports

import (
	"fleet-formation/pkg/fleet/app"
	"fleet-formation/pkg/skysign_proto"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleFleetIDGaveEvent(t *testing.T) {
	a := assert.New(t)

	service := manageFleetServiceMock{}

	service.On("CreateFleet", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFleet: &service,
		},
	}

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FleetIDGaveEvent{
		FleetId:          DefaultFleetID,
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFleetIDGaveEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, DefaultFleetID)
	a.Equal(service.NumberOfVehicles, DefaultFleetNumberOfVehicles)
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

	requestPb := &skysign_proto.FleetIDRemovedEvent{
		FleetId: DefaultFleetID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFleetIDRemovedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.ID, DefaultFleetID)
}

func TestHandleFleetCopiedEvent(t *testing.T) {
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

	handler := NewEventHandler(app)

	requestPb := &skysign_proto.FleetCopiedEvent{
		OriginalFleetId: DefaultFleetOriginalID,
		NewFleetId:      DefaultFleetNewID,
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFleetCopiedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.OriginalID, DefaultFleetOriginalID)
	a.Equal(service.NewID, DefaultFleetNewID)
}
