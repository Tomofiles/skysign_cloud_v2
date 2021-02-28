package flightplan

import (
	"errors"
	"flightplan/pkg/flightplan/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightplan not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightplan, error)
	GetAllOrigin(txmanager.Tx) ([]*Flightplan, error)
	GetByID(txmanager.Tx, ID) (*Flightplan, error)
	Save(txmanager.Tx, *Flightplan) error
	Delete(txmanager.Tx, ID) error
}
