package uuid

import (
	c "remote-communication/pkg/communication/domain/communication"
	"time"

	"github.com/google/uuid"
)

// CommunicationUUID .
type CommunicationUUID struct{}

// NewCommunicationUUID .
func NewCommunicationUUID() *CommunicationUUID {
	return &CommunicationUUID{}
}

// NewCommandID .
func (g *CommunicationUUID) NewCommandID() c.CommandID {
	uuid, _ := uuid.NewRandom()
	return c.CommandID(uuid.String())
}

// NewTime .
func (g *CommunicationUUID) NewTime() time.Time {
	now := time.Now()
	return now
}
