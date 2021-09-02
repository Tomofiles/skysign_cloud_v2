package flightplan

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
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
