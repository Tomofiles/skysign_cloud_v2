package rabbitmq

import (
	"context"

	"github.com/Tomofiles/skysign_cloud_v2/collection-analysis/pkg/action/app"

	"github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const (
	// TelemetryUpdatedEventExchangeName .
	TelemetryUpdatedEventExchangeName = "communication.telemetry_updated_event"
	// TelemetryUpdatedEventQueueName .
	TelemetryUpdatedEventQueueName = "action.telemetry_updated_event"
)

// TelemetryUpdatedEventHandler .
type TelemetryUpdatedEventHandler interface {
	HandleTelemetryUpdatedEvent(ctx context.Context, event []byte) error
}

type telemetryUpdatedEventHandler struct {
	app app.Application
}

// NewTelemetryUpdatedEventHandler .
func NewTelemetryUpdatedEventHandler(application app.Application) TelemetryUpdatedEventHandler {
	return &telemetryUpdatedEventHandler{app: application}
}

// HandleTelemetryUpdatedEvent .
func (h *telemetryUpdatedEventHandler) HandleTelemetryUpdatedEvent(
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

type pushTelemetryCommandHolder struct {
	event *skysign_proto.TelemetryUpdatedEvent
}

func (h *pushTelemetryCommandHolder) GetCommunicationID() string {
	return h.event.CommunicationId
}
func (h *pushTelemetryCommandHolder) GetLatitudeDegree() float64 {
	return h.event.Telemetry.Latitude
}
func (h *pushTelemetryCommandHolder) GetLongitudeDegree() float64 {
	return h.event.Telemetry.Longitude
}
func (h *pushTelemetryCommandHolder) GetAltitudeM() float64 {
	return h.event.Telemetry.Altitude
}
func (h *pushTelemetryCommandHolder) GetRelativeAltitudeM() float64 {
	return h.event.Telemetry.RelativeAltitude
}
func (h *pushTelemetryCommandHolder) GetSpeedMS() float64 {
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
