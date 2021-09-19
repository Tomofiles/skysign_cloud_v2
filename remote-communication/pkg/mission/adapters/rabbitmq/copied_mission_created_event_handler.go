package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/app"
	"github.com/Tomofiles/skysign_cloud_v2/remote-communication/pkg/mission/service"

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

// CopiedMissionCreatedEventHandler .
type CopiedMissionCreatedEventHandler interface {
	HandleCopiedMissionCreatedEvent(ctx context.Context, event []byte) error
}

// copiedMissionCreatedEventHandler .
type copiedMissionCreatedEventHandler struct {
	app app.Application
}

// NewCopiedMissionCreatedEventHandler .
func NewCopiedMissionCreatedEventHandler(application app.Application) CopiedMissionCreatedEventHandler {
	return &copiedMissionCreatedEventHandler{app: application}
}

// HandleCommunicationIDGaveEvent .
func (h *copiedMissionCreatedEventHandler) HandleCopiedMissionCreatedEvent(
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
				LatitudeDegree:    w.Latitude,
				LongitudeDegree:   w.Longitude,
				RelativeAltitudeM: w.RelativeAltitude,
				SpeedMS:           w.Speed,
			},
		)
	}
	return waypoints
}

type waypoint struct {
	LatitudeDegree, LongitudeDegree, RelativeAltitudeM, SpeedMS float64
}

func (v *waypoint) GetLatitudeDegree() float64 {
	return v.LatitudeDegree
}

func (v *waypoint) GetLongitudeDegree() float64 {
	return v.LongitudeDegree
}

func (v *waypoint) GetRelativeAltitudeM() float64 {
	return v.RelativeAltitudeM
}

func (v *waypoint) GetSpeedMS() float64 {
	return v.SpeedMS
}
