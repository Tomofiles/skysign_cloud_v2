package vehicle

import (
	"vehicle/pkg/vehicle/domain/event"
	"vehicle/pkg/vehicle/domain/txmanager"
)

// CreateNewVehicle .
func CreateNewVehicle(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	communicationID string,
) (string, error) {
	vehicle := NewInstance(gen)

	vehicle.SetPublisher(pub)

	// 生成直後のためエラーは発生しない想定
	vehicle.NameVehicle(name)
	vehicle.GiveCommunication(CommunicationID(communicationID))

	if err := repo.Save(tx, vehicle); err != nil {
		return "", err
	}

	return string(vehicle.GetID()), nil
}
