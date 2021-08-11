package communication

// MissionID .
type MissionID string

// UploadMission .
type UploadMission struct {
	commandID CommandID
	missionID MissionID
}

// NewUploadMission .
func NewUploadMission(commandID CommandID, missionID MissionID) *UploadMission {
	return nil
}
