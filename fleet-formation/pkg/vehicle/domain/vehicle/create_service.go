package vehicle

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// CreateNewVehicle .
func CreateNewVehicle(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	communicationID CommunicationID,
) (string, error) {
	vehicle := NewInstance(gen)

	vehicle.SetPublisher(pub)

	// 生成直後のためエラーは発生しない想定
	vehicle.NameVehicle(name)
	vehicle.GiveCommunication(communicationID)

	if err := repo.Save(tx, vehicle); err != nil {
		return "", err
	}

	return string(vehicle.GetID()), nil
}
