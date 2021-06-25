package mission

import (
	"errors"
	"mission/pkg/mission/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("mission not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Mission, error)
	GetAllOrigin(txmanager.Tx) ([]*Mission, error)
	GetByID(txmanager.Tx, ID) (*Mission, error)
	Save(txmanager.Tx, *Mission) error
	Delete(txmanager.Tx, ID) error
}
