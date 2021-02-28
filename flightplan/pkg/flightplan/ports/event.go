package ports

import (
	"context"
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// FlightplanCreatedEventExchangeName .
	FlightplanCreatedEventExchangeName = "flightplan.flightplan_created_event"
	// FlightplanDeletedEventExchangeName .
	FlightplanDeletedEventExchangeName = "flightplan.flightplan_deleted_event"
	// FlightplanCopiedEventExchangeName .
	FlightplanCopiedEventExchangeName = "flightplan.flightplan_copied_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleCreatedEvent .
func (h *EventHandler) HandleCreatedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FlightplanCreatedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightplanCreatedEventExchangeName, eventPb.String())

	requestDpo := flightplanIDRequestDpoHolder{id: eventPb.GetFlightplanId()}
	if ret := h.app.Services.ManageFleet.CreateFleet(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

// HandleDeletedEvent .
func (h *EventHandler) HandleDeletedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FlightplanDeletedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightplanDeletedEventExchangeName, eventPb.String())

	requestDpo := flightplanIDRequestDpoHolder{id: eventPb.GetFlightplanId()}
	if ret := h.app.Services.ManageFleet.DeleteFleet(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

// HandleCopiedEvent .
func (h *EventHandler) HandleCopiedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.FlightplanCopiedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightplanCopiedEventExchangeName, eventPb.String())

	requestDpo := copyRequestDpoHolder{
		originalID: eventPb.GetOriginalFlightplanId(),
		newID:      eventPb.GetNewFlightplanId(),
	}
	if ret := h.app.Services.ManageFleet.CarbonCopyFleet(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

type flightplanIDRequestDpoHolder struct {
	id string
}

func (h *flightplanIDRequestDpoHolder) GetFlightplanID() string {
	return h.id
}

type copyRequestDpoHolder struct {
	originalID string
	newID      string
}

func (h *copyRequestDpoHolder) GetOriginalID() string {
	return h.originalID
}
func (h *copyRequestDpoHolder) GetNewID() string {
	return h.newID
}
