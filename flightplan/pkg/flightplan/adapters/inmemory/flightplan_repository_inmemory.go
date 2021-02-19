package inmemory

import (
	fpl "flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// FlightplanRepository .
type FlightplanRepository struct {
	flightplans []*fpl.Flightplan
}

// GetAll .
func (r *FlightplanRepository) GetAll(
	tx txmanager.Tx,
) ([]*fpl.Flightplan, error) {
	return r.flightplans, nil
}

// GetByID .
func (r *FlightplanRepository) GetByID(
	tx txmanager.Tx,
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
func (r *FlightplanRepository) Save(
	tx txmanager.Tx,
	flightplan *fpl.Flightplan,
) error {
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
func (r *FlightplanRepository) Delete(
	tx txmanager.Tx,
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
