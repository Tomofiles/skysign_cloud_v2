package ports

import (
	"context"
	"vehicle/pkg/skysign_proto"
	"vehicle/pkg/vehicle/app"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// VehicleCopiedWhenFlightplanCopiedEventExchangeName .
	VehicleCopiedWhenFlightplanCopiedEventExchangeName = "fleet.vehicle_copied_when_flightplan_copied_event"
	// VehicleCopiedWhenFlightplanCopiedEventQueueName .
	VehicleCopiedWhenFlightplanCopiedEventQueueName = "vehicle.vehicle_copied_when_flightplan_copied_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleVehicleCopiedWhenFlightplanCopiedEvent .
func (h *EventHandler) HandleVehicleCopiedWhenFlightplanCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.VehicleCopiedWhenFlightplanCopiedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", VehicleCopiedWhenFlightplanCopiedEventQueueName, eventPb.String())

	requestDpo := copyRequestDpoHolder{
		originalID:   eventPb.GetOriginalVehicleId(),
		newID:        eventPb.GetNewVehicleId(),
		flightplanID: eventPb.GetFlightplanId(),
	}
	if ret := h.app.Services.ManageVehicle.CarbonCopyVehicle(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

type copyRequestDpoHolder struct {
	originalID   string
	newID        string
	flightplanID string
}

func (h *copyRequestDpoHolder) GetOriginalID() string {
	return h.originalID
}
func (h *copyRequestDpoHolder) GetNewID() string {
	return h.newID
}
func (h *copyRequestDpoHolder) GetFlightplanID() string {
	return h.flightplanID
}
