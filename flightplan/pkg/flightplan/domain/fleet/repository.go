package fleet

import (
	"context"
	"flightplan/pkg/flightplan/domain/flightplan"
)

// Repository .
type Repository interface {
	GetByFlightplanID(context.Context, flightplan.ID) (*Fleet, error)
	Save(context.Context, *Fleet) error
	DeleteByFlightplanID(context.Context, flightplan.ID) error
}
