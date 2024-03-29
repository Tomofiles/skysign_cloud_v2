package flightplan

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// UpdateFlightplan .
func UpdateFlightplan(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	name string,
	description string,
) error {
	flightplan, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	if err := flightplan.NameFlightplan(name); err != nil {
		return err
	}
	if err := flightplan.ChangeDescription(description); err != nil {
		return err
	}

	if ret := repo.Save(tx, flightplan); ret != nil {
		return ret
	}

	return nil
}
