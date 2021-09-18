package grpc

import (
	"flight-operation/pkg/flightplan/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	service := changeFlightplanServiceMock{}

	service.On("ChangeNumberOfVehicles", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ChangeFlightplan: &service,
		},
	}

	grpc := NewChangeFlightplanServiceServer(app)

	request := &skysign_proto.ChangeNumberOfVehiclesRequest{
		Id:               DefaultFlightplanID,
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}
	response, err := grpc.ChangeNumberOfVehicles(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ChangeNumberOfVehiclesResponse{
		Id:               DefaultFlightplanID,
		NumberOfVehicles: DefaultFleetNumberOfVehicles,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
