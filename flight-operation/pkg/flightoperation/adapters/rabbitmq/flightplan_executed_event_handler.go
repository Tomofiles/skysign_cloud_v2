package rabbitmq

import (
	"context"
	"flight-operation/pkg/flightoperation/app"
	"flight-operation/pkg/flightoperation/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FlightplanExecutedEventExchangeName .
	FlightplanExecutedEventExchangeName = "flightplan.flightplan_executed_event"
	// FlightplanExecutedEventQueueName .
	FlightplanExecutedEventQueueName = "flightoperation.flightplan_executed_event"
)

// FlightplanExecutedEventHandler .
type FlightplanExecutedEventHandler interface {
	HandleFlightplanExecutedEvent(ctx context.Context, event []byte) error
}

type flightplanExecutedEventHandler struct {
	app app.Application
}

// NewFlightplanExecutedEventHandler .
func NewFlightplanExecutedEventHandler(application app.Application) FlightplanExecutedEventHandler {
	return &flightplanExecutedEventHandler{app: application}
}

// HandleFlightplanExecutedEvent .
func (h *flightplanExecutedEventHandler) HandleFlightplanExecutedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FlightplanExecutedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightplanExecutedEventQueueName, eventPb.String())

	command := createFlightoperationCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageFlightoperation.CreateFlightoperation(&command); ret != nil {
		return ret
	}
	return nil
}

type createFlightoperationCommandHolder struct {
	event *skysign_proto.FlightplanExecutedEvent
}

func (h *createFlightoperationCommandHolder) GetFlightoperation() service.Flightoperation {
	return &flightoperation{
		event: h.event,
	}
}

type flightoperation struct {
	event *skysign_proto.FlightplanExecutedEvent
}

// GetID .
func (f *flightoperation) GetID() string {
	return f.event.Flightplan.Id
}

// GetName .
func (f *flightoperation) GetName() string {
	return f.event.Flightplan.Name
}

// GetDescription .
func (f *flightoperation) GetDescription() string {
	return f.event.Flightplan.Description
}

// GetFleetID .
func (f *flightoperation) GetFleetID() string {
	return f.event.Flightplan.FleetId
}
