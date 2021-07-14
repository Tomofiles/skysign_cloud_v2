package flightplan

import (
	"flightplan/pkg/flightplan/domain/event"
	"flightplan/pkg/flightplan/domain/txmanager"
)

// ExecuteFlightplan .
func ExecuteFlightplan(
	tx txmanager.Tx,
	repo Repository,
	pub event.Publisher,
	id ID,
) error {
	flightplan, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	pub.Publish(&FlightplanExecutedEvent{
		ID:          flightplan.GetID(),
		Name:        flightplan.GetName(),
		Description: flightplan.GetDescription(),
		FleetID:     flightplan.GetFleetID(),
	})
	return nil
}
