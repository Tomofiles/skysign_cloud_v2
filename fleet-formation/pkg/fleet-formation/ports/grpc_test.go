package ports

import (
	"fleet-formation/pkg/fleet-formation/app"
	"fleet-formation/pkg/skysign_proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
