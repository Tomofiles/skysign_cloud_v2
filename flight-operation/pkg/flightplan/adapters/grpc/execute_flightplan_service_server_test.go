package grpc

import (
	"flight-operation/pkg/flightplan/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecuteFlightplan(t *testing.T) {
	a := assert.New(t)

	service := executeFlightplanServiceMock{}

	service.On("ExecuteFlightplan", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ExecuteFlightplan: &service,
		},
	}

	grpc := NewExecuteFlightplanServiceServer(app)

	request := &skysign_proto.ExecuteFlightplanRequest{
		Id: DefaultFlightplanID,
	}
	response, err := grpc.ExecuteFlightplan(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ExecuteFlightplanResponse{
		Id: DefaultFlightplanID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
