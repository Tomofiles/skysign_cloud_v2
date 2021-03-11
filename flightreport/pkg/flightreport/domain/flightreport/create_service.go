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
	newID := gen.NewFlightoperationID()

	flightreport := NewInstance(gen, newID)

	if err := repo.Save(tx, flightreport); err != nil {
		return err
	}

	return nil
}
