package ports

import (
	"context"
	"fleet-formation/pkg/mission/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// MissionCopiedEventExchangeName .
	MissionCopiedEventExchangeName = "fleet.mission_copied_event"
	// MissionCopiedEventQueueName .
	MissionCopiedEventQueueName = "mission.mission_copied_event"
)

// EventHandler .
type EventHandler interface {
	HandleMissionCopiedEvent(ctx context.Context, event []byte) error
}

// eventHandler .
type eventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) *eventHandler {
	return &eventHandler{app: application}
}

// HandleMissionCopiedEvent .
func (h *eventHandler) HandleMissionCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.MissionCopiedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", MissionCopiedEventQueueName, eventPb.String())

	command := &copyCommand{
		originalID: eventPb.GetOriginalMissionId(),
		newID:      eventPb.GetNewMissionId(),
	}
	if ret := h.app.Services.ManageMission.CarbonCopyMission(
		command,
		func(uploadID string) { /* 処理なし */ },
	); ret != nil {
		return ret
	}
	return nil
}

type copyCommand struct {
	originalID string
	newID      string
}

func (h *copyCommand) GetOriginalID() string {
	return h.originalID
}
func (h *copyCommand) GetNewID() string {
	return h.newID
}
