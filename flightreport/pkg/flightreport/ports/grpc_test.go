package ports

import (
	"flightreport/pkg/flightreport/app"
	"flightreport/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	flightreports := []flightreportMock{
		{
			id:                string(DefaultFlightreportID),
			flightoperationID: string(DefaultFlightoperationID),
		},
	}
	service.On("ListFlightreports", mock.Anything).Return(flightreports, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
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
				Id:                string(DefaultFlightreportID),
				FlightoperationId: string(DefaultFlightoperationID),
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightreportID1    = DefaultFlightreportID + "-1"
		DefaultFlightreportID2    = DefaultFlightreportID + "-2"
		DefaultFlightreportID3    = DefaultFlightreportID + "-3"
		DefaultFlightoperationID1 = DefaultFlightoperationID + "-1"
		DefaultFlightoperationID2 = DefaultFlightoperationID + "-2"
		DefaultFlightoperationID3 = DefaultFlightoperationID + "-3"
	)

	service := manageFlightreportServiceMock{}

	flightreports := []flightreportMock{
		{
			id:                string(DefaultFlightreportID1),
			flightoperationID: string(DefaultFlightoperationID1),
		},
		{
			id:                string(DefaultFlightreportID2),
			flightoperationID: string(DefaultFlightoperationID2),
		},
		{
			id:                string(DefaultFlightreportID3),
			flightoperationID: string(DefaultFlightoperationID3),
		},
	}
	service.On("ListFlightreports", mock.Anything).Return(flightreports, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
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
				Id:                string(DefaultFlightreportID1),
				FlightoperationId: string(DefaultFlightoperationID1),
			},
			{
				Id:                string(DefaultFlightreportID2),
				FlightoperationId: string(DefaultFlightoperationID2),
			},
			{
				Id:                string(DefaultFlightreportID3),
				FlightoperationId: string(DefaultFlightoperationID3),
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	flightreports := []flightreportMock{}
	service.On("ListFlightreports", mock.Anything).Return(flightreports, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
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

func TestGetFlightreport(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	flightreport := flightreportMock{
		id:                string(DefaultFlightreportID),
		flightoperationID: string(DefaultFlightoperationID),
	}
	service.On("GetFlightreport", mock.Anything, mock.Anything).Return(flightreport, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetFlightreportRequest{
		Id: string(DefaultFlightreportID),
	}
	response, err := grpc.GetFlightreport(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightreport{
		Id:                string(DefaultFlightreportID),
		FlightoperationId: string(DefaultFlightoperationID),
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateFlightreport(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	service.On("CreateFlightreport", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.CreateFlightreportRequest{
		FlightoperationId: string(DefaultFlightoperationID),
	}
	response, err := grpc.CreateFlightreport(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
