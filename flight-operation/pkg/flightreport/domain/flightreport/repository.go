package flightreport

import (
	"errors"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightreport not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightreport, error)
	GetByID(txmanager.Tx, ID) (*Flightreport, error)
	Save(txmanager.Tx, *Flightreport) error
}
