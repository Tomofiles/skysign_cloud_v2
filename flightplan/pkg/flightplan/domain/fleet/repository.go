package fleet

import (
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// Repository .
type Repository interface {
	GetByFlightplanID(txmanager.Tx, flightplan.ID) (*Fleet, error)
	Save(txmanager.Tx, *Fleet) error
	DeleteByFlightplanID(txmanager.Tx, flightplan.ID) error
}
