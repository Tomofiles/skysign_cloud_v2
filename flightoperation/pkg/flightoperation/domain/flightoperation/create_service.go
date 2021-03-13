package flightoperation

import (
	"flightoperation/pkg/flightoperation/domain/event"
	"flightoperation/pkg/flightoperation/domain/txmanager"
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
		ID:         flightoperation.GetID(),
		OriginalID: originalID,
		NewID:      newID,
	})
	return nil
}
