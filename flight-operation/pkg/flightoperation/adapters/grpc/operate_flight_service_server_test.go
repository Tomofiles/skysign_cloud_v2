package grpc

import (
	"flight-operation/pkg/flightoperation/app"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleFlightoperationsListFlightoperations(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	flightoperations := []flightoperationMock{
		{
			id:          string(DefaultID),
			name:        DefaultName,
			description: DefaultDescription,
			fleetID:     string(DefaultFleetID),
		},
	}
	service.On("ListFlightoperations", mock.Anything).Return(flightoperations, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewOperateFlightServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightoperations(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightoperationsResponses{
		Flightoperations: []*skysign_proto.Flightoperation{
			{
				Id:          string(DefaultID),
				Name:        DefaultName,
				Description: DefaultDescription,
				FleetId:     string(DefaultFleetID),
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleFlightoperationsListFlightoperations(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultID1          = DefaultID + "-1"
		DefaultID2          = DefaultID + "-2"
		DefaultID3          = DefaultID + "-3"
		DefaultName1        = DefaultName + "-1"
		DefaultName2        = DefaultName + "-2"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID1     = DefaultFleetID + "-1"
		DefaultFleetID2     = DefaultFleetID + "-2"
		DefaultFleetID3     = DefaultFleetID + "-3"
	)

	service := manageFlightoperationServiceMock{}

	flightoperations := []flightoperationMock{
		{
			id:          string(DefaultID1),
			name:        DefaultName1,
			description: DefaultDescription1,
			fleetID:     string(DefaultFleetID1),
		},
		{
			id:          string(DefaultID2),
			name:        DefaultName2,
			description: DefaultDescription2,
			fleetID:     string(DefaultFleetID2),
		},
		{
			id:          string(DefaultID3),
			name:        DefaultName3,
			description: DefaultDescription3,
			fleetID:     string(DefaultFleetID3),
		},
	}
	service.On("ListFlightoperations", mock.Anything).Return(flightoperations, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewOperateFlightServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightoperations(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightoperationsResponses{
		Flightoperations: []*skysign_proto.Flightoperation{
			{
				Id:          string(DefaultID1),
				Name:        DefaultName1,
				Description: DefaultDescription1,
				FleetId:     string(DefaultFleetID1),
			},
			{
				Id:          string(DefaultID2),
				Name:        DefaultName2,
				Description: DefaultDescription2,
				FleetId:     string(DefaultFleetID2),
			},
			{
				Id:          string(DefaultID3),
				Name:        DefaultName3,
				Description: DefaultDescription3,
				FleetId:     string(DefaultFleetID3),
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneFlightoperationsListFlightoperations(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	flightoperations := []flightoperationMock{}
	service.On("ListFlightoperations", mock.Anything).Return(flightoperations, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewOperateFlightServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightoperations(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightoperationsResponses{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestGetFlightoperation(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	flightoperation := flightoperationMock{
		id:          string(DefaultID),
		name:        DefaultName,
		description: DefaultDescription,
		fleetID:     string(DefaultFleetID),
	}
	service.On("GetFlightoperation", mock.Anything, mock.Anything).Return(flightoperation, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewOperateFlightServiceServer(app)

	request := &skysign_proto.GetFlightoperationRequest{
		Id: string(DefaultID),
	}
	response, err := grpc.GetFlightoperation(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightoperation{
		Id:          string(DefaultID),
		Name:        DefaultName,
		Description: DefaultDescription,
		FleetId:     string(DefaultFleetID),
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCompleteFlightoperation(t *testing.T) {
	a := assert.New(t)

	service := operateFlightoperationServiceMock{}

	service.On("CompleteFlightoperation", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			OperateFlightoperation: &service,
		},
	}

	grpc := NewOperateFlightServiceServer(app)

	request := &skysign_proto.CompleteFlightoperationRequest{
		Id: string(DefaultID),
	}
	response, err := grpc.CompleteFlightoperation(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
