package ports

import (
	"context"
	"fleet-formation/pkg/fleet/app"
	"fleet-formation/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FleetIDGaveEventExchangeName .
	FleetIDGaveEventExchangeName = "flightplan.fleet_id_gave_event"
	// FleetIDGaveEventQueueName .
	FleetIDGaveEventQueueName = "fleet.fleet_id_gave_event"
	// FleetIDRemovedEventExchangeName .
	FleetIDRemovedEventExchangeName = "flightplan.fleet_id_removed_event"
	// FleetIDRemovedEventQueueName .
	FleetIDRemovedEventQueueName = "fleet.fleet_id_removed_event"
	// FleetCopiedEventExchangeName .
	FleetCopiedEventExchangeName = "flightoperation.fleet_copied_event"
	// FleetCopiedEventQueueName .
	FleetCopiedEventQueueName = "fleet.fleet_copied_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleFleetIDGaveEvent .
func (h *EventHandler) HandleFleetIDGaveEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FleetIDGaveEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FleetIDGaveEventQueueName, eventPb.String())

	command := createFleetCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageFleet.CreateFleet(&command); ret != nil {
		return ret
	}
	return nil
}

// HandleFleetIDRemovedEvent .
func (h *EventHandler) HandleFleetIDRemovedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FleetIDRemovedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FleetIDRemovedEventQueueName, eventPb.String())

	command := deleteFleetCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageFleet.DeleteFleet(&command); ret != nil {
		return ret
	}
	return nil
}

// HandleFleetCopiedEvent .
func (h *EventHandler) HandleFleetCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FleetCopiedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FleetCopiedEventQueueName, eventPb.String())

	command := copyCommandHolder{
		originalID: eventPb.GetOriginalFleetId(),
		newID:      eventPb.GetNewFleetId(),
	}
	if ret := h.app.Services.ManageFleet.CarbonCopyFleet(&command); ret != nil {
		return ret
	}
	return nil
}

type createFleetCommandHolder struct {
	event *skysign_proto.FleetIDGaveEvent
}

func (h *createFleetCommandHolder) GetID() string {
	return h.event.FleetId
}

func (h *createFleetCommandHolder) GetNumberOfVehicles() int {
	return int(h.event.NumberOfVehicles)
}

type deleteFleetCommandHolder struct {
	event *skysign_proto.FleetIDRemovedEvent
}

func (h *deleteFleetCommandHolder) GetID() string {
	return h.event.FleetId
}

type copyCommandHolder struct {
	originalID string
	newID      string
}

func (h *copyCommandHolder) GetOriginalID() string {
	return h.originalID
}
func (h *copyCommandHolder) GetNewID() string {
	return h.newID
}
