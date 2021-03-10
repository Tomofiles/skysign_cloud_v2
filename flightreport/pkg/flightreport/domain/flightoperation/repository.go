package flightoperation

import (
	"errors"
	"flightreport/pkg/flightreport/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightoperation not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightoperation, error)
	GetByID(txmanager.Tx, ID) (*Flightoperation, error)
	Save(txmanager.Tx, *Flightoperation) error
}
