package flightreport

import (
	"flightreport/pkg/flightreport/domain/txmanager"
)

// CreateNewFlightreport .
func CreateNewFlightreport(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	originalID FlightoperationID,
) error {
	flightreport := NewInstance(gen, originalID)

	if err := repo.Save(tx, flightreport); err != nil {
		return err
	}

	return nil
}
