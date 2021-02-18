package flightplan

import (
	"errors"
	"flightplan/pkg/flightplan/event"
	"flightplan/pkg/flightplan/txmanager"
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
		return errors.New("flightplan not found")
	}

	if err := repo.Delete(tx, id); err != nil {
		return err
	}

	pub.Publish(DeletedEvent{id: flightplan.GetID()})
	return nil
}
