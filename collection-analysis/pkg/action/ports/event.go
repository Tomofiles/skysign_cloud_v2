package ports

import (
	"collection-analysis/pkg/action/app"
	"collection-analysis/pkg/skysign_proto"
	"context"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// CopiedVehicleCreatedEventExchangeName .
	CopiedVehicleCreatedEventExchangeName = "vehicle.copied_vehicle_created_event"
	// CopiedVehicleCreatedEventQueueName .
	CopiedVehicleCreatedEventQueueName = "action.copied_vehicle_created_event"
	// FlightoperationCompletedEventExchangeName .
	FlightoperationCompletedEventExchangeName = "flightoperation.flightoperation_completed_event"
	// FlightoperationCompletedEventQueueName .
	FlightoperationCompletedEventQueueName = "action.flightoperation_completed_event"
	// TelemetryUpdatedEventExchangeName .
	TelemetryUpdatedEventExchangeName = "communication.telemetry_updated_event"
	// TelemetryUpdatedEventQueueName .
	TelemetryUpdatedEventQueueName = "action.telemetry_updated_event"
)

// EventHandler .
type EventHandler struct {
	app app.Application
}

// NewEventHandler .
func NewEventHandler(application app.Application) EventHandler {
	return EventHandler{app: application}
}

// HandleCopiedVehicleCreatedEvent .
func (h *EventHandler) HandleCopiedVehicleCreatedEvent(
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

	command := completeCommandHolder{event: &eventPb}
	if ret := h.app.Services.OperateAction.CompleteAction(&command); ret != nil {
		return ret
	}
	return nil
}

// HandleTelemetryUpdatedEvent .
func (h *EventHandler) HandleTelemetryUpdatedEvent(
	ctx context.Context,
	event []byte,
) error {
	eventPb := skysign_proto.TelemetryUpdatedEvent{}
	if err := proto.Unmarshal(event, &eventPb); err != nil {
		return err
	}

	glog.Infof("RECEIVE , Event: %s, Message: %s", TelemetryUpdatedEventQueueName, eventPb.String())

	command := pushTelemetryCommandHolder{event: &eventPb}
	if ret := h.app.Services.OperateAction.PushTelemetry(&command); ret != nil {
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

type completeCommandHolder struct {
	event *skysign_proto.FlightoperationCompletedEvent
}

func (h *completeCommandHolder) GetFleetID() string {
	return h.event.Flightoperation.FleetId
}

type pushTelemetryCommandHolder struct {
	event *skysign_proto.TelemetryUpdatedEvent
}

func (h *pushTelemetryCommandHolder) GetCommunicationID() string {
	return h.event.CommunicationId
}
func (h *pushTelemetryCommandHolder) GetLatitude() float64 {
	return h.event.Telemetry.Latitude
}
func (h *pushTelemetryCommandHolder) GetLongitude() float64 {
	return h.event.Telemetry.Longitude
}
func (h *pushTelemetryCommandHolder) GetAltitude() float64 {
	return h.event.Telemetry.Altitude
}
func (h *pushTelemetryCommandHolder) GetRelativeAltitude() float64 {
	return h.event.Telemetry.RelativeAltitude
}
func (h *pushTelemetryCommandHolder) GetSpeed() float64 {
	return h.event.Telemetry.Speed
}
func (h *pushTelemetryCommandHolder) GetArmed() bool {
	return h.event.Telemetry.Armed
}
func (h *pushTelemetryCommandHolder) GetFlightMode() string {
	return h.event.Telemetry.FlightMode
}
func (h *pushTelemetryCommandHolder) GetOrientationX() float64 {
	return h.event.Telemetry.OrientationX
}
func (h *pushTelemetryCommandHolder) GetOrientationY() float64 {
	return h.event.Telemetry.OrientationY
}
func (h *pushTelemetryCommandHolder) GetOrientationZ() float64 {
	return h.event.Telemetry.OrientationZ
}
func (h *pushTelemetryCommandHolder) GetOrientationW() float64 {
	return h.event.Telemetry.OrientationW
}
