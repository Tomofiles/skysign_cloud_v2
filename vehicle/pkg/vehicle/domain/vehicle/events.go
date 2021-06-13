package vehicle

// CommunicationIdGaveEvent .
type CommunicationIdGaveEvent struct {
	CommunicationID CommunicationID
}

// GetCommunicationID .
func (e *CommunicationIdGaveEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}

// CommunicationIdRemovedEvent .
type CommunicationIdRemovedEvent struct {
	CommunicationID CommunicationID
}

// GetCommunicationID .
func (e *CommunicationIdRemovedEvent) GetCommunicationID() string {
	return string(e.CommunicationID)
}
