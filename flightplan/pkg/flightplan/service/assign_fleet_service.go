package service

import (
	"context"
	"errors"
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/event"
)

// AssignFleetService .
type AssignFleetService struct {
	gen  fleet.Generator
	repo fleet.Repository
	pub  event.Publisher
}

// NewAssignFleetService .
func NewAssignFleetService(
	gen fleet.Generator,
	repo fleet.Repository,
	pub event.Publisher,
) AssignFleetService {
	return AssignFleetService{
		gen:  gen,
		repo: repo,
		pub:  pub,
	}
}

// ChangeNumberOfVehicles .
func (s *AssignFleetService) ChangeNumberOfVehicles(
	requestDpo ChangeNumberOfVehiclesRequestDpo,
	responseDpo ChangeNumberOfVehiclesResponseDpo,
) error {
	ctx := context.Background()

	oldFleet, err := s.repo.GetByFlightplanID(
		ctx,
		flightplan.ID(requestDpo.GetId()),
	)
	if err != nil {
		return err
	}
	if oldFleet == nil {
		return errors.New("fleet not found")
	}

	err = s.repo.DeleteByFlightplanID(
		ctx,
		flightplan.ID(requestDpo.GetId()),
	)
	if err != nil {
		return err
	}

	newFleet := fleet.NewInstance(
		s.gen,
		flightplan.ID(requestDpo.GetId()),
		requestDpo.GetNumberOfVehicles())
	for _, assignmentID := range newFleet.GetAllAssignmentID() {
		newFleet.AddNewEvent(assignmentID)
	}
	ret := s.repo.Save(ctx, newFleet)
	if ret != nil {
		return ret
	}

	responseDpo(requestDpo.GetId(), requestDpo.GetNumberOfVehicles())
	return nil
}

// GetAssignments .
func (s *AssignFleetService) GetAssignments(
	requestDpo GetAssignmentsRequestDpo,
	responseEachDpo GetAssignmentsResponseDpo,
) error {
	ctx := context.Background()

	fleet, err := s.repo.GetByFlightplanID(
		ctx,
		flightplan.ID(requestDpo.GetId()),
	)
	if err != nil {
		return err
	}
	if fleet == nil {
		return errors.New("fleet not found")
	}

	var assignments []*assignmentVehicle
	fleet.ProvideAssignmentsInterest(
		func(assignmentID string, vehicleID string) {
			assignments = append(
				assignments,
				&assignmentVehicle{
					assignmentID: assignmentID,
					vehicleID:    vehicleID,
				},
			)
		},
		func(eventID string, assignmentID string, missionID string) {
			for _, av := range assignments {
				if av.assignmentID == assignmentID {
					av.events = append(
						av.events,
						&eventMission{
							eventID:   eventID,
							missionID: missionID,
						},
					)
				}
			}
		},
	)
	for _, a := range assignments {
		for _, e := range a.events {
			responseEachDpo(
				e.eventID,
				a.assignmentID,
				a.vehicleID,
				e.missionID)
		}
	}
	return nil
}

// UpdateAssignment .
func (s *AssignFleetService) UpdateAssignment(
	requestDpo UpdateAssignmentRequestDpo,
	responseDpo UpdateAssignmentResponseDpo,
) error {
	ctx := context.Background()

	aFleet, err := s.repo.GetByFlightplanID(
		ctx,
		flightplan.ID(requestDpo.GetId()),
	)
	if err != nil {
		return err
	}
	if aFleet == nil {
		return errors.New("fleet not found")
	}

	aFleet.AssignVehicle(
		fleet.AssignmentID(requestDpo.GetAssignmentId()),
		fleet.VehicleID(requestDpo.GetVehicleId()),
	)
	aFleet.AssignMission(
		fleet.EventID(requestDpo.GetId()),
		fleet.MissionID(requestDpo.GetMissionId()),
	)
	ret := s.repo.Save(ctx, aFleet)
	if ret != nil {
		return ret
	}

	responseDpo(
		requestDpo.GetId(),
		requestDpo.GetAssignmentId(),
		requestDpo.GetVehicleId(),
		requestDpo.GetMissionId(),
	)
	return nil
}

// ChangeNumberOfVehiclesRequestDpo .
type ChangeNumberOfVehiclesRequestDpo interface {
	GetId() string
	GetNumberOfVehicles() int32
}

// ChangeNumberOfVehiclesResponseDpo .
type ChangeNumberOfVehiclesResponseDpo = func(id string, numberOfVehicles int32)

// GetAssignmentsRequestDpo .
type GetAssignmentsRequestDpo interface {
	GetId() string
}

// GetAssignmentsResponseDpo .
type GetAssignmentsResponseDpo = func(id, assignmentId, vehicleId, missionId string)

type assignmentVehicle struct {
	assignmentID string
	vehicleID    string
	events       []*eventMission
}
type eventMission struct {
	eventID   string
	missionID string
}

// UpdateAssignmentRequestDpo .
type UpdateAssignmentRequestDpo interface {
	GetId() string
	GetAssignmentId() string
	GetVehicleId() string
	GetMissionId() string
}

// UpdateAssignmentResponseDpo .
type UpdateAssignmentResponseDpo = func(id, assignmentId, vehicleId, missionId string)
