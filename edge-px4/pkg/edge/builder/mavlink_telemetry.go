package builder

import (
	"context"
	"fmt"

	mavlink_telemetry_adapter "github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/adapters/mavlink/telemetry"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/common"
	"github.com/Tomofiles/skysign_cloud_v2/edge-px4/pkg/edge/domain/model"

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
		return nil, fmt.Errorf("connectionState adapter error: %w", err)
	}
	positionStream, err := mavlink_telemetry_adapter.AdapterPosition(ctx, gr, support)
	if err != nil {
		return nil, fmt.Errorf("position adapter error: %w", err)
	}
	quaternionStream, err := mavlink_telemetry_adapter.AdapterQuaternion(ctx, gr, support)
	if err != nil {
		return nil, fmt.Errorf("quaternion adapter error: %w", err)
	}
	velocityStream, err := mavlink_telemetry_adapter.AdapterVelocity(ctx, gr, support)
	if err != nil {
		return nil, fmt.Errorf("velocity adapter error: %w", err)
	}
	armedStream, err := mavlink_telemetry_adapter.AdapterArmed(ctx, gr, support)
	if err != nil {
		return nil, fmt.Errorf("armed adapter error: %w", err)
	}
	flightModeStream, err := mavlink_telemetry_adapter.AdapterFlightMode(ctx, gr, support)
	if err != nil {
		return nil, fmt.Errorf("flightMode adapter error: %w", err)
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
