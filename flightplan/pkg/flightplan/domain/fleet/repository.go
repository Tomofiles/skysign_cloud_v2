package fleet

import (
	"errors"
	"flightplan/pkg/flightplan/domain/flightplan"
	"flightplan/pkg/flightplan/domain/txmanager"
)

var (
	// ErrSave .
	ErrSave = errors.New("fleet save error")
	// ErrNotFound .
	ErrNotFound = errors.New("fleet not found")
	// ErrGet .
	ErrGet = errors.New("fleet get error")
	// ErrDelete .
	ErrDelete = errors.New("fleet delete error")
)

// Repository .
type Repository interface {
	GetByFlightplanID(txmanager.Tx, flightplan.ID) (*Fleet, error)
	Save(txmanager.Tx, *Fleet) error
	DeleteByFlightplanID(txmanager.Tx, flightplan.ID) error
}
