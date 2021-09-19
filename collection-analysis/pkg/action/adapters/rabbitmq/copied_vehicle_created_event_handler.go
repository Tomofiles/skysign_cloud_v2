package rabbitmq

import (
	"collection-analysis/pkg/action/app"
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// CopiedVehicleCreatedEventExchangeName .
	CopiedVehicleCreatedEventExchangeName = "vehicle.copied_vehicle_created_event"
	// CopiedVehicleCreatedEventQueueName .
	CopiedVehicleCreatedEventQueueName = "action.copied_vehicle_created_event"
)

// CopiedVehicleCreatedEventHandler .
type CopiedVehicleCreatedEventHandler interface {
	HandleCopiedVehicleCreatedEvent(ctx context.Context, event []byte) error
}

type copiedVehicleCreatedEventHandler struct {
	app app.Application
}

// NewCopiedVehicleCreatedEventHandler .
func NewCopiedVehicleCreatedEventHandler(application app.Application) CopiedVehicleCreatedEventHandler {
	return &copiedVehicleCreatedEventHandler{app: application}
}

// HandleCopiedVehicleCreatedEvent .
func (h *copiedVehicleCreatedEventHandler) HandleCopiedVehicleCreatedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.CopiedVehicleCreatedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", CopiedVehicleCreatedEventQueueName, eventPb.String())

	command := createCommandHolder{event: &eventPb}
	if ret := h.app.Services.ManageAction.CreateAction(&command); ret != nil {
		return ret
	}
	return nil
}

type createCommandHolder struct {
	event *skysign_proto.CopiedVehicleCreatedEvent
}

func (h *createCommandHolder) GetID() string {
	return h.event.VehicleId
}

func (h *createCommandHolder) GetCommunicationID() string {
	return h.event.CommunicationId
}

func (h *createCommandHolder) GetFleetID() string {
	return h.event.FleetId
}
