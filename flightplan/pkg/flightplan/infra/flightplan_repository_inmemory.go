package infra

import (
	"context"
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"log"
)

// InmemoryFlightplanRepository .
type InmemoryFlightplanRepository struct {
	flightplans []*fpl.Flightplan
}

// GetAll .
func (r *InmemoryFlightplanRepository) GetAll(
	ctx context.Context,
) ([]*fpl.Flightplan, error) {
	return r.flightplans, nil
}

// GetByID .
func (r *InmemoryFlightplanRepository) GetByID(
	ctx context.Context,
	id fpl.ID,
) (*fpl.Flightplan, error) {
	for _, f := range r.flightplans {
		if f.GetID() == id {
			return f, nil
		}
	}
	return nil, nil
}

// Save .
func (r *InmemoryFlightplanRepository) Save(
	ctx context.Context,
	flightplan *fpl.Flightplan,
) error {
	log.Println(flightplan)
	var flightplans []*fpl.Flightplan
	for _, f := range r.flightplans {
		if f.GetID() != flightplan.GetID() {
			flightplans = append(flightplans, f)
		}
	}
	flightplans = append(flightplans, flightplan)
	r.flightplans = flightplans
	return nil
}

// Delete .
func (r *InmemoryFlightplanRepository) Delete(
	ctx context.Context,
	id fpl.ID,
) error {
	var flightplans []*fpl.Flightplan
	for _, f := range r.flightplans {
		if f.GetID() != id {
			flightplans = append(flightplans, f)
		}
	}
	r.flightplans = flightplans
	return nil
}
