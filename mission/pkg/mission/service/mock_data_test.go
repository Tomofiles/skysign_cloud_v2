package service

import (
	"context"
	"mission/pkg/mission/domain/event"
	m "mission/pkg/mission/domain/mission"
	"mission/pkg/mission/domain/txmanager"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = m.ID("mission-id")
const DefaultMissionVersion = m.Version("version")
const DefaultMissionName = "mission-name"
const DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM float64 = 10

type missionRepositoryMock struct {
	mock.Mock
}

func (r *missionRepositoryMock) GetAll(
	tx txmanager.Tx,
) ([]*m.Mission, error) {
	ret := r.Called()
	var f []*m.Mission
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*m.Mission)
	}
	return f, ret.Error(1)
}

func (r *missionRepositoryMock) GetAllOrigin(
	tx txmanager.Tx,
) ([]*m.Mission, error) {
	ret := r.Called()
	var f []*m.Mission
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).([]*m.Mission)
	}
	return f, ret.Error(1)
}

func (r *missionRepositoryMock) GetByID(
	tx txmanager.Tx,
	id m.ID,
) (*m.Mission, error) {
	ret := r.Called(id)
	var f *m.Mission
	if ret.Get(0) == nil {
		f = nil
	} else {
		f = ret.Get(0).(*m.Mission)
	}
	return f, ret.Error(1)
}

func (r *missionRepositoryMock) Save(
	tx txmanager.Tx,
	v *m.Mission,
) error {
	ret := r.Called(v)
	return ret.Error(0)
}

func (r *missionRepositoryMock) Delete(
	tx txmanager.Tx,
	id m.ID,
) error {
	ret := r.Called(id)
	return ret.Error(0)
}

type generatorMock struct {
	m.Generator
	id           m.ID
	versions     []m.Version
	versionIndex int
}

func (gen *generatorMock) NewID() m.ID {
	return gen.id
}
func (gen *generatorMock) NewVersion() m.Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
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

func (psm *pubSubManagerMock) SetConsumer(ctx context.Context, exchangeName, queueName string, handler event.Handler) error {
	return nil
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

// Mission構成オブジェクトモック
type missionComponentMock struct {
	ID           string
	Name         string
	Navigation   navigationComponentMock
	IsCarbonCopy bool
	Version      string
}

func (v *missionComponentMock) GetID() string {
	return v.ID
}

func (v *missionComponentMock) GetName() string {
	return v.Name
}

func (v *missionComponentMock) GetNavigation() m.NavigationComponent {
	return &v.Navigation
}

func (v *missionComponentMock) GetIsCarbonCopy() bool {
	return v.IsCarbonCopy
}

func (v *missionComponentMock) GetVersion() string {
	return v.Version
}

// Navigation構成オブジェクトモック
type navigationComponentMock struct {
	TakeoffPointGroundHeightWGS84EllipsoidM float64
	Waypoints                               []waypointComponentMock
}

func (v *navigationComponentMock) GetTakeoffPointGroundHeightWGS84EllipsoidM() float64 {
	return v.TakeoffPointGroundHeightWGS84EllipsoidM
}

func (v *navigationComponentMock) GetWaypoints() []m.WaypointComponent {
	var waypoints []m.WaypointComponent
	for _, w := range v.Waypoints {
		waypoints = append(
			waypoints,
			&waypointComponentMock{
				PointOrder:      w.PointOrder,
				LatitudeDegree:  w.LatitudeDegree,
				LongitudeDegree: w.LongitudeDegree,
				RelativeHeightM: w.RelativeHeightM,
				SpeedMS:         w.SpeedMS,
			},
		)
	}
	return waypoints
}

// Waypoint構成オブジェクトモック
type waypointComponentMock struct {
	PointOrder                                                int
	LatitudeDegree, LongitudeDegree, RelativeHeightM, SpeedMS float64
}

func (v *waypointComponentMock) GetPointOrder() int {
	return v.PointOrder
}

func (v *waypointComponentMock) GetLatitudeDegree() float64 {
	return v.LatitudeDegree
}

func (v *waypointComponentMock) GetLongitudeDegree() float64 {
	return v.LongitudeDegree
}

func (v *waypointComponentMock) GetRelativeHeightM() float64 {
	return v.RelativeHeightM
}

func (v *waypointComponentMock) GetSpeedMS() float64 {
	return v.SpeedMS
}

type missionCommandMock struct {
	Mission missionMock
}

func (v *missionCommandMock) GetID() string {
	return v.Mission.ID
}

func (v *missionCommandMock) GetMission() Mission {
	return &v.Mission
}

type missionMock struct {
	ID         string
	Name       string
	Navigation navigationMock
}

func (v *missionMock) GetID() string {
	return v.ID
}

func (v *missionMock) GetName() string {
	return v.Name
}

func (v *missionMock) GetNavigation() Navigation {
	return &v.Navigation
}

type navigationMock struct {
	TakeoffPointGroundHeight float64
	Waypoints                []waypointMock
}

func (v *navigationMock) GetTakeoffPointGroundHeight() float64 {
	return v.TakeoffPointGroundHeight
}

func (v *navigationMock) GetWaypoints() []Waypoint {
	var waypoints []Waypoint
	for _, w := range v.Waypoints {
		waypoints = append(
			waypoints,
			&waypointMock{
				Latitude:       w.Latitude,
				Longitude:      w.Longitude,
				RelativeHeight: w.RelativeHeight,
				Speed:          w.Speed,
			},
		)
	}
	return waypoints
}

type waypointMock struct {
	Latitude, Longitude, RelativeHeight, Speed float64
}

func (v *waypointMock) GetLatitude() float64 {
	return v.Latitude
}

func (v *waypointMock) GetLongitude() float64 {
	return v.Longitude
}

func (v *waypointMock) GetRelativeHeight() float64 {
	return v.RelativeHeight
}

func (v *waypointMock) GetSpeed() float64 {
	return v.Speed
}

type missionIDCommandMock struct {
	ID string
}

func (f *missionIDCommandMock) GetID() string {
	return f.ID
}

type carbonCopyCommandMock struct {
	OriginalID string
	NewID      string
}

func (f *carbonCopyCommandMock) GetOriginalID() string {
	return f.OriginalID
}

func (f *carbonCopyCommandMock) GetNewID() string {
	return f.NewID
}
