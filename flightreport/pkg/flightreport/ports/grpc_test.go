package ports

import (
	"flightreport/pkg/flightreport/app"
	"flightreport/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleFlightoperationsListFlightoperations(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	flightoperations := []flightoperationMock{
		{
			id:           string(DefaultFlightoperationID),
			flightplanID: string(DefaultFlightplanID),
		},
	}
	service.On("ListFlightoperations", mock.Anything).Return(flightoperations, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightreports(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightreportsResponses{
		Flightreports: []*skysign_proto.Flightreport{
			{
				Id:                string(DefaultFlightoperationID),
				FlightoperationId: string(DefaultFlightplanID),
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleFlightoperationsListFlightoperations(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightoperationID1 = DefaultFlightoperationID + "-1"
		DefaultFlightoperationID2 = DefaultFlightoperationID + "-2"
		DefaultFlightoperationID3 = DefaultFlightoperationID + "-3"
		DefaultFlightplanID1      = DefaultFlightplanID + "-1"
		DefaultFlightplanID2      = DefaultFlightplanID + "-2"
		DefaultFlightplanID3      = DefaultFlightplanID + "-3"
	)

	service := manageFlightoperationServiceMock{}

	flightoperations := []flightoperationMock{
		{
			id:           string(DefaultFlightoperationID1),
			flightplanID: string(DefaultFlightplanID1),
		},
		{
			id:           string(DefaultFlightoperationID2),
			flightplanID: string(DefaultFlightplanID2),
		},
		{
			id:           string(DefaultFlightoperationID3),
			flightplanID: string(DefaultFlightplanID3),
		},
	}
	service.On("ListFlightoperations", mock.Anything).Return(flightoperations, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightreports(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightreportsResponses{
		Flightreports: []*skysign_proto.Flightreport{
			{
				Id:                string(DefaultFlightoperationID1),
				FlightoperationId: string(DefaultFlightplanID1),
			},
			{
				Id:                string(DefaultFlightoperationID2),
				FlightoperationId: string(DefaultFlightplanID2),
			},
			{
				Id:                string(DefaultFlightoperationID3),
				FlightoperationId: string(DefaultFlightplanID3),
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

	grpc := NewGrpcServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightreports(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightreportsResponses{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestGetFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	flightoperation := flightoperationMock{
		id:           string(DefaultFlightoperationID),
		flightplanID: string(DefaultFlightplanID),
	}
	service.On("GetFlightoperation", mock.Anything, mock.Anything).Return(flightoperation, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetFlightreportRequest{
		Id: string(DefaultFlightoperationID),
	}
	response, err := grpc.GetFlightreport(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightreport{
		Id:                string(DefaultFlightoperationID),
		FlightoperationId: string(DefaultFlightplanID),
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightoperationServiceMock{}

	service.On("CreateFlightoperation", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightoperation: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.CreateFlightreportRequest{
		FlightplanId: string(DefaultFlightplanID),
	}
	response, err := grpc.CreateFlightreport(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
