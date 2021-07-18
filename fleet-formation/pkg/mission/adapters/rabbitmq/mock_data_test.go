package rabbitmq

import (
	"context"
	crm "fleet-formation/pkg/common/adapters/rabbitmq"
	"fleet-formation/pkg/mission/domain/mission"

	"github.com/stretchr/testify/mock"
)

const DefaultMissionID = mission.ID("mission-id")
const DefaultMissionVersion = mission.Version("version")
const DefaultMissionName = "mission-name"
const DefaultMissionTakeoffPointGroundHeightWGS84EllipsoidM float64 = 10
const DefaultMissionUploadID = mission.UploadID("upload-id")

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
	TakeoffPointGroundHeightWGS84EllipsoidM float64
	Waypoints                               []waypointComponentMock
	UploadID                                string
}

func (v *navigationComponentMock) GetTakeoffPointGroundHeightWGS84EllipsoidM() float64 {
	return v.TakeoffPointGroundHeightWGS84EllipsoidM
}

func (v *navigationComponentMock) GetWaypoints() []mission.WaypointComponent {
	var waypoints []mission.WaypointComponent
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

// GetUploadID .
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

type publishHandlerMock struct {
	publishHandlers []func(ch crm.Channel, e interface{})
}

func (h *publishHandlerMock) SetPublishHandler(handler func(ch crm.Channel, e interface{})) error {
	h.publishHandlers = append(h.publishHandlers, handler)
	return nil
}
