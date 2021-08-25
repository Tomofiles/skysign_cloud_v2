package communication

import (
	"remote-communication/pkg/common/domain/event"
	"remote-communication/pkg/common/domain/txmanager"
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
	communication, err := repo.GetByID(gen, id)
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
	communication, err := repo.GetByID(gen, id)
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
