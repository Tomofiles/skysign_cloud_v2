package api

import (
	"context"
	"flightplan/pkg/flightplan/app"
	"flightplan/pkg/flightplan/domain/flightplan"
	"log"
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
	event flightplan.CreatedEvent,
) {
	requestDpo := requestDpoHolder{id: event.GetFlightplanID()}
	ret := h.app.Services.ManageFleet.CreateFleet(&requestDpo)
	if ret != nil {
		log.Println("handle created event error")
	}
}

// HandleDeletedEvent .
func (h *EventHandler) HandleDeletedEvent(
	ctx context.Context,
	event flightplan.DeletedEvent,
) {
	requestDpo := requestDpoHolder{id: event.GetFlightplanID()}
	ret := h.app.Services.ManageFleet.DeleteFleet(&requestDpo)
	if ret != nil {
		log.Println("handle created event error")
	}
}

type requestDpoHolder struct {
	id string
}

func (h *requestDpoHolder) GetFlightplanID() string {
	return h.id
}
