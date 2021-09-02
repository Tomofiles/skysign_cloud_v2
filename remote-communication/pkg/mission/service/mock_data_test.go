package service

import (
	m "remote-communication/pkg/mission/domain/mission"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/event"
	"github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = m.ID("mission-id")

type repositoryMock struct {
	mock.Mock

	saveMissions []*m.Mission
}

func (rm *repositoryMock) GetByID(tx txmanager.Tx, id m.ID) (*m.Mission, error) {
	ret := rm.Called(id)
	var v *m.Mission
	if ret.Get(0) == nil {
		v = nil
	} else {
		v = ret.Get(0).(*m.Mission)
	}
	return v, ret.Error(1)
}
func (rm *repositoryMock) Save(tx txmanager.Tx, v *m.Mission) error {
	ret := rm.Called(v)
	if ret.Error(0) == nil {
		rm.saveMissions = append(rm.saveMissions, v)
	}
	return ret.Error(0)
}
func (rm *repositoryMock) Delete(tx txmanager.Tx, id m.ID) error {
	ret := rm.Called(id)
	return ret.Error(0)
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

type missionCommandMock struct {
	ID        string
	Waypoints []Waypoint
}

func (c *missionCommandMock) GetID() string {
	return c.ID
}

func (c *missionCommandMock) GetWaypoints() []Waypoint {
	return c.Waypoints
}

type missionIDCommandMock struct {
	ID string
}

func (c *missionIDCommandMock) GetID() string {
	return c.ID
}
