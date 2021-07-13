package ports

import (
	f "flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/service"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightplanID = "flightplan-id"
const DefaultFlightplanName = "flightplan-name"
const DefaultFlightplanDescription = "flightplan-description"
const DefaultFlightplanFleetID = "fleet-id"
const DefaultFleetNumberOfVehicles = 3
const DefaultFleetID = "fleet-id"
const DefaultFleetAssignmentID = "assignment-id"
const DefaultFleetEventID = "event-id"
const DefaultFleetVehicleID = "vehicle-id"
const DefaultFleetMissionID = "mission-id"

type manageFlightplanServiceMock struct {
	mock.Mock
	OriginalID string
	NewID      string
}

func (s *manageFlightplanServiceMock) GetFlightplan(
	command service.GetFlightplanCommand,
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightplanPresentationModel)
		retrievedModel(f)
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) ListFlightplans(
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if models := ret.Get(0); models != nil {
		for _, f := range models.([]service.FlightplanPresentationModel) {
			retrievedModel(f)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) CreateFlightplan(
	command service.CreateFlightplanCommand,
	createdID service.CreatedID,
	fleetID service.FleetID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightplanPresentationModel)
		createdID(f.GetFlightplan().GetID())
		fleetID(f.GetFlightplan().GetFleetID())
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) UpdateFlightplan(
	command service.UpdateFlightplanCommand,
	fleetID service.FleetID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.FlightplanPresentationModel)
		fleetID(f.GetFlightplan().GetFleetID())
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) DeleteFlightplan(
	command service.DeleteFlightplanCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

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

type changeFlightplanServiceMock struct {
	mock.Mock
}

func (s *changeFlightplanServiceMock) ChangeNumberOfVehicles(
	command service.ChangeNumberOfVehiclesCommand,
) error {
	ret := s.Called()
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

type flightplanModelMock struct {
	flightplan *f.Flightplan
}

func (f *flightplanModelMock) GetFlightplan() service.Flightplan {
	return &flightplanMock{
		flightplan: f.flightplan,
	}
}

type flightplanMock struct {
	flightplan *f.Flightplan
}

func (f *flightplanMock) GetID() string {
	return string(f.flightplan.GetID())
}

func (f *flightplanMock) GetName() string {
	return f.flightplan.GetName()
}

func (f *flightplanMock) GetDescription() string {
	return f.flightplan.GetDescription()
}

func (f *flightplanMock) GetFleetID() string {
	return string(f.flightplan.GetFleetID())
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

// Flightplan構成オブジェクトモック
type flightplanComponentMock struct {
	ID          string
	Name        string
	Description string
	FleetID     string
	Version     string
}

func (f *flightplanComponentMock) GetID() string {
	return f.ID
}

func (f *flightplanComponentMock) GetName() string {
	return f.Name
}

func (f *flightplanComponentMock) GetDescription() string {
	return f.Description
}

func (f *flightplanComponentMock) GetFleetID() string {
	return f.FleetID
}

func (f *flightplanComponentMock) GetVersion() string {
	return f.Version
}
