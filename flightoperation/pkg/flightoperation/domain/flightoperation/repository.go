package flightoperation

import (
	"errors"
	"flightoperation/pkg/flightoperation/domain/txmanager"
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
