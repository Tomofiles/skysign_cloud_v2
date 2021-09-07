package builder

import (
	"context"
	mavlink_telemetry_adapter "edge-px4/pkg/edge/adapters/mavlink/telemetry"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"

	"google.golang.org/grpc"
)

// TelemetryStream .
type TelemetryStream struct {
	ConnectionStateStream <-chan *model.ConnectionState
	PositionStream        <-chan *model.Position
	QuaternionStream      <-chan *model.Quaternion
	VelocityStream        <-chan *model.Velocity
	ArmedStream           <-chan *model.Armed
	FlightModeStream      <-chan *model.FlightMode
}

// MavlinkTelemetry .
func MavlinkTelemetry(
	ctx context.Context,
	gr *grpc.ClientConn,
	support common.Support,
) (*TelemetryStream, error) {
	connectionStateStream, err := mavlink_telemetry_adapter.AdapterConnectionState(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink connectionStateStream adapter error: %v", err)
		return nil, err
	}
	positionStream, err := mavlink_telemetry_adapter.AdapterPosition(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink position adapter error: %v", err)
		return nil, err
	}
	quaternionStream, err := mavlink_telemetry_adapter.AdapterQuaternion(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink quaternion adapter error: %v", err)
		return nil, err
	}
	velocityStream, err := mavlink_telemetry_adapter.AdapterVelocity(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink velocity adapter error: %v", err)
		return nil, err
	}
	armedStream, err := mavlink_telemetry_adapter.AdapterArmed(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink armed adapter error: %v", err)
		return nil, err
	}
	flightModeStream, err := mavlink_telemetry_adapter.AdapterFlightMode(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink flightMode adapter error: %v", err)
		return nil, err
	}

	telemetryStream := &TelemetryStream{
		ConnectionStateStream: connectionStateStream,
		PositionStream:        positionStream,
		QuaternionStream:      quaternionStream,
		VelocityStream:        velocityStream,
		ArmedStream:           armedStream,
		FlightModeStream:      flightModeStream,
	}
	return telemetryStream, nil
}
