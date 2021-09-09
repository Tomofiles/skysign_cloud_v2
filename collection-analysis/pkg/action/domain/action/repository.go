package action

import (
	"errors"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("action not found")
)

// Repository .
type Repository interface {
	GetByID(txmanager.Tx, ID) (*Action, error)
	GetAllActiveByFleetID(txmanager.Tx, FleetID) ([]*Action, error)
	GetActiveByCommunicationID(txmanager.Tx, CommunicationID) (*Action, error)
	Save(txmanager.Tx, *Action) error
}
