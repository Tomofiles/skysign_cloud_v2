package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/adapters/proto"
	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/vehicle/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	proto_lib "google.golang.org/protobuf/proto"
)

const (
	// VehicleCopiedEventExchangeName .
	VehicleCopiedEventExchangeName = "fleet.vehicle_copied_event"
	// VehicleCopiedEventQueueName .
	VehicleCopiedEventQueueName = "vehicle.vehicle_copied_event"
)

// VehicleCopiedEventHandler .
type VehicleCopiedEventHandler interface {
	HandleVehicleCopiedEvent(ctx context.Context, event []byte) error
}

// vehicleCopiedEventHandler .
type vehicleCopiedEventHandler struct {
	app app.Application
}

// NewVehicleCopiedEventHandler .
func NewVehicleCopiedEventHandler(application app.Application) VehicleCopiedEventHandler {
	return &vehicleCopiedEventHandler{app: application}
}

// HandleVehicleCopiedEvent .
func (h *vehicleCopiedEventHandler) HandleVehicleCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.VehicleCopiedEvent{}
	if err := proto_lib.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", VehicleCopiedEventQueueName, eventPb.String())

	if ret := proto.ValidateVehicleCopiedEvent(&eventPb); ret != nil {
		return ret
	}

	command := copyCommand{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageVehicle.CarbonCopyVehicle(&command); ret != nil {
		return ret
	}
	return nil
}

type copyCommand struct {
	event *skysign_proto.VehicleCopiedEvent
}

func (h *copyCommand) GetOriginalID() string {
	return h.event.GetOriginalVehicleId()
}
func (h *copyCommand) GetNewID() string {
	return h.event.GetNewVehicleId()
}
func (h *copyCommand) GetFleetID() string {
	return h.event.GetFleetId()
}
