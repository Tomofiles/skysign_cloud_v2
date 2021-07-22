package flightplan

import (
	"flight-operation/pkg/common/domain/event"
	"flight-operation/pkg/common/domain/txmanager"
)

// CreateNewFlightplan .
func CreateNewFlightplan(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	description string,
) (ID, FleetID, error) {
	flightplan := NewInstance(gen)

	flightplan.SetPublisher(pub)

	// 生成直後のためエラーは発生しない想定
	flightplan.NameFlightplan(name)
	flightplan.ChangeDescription(description)
	flightplan.ChangeNumberOfVehicles(0) // デフォルト機体数は0

	if err := repo.Save(tx, flightplan); err != nil {
		return ID(""), FleetID(""), err
	}

	return flightplan.GetID(), flightplan.GetFleetID(), nil
}
