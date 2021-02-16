package service

import (
	"context"
	"errors"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/event"
)

// ManageFlightplanService .
type ManageFlightplanService struct {
	gen  flightplan.Generator
	repo flightplan.Repository
	pub  event.Publisher
}

// NewManageFlightplanService .
func NewManageFlightplanService(
	gen flightplan.Generator,
	repo flightplan.Repository,
	pub event.Publisher,
) ManageFlightplanService {
	return ManageFlightplanService{
		gen:  gen,
		repo: repo,
		pub:  pub,
	}
}

// GetFlightplan .
func (s *ManageFlightplanService) GetFlightplan(
	requestDpo GetFlightplanRequestDpo,
	responseDpo GetFlightplanResponseDpo,
) error {
	ctx := context.Background()

	flightplan, err := s.repo.GetByID(ctx, flightplan.ID(requestDpo.GetId()))
	if err != nil {
		return err
	}
	if flightplan == nil {
		return errors.New("flightplan not found")
	}

	responseDpo(
		string(flightplan.GetID()),
		flightplan.GetName(),
		flightplan.GetDescription(),
	)
	return nil
}

// ListFlightplans .
func (s *ManageFlightplanService) ListFlightplans(
	responseEachDpo ListFlightplansResponseDpo,
) error {
	ctx := context.Background()

	flightplans, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, f := range flightplans {
		responseEachDpo(
			string(f.GetID()),
			f.GetName(),
			f.GetDescription(),
		)
	}
	return nil
}

// CreateFlightplan .
func (s *ManageFlightplanService) CreateFlightplan(
	requestDpo CreateFlightplanRequestDpo,
	responseDpo CreateFlightplanResponseDpo,
) error {
	ctx := context.Background()

	id, ret := flightplan.CreateNewFlightplan(
		ctx,
		s.gen,
		s.repo,
		s.pub,
		requestDpo.GetName(),
		requestDpo.GetDescription())
	if ret != nil {
		return ret
	}

	responseDpo(id, requestDpo.GetName(), requestDpo.GetDescription())
	return nil
}

// UpdateFlightplan .
func (s *ManageFlightplanService) UpdateFlightplan(
	requestDpo UpdateFlightplanRequestDpo,
	responseDpo UpdateFlightplanResponseDpo,
) error {
	ctx := context.Background()

	flightplan, err := s.repo.GetByID(ctx, flightplan.ID(requestDpo.GetId()))
	if err != nil {
		return err
	}
	if flightplan == nil {
		return errors.New("flightplan not found")
	}

	flightplan.NameFlightplan(requestDpo.GetName())
	flightplan.ChangeDescription(requestDpo.GetDescription())

	ret := s.repo.Save(ctx, flightplan)
	if ret != nil {
		return ret
	}

	responseDpo(
		string(flightplan.GetID()),
		flightplan.GetName(),
		flightplan.GetDescription(),
	)
	return nil
}

// DeleteFlightplan .
func (s *ManageFlightplanService) DeleteFlightplan(
	requestDpo DeleteFlightplanRequestDpo,
) error {
	ctx := context.Background()

	ret := flightplan.DeleteFlightplan(
		ctx,
		s.repo,
		s.pub,
		flightplan.ID(requestDpo.GetId()))
	if ret != nil {
		return ret
	}

	return nil
}

// CreateFlightplanRequestDpo .
type CreateFlightplanRequestDpo interface {
	GetId() string
	GetName() string
	GetDescription() string
}

// CreateFlightplanResponseDpo .
type CreateFlightplanResponseDpo = func(id, name, description string)

// UpdateFlightplanRequestDpo .
type UpdateFlightplanRequestDpo interface {
	GetId() string
	GetName() string
	GetDescription() string
}

// UpdateFlightplanResponseDpo .
type UpdateFlightplanResponseDpo = func(id, name, description string)

// GetFlightplanRequestDpo .
type GetFlightplanRequestDpo interface {
	GetId() string
}

// GetFlightplanResponseDpo .
type GetFlightplanResponseDpo = func(id, name, description string)

// ListFlightplansResponseDpo .
type ListFlightplansResponseDpo = func(id, name, description string)

// DeleteFlightplanRequestDpo .
type DeleteFlightplanRequestDpo interface {
	GetId() string
}
