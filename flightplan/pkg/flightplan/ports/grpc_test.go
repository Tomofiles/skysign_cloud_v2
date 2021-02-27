package ports

import (
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSingleFlightplansListFlightplans(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplans := []flightplanMock{
		{
			id:          DefaultFlightplanID,
			name:        DefaultFlightplanName,
			description: DefaultFlightplanDescription,
		},
	}
	service.On("ListFlightplans", mock.Anything).Return(flightplans, nil)

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
		DefaultFlightplanID2          = DefaultFlightplanID + "-2"
		DefaultFlightplanName2        = DefaultFlightplanName + "-2"
		DefaultFlightplanDescription2 = DefaultFlightplanDescription + "-2"
		DefaultFlightplanID3          = DefaultFlightplanID + "-3"
		DefaultFlightplanName3        = DefaultFlightplanName + "-3"
		DefaultFlightplanDescription3 = DefaultFlightplanDescription + "-3"
	)

	service := manageFlightplanServiceMock{}

	flightplans := []flightplanMock{
		{
			id:          DefaultFlightplanID1,
			name:        DefaultFlightplanName1,
			description: DefaultFlightplanDescription1,
		},
		{
			id:          DefaultFlightplanID2,
			name:        DefaultFlightplanName2,
			description: DefaultFlightplanDescription2,
		},
		{
			id:          DefaultFlightplanID3,
			name:        DefaultFlightplanName3,
			description: DefaultFlightplanDescription3,
		},
	}
	service.On("ListFlightplans", mock.Anything).Return(flightplans, nil)

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
			},
			{
				Id:          DefaultFlightplanID2,
				Name:        DefaultFlightplanName2,
				Description: DefaultFlightplanDescription2,
			},
			{
				Id:          DefaultFlightplanID3,
				Name:        DefaultFlightplanName3,
				Description: DefaultFlightplanDescription3,
			},
		},
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestNoneFlightplansListFlightplans(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplans := []flightplanMock{}
	service.On("ListFlightplans", mock.Anything).Return(flightplans, nil)

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

	flightplan := flightplanMock{
		id:          DefaultFlightplanID,
		name:        DefaultFlightplanName,
		description: DefaultFlightplanDescription,
	}
	service.On("GetFlightplan", mock.Anything, mock.Anything).Return(flightplan, nil)

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
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestCreateFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplan := flightplanMock{
		id:          DefaultFlightplanID,
		name:        DefaultFlightplanName,
		description: DefaultFlightplanDescription,
	}
	service.On("CreateFlightplan", mock.Anything, mock.Anything).Return(flightplan, nil)

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
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}

func TestUpdateFlightplan(t *testing.T) {
	a := assert.New(t)

	service := manageFlightplanServiceMock{}

	flightplan := flightplanMock{
		id:          DefaultFlightplanID,
		name:        DefaultFlightplanName,
		description: DefaultFlightplanDescription,
	}
	service.On("UpdateFlightplan", mock.Anything, mock.Anything).Return(flightplan, nil)

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

	service := assignFleetServiceMock{}

	changeNumberOfVehicles := changeNumberOfVehiclesMock{
		flightplanID:     DefaultFlightplanID,
		numberOfVehicles: DefaultFleetNumberOfVehicles,
	}

	service.On("ChangeNumberOfVehicles", mock.Anything, mock.Anything).Return(changeNumberOfVehicles, nil)

	app := app.Application{
		Services: app.Services{
			AssignFleet: &service,
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
			id:           DefaultFleetEventID,
			assignmentID: DefaultFleetAssignmentID,
			vehicleID:    DefaultFleetVehicleID,
			missionID:    DefaultFleetMissionID,
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
		Id: DefaultFlightplanID,
	}
	response, err := grpc.GetAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetAssignmentsResponse{
		Id: DefaultFlightplanID,
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
			id:           DefaultFleetEventID1,
			assignmentID: DefaultFleetAssignmentID1,
			vehicleID:    DefaultFleetVehicleID1,
			missionID:    DefaultFleetMissionID1,
		},
		{
			id:           DefaultFleetEventID2,
			assignmentID: DefaultFleetAssignmentID2,
			vehicleID:    DefaultFleetVehicleID2,
			missionID:    DefaultFleetMissionID2,
		},
		{
			id:           DefaultFleetEventID3,
			assignmentID: DefaultFleetAssignmentID3,
			vehicleID:    DefaultFleetVehicleID3,
			missionID:    DefaultFleetMissionID3,
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
		Id: DefaultFlightplanID,
	}
	response, err := grpc.GetAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetAssignmentsResponse{
		Id: DefaultFlightplanID,
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
		Id: DefaultFlightplanID,
	}
	response, err := grpc.GetAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.GetAssignmentsResponse{
		Id: DefaultFlightplanID,
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
		Id: DefaultFlightplanID,
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
		Id: DefaultFlightplanID,
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
		Id: DefaultFlightplanID,
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
		Id: DefaultFlightplanID,
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
		Id:          DefaultFlightplanID,
		Assignments: []*skysign_proto.Assignment{},
	}
	response, err := grpc.UpdateAssignments(
		nil,
		request,
	)

	expectResponse := &skysign_proto.UpdateAssignmentsResponse{
		Id: DefaultFlightplanID,
	}

	a.Nil(err)
	a.Equal(response, expectResponse)
}
