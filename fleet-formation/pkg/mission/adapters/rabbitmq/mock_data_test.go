package rabbitmq

import (
	"context"
	"fleet-formation/pkg/mission/domain/mission"
	m "fleet-formation/pkg/mission/domain/mission"
	"fleet-formation/pkg/mission/service"

	crm "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/rabbitmq"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = mission.ID("mission-id")
const DefaultMissionVersion = mission.Version("version")
const DefaultMissionName = "mission-name"
const DefaultMissionTakeoffPointGroundAltitudeM float64 = 10
const DefaultMissionUploadID = mission.UploadID("upload-id")

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
	return ret.Error(0)
}

func (s *manageMissionServiceMock) ListMissions(
	retrievedModel service.RetrievedModel,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageMissionServiceMock) CreateMission(
	command service.CreateMissionCommand,
	createdID service.CreatedID,
	uploadID service.UploadID,
) error {
	ret := s.Called()
	return ret.Error(0)
}

func (s *manageMissionServiceMock) UpdateMission(
	command service.UpdateMissionCommand,
	uploadID service.UploadID,
) error {
	ret := s.Called()
	return ret.Error(0)
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
		func(pointOrder int, latitudeDegree, longitudeDegree, relativeAltitudeM, speedMS float64) {
			waypoints = append(
				waypoints,
				waypointMock{
					latitudeDegree:    latitudeDegree,
					longitudeDegree:   longitudeDegree,
					relativeAltitudeM: relativeAltitudeM,
					speedMS:           speedMS,
				},
			)
		},
	)
	navigation := &navigationMock{
		takeoffPointGroundAltitudeM: f.mission.GetNavigation().GetTakeoffPointGroundAltitudeM(),
		waypoints:                   waypoints,
		uploadID:                    string(f.mission.GetNavigation().GetUploadID()),
	}
	return navigation
}

type navigationMock struct {
	takeoffPointGroundAltitudeM float64
	waypoints                   []waypointMock
	uploadID                    string
}

func (f *navigationMock) GetTakeoffPointGroundAltitudeM() float64 {
	return f.takeoffPointGroundAltitudeM
}

func (f *navigationMock) GetWaypoints() []service.Waypoint {
	waypoints := []service.Waypoint{}
	for _, w := range f.waypoints {
		waypoints = append(
			waypoints,
			&waypointMock{
				latitudeDegree:    w.latitudeDegree,
				longitudeDegree:   w.longitudeDegree,
				relativeAltitudeM: w.relativeAltitudeM,
				speedMS:           w.speedMS,
			},
		)
	}
	return waypoints
}

func (v *navigationMock) GetUploadID() string {
	return v.uploadID
}

type waypointMock struct {
	latitudeDegree    float64
	longitudeDegree   float64
	relativeAltitudeM float64
	speedMS           float64
}

func (f *waypointMock) GetLatitudeDegree() float64 {
	return f.latitudeDegree
}

func (f *waypointMock) GetLongitudeDegree() float64 {
	return f.longitudeDegree
}

func (f *waypointMock) GetRelativeAltitudeM() float64 {
	return f.relativeAltitudeM
}

func (f *waypointMock) GetSpeedMS() float64 {
	return f.speedMS
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

func (v *missionComponentMock) GetNavigation() mission.NavigationComponent {
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
	TakeoffPointGroundAltitudeM float64
	Waypoints                   []waypointComponentMock
	UploadID                    string
}

func (v *navigationComponentMock) GetTakeoffPointGroundAltitudeM() float64 {
	return v.TakeoffPointGroundAltitudeM
}

func (v *navigationComponentMock) GetWaypoints() []mission.WaypointComponent {
	var waypoints []mission.WaypointComponent
	for _, w := range v.Waypoints {
		waypoints = append(
			waypoints,
			&waypointComponentMock{
				PointOrder:        w.PointOrder,
				LatitudeDegree:    w.LatitudeDegree,
				LongitudeDegree:   w.LongitudeDegree,
				RelativeAltitudeM: w.RelativeAltitudeM,
				SpeedMS:           w.SpeedMS,
			},
		)
	}
	return waypoints
}

// GetUploadID .
func (v *navigationComponentMock) GetUploadID() string {
	return v.UploadID
}

// Waypoint構成オブジェクトモック
type waypointComponentMock struct {
	PointOrder                                                  int
	LatitudeDegree, LongitudeDegree, RelativeAltitudeM, SpeedMS float64
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

func (v *waypointComponentMock) GetRelativeAltitudeM() float64 {
	return v.RelativeAltitudeM
}

func (v *waypointComponentMock) GetSpeedMS() float64 {
	return v.SpeedMS
}

type channelMockPublish struct {
	mock.Mock
	message          crm.Message
	messageCallCount int
	isClose          bool
}

func (ch *channelMockPublish) FanoutExchangeDeclare(exchange string) error {
	ret := ch.Called()
	return ret.Error(0)
}

func (ch *channelMockPublish) QueueDeclareAndBind(exchange, queue string) error {
	panic("implement me")
}

func (ch *channelMockPublish) Publish(queue string, message crm.Message) error {
	ret := ch.Called()
	ch.message = message
	ch.messageCallCount++
	return ret.Error(0)
}

func (ch *channelMockPublish) Consume(ctx context.Context, queue string) (<-chan crm.Message, error) {
	panic("implement me")
}

func (ch *channelMockPublish) Close() error {
	panic("implement me")
}

type pubSubManagerMock struct {
	consumers       []consumer
	publishHandlers []func(ch crm.Channel, e interface{})
}

func (h *pubSubManagerMock) SetConsumer(ctx context.Context, exchangeName, queueName string, handler func([]byte)) error {
	h.consumers = append(
		h.consumers,
		consumer{
			exchangeName: exchangeName,
			queueName:    queueName,
			handler:      handler,
		})
	return nil
}

func (h *pubSubManagerMock) SetPublishHandler(handler func(ch crm.Channel, e interface{})) error {
	h.publishHandlers = append(h.publishHandlers, handler)
	return nil
}

type consumer struct {
	exchangeName, queueName string
	handler                 func([]byte)
}
