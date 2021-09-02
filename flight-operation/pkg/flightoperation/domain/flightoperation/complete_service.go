package flightoperation

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// CompleteFlightoperation .
func CompleteFlightoperation(
	tx txmanager.Tx,
	repo Repository,
	pub event.Publisher,
	id ID,
) error {
	flightoperation, err := repo.GetByID(tx, id)
	if err != nil {
		return err
	}

	flightoperation.SetPublisher(pub)
	if err := flightoperation.Complete(); err != nil {
		return err
	}

	if err := repo.Save(tx, flightoperation); err != nil {
		return err
	}

	return nil
}
