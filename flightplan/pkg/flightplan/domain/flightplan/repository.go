package flightplan

import "context"

// Repository .
type Repository interface {
	GetAll(context.Context) ([]*Flightplan, error)
	GetByID(context.Context, ID) (*Flightplan, error)
	Save(context.Context, *Flightplan) error
	Delete(context.Context, ID) error
}
