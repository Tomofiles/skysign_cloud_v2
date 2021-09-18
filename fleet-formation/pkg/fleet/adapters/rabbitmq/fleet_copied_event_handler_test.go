package rabbitmq

import (
	"fleet-formation/pkg/fleet/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

	handler := NewFleetCopiedEventHandler(app)

	requestPb := &skysign_proto.FleetCopiedEvent{
		OriginalFleetId: string(DefaultFleetOriginalID),
		NewFleetId:      string(DefaultFleetNewID),
	}
	requestBin, _ := proto.Marshal(requestPb)
	err := handler.HandleFleetCopiedEvent(
		nil,
		requestBin,
	)

	a.Nil(err)
	a.Equal(service.OriginalID, string(DefaultFleetOriginalID))
	a.Equal(service.NewID, string(DefaultFleetNewID))
}
