package flightreport

import (
	"flightreport/pkg/flightreport/domain/txmanager"
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
