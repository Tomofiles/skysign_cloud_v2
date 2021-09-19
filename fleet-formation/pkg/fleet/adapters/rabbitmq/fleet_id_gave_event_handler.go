package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/fleet-formation/pkg/fleet/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FleetIDGaveEventExchangeName .
	FleetIDGaveEventExchangeName = "flightplan.fleet_id_gave_event"
	// FleetIDGaveEventQueueName .
	FleetIDGaveEventQueueName = "fleet.fleet_id_gave_event"
)

// FleetIDGaveEventHandler .
type FleetIDGaveEventHandler interface {
	HandleFleetIDGaveEvent(ctx context.Context, event []byte) error
}

// fleetIDGaveEventHandler .
type fleetIDGaveEventHandler struct {
	app app.Application
}

// NewFleetIDGaveEventHandler .
func NewFleetIDGaveEventHandler(application app.Application) FleetIDGaveEventHandler {
	return &fleetIDGaveEventHandler{app: application}
}

// HandleFleetIDGaveEvent .
func (h *fleetIDGaveEventHandler) HandleFleetIDGaveEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FleetIdGaveEvent{}
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

type createFleetCommandHolder struct {
	event *skysign_proto.FleetIdGaveEvent
}

func (h *createFleetCommandHolder) GetID() string {
	return h.event.FleetId
}

func (h *createFleetCommandHolder) GetNumberOfVehicles() int {
	return int(h.event.NumberOfVehicles)
}
