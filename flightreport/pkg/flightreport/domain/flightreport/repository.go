package flightreport

import (
	"errors"
	"flightreport/pkg/flightreport/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightreport not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightreport, error)
	GetByID(txmanager.Tx, ID) (*Flightreport, error)
	Save(txmanager.Tx, *Flightreport) error
}
