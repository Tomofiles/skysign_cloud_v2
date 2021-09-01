package flightoperation

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// CreateNewFlightoperation .
func CreateNewFlightoperation(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	name string,
	description string,
	originalID FleetID,
) error {
	newID := gen.NewFleetID()

	flightoperation := NewInstance(gen, newID)

	flightoperation.SetPublisher(pub)

	// 生成直後のためエラーは発生しない想定
	flightoperation.NameFlightoperation(name)
	flightoperation.ChangeDescription(description)

	if err := repo.Save(tx, flightoperation); err != nil {
		return err
	}

	pub.Publish(FleetCopiedEvent{
		OriginalID: originalID,
		NewID:      newID,
	})
	return nil
}
