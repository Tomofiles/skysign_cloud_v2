package flightreport

import (
	"flightreport/pkg/flightreport/domain/event"
	"flightreport/pkg/flightreport/domain/txmanager"
)

// CreateNewFlightreport .
func CreateNewFlightreport(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	originalID FlightoperationID,
) error {
	newID := gen.NewFlightoperationID()

	flightreport := NewInstance(gen, newID)

	if err := repo.Save(tx, flightreport); err != nil {
		return err
	}

	pub.Publish(CreatedEvent{
		ID:                flightreport.GetID(),
		FlightoperationID: newID,
	})
	pub.Publish(FlightoperationCopiedWhenCreatedEvent{
		OriginalID: originalID,
		NewID:      newID,
	})
	return nil
}
