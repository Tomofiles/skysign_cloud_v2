package ports

import (
	"context"
	"remote-communication/pkg/mission/app"
	"remote-communication/pkg/mission/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// CopiedMissionCreatedEventExchangeName .
	CopiedMissionCreatedEventExchangeName = "mission.copied_mission_created_event"
	// CopiedMissionCreatedEventQueueName .
	CopiedMissionCreatedEventQueueName = "uploadmission.copied_mission_created_event"
)

// EventHandler .
type EventHandler interface {
	HandleCopiedMissionCreatedEvent(ctx context.Context, event []byte) error
}

// eventHandler .
type eventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) *eventHandler {
	return &eventHandler{app: application}
}

// HandleCommunicationIDGaveEvent .
func (h *eventHandler) HandleCopiedMissionCreatedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.CopiedMissionCreatedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", CopiedMissionCreatedEventQueueName, eventPb.String())

	command := createMissionCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageMission.CreateMission(&command); ret != nil {
		return ret
	}
	return nil
}

type createMissionCommandHolder struct {
	event *skysign_proto.CopiedMissionCreatedEvent
}

func (h *createMissionCommandHolder) GetID() string {
	return h.event.Mission.Navigation.UploadId
}

func (h *createMissionCommandHolder) GetWaypoints() []service.Waypoint {
	var waypoints []service.Waypoint
	for _, w := range h.event.Mission.Navigation.Waypoints {
		waypoints = append(
			waypoints,
			&waypoint{
				Latitude:         w.Latitude,
				Longitude:        w.Longitude,
				RelativeAltitude: w.RelativeHeight,
				Speed:            w.Speed,
			},
		)
	}
	return waypoints
}

type waypoint struct {
	Latitude, Longitude, RelativeAltitude, Speed float64
}

func (v *waypoint) GetLatitude() float64 {
	return v.Latitude
}

func (v *waypoint) GetLongitude() float64 {
	return v.Longitude
}

func (v *waypoint) GetRelativeAltitude() float64 {
	return v.RelativeAltitude
}

func (v *waypoint) GetSpeed() float64 {
	return v.Speed
}
