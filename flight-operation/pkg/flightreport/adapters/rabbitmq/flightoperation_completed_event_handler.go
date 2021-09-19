package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/app"
	"github.com/Tomofiles/skysign_cloud_v2/flight-operation/pkg/flightreport/service"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FlightoperationCompletedEventExchangeName .
	FlightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"
	// FlightoperationCompletedEventQueueName .
	FlightoperationCompletedEventQueueName = "flightreport.flightoperation_completed_event"
)

// FlightoperationCompletedEventHandler .
type FlightoperationCompletedEventHandler interface {
	HandleFlightoperationCompletedEvent(ctx context.Context, event []byte) error
}

// flightoperationCompletedEventHandler .
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

	command := createFlightreportCommandHolder{
		event: &eventPb,
	}
	if ret := h.app.Services.ManageFlightreport.CreateFlightreport(&command); ret != nil {
		return ret
	}
	return nil
}

type createFlightreportCommandHolder struct {
	event *skysign_proto.FlightoperationCompletedEvent
}

func (h *createFlightreportCommandHolder) GetFlightreport() service.Flightreport {
	return &flightreport{
		event: h.event,
	}
}

type flightreport struct {
	event *skysign_proto.FlightoperationCompletedEvent
}

// GetID .
func (f *flightreport) GetID() string {
	return f.event.Flightoperation.Id
}

// GetName .
func (f *flightreport) GetName() string {
	return f.event.Flightoperation.Name
}

// GetDescription .
func (f *flightreport) GetDescription() string {
	return f.event.Flightoperation.Description
}

// GetFleetID .
func (f *flightreport) GetFleetID() string {
	return f.event.Flightoperation.FleetId
}
