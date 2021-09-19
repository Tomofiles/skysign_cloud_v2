package uuid

import (
	"time"

	c "github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/communication/domain/communication"

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
