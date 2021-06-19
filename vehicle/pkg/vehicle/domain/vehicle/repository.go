package vehicle

import (
	"errors"
	"vehicle/pkg/vehicle/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("vehicle not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Vehicle, error)
	GetAllOrigin(txmanager.Tx) ([]*Vehicle, error)
	GetByID(txmanager.Tx, ID) (*Vehicle, error)
	Save(txmanager.Tx, *Vehicle) error
	Delete(txmanager.Tx, ID) error
}
