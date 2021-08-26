package service

import (
	"remote-communication/pkg/common/domain/event"
	"remote-communication/pkg/common/domain/txmanager"
	c "remote-communication/pkg/communication/domain/communication"
	"time"

	"github.com/stretchr/testify/mock"
)

const DefaultCommunicationID = c.ID("communication-id")
const DefaultCommunicationCommandID = c.CommandID("command-id")
const DefaultCommunicationMissionID = c.MissionID("mission-id")

var DefaultCommunicationTime = time.Now()

type repositoryMock struct {
	mock.Mock

	saveCommunications []*c.Communication
}

func (rm *repositoryMock) GetByID(tx txmanager.Tx, id c.ID) (*c.Communication, error) {
	ret := rm.Called(id)
	var v *c.Communication
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*c.Communication)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMock) Save(tx txmanager.Tx, v *c.Communication) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveCommunications = append(rm.saveCommunications, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMock) Delete(tx txmanager.Tx, id c.ID) error {
	ret := rm.Called(id)
	return ret.Error(0)
}

type generatorMock struct {
	c.Generator
	commandIDs     []c.CommandID
	commandIDIndex int
	times          []time.Time
	timeIndex      int
}

func (gen *generatorMock) NewCommandID() c.CommandID {
	commandID := gen.commandIDs[gen.commandIDIndex]
	gen.commandIDIndex++
	return commandID
}
func (gen *generatorMock) NewTime() time.Time {
	time := gen.times[gen.timeIndex]
	gen.timeIndex++
	return time
}

type publisherMock struct {
	events  []interface{}
	isFlush bool
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	rm.isFlush = true
	return nil
}

type pubSubManagerMock struct {
	mock.Mock
}

func (psm *pubSubManagerMock) GetPublisher() (event.Publisher, event.ChannelClose, error) {
	ret := psm.Called()
	var pub event.Publisher
	var close func() error
	if ret.Get(0) == nil {
		pub = nil
	} else {
		pub = ret.Get(0).(event.Publisher)
		close = ret.Get(1).(func() error)
	}
	return pub, close, ret.Error(2)
}

type txManagerMock struct {
	isOpe, isEH error
}

func (txm *txManagerMock) Do(operation func(txmanager.Tx) error) error {
	txm.isOpe = operation(nil)
	return nil
}

func (txm *txManagerMock) DoAndEndHook(operation func(txmanager.Tx) error, endHook func() error) error {
	txm.isOpe = operation(nil)
	txm.isEH = endHook()
	return nil
}

type communicationIDCommandMock struct {
	ID string
}

func (c *communicationIDCommandMock) GetID() string {
	return c.ID
}

type pushCommandCommandMock struct {
	ID, Type string
}

func (c *pushCommandCommandMock) GetID() string {
	return c.ID
}

func (c *pushCommandCommandMock) GetType() string {
	return c.Type
}

type pushUploadMissionCommandMock struct {
	ID, MissionID string
}

func (c *pushUploadMissionCommandMock) GetID() string {
	return c.ID
}

func (c *pushUploadMissionCommandMock) GetMissionID() string {
	return c.MissionID
}

type pullCommandMock struct {
	ID, CommandID string
}

func (c *pullCommandMock) GetID() string {
	return c.ID
}

func (c *pullCommandMock) GetCommandID() string {
	return c.CommandID
}

type pushTelemetryCommandMock struct {
	ID        string
	Telemetry *telemetry
}

func (c *pushTelemetryCommandMock) GetID() string {
	return c.ID
}

func (c *pushTelemetryCommandMock) GetTelemetry() EdgeTelemetry {
	return c.Telemetry
}
