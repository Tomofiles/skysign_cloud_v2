package communication

import (
	"errors"
	"remote-communication/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("communication not found")
)

// Repository .
type Repository interface {
	GetByID(txmanager.Tx, ID) (*Communication, error)
	Save(txmanager.Tx, *Communication) error
	Delete(txmanager.Tx, ID) error
}
