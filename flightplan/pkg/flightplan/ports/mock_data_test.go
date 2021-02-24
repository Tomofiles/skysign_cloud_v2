package ports

import (
	"flightplan/pkg/flightplan/service"

	"github.com/stretchr/testify/mock"
)

const DefaultFlightplanID = "flightplan-id"
const DefaultFlightplanName = "flightplan-name"
const DefaultFlightplanDescription = "flightplan-description"
const DefaultFleetNumberOfVehicles int32 = 3
const DefaultFleetAssignmentID = "assignment-id"
const DefaultFleetEventID = "event-id"
const DefaultFleetVehicleID = "vehicle-id"
const DefaultFleetMissionID = "mission-id"

type manageFlightplanServiceMock struct {
	mock.Mock
}

func (s *manageFlightplanServiceMock) GetFlightplan(
	requestDpo service.GetFlightplanRequestDpo,
	responseDpo service.GetFlightplanResponseDpo,
) error {
	ret := s.Called()
	if flightplan := ret.Get(0); flightplan != nil {
		f := flightplan.(flightplanMock)
		responseDpo(f.id, f.name, f.description)
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) ListFlightplans(
	responseEachDpo service.ListFlightplansResponseDpo,
) error {
	ret := s.Called()
	if flightplans := ret.Get(0); flightplans != nil {
		for _, f := range flightplans.([]flightplanMock) {
			responseEachDpo(f.id, f.name, f.description)
		}
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) CreateFlightplan(
	requestDpo service.CreateFlightplanRequestDpo,
	responseDpo service.CreateFlightplanResponseDpo,
) error {
	ret := s.Called()
	if flightplan := ret.Get(0); flightplan != nil {
		f := flightplan.(flightplanMock)
		responseDpo(f.id, f.name, f.description)
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) UpdateFlightplan(
	requestDpo service.UpdateFlightplanRequestDpo,
	responseDpo service.UpdateFlightplanResponseDpo,
) error {
	ret := s.Called()
	if flightplan := ret.Get(0); flightplan != nil {
		f := flightplan.(flightplanMock)
		responseDpo(f.id, f.name, f.description)
	}
	return ret.Error(1)
}

func (s *manageFlightplanServiceMock) DeleteFlightplan(
	requestDpo service.DeleteFlightplanRequestDpo,
) error {
	ret := s.Called()
	return ret.Error(0)
}

type manageFleetServiceMock struct {
	mock.Mock
	ID string
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

type flightplanMock struct {
	id, name, description string
}

type changeNumberOfVehiclesMock struct {
	flightplanID     string
	numberOfVehicles int32
}

type assignmentMock struct {
	id, assignmentID, vehicleID, missionID string
}
