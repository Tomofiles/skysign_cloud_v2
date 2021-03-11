package ports

import (
	"context"
	"flightreport/pkg/flightreport/app"
	"flightreport/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FlightoperationCompletedEventExchangeName .
	FlightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"
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

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightoperationCompletedEventExchangeName, eventPb.String())

	requestDpo := flightoperationIDRequestDpoHolder{id: eventPb.GetFlightoperationId()}
	if ret := h.app.Services.ManageFlightreport.CreateFlightreport(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

type flightoperationIDRequestDpoHolder struct {
	id string
}

func (h *flightoperationIDRequestDpoHolder) GetFlightoperationID() string {
	return h.id
}
