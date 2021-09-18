package rabbitmq

import (
	"context"
	"fleet-formation/pkg/fleet/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FleetIDRemovedEventExchangeName .
	FleetIDRemovedEventExchangeName = "flightplan.fleet_id_removed_event"
	// FleetIDRemovedEventQueueName .
	FleetIDRemovedEventQueueName = "fleet.fleet_id_removed_event"
)

// FleetIDRemovedEventHandler .
type FleetIDRemovedEventHandler interface {
	HandleFleetIDRemovedEvent(ctx context.Context, event []byte) error
}

// fleetIDRemovedEventHandler .
type fleetIDRemovedEventHandler struct {
	app app.Application
}

// NewFleetIDRemovedEventHandler .
func NewFleetIDRemovedEventHandler(application app.Application) *fleetIDRemovedEventHandler {
	return &fleetIDRemovedEventHandler{app: application}
}

// HandleFleetIDRemovedEvent .
func (h *fleetIDRemovedEventHandler) HandleFleetIDRemovedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FleetIdRemovedEvent{}
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

type deleteFleetCommandHolder struct {
	event *skysign_proto.FleetIdRemovedEvent
}

func (h *deleteFleetCommandHolder) GetID() string {
	return h.event.FleetId
}
