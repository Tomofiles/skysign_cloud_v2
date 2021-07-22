package flightplan

import (
	"errors"
	"flight-operation/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightplan not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightplan, error)
	GetByID(txmanager.Tx, ID) (*Flightplan, error)
	Save(txmanager.Tx, *Flightplan) error
	Delete(txmanager.Tx, ID) error
}
