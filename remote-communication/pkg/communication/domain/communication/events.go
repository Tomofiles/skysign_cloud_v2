package communication

// TelemetryUpdatedEvent .
type TelemetryUpdatedEvent struct {
	CommunicationID ID
	Telemetry       TelemetrySnapshot
}

// GetID .
func (e *TelemetryUpdatedEvent) GetID() string {
	return string(e.CommunicationID)
}

// GetTelemetry .
func (e *TelemetryUpdatedEvent) GetTelemetry() TelemetrySnapshot {
	return e.Telemetry
}
