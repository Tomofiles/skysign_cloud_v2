package cloudlink

import (
	"fmt"
)

const DefaultEdgeVehicleID = "vehicle-id"
const DefaultEdgeCommandID = "command-id"
const DefaultEdgeMissionID = "mission-id"

type supportMock struct {
	messages []string
}

func (m *supportMock) NotifyInfo(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}
