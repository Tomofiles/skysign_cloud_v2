package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FlightoperationCompletedEventExchangeName .
	FlightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"
	// FlightoperationCompletedEventQueueName .
	FlightoperationCompletedEventQueueName = "action.flightoperation_completed_event"
)

// FlightoperationCompletedEventHandler .
type FlightoperationCompletedEventHandler interface {
	HandleFlightoperationCompletedEvent(ctx context.Context, event []byte) error
}

type flightoperationCompletedEventHandler struct {
	app app.Application
}

// NewFlightoperationCompletedEventHandler .
func NewFlightoperationCompletedEventHandler(application app.Application) FlightoperationCompletedEventHandler {
	return &flightoperationCompletedEventHandler{app: application}
}

// HandleFlightoperationCompletedEvent .
func (h *flightoperationCompletedEventHandler) HandleFlightoperationCompletedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FlightoperationCompletedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightoperationCompletedEventQueueName, eventPb.String())

	command := completeCommandHolder{event: &eventPb}
	if ret := h.app.Services.OperateAction.CompleteAction(&command); ret != nil {
		return ret
	}
	return nil
}

type completeCommandHolder struct {
	event *skysign_proto.FlightoperationCompletedEvent
}

func (h *completeCommandHolder) GetFleetID() string {
	return h.event.Flightoperation.FleetId
}
