package service

import (
	"context"
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

	var assignments []assignmentVehicle
	fleet.ProvideAssignmentsInterest(
		func(assignmentID string, vehicleID string) {
			assignments = append(
				assignments,
				assignmentVehicle{
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
						eventMission{
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

// GetAssignmentsRequestDpo .
type GetAssignmentsRequestDpo interface {
	GetId() string
}

// GetAssignmentsResponseDpo .
type GetAssignmentsResponseDpo = func(id, assignmentId, vehicleId, missionId string)

type assignmentVehicle struct {
	assignmentID string
	vehicleID    string
	events       []eventMission
}
type eventMission struct {
	eventID   string
	missionID string
}
