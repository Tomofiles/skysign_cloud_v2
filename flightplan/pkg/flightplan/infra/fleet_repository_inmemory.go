package infra

import (
	fl "flightplan/pkg/flightplan/domain/fleet"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/txmanager"
	"log"
)

// InmemoryFleetRepository .
type InmemoryFleetRepository struct {
	fleets []*fl.Fleet
}

// GetByFlightplanID .
func (r *InmemoryFleetRepository) GetByFlightplanID(
	tx txmanager.Tx,
	flightplanID flightplan.ID,
) (*fl.Fleet, error) {
	for _, f := range r.fleets {
		if f.GetFlightplanID() == flightplanID {
			return f, nil
		}
	}
	return nil, nil
}

// Save .
func (r *InmemoryFleetRepository) Save(
	tx txmanager.Tx,
	fleet *fl.Fleet,
) error {
	log.Println(fleet)
	var fleets []*fl.Fleet
	for _, f := range r.fleets {
		if f.GetID() != fleet.GetID() {
			fleets = append(fleets, f)
		}
	}
	fleets = append(fleets, fleet)
	r.fleets = fleets
	return nil
}

// DeleteByFlightplanID .
func (r *InmemoryFleetRepository) DeleteByFlightplanID(
	tx txmanager.Tx,
	flightplanID flightplan.ID,
) error {
	var fleets []*fl.Fleet
	for _, f := range r.fleets {
		if f.GetFlightplanID() != flightplanID {
			fleets = append(fleets, f)
		}
	}
	r.fleets = fleets
	return nil
}