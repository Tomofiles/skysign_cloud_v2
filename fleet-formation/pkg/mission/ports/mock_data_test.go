package ports

import (
	m "fleet-formation/pkg/mission/domain/mission"
	"fleet-formation/pkg/mission/service"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = "mission-id"
const DefaultMissionName = "mission-name"
const DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM float64 = 10
const DefaultMissionVersion = m.Version("version")
const DefaultMissionUploadID = m.UploadID("upload-id")

type manageMissionServiceMock struct {
	mock.Mock
	OriginalID string
	NewID      string
}

func (s *manageMissionServiceMock) GetMission(
	command service.GetMissionCommand,
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.MissionPresentationModel)
		retrievedModel(f)
	}
	return ret.Error(1)
}

func (s *manageMissionServiceMock) ListMissions(
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	if models := ret.Get(0); models != nil {
		for _, f := range models.([]service.MissionPresentationModel) {
			retrievedModel(f)
		}
	}
	return ret.Error(1)
}

func (s *manageMissionServiceMock) CreateMission(
	command service.CreateMissionCommand,
	createdID service.CreatedID,
	uploadID service.UploadID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.MissionPresentationModel)
		createdID(f.GetMission().GetID())
		uploadID(f.GetMission().GetNavigation().GetUploadID())
	}
	return ret.Error(1)
}

func (s *manageMissionServiceMock) UpdateMission(
	command service.UpdateMissionCommand,
	uploadID service.UploadID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.MissionPresentationModel)
		uploadID(f.GetMission().GetNavigation().GetUploadID())
	}
	return ret.Error(1)
}

func (s *manageMissionServiceMock) DeleteMission(
	command service.DeleteMissionCommand,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageMissionServiceMock) CarbonCopyMission(
	command service.CarbonCopyMissionCommand,
	uploadID service.UploadID,
) error {
	ret := s.Called()
	if model := ret.Get(0); model != nil {
		f := model.(service.MissionPresentationModel)
		uploadID(f.GetMission().GetNavigation().GetUploadID())
	}
	s.OriginalID = command.GetOriginalID()
	s.NewID = command.GetNewID()
	return ret.Error(1)
}

type missionModelMock struct {
	mission *m.Mission
}

func (f *missionModelMock) GetMission() service.Mission {
	return &missionMock{
		mission: f.mission,
	}
}

type missionMock struct {
	mission *m.Mission
}

func (f *missionMock) GetID() string {
	return string(f.mission.GetID())
}

func (f *missionMock) GetName() string {
	return f.mission.GetName()
}

func (f *missionMock) GetNavigation() service.Navigation {
	waypoints := []waypointMock{}
	f.mission.GetNavigation().ProvideWaypointsInterest(
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeHeightM, speedMS float64) {
			waypoints = append(
				waypoints,
				waypointMock{
					latitude:       latitudeDegree,
					longitude:      longitudeDegree,
					relativeHeight: relativeHeightM,
					speed:          speedMS,
				},
			)
		},
	)
	navigation := &navigationMock{
		takeoffPointGroundHeight: f.mission.GetNavigation().GetTakeoffPointGroundHeightWGS84EllipsoidM(),
		waypoints:                waypoints,
		uploadID:                 string(f.mission.GetNavigation().GetUploadID()),
	}
	return navigation
}

type navigationMock struct {
	takeoffPointGroundHeight float64
	waypoints                []waypointMock
	uploadID                 string
}

func (f *navigationMock) GetTakeoffPointGroundHeight() float64 {
	return f.takeoffPointGroundHeight
}

func (f *navigationMock) GetWaypoints() []service.Waypoint {
	waypoints := []service.Waypoint{}
	for _, w := range f.waypoints {
		waypoints = append(
			waypoints,
			&waypointMock{
				latitude:       w.latitude,
				longitude:      w.longitude,
				relativeHeight: w.relativeHeight,
				speed:          w.speed,
			},
		)
	}
	return waypoints
}

func (v *navigationMock) GetUploadID() string {
	return v.uploadID
}

type waypointMock struct {
	latitude       float64
	longitude      float64
	relativeHeight float64
	speed          float64
}

func (f *waypointMock) GetLatitude() float64 {
	return f.latitude
}

func (f *waypointMock) GetLongitude() float64 {
	return f.longitude
}

func (f *waypointMock) GetRelativeHeight() float64 {
	return f.relativeHeight
}

func (f *waypointMock) GetSpeed() float64 {
	return f.speed
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
	UploadID                                string
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

func (v *navigationComponentMock) GetUploadID() string {
	return v.UploadID
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
