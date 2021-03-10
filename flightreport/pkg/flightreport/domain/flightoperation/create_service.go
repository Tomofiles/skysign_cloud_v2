package flightoperation

import (
	"flightreport/pkg/flightreport/domain/event"
	"flightreport/pkg/flightreport/domain/txmanager"
)

// CreateNewFlightoperation .
func CreateNewFlightoperation(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	originalID FlightplanID,
) error {
	newID := gen.NewFlightplanID()

	flightoperation := NewInstance(gen, newID)

	if err := repo.Save(tx, flightoperation); err != nil {
		return err
	}

	pub.Publish(CreatedEvent{
		ID:           flightoperation.GetID(),
		FlightplanID: newID,
	})
	pub.Publish(FlightplanCopiedWhenCreatedEvent{
		OriginalID: originalID,
		NewID:      newID,
	})
	return nil
}
