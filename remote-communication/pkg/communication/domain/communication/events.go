package communication

// TelemetryUpdatedEvent .
type TelemetryUpdatedEvent struct {
	CommunicationID ID
	Telemetry       TelemetrySnapshot
}
