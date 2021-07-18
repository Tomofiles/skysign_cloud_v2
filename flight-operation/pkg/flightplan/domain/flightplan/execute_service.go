package flightplan

import (
	"flight-operation/pkg/common/domain/event"
	"flight-operation/pkg/common/domain/txmanager"
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

	pub.Publish(FlightplanExecutedEvent{
		ID:          flightplan.GetID(),
		Name:        flightplan.GetName(),
		Description: flightplan.GetDescription(),
		FleetID:     flightplan.GetFleetID(),
	})
	return nil
}
