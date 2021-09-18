package grpc

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
					ID:          string(DefaultFlightreportID),
					Name:        DefaultFlightreportName,
					Description: DefaultFlightreportDescription,
					FleetID:     string(DefaultFlightreportFleetID),
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

	grpc := NewReportFlightServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightreports(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightreportsResponses{
		Flightreports: []*skysign_proto.Flightreport{
			{
				Id:          string(DefaultFlightreportID),
				Name:        DefaultFlightreportName,
				Description: DefaultFlightreportDescription,
				FleetId:     string(DefaultFlightreportFleetID),
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleFlightreportsListFlightreports(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightreportID1          = string(DefaultFlightreportID) + "-1"
		DefaultFlightreportName1        = DefaultFlightreportName + "-1"
		DefaultFlightreportDescription1 = DefaultFlightreportDescription + "-1"
		DefaultFlightreportFleetID1     = string(DefaultFlightreportFleetID) + "-1"
		DefaultFlightreportID2          = string(DefaultFlightreportID) + "-2"
		DefaultFlightreportName2        = DefaultFlightreportName + "-2"
		DefaultFlightreportDescription2 = DefaultFlightreportDescription + "-2"
		DefaultFlightreportFleetID2     = string(DefaultFlightreportFleetID) + "-2"
		DefaultFlightreportID3          = string(DefaultFlightreportID) + "-3"
		DefaultFlightreportName3        = DefaultFlightreportName + "-3"
		DefaultFlightreportDescription3 = DefaultFlightreportDescription + "-3"
		DefaultFlightreportFleetID3     = string(DefaultFlightreportFleetID) + "-3"
	)

	service := manageFlightreportServiceMock{}

	flightreportModels := []s.FlightreportPresentationModel{
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultFlightreportID1),
					Name:        DefaultFlightreportName1,
					Description: DefaultFlightreportDescription1,
					FleetID:     string(DefaultFlightreportFleetID1),
				},
			),
		},
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultFlightreportID2),
					Name:        DefaultFlightreportName2,
					Description: DefaultFlightreportDescription2,
					FleetID:     string(DefaultFlightreportFleetID2),
				},
			),
		},
		&flightreportModelMock{
			flightreport: frep.AssembleFrom(
				nil,
				&flightreportComponentMock{
					ID:          string(DefaultFlightreportID3),
					Name:        DefaultFlightreportName3,
					Description: DefaultFlightreportDescription3,
					FleetID:     string(DefaultFlightreportFleetID3),
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

	grpc := NewReportFlightServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightreports(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightreportsResponses{
		Flightreports: []*skysign_proto.Flightreport{
			{
				Id:          string(DefaultFlightreportID1),
				Name:        DefaultFlightreportName1,
				Description: DefaultFlightreportDescription1,
				FleetId:     string(DefaultFlightreportFleetID1),
			},
			{
				Id:          string(DefaultFlightreportID2),
				Name:        DefaultFlightreportName2,
				Description: DefaultFlightreportDescription2,
				FleetId:     string(DefaultFlightreportFleetID2),
			},
			{
				Id:          string(DefaultFlightreportID3),
				Name:        DefaultFlightreportName3,
				Description: DefaultFlightreportDescription3,
				FleetId:     string(DefaultFlightreportFleetID3),
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

	grpc := NewReportFlightServiceServer(app)

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
				ID:          string(DefaultFlightreportID),
				Name:        DefaultFlightreportName,
				Description: DefaultFlightreportDescription,
				FleetID:     string(DefaultFlightreportFleetID),
			},
		),
	}
	service.On("GetFlightreport", mock.Anything, mock.Anything).Return(flightreportModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightreport: &service,
		},
	}

	grpc := NewReportFlightServiceServer(app)

	request := &skysign_proto.GetFlightreportRequest{
		Id: string(DefaultFlightreportID),
	}
	response, err := grpc.GetFlightreport(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightreport{
		Id:          string(DefaultFlightreportID),
		Name:        DefaultFlightreportName,
		Description: DefaultFlightreportDescription,
		FleetId:     string(DefaultFlightreportFleetID),
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
