package flightplan

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// CreateNewFlightplan .
func CreateNewFlightplan(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	description string,
) (string, error) {
	flightplan := NewInstance(gen)

	flightplan.NameFlightplan(name)
	flightplan.ChangeDescription(description)

	if err := repo.Save(tx, flightplan); err != nil {
		return "", err
	}

	pub.Publish(CreatedEvent{ID: flightplan.GetID()})
	return string(flightplan.GetID()), nil
}
