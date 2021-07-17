package ports

import (
	"context"
	"fleet-formation/pkg/skysign_proto"
	"fleet-formation/pkg/vehicle/app"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// VehicleCopiedEventExchangeName .
	VehicleCopiedEventExchangeName = "fleet.vehicle_copied_event"
	// VehicleCopiedEventQueueName .
	VehicleCopiedEventQueueName = "vehicle.vehicle_copied_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleVehicleCopiedEvent .
func (h *EventHandler) HandleVehicleCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.VehicleCopiedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", VehicleCopiedEventQueueName, eventPb.String())

	requestDpo := copyRequestDpoHolder{
		originalID: eventPb.GetOriginalVehicleId(),
		newID:      eventPb.GetNewVehicleId(),
		fleetID:    eventPb.GetFleetId(),
	}
	if ret := h.app.Services.ManageVehicle.CarbonCopyVehicle(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

type copyRequestDpoHolder struct {
	originalID string
	newID      string
	fleetID    string
}

func (h *copyRequestDpoHolder) GetOriginalID() string {
	return h.originalID
}
func (h *copyRequestDpoHolder) GetNewID() string {
	return h.newID
}
func (h *copyRequestDpoHolder) GetFleetID() string {
	return h.fleetID
}
