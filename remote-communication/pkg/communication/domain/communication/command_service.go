package communication

import (
	"remote-communication/pkg/common/domain/event"
	"remote-communication/pkg/common/domain/txmanager"
)

// PushCommandService .
func PushCommandService(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	cType CommandType,
) (CommandID, error) {
	communication, err := repo.GetByID(tx, id)
	if err != nil {
		return "", err
	}

	commandID := communication.PushCommand(cType)

	if ret := repo.Save(tx, communication); ret != nil {
		return "", ret
	}

	return commandID, nil
}

// PullCommandService .
func PullCommandService(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	commandID CommandID,
) (CommandType, error) {
	communication, err := repo.GetByID(tx, id)
	if err != nil {
		return "", err
	}

	cType, err := communication.PullCommandByID(commandID)
	if err != nil {
		return "", err
	}

	if ret := repo.Save(tx, communication); ret != nil {
		return "", ret
	}

	return cType, nil
}
