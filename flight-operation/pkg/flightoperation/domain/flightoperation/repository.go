package flightoperation

import (
	"errors"
	"flight-operation/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightoperation not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightoperation, error)
	GetAllOperating(txmanager.Tx) ([]*Flightoperation, error)
	GetByID(txmanager.Tx, ID) (*Flightoperation, error)
	Save(txmanager.Tx, *Flightoperation) error
}
