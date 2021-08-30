package communication

// IsFollowArmCommandPushPolicy .
func IsFollowArmCommandPushPolicy(cType CommandType, communication *Communication) bool {
	if cType == CommandTypeTAKEOFF || cType == CommandTypeSTART {
		if communication.telemetry.IsDisarmed() {
			return true
		}
	}
	return false
}
