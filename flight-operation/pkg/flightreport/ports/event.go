package ports

import (
	"context"
	"flight-operation/pkg/flightreport/app"
	"flight-operation/pkg/flightreport/service"
	"flight-operation/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FlightoperationCompletedEventExchangeName .
	FlightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"
	// FlightoperationCompletedEventQueueName .
	FlightoperationCompletedEventQueueName = "flightreport.flightoperation_completed_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleFlightoperationCompletedEvent .
func (h *EventHandler) HandleFlightoperationCompletedEvent(
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
