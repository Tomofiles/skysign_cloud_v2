package grpc

import (
	"testing"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/app"
	f "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/domain/flightplan"
	s "github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightplan/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleFlightplansListFlightplans(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplanModels := []s.FlightplanPresentationModel{
		&flightplanModelMock{
			flightplan: f.AssembleFrom(
				nil,
				&flightplanComponentMock{
					ID:          string(DefaultFlightplanID),
					Name:        DefaultFlightplanName,
					Description: DefaultFlightplanDescription,
					FleetID:     DefaultFlightplanFleetID,
				},
			),
		},
	}
	service.On("ListFlightplans", mock.Anything).Return(flightplanModels, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightplans(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightplansResponses{
		Flightplans: []*skysign_proto.Flightplan{
			{
				Id:          DefaultFlightplanID,
				Name:        DefaultFlightplanName,
				Description: DefaultFlightplanDescription,
				FleetId:     DefaultFlightplanFleetID,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleFlightplansListFlightplans(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFlightplanID1          = DefaultFlightplanID + "-1"
		DefaultFlightplanName1        = DefaultFlightplanName + "-1"
		DefaultFlightplanDescription1 = DefaultFlightplanDescription + "-1"
		DefaultFlightplanFleetID1     = DefaultFlightplanFleetID + "-1"
		DefaultFlightplanID2          = DefaultFlightplanID + "-2"
		DefaultFlightplanName2        = DefaultFlightplanName + "-2"
		DefaultFlightplanDescription2 = DefaultFlightplanDescription + "-2"
		DefaultFlightplanFleetID2     = DefaultFlightplanFleetID + "-2"
		DefaultFlightplanID3          = DefaultFlightplanID + "-3"
		DefaultFlightplanName3        = DefaultFlightplanName + "-3"
		DefaultFlightplanDescription3 = DefaultFlightplanDescription + "-3"
		DefaultFlightplanFleetID3     = DefaultFlightplanFleetID + "-3"
	)

	service := manageFlightplanServiceMock{}

	flightplanModels := []s.FlightplanPresentationModel{
		&flightplanModelMock{
			flightplan: f.AssembleFrom(
				nil,
				&flightplanComponentMock{
					ID:          string(DefaultFlightplanID1),
					Name:        DefaultFlightplanName1,
					Description: DefaultFlightplanDescription1,
					FleetID:     DefaultFlightplanFleetID1,
				},
			),
		},
		&flightplanModelMock{
			flightplan: f.AssembleFrom(
				nil,
				&flightplanComponentMock{
					ID:          string(DefaultFlightplanID2),
					Name:        DefaultFlightplanName2,
					Description: DefaultFlightplanDescription2,
					FleetID:     DefaultFlightplanFleetID2,
				},
			),
		},
		&flightplanModelMock{
			flightplan: f.AssembleFrom(
				nil,
				&flightplanComponentMock{
					ID:          string(DefaultFlightplanID3),
					Name:        DefaultFlightplanName3,
					Description: DefaultFlightplanDescription3,
					FleetID:     DefaultFlightplanFleetID3,
				},
			),
		},
	}
	service.On("ListFlightplans", mock.Anything).Return(flightplanModels, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightplans(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightplansResponses{
		Flightplans: []*skysign_proto.Flightplan{
			{
				Id:          DefaultFlightplanID1,
				Name:        DefaultFlightplanName1,
				Description: DefaultFlightplanDescription1,
				FleetId:     DefaultFlightplanFleetID1,
			},
			{
				Id:          DefaultFlightplanID2,
				Name:        DefaultFlightplanName2,
				Description: DefaultFlightplanDescription2,
				FleetId:     DefaultFlightplanFleetID2,
			},
			{
				Id:          DefaultFlightplanID3,
				Name:        DefaultFlightplanName3,
				Description: DefaultFlightplanDescription3,
				FleetId:     DefaultFlightplanFleetID3,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneFlightplansListFlightplans(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplanModels := []s.FlightplanPresentationModel{}
	service.On("ListFlightplans", mock.Anything).Return(flightplanModels, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.Empty{}
	response, err := grpc.ListFlightplans(
		nil,
		request,
	)

	expectResponse := &skysign_proto.ListFlightplansResponses{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestGetFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplanModel := &flightplanModelMock{
		flightplan: f.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          string(DefaultFlightplanID),
				Name:        DefaultFlightplanName,
				Description: DefaultFlightplanDescription,
				FleetID:     DefaultFlightplanFleetID,
			},
		),
	}
	service.On("GetFlightplan", mock.Anything, mock.Anything).Return(flightplanModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.GetFlightplanRequest{
		Id: DefaultFlightplanID,
	}
	response, err := grpc.GetFlightplan(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightplan{
		Id:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetId:     DefaultFlightplanFleetID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplanModel := &flightplanModelMock{
		flightplan: f.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          string(DefaultFlightplanID),
				Name:        DefaultFlightplanName,
				Description: DefaultFlightplanDescription,
				FleetID:     DefaultFlightplanFleetID,
			},
		),
	}
	service.On("CreateFlightplan", mock.Anything, mock.Anything).Return(flightplanModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.Flightplan{
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
	}
	response, err := grpc.CreateFlightplan(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightplan{
		Id:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetId:     DefaultFlightplanFleetID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestUpdateFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplanModel := &flightplanModelMock{
		flightplan: f.AssembleFrom(
			nil,
			&flightplanComponentMock{
				ID:          string(DefaultFlightplanID),
				Name:        DefaultFlightplanName,
				Description: DefaultFlightplanDescription,
				FleetID:     DefaultFlightplanFleetID,
			},
		),
	}
	service.On("UpdateFlightplan", mock.Anything, mock.Anything).Return(flightplanModel, nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.Flightplan{
		Id:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
	}
	response, err := grpc.UpdateFlightplan(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Flightplan{
		Id:          DefaultFlightplanID,
		Name:        DefaultFlightplanName,
		Description: DefaultFlightplanDescription,
		FleetId:     DefaultFlightplanFleetID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestDeleteFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	service.On("DeleteFlightplan", mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ManageFlightplan: &service,
		},
	}

	grpc := NewManageFlightplanServiceServer(app)

	request := &skysign_proto.DeleteFlightplanRequest{
		Id: DefaultFlightplanID,
	}
	response, err := grpc.DeleteFlightplan(
		nil,
		request,
	)

	expectResponse := &skysign_proto.Empty{}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
