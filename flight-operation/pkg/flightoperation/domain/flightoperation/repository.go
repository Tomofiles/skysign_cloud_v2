package flightoperation

import (
	"errors"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

var (
	// ErrNotFound .
	ErrNotFound = errors.New("flightoperation not found")
)

// Repository .
type Repository interface {
	GetAll(txmanager.Tx) ([]*Flightoperation, error)
	GetAllOperating(txmanager.Tx) ([]*Flightoperation, error)
	GetByID(txmanager.Tx, ID) (*Flightoperation, error)
	Save(txmanager.Tx, *Flightoperation) error
}
