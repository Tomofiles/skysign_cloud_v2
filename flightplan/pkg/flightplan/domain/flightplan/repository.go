package flightplan

import "context"

// Repository .
type Repository interface {
	GetByID(context.Context, ID) (*Flightplan, error)
	Save(context.Context, *Flightplan) error
	Delete(context.Context, ID) error
}
