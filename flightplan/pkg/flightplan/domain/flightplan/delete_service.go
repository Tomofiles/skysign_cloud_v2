package flightplan

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// DeleteFlightplan .
func DeleteFlightplan(
	tx txmanager.Tx,
	repo Repository,
	pub event.Publisher,
	id ID,
) error {
	flightplan, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}
	if flightplan == nil {
		return ErrNotFound
	}

	if err := repo.Delete(tx, id); err != nil {
		return err
	}

	pub.Publish(DeletedEvent{ID: flightplan.GetID()})
	return nil
}
