package mission

import "errors"

const DefaultID = ID("mission-id")
const DefaultVersion = Version("version")
const DefaultName = "mission-name"
const DefaultTakeoffPointGroundHeightWGS84EllipsoidM float64 = 10

var (
	ErrSave   = errors.New("save error")
	ErrGet    = errors.New("get error")
	ErrDelete = errors.New("delete error")
)

// Mission用汎用ジェネレータモック
type generatorMock struct {
	Generator
	id           ID
	versions     []Version
	versionIndex int
}

func (gen *generatorMock) NewID() ID {
	return gen.id
}
func (gen *generatorMock) NewVersion() Version {
	version := gen.versions[gen.versionIndex]
	gen.versionIndex++
	return version
}

// Mission用汎用パブリッシャモック
type publisherMock struct {
	events []interface{}
}

func (rm *publisherMock) Publish(event interface{}) {
	rm.events = append(rm.events, event)
}

func (rm *publisherMock) Flush() error {
	return nil
}

// Mission構成オブジェクトモック
type missionComponentMock struct {
	id           string
	name         string
	navigation   navigationComponentMock
	isCarbonCopy bool
	version      string
}

func (v *missionComponentMock) GetID() string {
	return v.id
}

func (v *missionComponentMock) GetName() string {
	return v.name
}

func (v *missionComponentMock) GetNavigation() NavigationComponent {
	return &v.navigation
}

func (v *missionComponentMock) GetIsCarbonCopy() bool {
	return v.isCarbonCopy
}

func (v *missionComponentMock) GetVersion() string {
	return v.version
}

// Navigation構成オブジェクトモック
type navigationComponentMock struct {
	takeoffPointGroundHeightWGS84EllipsoidM float64
	waypoints                               []waypointComponentMock
}

func (v *navigationComponentMock) GetTakeoffPointGroundHeightWGS84EllipsoidM() float64 {
	return v.takeoffPointGroundHeightWGS84EllipsoidM
}

func (v *navigationComponentMock) GetWaypoints() []WaypointComponent {
	var waypoints []WaypointComponent
	for _, w := range v.waypoints {
		waypoints = append(
			waypoints,
			&waypointComponentMock{
				order:           w.order,
				latitudeDegree:  w.latitudeDegree,
				longitudeDegree: w.longitudeDegree,
				relativeHeightM: w.relativeHeightM,
				speedMS:         w.speedMS,
			},
		)
	}
	return waypoints
}

// Waypoint構成オブジェクトモック
type waypointComponentMock struct {
	order                                                     int
	latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64
}

func (v *waypointComponentMock) GetOrder() int {
	return v.order
}

func (v *waypointComponentMock) GetLatitudeDegree() float64 {
	return v.latitudeDegree
}

func (v *waypointComponentMock) GetLongitudeDegree() float64 {
	return v.longitudeDegree
}

func (v *waypointComponentMock) GetRelativeHeightM() float64 {
	return v.relativeHeightM
}

func (v *waypointComponentMock) GetSpeedMS() float64 {
	return v.speedMS
}
