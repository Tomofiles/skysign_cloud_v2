package builder

import (
	"context"
	"edge/pkg/edge"
	mavlink "edge/pkg/edge/adapters/mavlink/telemetry"
	"edge/pkg/edge/common"

	"google.golang.org/grpc"
)

// TelemetryStream .
type TelemetryStream struct {
	ConnectionStateStream <-chan *edge.ConnectionState
	PositionStream        <-chan *edge.Position
	QuaternionStream      <-chan *edge.Quaternion
	VelocityStream        <-chan *edge.Velocity
	ArmedStream           <-chan *edge.Armed
	FlightModeStream      <-chan *edge.FlightMode
}

// MavlinkTelemetry .
func MavlinkTelemetry(
	ctx context.Context,
	gr *grpc.ClientConn,
	support common.Support,
) (*TelemetryStream, error) {
	connectionStateStream, err := mavlink.AdapterConnectionState(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink connState adapter error: %v", err)
		return nil, err
	}
	positionStream, err := mavlink.AdapterPosition(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink position adapter error: %v", err)
		return nil, err
	}
	quaternionStream, err := mavlink.AdapterQuaternion(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink quaternion adapter error: %v", err)
		return nil, err
	}
	velocityStream, err := mavlink.AdapterVelocity(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink velocity adapter error: %v", err)
		return nil, err
	}
	armedStream, err := mavlink.AdapterArmed(ctx, gr, support)
	if err != nil {
		support.NotifyError("mavlink armed adapter error: %v", err)
		return nil, err
	}
	flightModeStream, err := mavlink.AdapterFlightMode(ctx, gr, support)
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
