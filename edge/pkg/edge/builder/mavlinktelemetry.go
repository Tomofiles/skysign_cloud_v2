package builder

import (
	"context"
	mavlink "edge/pkg/edge/adapters/mavlink/telemetry"
	"edge/pkg/edge/common"
	"edge/pkg/edge/telemetry"
	"log"

	"google.golang.org/grpc"
)

// MavlinkTelemetry .
func MavlinkTelemetry(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan interface{}, telemetry.Telemetry, error) {
	connStateStream, err := mavlink.AdapterConnectionState(ctx, gr, support)
	if err != nil {
		log.Println("mavlink connState adapter error:", err)
		return nil, nil, err
	}
	positionStream, err := mavlink.AdapterPosition(ctx, gr, support)
	if err != nil {
		log.Println("mavlink position adapter error:", err)
		return nil, nil, err
	}
	quaternionStream, err := mavlink.AdapterQuaternion(ctx, gr, support)
	if err != nil {
		log.Println("mavlink quaternion adapter error:", err)
		return nil, nil, err
	}
	velocityStream, err := mavlink.AdapterVelocity(ctx, gr, support)
	if err != nil {
		log.Println("mavlink velocity adapter error:", err)
		return nil, nil, err
	}
	armedStream, err := mavlink.AdapterArmed(ctx, gr, support)
	if err != nil {
		log.Println("mavlink armed adapter error:", err)
		return nil, nil, err
	}
	flightModeStream, err := mavlink.AdapterFlightMode(ctx, gr, support)
	if err != nil {
		log.Println("mavlink flightMode adapter error:", err)
		return nil, nil, err
	}

	telemetry := telemetry.NewTelemetry()
	updateExit := telemetry.Updater(
		ctx.Done(),
		connStateStream,
		positionStream,
		quaternionStream,
		velocityStream,
		armedStream,
		flightModeStream,
	)

	return updateExit, telemetry, nil
}
