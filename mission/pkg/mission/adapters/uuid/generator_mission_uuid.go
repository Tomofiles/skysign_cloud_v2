package uuid

import (
	"mission/pkg/mission/domain/mission"

	"github.com/google/uuid"
)

// MissionUUID .
type MissionUUID struct{}

// NewMissionUUID .
func NewMissionUUID() *MissionUUID {
	return &MissionUUID{}
}

// NewID .
func (g *MissionUUID) NewID() mission.ID {
	uuid, _ := uuid.NewRandom()
	return mission.ID(uuid.String())
}

// NewVersion .
func (g *MissionUUID) NewVersion() mission.Version {
	uuid, _ := uuid.NewRandom()
	return mission.Version(uuid.String())
}
