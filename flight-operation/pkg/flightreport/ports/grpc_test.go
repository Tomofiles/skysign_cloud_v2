package ports

import (
	"flight-operation/pkg/flightreport/app"
	frep "flight-operation/pkg/flightreport/domain/flightreport"
	s "flight-operation/pkg/flightreport/service"
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	flightreportModels := []s.FlightreportPresentationModel{
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultID),
					Name:        DefaultName,
					Description: DefaultDescription,
					FleetID:     string(DefaultFleetID),
				},
			),
		},
	}
	service.On("ListFlightreports", mock.Anything).Return(flightreportModels, nil)

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

func TestMultipleFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultID1          = string(DefaultID) + "-1"
		DefaultName1        = DefaultName + "-1"
		DefaultDescription1 = DefaultDescription + "-1"
		DefaultFleetID1     = string(DefaultFleetID) + "-1"
		DefaultID2          = string(DefaultID) + "-2"
		DefaultName2        = DefaultName + "-2"
		DefaultDescription2 = DefaultDescription + "-2"
		DefaultFleetID2     = string(DefaultFleetID) + "-2"
		DefaultID3          = string(DefaultID) + "-3"
		DefaultName3        = DefaultName + "-3"
		DefaultDescription3 = DefaultDescription + "-3"
		DefaultFleetID3     = string(DefaultFleetID) + "-3"
	)

	service := manageFlightreportServiceMock{}

	flightreportModels := []s.FlightreportPresentationModel{
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultID1),
					Name:        DefaultName1,
					Description: DefaultDescription1,
					FleetID:     string(DefaultFleetID1),
				},
			),
		},
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultID2),
					Name:        DefaultName2,
					Description: DefaultDescription2,
					FleetID:     string(DefaultFleetID2),
				},
			),
		},
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultID3),
					Name:        DefaultName3,
					Description: DefaultDescription3,
					FleetID:     string(DefaultFleetID3),
				},
			),
		},
	}
	service.On("ListFlightreports", mock.Anything).Return(flightreportModels, nil)

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

func TestNoneFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	service := manageFlightreportServiceMock{}

	flightreportModels := []s.FlightreportPresentationModel{}
	service.On("ListFlightreports", mock.Anything).Return(flightreportModels, nil)

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

	flightreportModel := &flightreportModelMock{
		flightreport: frep.AssembleFrom(
			nil,
			&flightreportComponentMock{
				ID:          string(DefaultID),
				Name:        DefaultName,
				Description: DefaultDescription,
				FleetID:     string(DefaultFleetID),
			},
		),
	}
	service.On("GetFlightreport", mock.Anything, mock.Anything).Return(flightreportModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetFlightreportRequest{
		Id: string(DefaultID),
	}
	response, err := grpc.GetFlightreport(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightreport{
		Id:          string(DefaultID),
		Name:        DefaultName,
		Description: DefaultDescription,
		FleetId:     string(DefaultFleetID),
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
