package communication

import (
	"errors"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
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
