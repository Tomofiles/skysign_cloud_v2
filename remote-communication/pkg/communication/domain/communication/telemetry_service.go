package communication

import (
	"remote-communication/pkg/common/domain/event"
	"remote-communication/pkg/common/domain/txmanager"
)

// PushTelemetryService .
func PushTelemetryService(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
	snapshot TelemetrySnapshot,
) ([]CommandID, error) {
	communication, err := repo.GetByID(gen, id)
	if err != nil {
		return []CommandID{}, err
	}

	communication.PushTelemetry(snapshot)
	commandIDs := communication.GetCommandIDs()

	if ret := repo.Save(tx, communication); ret != nil {
		return []CommandID{}, ret
	}

	return commandIDs, nil
}

// PullTelemetryService .
func PullTelemetryService(
	tx txmanager.Tx,
	gen Generator,
	repo Repository,
	pub event.Publisher,
	id ID,
) (TelemetrySnapshot, error) {
	communication, err := repo.GetByID(gen, id)
	if err != nil {
		return TelemetrySnapshot{}, err
	}

	snapshot := communication.PullTelemetry()

	return snapshot, nil
}
