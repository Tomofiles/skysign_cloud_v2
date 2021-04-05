package action

import (
	"action/pkg/action/domain/txmanager"
	"errors"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("action not found")
)

// Repository .
type Repository interface {
	GetByID(txmanager.Tx, ID) (*Action, error)
	GetAllActiveByFlightplanID(txmanager.Tx, FlightplanID) ([]*Action, error)
	GetActiveByCommunicationID(txmanager.Tx, CommunicationID) (*Action, error)
	Save(txmanager.Tx, *Action) error
}
