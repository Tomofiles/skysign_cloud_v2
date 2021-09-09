package flightreport

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// CreateNewFlightreport .
func CreateNewFlightreport(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	name string,
	description string,
	fleetID FleetID,
) error {
	flightreport := NewInstance(gen, name, description, fleetID)

	if err := repo.Save(tx, flightreport); err != nil {
		return err
	}

	return nil
}
