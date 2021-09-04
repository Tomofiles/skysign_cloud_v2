package telemetry

import (
	"fmt"
)

const DefaultEdgeVehicleID = "vehicle-id"

type supportMock struct {
	messages []string
}

func (m *supportMock) NotifyInfo(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}
