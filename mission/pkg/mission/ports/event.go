package ports

import (
	"context"
	"mission/pkg/mission/app"
	"mission/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// MissionCopiedWhenFlightplanCopiedEventExchangeName .
	MissionCopiedWhenFlightplanCopiedEventExchangeName = "fleet.mission_copied_when_flightplan_copied_event"
	// MissionCopiedWhenFlightplanCopiedEventQueueName .
	MissionCopiedWhenFlightplanCopiedEventQueueName = "mission.mission_copied_when_flightplan_copied_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleMissionCopiedWhenFlightplanCopiedEvent .
func (h *EventHandler) HandleMissionCopiedWhenFlightplanCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.MissionCopiedWhenFlightplanCopiedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", MissionCopiedWhenFlightplanCopiedEventQueueName, eventPb.String())

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
