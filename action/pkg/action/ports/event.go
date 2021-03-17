package ports

import (
	"action/pkg/action/app"
	"action/pkg/skysign_proto"
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

	glog.Infof("RECEIVE , Event: %s, Message: %s", CopiedVehicleCreatedEventExchangeName, eventPb.String())

	requestDpo := createRequestDpoHolder{event: &eventPb}
	if ret := h.app.Services.ManageAction.CreateAction(&requestDpo); ret != nil {
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

	glog.Infof("RECEIVE , Event: %s, Message: %s", FlightoperationCompletedEventExchangeName, eventPb.String())

	requestDpo := completeRequestDpoHolder{event: &eventPb}
	if ret := h.app.Services.OperateAction.CompleteAction(&requestDpo); ret != nil {
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

	glog.Infof("RECEIVE , Event: %s, Message: %s", TelemetryUpdatedEventExchangeName, eventPb.String())

	requestDpo := pushTelemetryRequestDpoHolder{event: &eventPb}
	if ret := h.app.Services.OperateAction.PushTelemetry(&requestDpo); ret != nil {
		return ret
	}
	return nil
}

type createRequestDpoHolder struct {
	event *skysign_proto.CopiedVehicleCreatedEvent
}

func (h *createRequestDpoHolder) GetID() string {
	return h.event.VehicleId
}

func (h *createRequestDpoHolder) GetCommunicationID() string {
	return h.event.CommunicationId
}

func (h *createRequestDpoHolder) GetFlightplanID() string {
	return h.event.FlightplanId
}

type completeRequestDpoHolder struct {
	event *skysign_proto.FlightoperationCompletedEvent
}

func (h *completeRequestDpoHolder) GetFlightplanID() string {
	return h.event.FlightplanId
}

type pushTelemetryRequestDpoHolder struct {
	event *skysign_proto.TelemetryUpdatedEvent
}

func (h *pushTelemetryRequestDpoHolder) GetCommunicationID() string {
	return h.event.CommunicationId
}
func (h *pushTelemetryRequestDpoHolder) GetLatitude() float64 {
	return h.event.Telemetry.Latitude
}
func (h *pushTelemetryRequestDpoHolder) GetLongitude() float64 {
	return h.event.Telemetry.Longitude
}
func (h *pushTelemetryRequestDpoHolder) GetAltitude() float64 {
	return h.event.Telemetry.Altitude
}
func (h *pushTelemetryRequestDpoHolder) GetRelativeAltitude() float64 {
	return h.event.Telemetry.RelativeAltitude
}
func (h *pushTelemetryRequestDpoHolder) GetSpeed() float64 {
	return h.event.Telemetry.Speed
}
func (h *pushTelemetryRequestDpoHolder) GetArmed() bool {
	return h.event.Telemetry.Armed
}
func (h *pushTelemetryRequestDpoHolder) GetFlightMode() string {
	return h.event.Telemetry.FlightMode
}
func (h *pushTelemetryRequestDpoHolder) GetOrientationX() float64 {
	return h.event.Telemetry.OrientationX
}
func (h *pushTelemetryRequestDpoHolder) GetOrientationY() float64 {
	return h.event.Telemetry.OrientationY
}
func (h *pushTelemetryRequestDpoHolder) GetOrientationZ() float64 {
	return h.event.Telemetry.OrientationZ
}
func (h *pushTelemetryRequestDpoHolder) GetOrientationW() float64 {
	return h.event.Telemetry.OrientationW
}
