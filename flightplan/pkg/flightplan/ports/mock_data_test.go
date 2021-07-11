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
const DefaultFleetNumberOfVehicles int32 = 3
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
	ID         string
	OriginalID string
	NewID      string
}

func (s *manageFleetServiceMock) CreateFleet(
	requestDpo service.CreateFleetRequestDpo,
) error {
	ret := s.Called()
	s.ID = requestDpo.GetFlightplanID()
	return ret.Error(0)
}

func (s *manageFleetServiceMock) DeleteFleet(
	requestDpo service.DeleteFleetRequestDpo,
) error {
	ret := s.Called()
	s.ID = requestDpo.GetFlightplanID()
	return ret.Error(0)
}

func (s *manageFleetServiceMock) CarbonCopyFleet(
	requestDpo service.CarbonCopyFleetRequestDpo,
) error {
	ret := s.Called()
	s.OriginalID = requestDpo.GetOriginalID()
	s.NewID = requestDpo.GetNewID()
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

func (s *assignFleetServiceMock) ChangeNumberOfVehicles(
	requestDpo service.ChangeNumberOfVehiclesRequestDpo,
	responseDpo service.ChangeNumberOfVehiclesResponseDpo,
) error {
	ret := s.Called()
	if changeNumberOfVehicles := ret.Get(0); changeNumberOfVehicles != nil {
		f := changeNumberOfVehicles.(changeNumberOfVehiclesMock)
		responseDpo(f.flightplanID, f.numberOfVehicles)
	}
	return ret.Error(1)
}

func (s *assignFleetServiceMock) GetAssignments(
	requestDpo service.GetAssignmentsRequestDpo,
	responseEachDpo service.GetAssignmentsResponseDpo,
) error {
	ret := s.Called()
	if assignments := ret.Get(0); assignments != nil {
		for _, a := range assignments.([]assignmentMock) {
			responseEachDpo(a.id, a.assignmentID, a.vehicleID, a.missionID)
		}
	}
	return ret.Error(1)
}

func (s *assignFleetServiceMock) UpdateAssignment(
	requestDpo service.UpdateAssignmentRequestDpo,
	responseDpo service.UpdateAssignmentResponseDpo,
) error {
	ret := s.Called()
	responseDpo(
		requestDpo.GetEventID(),
		requestDpo.GetAssignmentID(),
		requestDpo.GetVehicleID(),
		requestDpo.GetMissionID(),
	)
	return ret.Error(0)
}

// type flightplanMock struct {
// 	id, name, description, fleetID string
// }

type changeNumberOfVehiclesMock struct {
	flightplanID     string
	numberOfVehicles int32
}

type assignmentMock struct {
	id, assignmentID, vehicleID, missionID string
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
