package communication

import (
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"
)

// PushUploadMissionService .
func PushUploadMissionService(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	missionID MissionID,
) (CommandID, error) {
	communication, err := repo.GetByID(tx, id)
	if err != nil {
		return "", err
	}

	commandID := communication.PushUploadMission(missionID)

	if ret := repo.Save(tx, communication); ret != nil {
		return "", ret
	}

	return commandID, nil
}

// PullUploadMissionService .
func PullUploadMissionService(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	commandID CommandID,
) (MissionID, error) {
	communication, err := repo.GetByID(tx, id)
	if err != nil {
		return "", err
	}

	missionID, err := communication.PullUploadMissionByID(commandID)
	if err != nil {
		return "", err
	}

	if ret := repo.Save(tx, communication); ret != nil {
		return "", ret
	}

	return missionID, nil
}
