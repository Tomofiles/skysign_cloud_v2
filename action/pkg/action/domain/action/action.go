package action

import "errors"

// ID .
type ID string

// FleetID .
type FleetID string

// CommunicationID .
type CommunicationID string

const (
	// Active .
	Active = false
	// Completed .
	Completed = true
)

var (
	// ErrCannotChange .
	ErrCannotChange = errors.New("cannnot change completed action")
)

// Action .
type Action struct {
	id              ID
	communicationID CommunicationID
	fleetID         FleetID
	isCompleted     bool
	trajectory      Trajectory
}

// GetID .
func (a *Action) GetID() ID {
	return a.id
}

// GetCommunicationID .
func (a *Action) GetCommunicationID() CommunicationID {
	return a.communicationID
}

// GetFleetID .
func (a *Action) GetFleetID() FleetID {
	return a.fleetID
}

// PushTelemetry .
func (a *Action) PushTelemetry(snapshot TelemetrySnapshot) error {
	if a.isCompleted {
		return ErrCannotChange
	}
	if snapshot.IsDisarmed() {
		// telemetryがdisarmedの場合何もしない
		return nil
	}
	a.trajectory = a.trajectory.Extension(snapshot)

	return nil
}

// Complete .
func (a *Action) Complete() error {
	if a.isCompleted {
		return ErrCannotChange
	}
	a.isCompleted = Completed

	return nil
}

// ProvideTrajectoryInterest .
func (a *Action) ProvideTrajectoryInterest(trajectoryPoint func(snapshot TelemetrySnapshot)) {
	for _, snapshot := range a.trajectory.TakeASnapshots() {
		trajectoryPoint(snapshot)
	}
}
