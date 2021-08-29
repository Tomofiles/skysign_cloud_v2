package mission

import (
	"errors"
	"remote-communication/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("mission not found")
)

// Repository .
type Repository interface {
	GetByID(txmanager.Tx, ID) (*Mission, error)
	Save(txmanager.Tx, *Mission) error
	Delete(txmanager.Tx, ID) error
}
