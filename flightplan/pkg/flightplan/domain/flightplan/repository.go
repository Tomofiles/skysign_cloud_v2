package flightplan

import (
	"flightplan/pkg/flightplan/txmanager"
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightplan, error)
	GetByID(txmanager.Tx, ID) (*Flightplan, error)
	Save(txmanager.Tx, *Flightplan) error
	Delete(txmanager.Tx, ID) error
}
