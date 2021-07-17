package ports

import (
	"flightplan/pkg/flightplan/app"
	f "flightplan/pkg/flightplan/domain/flightplan"
	s "flightplan/pkg/flightplan/service"
	"flightplan/pkg/skysign_proto"
	"testing"

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

	grpc := NewGrpcServer(app)

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

	grpc := NewGrpcServer(app)

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

	grpc := NewGrpcServer(app)

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

	grpc := NewGrpcServer(app)

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

	grpc := NewGrpcServer(app)

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

	grpc := NewGrpcServer(app)

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

	grpc := NewGrpcServer(app)

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

func TestChangeNumberOfVehicles(t *testing.T) {
	a := assert.New(t)

	service := changeFlightplanServiceMock{}

	service.On("ChangeNumberOfVehicles", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ChangeFlightplan: &service,
		},
	}

	grpc := NewGrpcServer(app)

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

func TestSingleAssignmentsGetAssignments(t *testing.T) {
	a := assert.New(t)

	service := assignFleetServiceMock{}

	assignments := []assignmentMock{
		{
			ID:           DefaultFleetID,
			EventID:      DefaultFleetEventID,
			AssignmentID: DefaultFleetAssignmentID,
			VehicleID:    DefaultFleetVehicleID,
			MissionID:    DefaultFleetMissionID,
		},
	}

	service.On("GetAssignments", mock.Anything, mock.Anything).Return(assignments, nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetAssignmentsRequest{
		Id: DefaultFleetID,
	}
	response, err := grpc.GetAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetAssignmentsResponse{
		Id: DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{
			{
				Id:           DefaultFleetEventID,
				AssignmentId: DefaultFleetAssignmentID,
				VehicleId:    DefaultFleetVehicleID,
				MissionId:    DefaultFleetMissionID,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleAssignmentsGetAssignments(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetAssignmentID1 = DefaultFleetAssignmentID + "-1"
		DefaultFleetAssignmentID2 = DefaultFleetAssignmentID + "-2"
		DefaultFleetAssignmentID3 = DefaultFleetAssignmentID + "-3"
		DefaultFleetEventID1      = DefaultFleetEventID + "-1"
		DefaultFleetEventID2      = DefaultFleetEventID + "-2"
		DefaultFleetEventID3      = DefaultFleetEventID + "-3"
		DefaultFleetVehicleID1    = DefaultFleetVehicleID + "-1"
		DefaultFleetVehicleID2    = DefaultFleetVehicleID + "-2"
		DefaultFleetVehicleID3    = DefaultFleetVehicleID + "-3"
		DefaultFleetMissionID1    = DefaultFleetMissionID + "-1"
		DefaultFleetMissionID2    = DefaultFleetMissionID + "-2"
		DefaultFleetMissionID3    = DefaultFleetMissionID + "-3"
	)

	service := assignFleetServiceMock{}

	assignments := []assignmentMock{
		{
			ID:           DefaultFleetID,
			EventID:      DefaultFleetEventID1,
			AssignmentID: DefaultFleetAssignmentID1,
			VehicleID:    DefaultFleetVehicleID1,
			MissionID:    DefaultFleetMissionID1,
		},
		{
			ID:           DefaultFleetID,
			EventID:      DefaultFleetEventID2,
			AssignmentID: DefaultFleetAssignmentID2,
			VehicleID:    DefaultFleetVehicleID2,
			MissionID:    DefaultFleetMissionID2,
		},
		{
			ID:           DefaultFleetID,
			EventID:      DefaultFleetEventID3,
			AssignmentID: DefaultFleetAssignmentID3,
			VehicleID:    DefaultFleetVehicleID3,
			MissionID:    DefaultFleetMissionID3,
		},
	}

	service.On("GetAssignments", mock.Anything, mock.Anything).Return(assignments, nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetAssignmentsRequest{
		Id: DefaultFleetID,
	}
	response, err := grpc.GetAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetAssignmentsResponse{
		Id: DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{
			{
				Id:           DefaultFleetEventID1,
				AssignmentId: DefaultFleetAssignmentID1,
				VehicleId:    DefaultFleetVehicleID1,
				MissionId:    DefaultFleetMissionID1,
			},
			{
				Id:           DefaultFleetEventID2,
				AssignmentId: DefaultFleetAssignmentID2,
				VehicleId:    DefaultFleetVehicleID2,
				MissionId:    DefaultFleetMissionID2,
			},
			{
				Id:           DefaultFleetEventID3,
				AssignmentId: DefaultFleetAssignmentID3,
				VehicleId:    DefaultFleetVehicleID3,
				MissionId:    DefaultFleetMissionID3,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneAssignmentsGetAssignments(t *testing.T) {
	a := assert.New(t)

	service := assignFleetServiceMock{}

	assignments := []assignmentMock{}

	service.On("GetAssignments", mock.Anything, mock.Anything).Return(assignments, nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.GetAssignmentsRequest{
		Id: DefaultFleetID,
	}
	response, err := grpc.GetAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetAssignmentsResponse{
		Id: DefaultFleetID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestSingleAssignmentsUpdateAssignments(t *testing.T) {
	a := assert.New(t)

	service := assignFleetServiceMock{}

	service.On("UpdateAssignment", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.UpdateAssignmentsRequest{
		Id: DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{
			{
				Id:           DefaultFleetEventID,
				AssignmentId: DefaultFleetAssignmentID,
				VehicleId:    DefaultFleetVehicleID,
				MissionId:    DefaultFleetMissionID,
			},
		},
	}
	response, err := grpc.UpdateAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.UpdateAssignmentsResponse{
		Id: DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{
			{
				Id:           DefaultFleetEventID,
				AssignmentId: DefaultFleetAssignmentID,
				VehicleId:    DefaultFleetVehicleID,
				MissionId:    DefaultFleetMissionID,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestMultipleAssignmentsUpdateAssignments(t *testing.T) {
	a := assert.New(t)

	var (
		DefaultFleetEventID1      = DefaultFleetEventID + "-1"
		DefaultFleetAssignmentID1 = DefaultFleetAssignmentID + "-1"
		DefaultFleetVehicleID1    = DefaultFleetVehicleID + "-1"
		DefaultFleetMissionID1    = DefaultFleetMissionID + "-1"

		DefaultFleetEventID2      = DefaultFleetEventID + "-2"
		DefaultFleetAssignmentID2 = DefaultFleetAssignmentID + "-2"
		DefaultFleetVehicleID2    = DefaultFleetVehicleID + "-2"
		DefaultFleetMissionID2    = DefaultFleetMissionID + "-2"

		DefaultFleetEventID3      = DefaultFleetEventID + "-3"
		DefaultFleetAssignmentID3 = DefaultFleetAssignmentID + "-3"
		DefaultFleetVehicleID3    = DefaultFleetVehicleID + "-3"
		DefaultFleetMissionID3    = DefaultFleetMissionID + "-3"
	)

	service := assignFleetServiceMock{}

	service.On("UpdateAssignment", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.UpdateAssignmentsRequest{
		Id: DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{
			{
				Id:           DefaultFleetEventID1,
				AssignmentId: DefaultFleetAssignmentID1,
				VehicleId:    DefaultFleetVehicleID1,
				MissionId:    DefaultFleetMissionID1,
			},
			{
				Id:           DefaultFleetEventID2,
				AssignmentId: DefaultFleetAssignmentID2,
				VehicleId:    DefaultFleetVehicleID2,
				MissionId:    DefaultFleetMissionID2,
			},
			{
				Id:           DefaultFleetEventID3,
				AssignmentId: DefaultFleetAssignmentID3,
				VehicleId:    DefaultFleetVehicleID3,
				MissionId:    DefaultFleetMissionID3,
			},
		},
	}
	response, err := grpc.UpdateAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.UpdateAssignmentsResponse{
		Id: DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{
			{
				Id:           DefaultFleetEventID1,
				AssignmentId: DefaultFleetAssignmentID1,
				VehicleId:    DefaultFleetVehicleID1,
				MissionId:    DefaultFleetMissionID1,
			},
			{
				Id:           DefaultFleetEventID2,
				AssignmentId: DefaultFleetAssignmentID2,
				VehicleId:    DefaultFleetVehicleID2,
				MissionId:    DefaultFleetMissionID2,
			},
			{
				Id:           DefaultFleetEventID3,
				AssignmentId: DefaultFleetAssignmentID3,
				VehicleId:    DefaultFleetVehicleID3,
				MissionId:    DefaultFleetMissionID3,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneAssignmentsUpdateAssignments(t *testing.T) {
	a := assert.New(t)

	service := assignFleetServiceMock{}

	service.On("UpdateAssignment", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
		},
	}

	grpc := NewGrpcServer(app)

	request := &skysign_proto.UpdateAssignmentsRequest{
		Id:          DefaultFleetID,
		Assignments: []*skysign_proto.Assignment{},
	}
	response, err := grpc.UpdateAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.UpdateAssignmentsResponse{
		Id: DefaultFleetID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestExecuteFlightplan(t *testing.T) {
	a := assert.New(t)

	service := executeFlightplanServiceMock{}

	service.On("ExecuteFlightplan", mock.Anything, mock.Anything).Return(nil)

	app := app.Application{
		Services: app.Services{
			ExecuteFlightplan: &service,
		},
	}

	grpc := NewGrpcServer(app)

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
