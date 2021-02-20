package flightplan

import (
	"errors"
	"flightplan/pkg/flightplan/domain/txmanager"
)

var (
	// ErrSave .
	ErrSave = errors.New("flightplan save error")
	// ErrNotFound .
	ErrNotFound = errors.New("flightplan not found")
	// ErrGet .
	ErrGet = errors.New("flightplan get error")
	// ErrDelete .
	ErrDelete = errors.New("flightplan delete error")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightplan, error)
	GetByID(txmanager.Tx, ID) (*Flightplan, error)
	Save(txmanager.Tx, *Flightplan) error
	Delete(txmanager.Tx, ID) error
}
