package mavlink

import (
	"context"
	"io"

	"edge/pkg/edge"
	"edge/pkg/edge/common"
	mavsdk_rpc_telemetry "edge/pkg/protos/telemetry"

	"google.golang.org/grpc"
)

// AdapterVelocity .
func AdapterVelocity(ctx context.Context, gr *grpc.ClientConn, support common.Support) (<-chan *edge.Velocity, error) {
	telemetry := mavsdk_rpc_telemetry.NewTelemetryServiceClient(gr)

	velocityReceiver, err := AdapterVelocityInternal(ctx, support, telemetry)
	if err != nil {
		return nil, err
	}

	velocityStream := AdapterVelocitySubscriber(velocityReceiver, support)

	return velocityStream, nil
}

// AdapterVelocityInternal .
func AdapterVelocityInternal(
	ctx context.Context,
	support common.Support,
	telemetry mavsdk_rpc_telemetry.TelemetryServiceClient,
) (velocityReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient, err error) {
	defer func() {
		if err != nil {
			support.NotifyError("velocity telemetry error: %v", err)
		}
	}()

	velocityRequest := mavsdk_rpc_telemetry.SubscribeGroundSpeedNedRequest{}
	velocityReceiver, err = telemetry.SubscribeGroundSpeedNed(ctx, &velocityRequest)

	return
}

// AdapterVelocitySubscriber .
func AdapterVelocitySubscriber(
	velocityReceiver mavsdk_rpc_telemetry.TelemetryService_SubscribeGroundSpeedNedClient,
	support common.Support,
) <-chan *edge.Velocity {
	velocityStream := make(chan *edge.Velocity)

	go func() {
		defer func() {
			if err := velocityReceiver.CloseSend(); err != nil {
				support.NotifyError("velocity telemetry error: %v", err)
			}
		}()
		defer close(velocityStream)
		func() {
			var err error
			defer func() {
				if err != nil {
					support.NotifyError("velocity receive error: %v", err)
					return
				}
				support.NotifyInfo("velocity receive finish")
			}()
			for {
				response, ret := velocityReceiver.Recv()
				if ret == io.EOF {
					return
				}
				if ret != nil {
					err = ret
					return
				}
				velocity := &edge.Velocity{
					North: float64(response.GetGroundSpeedNed().GetVelocityNorthMS()),
					East:  float64(response.GetGroundSpeedNed().GetVelocityEastMS()),
					Down:  float64(response.GetGroundSpeedNed().GetVelocityDownMS()),
				}
				velocityStream <- velocity
			}
		}()
	}()

	return velocityStream
}
