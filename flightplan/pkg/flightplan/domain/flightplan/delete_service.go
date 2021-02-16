package flightplan

import (
	"context"
	"errors"
	"flightplan/pkg/flightplan/event"
)

// DeleteFlightplan .
func DeleteFlightplan(
	ctx context.Context,
	repo Repository,
	pub event.Publisher,
	id ID,
) error {
	flightplan, err := repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if flightplan == nil {
		return errors.New("flightplan not found")
	}

	if err := repo.Delete(ctx, id); err != nil {
		return err
	}

	pub.Publish(DeletedEvent{id: flightplan.GetID()})
	return nil
}
