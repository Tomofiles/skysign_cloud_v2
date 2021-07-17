package ports

import (
	"fleet-formation/pkg/fleet/service"

	"github.com/stretchr/testify/mock"
)

const DefaultFleetNumberOfVehicles = 3
const DefaultFleetID = "fleet-id"
const DefaultFleetAssignmentID = "assignment-id"
const DefaultFleetEventID = "event-id"
const DefaultFleetVehicleID = "vehicle-id"
const DefaultFleetMissionID = "mission-id"

type manageFleetServiceMock struct {
	mock.Mock
	ID               string
	NumberOfVehicles int
	OriginalID       string
	NewID            string
}

func (s *manageFleetServiceMock) CreateFleet(
	command service.CreateFleetCommand,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	s.NumberOfVehicles = command.GetNumberOfVehicles()
	return ret.Error(0)
}

func (s *manageFleetServiceMock) DeleteFleet(
	command service.DeleteFleetCommand,
) error {
	ret := s.Called()
	s.ID = command.GetID()
	return ret.Error(0)
}

func (s *manageFleetServiceMock) CarbonCopyFleet(
	command service.CarbonCopyFleetCommand,
) error {
	ret := s.Called()
	s.OriginalID = command.GetOriginalID()
	s.NewID = command.GetNewID()
	return ret.Error(0)
}

type assignFleetServiceMock struct {
	mock.Mock
}

func (s *assignFleetServiceMock) GetAssignments(
	command service.GetAssignmentsCommand,
	model service.AssignmentRetrievedModel,
) error {
	ret := s.Called()
	if assignments := ret.Get(0); assignments != nil {
		for _, a := range assignments.([]assignmentMock) {
			model(
				&assignmentModelMock{
					&assignmentMock{
						ID:           a.GetID(),
						EventID:      a.GetEventID(),
						AssignmentID: a.GetAssignmentID(),
						VehicleID:    a.GetVehicleID(),
						MissionID:    a.GetMissionID(),
					},
				},
			)
		}
	}
	return ret.Error(1)
}

func (s *assignFleetServiceMock) UpdateAssignment(
	command service.UpdateAssignmentCommand,
	model service.AssignmentRetrievedModel,
) error {
	ret := s.Called()
	model(
		&assignmentModelMock{
			&assignmentMock{
				ID:           command.GetID(),
				EventID:      command.GetEventID(),
				AssignmentID: command.GetAssignmentID(),
				VehicleID:    command.GetVehicleID(),
				MissionID:    command.GetMissionID(),
			},
		},
	)
	return ret.Error(0)
}

type assignmentModelMock struct {
	assignment *assignmentMock
}

func (f *assignmentModelMock) GetAssignment() service.Assignment {
	return f.assignment
}

type assignmentMock struct {
	ID           string
	EventID      string
	AssignmentID string
	VehicleID    string
	MissionID    string
}

func (a *assignmentMock) GetID() string {
	return a.ID
}

func (a *assignmentMock) GetEventID() string {
	return a.EventID
}

func (a *assignmentMock) GetAssignmentID() string {
	return a.AssignmentID
}

func (a *assignmentMock) GetVehicleID() string {
	return a.VehicleID
}

func (a *assignmentMock) GetMissionID() string {
	return a.MissionID
}
