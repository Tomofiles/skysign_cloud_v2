package ports

import (
	"context"
	"mission/pkg/mission/app"
	"mission/pkg/skysign_proto"

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
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleMissionCopiedEvent .
func (h *EventHandler) HandleMissionCopiedEvent(
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
