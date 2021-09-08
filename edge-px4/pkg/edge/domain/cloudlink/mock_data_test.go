package cloudlink

import (
	"errors"
	"fmt"
	"time"
)

const DefaultEdgeVehicleID = "vehicle-id"
const DefaultEdgeCommandID = "command-id"
const DefaultEdgeMissionID = "mission-id"

var (
	ErrPushTelemetry     = errors.New("push telemetry error")
	ErrPullCommand       = errors.New("pull command error")
	ErrPullUploadMission = errors.New("pull upload mission error")
	ErrGetUploadMission  = errors.New("get upload mission error")
)

type supportMock struct {
	messages []string
}

func (m *supportMock) NotifyInfo(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

type tickerMock struct {
	isStop bool
}

func (m *tickerMock) Tick() <-chan time.Time {
	tickStream := make(chan time.Time)
	go func() {
		defer close(tickStream)
		// tickStream <- time.Now()
	}()
	return tickStream
}

func (m *tickerMock) Stop() {
	m.isStop = true
}
