package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FleetCopiedEventExchangeName .
	FleetCopiedEventExchangeName = "flightoperation.fleet_copied_event"
	// FleetCopiedEventQueueName .
	FleetCopiedEventQueueName = "fleet.fleet_copied_event"
)

// FleetCopiedEventHandler .
type FleetCopiedEventHandler interface {
	HandleFleetCopiedEvent(ctx context.Context, event []byte) error
}

// fleetCopiedEventHandler .
type fleetCopiedEventHandler struct {
	app app.Application
}

// NewFleetCopiedEventHandler .
func NewFleetCopiedEventHandler(application app.Application) FleetCopiedEventHandler {
	return &fleetCopiedEventHandler{app: application}
}

// HandleFleetCopiedEvent .
func (h *fleetCopiedEventHandler) HandleFleetCopiedEvent(
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
