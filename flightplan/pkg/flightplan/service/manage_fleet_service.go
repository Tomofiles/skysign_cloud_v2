package service

import (
	"context"
	"flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/event"
)

// ManageFleetService .
type ManageFleetService struct {
	gen  fleet.Generator
	repo fleet.Repository
	pub  event.Publisher
}

// NewManageFleetService .
func NewManageFleetService(
	gen fleet.Generator,
	repo fleet.Repository,
	pub event.Publisher,
) ManageFleetService {
	return ManageFleetService{
		gen:  gen,
		repo: repo,
		pub:  pub,
	}
}

// CreateFleet .
func (s *ManageFleetService) CreateFleet(
	requestDpo CreateFleetRequestDpo,
) error {
	ctx := context.Background()

	fleet := fleet.NewInstance(
		s.gen,
		flightplan.ID(requestDpo.GetFlightplanID()),
		0)
	ret := s.repo.Save(ctx, fleet)
	if ret != nil {
		return ret
	}

	return nil
}

// DeleteFleet .
func (s *ManageFleetService) DeleteFleet(
	requestDpo DeleteFleetRequestDpo,
) error {
	ctx := context.Background()

	ret := s.repo.DeleteByFlightplanID(ctx, flightplan.ID(requestDpo.GetFlightplanID()))
	if ret != nil {
		return ret
	}

	return nil
}

// CreateFleetRequestDpo .
type CreateFleetRequestDpo interface {
	GetFlightplanID() string
}

// DeleteFleetRequestDpo .
type DeleteFleetRequestDpo interface {
	GetFlightplanID() string
}
