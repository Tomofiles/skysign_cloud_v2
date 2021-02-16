package flightplan

import (
	"context"
	"flightplan/pkg/flightplan/event"
)

// CreateNewFlightplan .
func CreateNewFlightplan(
	ctx context.Context,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	description string,
) (string, error) {
	flightplan := NewInstance(gen)

	flightplan.NameFlightplan(name)
	flightplan.ChangeDescription(description)

	if err := repo.Save(ctx, flightplan); err != nil {
		return "", err
	}

	pub.Publish(CreatedEvent{id: flightplan.GetID()})
	return string(flightplan.GetID()), nil
}
