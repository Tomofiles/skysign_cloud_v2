package flightplan

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
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

	flightplan.SetPublisher(pub)
	if err := flightplan.RemoveFleetID(); err != nil {
		return err
	}

	if err := repo.Delete(tx, id); err != nil {
		return err
	}

	return nil
}
