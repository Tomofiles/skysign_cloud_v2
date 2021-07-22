package fleet

import (
	"errors"
	"fleet-formation/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("fleet not found")
)

// Repository .
type Repository interface {
	GetByID(txmanager.Tx, ID) (*Fleet, error)
	Save(txmanager.Tx, *Fleet) error
	Delete(txmanager.Tx, ID) error
}
